[![Docker Hub](https://img.shields.io/badge/dockerhub-images-important.svg?logo=docker)](https://hub.docker.com/u/kurtosistech)

# Kardinal

![Kardi B](https://kardinal.dev/_next/static/media/kardinal-orange.65ea335b.png)

Kardinal is a traffic control and data isolation layer that enables engineers to safely do development and QA work directly in production. Say goodbye to maintaining multiple environments and hello to faster, more efficient development workflows.

## Quick install

```bash
curl https://raw.githubusercontent.com/kurtosis-tech/kardinal/main/scripts/install_cli.sh -s | sh
```

## What is Kardinal?

Kardinal injects production data and service dependencies into your dev and test workflows safely and securely. Instead of spinning up ephemeral environments with mocked services, fake traffic, and fake data, developers using Kardinal can put their service directly into the production environment to see how it works... without risking the stability of that environment.

Key features:

- Develop and test directly in production without risk
- Catch bugs that "only appear in prod" faster
- Stop maintaining multiple environments - do it all in production
- Lighter-weight dev workflow: reuse deployed services
- Implement isolated dev sandbox flows with maximum dev-prod parity
- Control data and traffic access throughout the software development lifecycle with maturity gates

## How it Works

Kardinal uses traffic flow controls and a data isolation layer to protect production while you're developing. It achieves this by rethinking the idea of isolated "environments" and replacing them with isolated traffic flows within the production environment.

To use Kardinal, just drop the Kardinal sidecars into your production environment. Then run:

```bash
# Create a dev flow
kardinal create-dev-flow <service-name> <dev-image-tag>
```

This creates a dev flow for your service with access to all the data, traffic, and services in your production environment, while ensuring complete isolation and safety.

## Join the Beta Program

Our beta program for a select group of developers is coming soon! Be among the first to experience the future of software development:

- Email us at: [hello@kardinal.dev](mailto:hello@kardinal.dev) to learn more about joining the beta
- Get early access to Kardinal
- Provide valuable feedback to shape the future of the product

## Try the Kardinal Playground

Can't wait to get started? Check out our proof of concept:

- Visit the [Kardinal playground](https://github.com/kurtosis-tech/kardinal-playground/)
- Experience a simple demonstration of how Kardinal can enhance your development workflow
- Get a taste of developing directly in production, risk-free

## Get Support

Have questions or need assistance? We're here to help:

- Email us at: [hello@kardinal.dev](mailto:hello@kardinal.dev)
- Check out our documentation: [https://kardinal.dev/docs](https://kardinal.dev/docs)

---

## Architecture

Kardinal main components are the Kardinal CLI and the Kardinal Manager. The Kardinal CLI allows the user to manage the development flows. The Kardinal Manager retrieves the latest configuration from the Kardinal Cloud and applies changes to the K8S user services topology.

![kardinal-dev-overview](./img/kardinal-dev-overview.png?raw=true)

The Kardinal Cloud code is not open-source.

### Kardinal CLI

The Kardinal CLI is a standalone tool interacting with the Kardinal Cloud to manage the dev flows.

### Kardinal Manager

The Kardinal Manager retrieves the latest user services topology from the Kardinal Cloud and applies the changes by interacting with the Istio client and K8S client. The Manager manages traffic using Istio objects such as virtual services and destination rules. The Manager also updates the K8S services and deployments.

## Quickstart

## Deploying Kardinal on a Kubernetes Cluster

These instructions provide a guide for deploying Kardinal on any Kubernetes cluster, whether it's a local setup like Minikube, a managed cloud service, or your own self-hosted cluster. We'll use kubectl port-forwarding to access the services, which works universally across different Kubernetes setups.

### Prerequisites

- A Kubernetes cluster (e.g., Minikube, EKS, GKE, AKS, or any other Kubernetes distribution)
- kubectl installed and configured to access your cluster

### Steps

1. Install the Kardinal CLI:

```bash
curl https://raw.githubusercontent.com/kurtosis-tech/kardinal/main/scripts/install_cli.sh -s | sh
```

2. Install Istio (if not already installed):

If you don't already have Istio installed in your cluster, follow these steps to install it:

```bash
curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.22.1 TARGET_ARCH=x86_64 sh -
cd istio-1.22.1
export PATH=$PWD/bin:$PATH
echo 'export PATH=$PATH:'"$PWD/bin" >> ~/.bashrc
istioctl install --set profile=demo -y
cd ..
```

This will download Istio, add it to your PATH, and install it in your cluster. The output should look similar to this:

```
✔ Istio core installed
✔ Istiod installed
✔ Egress gateways installed
✔ Ingress gateways installed
✔ Installation complete
Making this installation the default for injection and validation.

Thank you for installing Istio 1.22.1.

To help improve the product, please consider participating in our survey at https://forms.gle/kWRLt3tAQLynXL5t9

Next steps:
  1. Ensure the following pods are running and ready:
    kubectl get pods -n istio-system
  2. Visit https://istio.io/latest/docs/setup/getting-started/ to learn about next steps.
```

If you already have Istio installed, you can skip this step.

3. Deploy the Kardinal Manager:

```bash
kardinal manager deploy kloud-kontrol
```

4. Note the tenant UUID generated during this process. You'll need this to check your traffic configuration.

5. Clone the Kardinal Playground repository to get the voting app demo:

```bash
git clone https://github.com/kurtosis-tech/kardinal-playground.git
cd kardinal-playground/voting-app-demo
```

6. Deploy the voting-app application with Kardinal:

```bash
kardinal deploy --docker-compose docker-compose.yaml
```

7. Check the initial Kardinal traffic configuration:
   Visit https://app.kardinal.dev/{your-tenant-id} (replace {your-tenant-id} with the UUID from step 4)
   You should see only the production version of your application in the traffic configuration.

8. Use the following script to set up port-forwarding for accessing the services. Save this script as `kardinal-port-forward.sh`:

```bash
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
        echo "Command failed. Retrying in $retry_delay seconds..."
        sleep $retry_delay
        retry_delay=$((retry_delay * 2))
        ((retries++))
    done

    echo "Max retries reached. Command failed."
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
```

9. Make the script executable:

```bash
chmod +x kardinal-port-forward.sh
```

10. Run the script to set up port-forwarding for the production version:

```bash
./kardinal-port-forward.sh prod
```

This will set up port-forwarding for the production version of the voting app.

11. Access the production application:
    - Production version: http://localhost:8090

12. To create a new development flow:

```bash
kardinal flow create voting-app-ui voting-app-ui-dev -d compose.yml
```

13. After creating the development flow, check the Kardinal traffic configuration again:
    Visit https://app.kardinal.dev/{your-tenant-id}
    You should now see both the production and development versions of your application in the traffic configuration.

14. Run the port-forwarding script again to include the new development version:

```bash
./kardinal-port-forward.sh all
```

Now you can access both the production and development versions:
   - Production version: http://localhost:8090
   - Development version: http://localhost:8091

15. To remove the development flow:

```bash
kardinal flow delete -d compose.yml
```

16. After deleting the development flow, check the Kardinal traffic configuration once more:
    Visit https://app.kardinal.dev/{your-tenant-id}
    You should now see only the production version of your application in the traffic configuration, confirming that the development flow has been removed.

17. Clean up:
    - Stop the port-forwarding: `pkill -f "kubectl port-forward.*voting-app"`
    - Remove Kardinal Manager: `kardinal manager remove`
    - Remove the voting-app: `kubectl delete ns prod`

By following these steps, you can deploy and manage Kardinal on any Kubernetes cluster, using kubectl port-forwarding to access the services. This method works universally across different Kubernetes setups, including Minikube, cloud-managed Kubernetes services, and self-hosted clusters. 

Remember to check the Kardinal traffic configuration at https://app.kardinal.dev/{your-tenant-id} before and after creating or deleting development flows to verify the changes in your application's topology.

### How to run Kardinal and use the voting app example to test the dev flow

#### Prerequisites

You will need the following tools installed (they will be already available if you are using the nix shell provided by this repository):

- A local Kubernetes cluster ([Minikube](https://minikube.sigs.k8s.io/docs/start/?arch=%2Fmacos%2Fx86-64%2Fstable%2Fbinary+download) used in this example)
- Istio resources installed in the local cluster (use the [getting started doc](https://istio.io/latest/docs/setup/getting-started/#download))

```bash
minikube start --driver=docker --cpus=10 --memory 8192 --disk-size 32g
minikube addons enable ingress
minikube addons enable metrics-server
istioctl install --set profile=default -y
minikube dashboard
```

- Both `prod.app.localhost` and `dev.app.localhost` defined in the host file

```bash
# Add these entries in the '/private/etc/hosts' file
127.0.0.1 prod.app.localhost
127.0.0.1 dev.app.localhost
```

#### Steps

##### Deploy the production voting app

1. Use the `kardinal` provided by the Nix shell (enter using `nix develop`) or follow [this to build and run the cli][run-build-cli]
2. Deploy `Kardinal Manager` in the local kubernetes cluster and set the `Kardinal Control` location (we are going to use the cloud version on these steps)

```bash
kardinal manager deploy kloud-kontrol
```

3. Copy the tenant UUID generated while running this command

```bash
# This log line will be printed in the terminal, copy the generated UUID
INFO[0000] Using tenant UUID 58d33536-3c9e-4110-aa83-bf112ae94a49
```

3. Deploy the voting-app application with Kardinal

```bash
kardinal deploy --docker-compose ./examples/voting-app/docker-compose.yaml
```

4. Check the current topology in the cloud Kontrol FE using this URL: https://app.kardinal.dev/{use-your-tenant-UUID-here}/traffic-configuration
5. Start the tunnel to access the services (you may have to provide you password for the underlying sudo access)

```bash
minikube tunnel
```

6. Open the [production page in the browser](http://prod.app.localhost/) to see the production `voting-app`

##### Deploy the voting app development version in the same cluster

1. Create a new flow to test a development `voting-app-ui-v2` version in production

```bash
kardinal flow create voting-app-ui voting-app-ui-v2 --docker-compose ./examples/voting-app/docker-compose.yaml
```

2. Check how the topology has changed, to reflect both prod and the dev version, in the cloud Kontrol FE using this URL: https://app.kardinal.dev/{use-your-tenant-UUID-here}/traffic-configuration
3. Open the [development voting-app-ui-v2 page in the browser](http://dev.app.localhost/) to see the development `voting-app-ui-v2`

##### Remove the voting app development version from the same cluster

1. Remove the flow created for the `voting-app-ui-v2`

```bash
kardinal flow delete --docker-compose ./examples/voting-app/docker-compose.yaml
```

2. Check the topology again to, it's showing only the production version as the beginning, in the cloud Kontrol FE using this URL: https://app.kardinal.dev/{use-your-tenant-UUID-here}/traffic-configuration
3. Open the [development voting-app-ui-v2 page in the browser](http://dev.app.localhost/) to check that it was successfully removed
4. Open the [production page in the browser](http://prod.app.localhost/) to check that it didn't change

##### Clean

1. Remove `Kardinal Manager` from the cluster

```bash
kardinal manager remove
```

2. Remove the `voting-app` application from the cluster

```bash
kubectl delete ns prod
```

## Development instructions

1. Enter the dev shell and start the local cluster:

```bash
nix develop
```

2. You're also likely to use a local k8s, in this case minikube is available to use:

```bash
kubectl config set-context minikube
minikube start --driver=docker --cpus=10 --memory 8192 --disk-size 32g
minikube addons enable ingress
minikube addons enable metrics-server
istioctl install --set profile=demo -y
minikube dashboard
```

On a second terminal, start the tunnel:

```bash
minikube tunnel
```

## Deploying Kardinal Manager to local cluster

Configure it by setting the following environment variables:

```bash
KARDINAL_MANAGER_CLUSTER_CONFIG_ENDPOINT=http://localhost:8080/tenant/{36e22127-3c9e-4110-aa83-af552cd94b88}/cluster-resources
KARDINAL_MANAGER_FETCHER_JOB_DURATION_SECONDS=10
```

or in the `kardinal-manager/deployment/k8s.yaml`:

```yaml
  env:
    - name: KARDINAL_MANAGER_CLUSTER_CONFIG_ENDPOINT
      # This is valid for reaching out the Kardinal Kontrol if this is running on the host
     value: "http://localhost:8080/tenant/36e22127-3c9e-4110-aa83-af552cd94b88/cluster-resources"
    - name: KARDINAL_MANAGER_FETCHER_JOB_DURATION_SECONDS
    value: "10"
```

NOTE: you can get your tenant UUID by running any CLI command

You can use tilt deploy and keeping the image hot-reloading:

```bash
tilt up
```

Or manually build it:

```bash
# First set the docker context to minikube
eval $(minikube docker-env)
docker load < $(nix build ./#kardinal-manager-container --no-link --print-out-paths)
kubectl apply -f kardinal-manager/deployment
```

## Deploying Redis Overlay Service to local cluster

Building and loading image into minikube:

```bash
# First set the docker context to minikube
eval $(minikube docker-env)
docker load < $(nix build ./#redis-proxy-overlay-container --no-link --print-out-paths)
```

To build and run the service directly:

```bash
nix run ./#redis-proxy-overlay
```

## Publishing multi-arch images

To publish multi-arch images, you can use the following command:

```bash
$(nix build .#publish-<SERVICE_NAME>-container --no-link --print-out-paths)/bin/push

# For instance, to publish the redis proxy overlay image:
$(nix build .#publish-redis-proxy-overlay-container --no-link --print-out-paths)/bin/push
```

## Running Kardinal CLI

To build and run the service directly:

```bash
nix run ./#kardinal-cli
```

### Regenerate gomod2nix.toml

You will need to do this every time a `go.mod` file is edited

```bash
nix develop
gomod2nix generate
```

<!--------------- ONLY LINKS BELOW THIS POINT ---------------------->

[run-build-cli]: #running-kardinal-cli
