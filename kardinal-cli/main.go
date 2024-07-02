package main

import "kardinal.cli/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		println("Error:", err.Error())
	}
}
