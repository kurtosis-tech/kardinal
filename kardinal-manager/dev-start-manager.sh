UUID=$(cat "$HOME/Library/Application Support/kardinal/fk-tenant-uuid")
KARDINAL_MANAGER_CLUSTER_CONFIG_ENDPOINT="http://localhost:8080/tenant/$UUID/cluster-resources" KARDINAL_MANAGER_FETCHER_JOB_DURATION_SECONDS=2 reflex -s -r '\.go$' -- go run kardinal-manager/main.go "$@"
