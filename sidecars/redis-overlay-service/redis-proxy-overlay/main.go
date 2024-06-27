package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/tidwall/redcon"
)

func main() {
	var mu sync.RWMutex
	items := make(map[string][]byte)
	var ps redcon.PubSub
	redisAddr := os.Getenv("REDIS_ADDR")
	addr := ":" + os.Getenv("PORT")

	go log.Printf("started server at %s", addr)

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := redcon.ListenAndServe(addr,
		func(conn redcon.Conn, cmd redcon.Command) {
			argsIndex := 0
			for argsIndex < len(cmd.Args) {
				log.Printf("Command arg %d: '%s'", argsIndex, string(cmd.Args[argsIndex]))
				argsIndex += 1
			}
			switch strings.ToLower(string(cmd.Args[0])) {
			default:
				conn.WriteError("ERR unknown command '" + string(cmd.Args[0]) + "'")
			case "publish":
				// Publish to all pub/sub subscribers and return the number of
				// messages that were sent.
				if len(cmd.Args) != 3 {
					conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
					return
				}
				count := ps.Publish(string(cmd.Args[1]), string(cmd.Args[2]))
				conn.WriteInt(count)
			case "subscribe", "psubscribe":
				// Subscribe to a pub/sub channel. The `Psubscribe` and
				// `Subscribe` operations will detach the connection from the
				// event handler and manage all network I/O for this connection
				// in the background.
				if len(cmd.Args) < 2 {
					conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
					return
				}
				command := strings.ToLower(string(cmd.Args[0]))
				for i := 1; i < len(cmd.Args); i++ {
					if command == "psubscribe" {
						ps.Psubscribe(conn, string(cmd.Args[i]))
					} else {
						ps.Subscribe(conn, string(cmd.Args[i]))
					}
				}
			case "detach":
				hconn := conn.Detach()
				log.Printf("connection has been detached")
				go func() {
					defer hconn.Close()
					hconn.WriteString("OK")
					hconn.Flush()
				}()
			case "ping":
				conn.WriteString("PONG")
			case "quit":
				conn.WriteString("OK")
				conn.Close()
			case "set":
				if len(cmd.Args) != 3 {
					conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
					return
				}
				mu.Lock()
				items[string(cmd.Args[1])] = cmd.Args[2]
				mu.Unlock()
				conn.WriteString("OK")
			case "incrby":
				if len(cmd.Args) != 3 {
					conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
					return
				}
				incrby, _ := strconv.Atoi(string(cmd.Args[2]))
				mu.RLock()
				val, ok := items[string(cmd.Args[1])]
				mu.RUnlock()
				var valInt int
				if ok {
					var err error
					valInt, err = strconv.Atoi(string(val[:]))
					if err != nil {
						conn.WriteError("Increment by a non integer")
						return
					}
					// #####################################################################
					// TODO: Unless upsteam get! Only to genereate calls to show up in Kiali
					rdb.Get(context.Background(), string(cmd.Args[1])).Result()
					// #####################################################################
				} else {
					// Get the upstream value if local write does not exist yet
					upstreamVal, err := rdb.Get(context.Background(), string(cmd.Args[1])).Result()
					if err == redis.Nil {
						log.Printf("%v does not exist", string(cmd.Args[1]))
						valInt = 0
					} else if err != nil {
						log.Printf("Proxied get command failed: '%v'", err)
						conn.WriteError("Proxied get command failed")
						return
					} else {
						log.Printf("got upval: '%v'", upstreamVal)
						valInt, err = strconv.Atoi(upstreamVal)
						if err != nil {
							log.Printf("Upstream value is not an Int: '%v'", err)
							conn.WriteError("Upstream value for key '" + string(cmd.Args[1]) + "' is not an Int")
							return
						}
					}
				}
				// Update the value in the local cache and return total value
				mu.Lock()
				items[string(cmd.Args[1])] = []byte(strconv.Itoa(valInt + incrby))
				mu.Unlock()
				conn.WriteInt(valInt + incrby)
			case "get":
				if len(cmd.Args) != 2 {
					conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
					return
				}
				mu.RLock()
				val, ok := items[string(cmd.Args[1])]
				mu.RUnlock()
				if !ok {
					res, err := rdb.Get(context.Background(), string(cmd.Args[1])).Result()
					if err == redis.Nil {
						conn.WriteNull()
						return
					}
					if err != nil {
						log.Printf("Proxied get command failed: '%v'", err)
						conn.WriteError("Proxied get command failed")
						return
					}
					mu.Lock()
					items[string(cmd.Args[1])] = []byte(res)
					mu.Unlock()
					conn.WriteString(res)
				} else {
					conn.WriteBulk(val)
				}
			case "exists":
				if len(cmd.Args) != 2 {
					conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
					return
				}
				mu.Lock()
				_, ok := items[string(cmd.Args[1])]
				mu.Unlock()
				if !ok {
					_, err := rdb.Get(context.Background(), string(cmd.Args[1])).Result()
					if err == redis.Nil {
						conn.WriteInt(0)
						return
					} else if err != nil {
						log.Printf("Proxied get command failed: '%v'", err)
						conn.WriteError("Proxied get command failed")
						return
					} else {
						conn.WriteInt(1)
						return
					}
				} else {
					conn.WriteInt(1)
				}
			case "del":
				if len(cmd.Args) != 2 {
					conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
					return
				}
				mu.Lock()
				_, ok := items[string(cmd.Args[1])]
				delete(items, string(cmd.Args[1]))
				mu.Unlock()
				if !ok {
					conn.WriteInt(0)
				} else {
					conn.WriteInt(1)
				}
			case "echo":
				if len(cmd.Args) != 2 {
					conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
					return
				}
				conn.WriteBulk(cmd.Args[1])
			case "client":
				conn.WriteString("OK")
			}
		},
		func(conn redcon.Conn) bool {
			// Use this function to accept or deny the connection.
			// log.Printf("accept: %s", conn.RemoteAddr())
			return true
		},
		func(conn redcon.Conn, err error) {
			// This is called when the connection has been closed
			// log.Printf("closed: %s, err: %v", conn.RemoteAddr(), err)
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
