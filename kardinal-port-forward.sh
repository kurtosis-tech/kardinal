#!/bin/bash

set -euo pipefail

MAX_RETRIES=5
INITIAL_RETRY_DELAY=2

check_pod_status() {
    local resource_name=$1
    local namespace=$2
    local status
    status=$(kubectl get pods -n "$namespace" | grep "^$resource_name" | awk '{print $3}')
    if [ "$status" = "Running" ]; then
        return 0
    elif [ -z "$status" ]; then
        echo "Resource $resource_name in namespace $namespace not found"
        return 1
    else
        echo "Resource $resource_name in namespace $namespace is not running (status: $status)"
        return 1
    fi
}

retry_with_exponential_backoff() {
    local cmd="$1"
    local retry_delay=$INITIAL_RETRY_DELAY
    local retries=0

    while [ $retries -lt $MAX_RETRIES ]; do
        if eval "$cmd"; then
            return 0
        fi
        echo "Port forwarding failed. Retrying in $retry_delay seconds..."
        sleep $retry_delay
        retry_delay=$((retry_delay * 2))
        ((retries++))
    done

    echo "Max retries reached. Port forwarding failed."
    return 1
}

forward_dev() {
    echo "🛠️ Forwarding dev version (voting-app-dev)..."
    if retry_with_exponential_backoff "check_pod_status 'voting-app-ui-dev' 'prod'"; then
        retry_with_exponential_backoff "kubectl port-forward -n prod deploy/voting-app-ui-dev 8091:80 > /dev/null 2>&1 &"
        echo "✅ Dev version forwarded to port 8091"
    else
        echo "❌ Failed to forward dev version: pod is not running after retries"
    fi
}

forward_prod() {
    echo "🚀 Forwarding prod version (voting-app-prod)..."
    if retry_with_exponential_backoff "check_pod_status 'voting-app-ui-prod' 'prod'"; then
        retry_with_exponential_backoff "kubectl port-forward -n prod svc/voting-app-ui 8090:80 > /dev/null 2>&1 &"
        echo "✅ Prod version forwarded to port 8090"
    else
        echo "❌ Failed to forward prod version: pod is not running after retries"
    fi
}

kill_existing_forwards() {
    echo "🔪 Killing existing port-forwards..."
    pkill -f "kubectl port-forward.*voting-app" || true
}

forward_all() {
    kill_existing_forwards
    forward_prod
    if kubectl get deploy -n prod voting-app-ui-dev &> /dev/null; then
        forward_dev
    else
        echo "⚠️ Dev version not found. Skipping dev forwarding."
    fi
}

print_usage() {
    echo "Usage: $0 [dev|prod|all]"
    echo "  dev  : Forward dev version (voting-app-dev) to port 8091 (if it exists)"
    echo "  prod : Forward prod version (voting-app-prod) to port 8090"
    echo "  all  : Forward all available versions (default if no argument is provided)"
}

main() {
    local command=${1:-all}
    
    case $command in
        dev)
            kill_existing_forwards
            if kubectl get deploy -n prod voting-app-ui-dev &> /dev/null; then
                forward_dev
                echo "🎉 Port forwarding complete!"
                echo "🔗 Dev app: http://localhost:8091"
            else
                echo "⚠️ Dev version not found. No forwarding performed."
            fi
            ;;
        prod)
            kill_existing_forwards
            forward_prod
            echo "🎉 Port forwarding complete!"
            echo "🔗 Prod app: http://localhost:8090"
            ;;
        all)
            forward_all
            echo "🎉 Port forwarding complete!"
            echo "🔗 Prod app: http://localhost:8090"
            if kubectl get deploy -n prod voting-app-ui-dev &> /dev/null; then
                echo "🔗 Dev app: http://localhost:8091"
            fi
            ;;
        *)
            print_usage
            exit 1
            ;;
    esac
}

# Call main function with all script arguments
main "$@"
