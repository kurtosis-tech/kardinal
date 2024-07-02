[![Docker Hub](https://img.shields.io/badge/dockerhub-images-important.svg?logo=docker)](https://hub.docker.com/u/kurtosistech)

# Kardinal

![Kardi B](https://kardinal.dev/_next/static/media/kardinal-orange.65ea335b.png)

Kardinal is a traffic control and data isolation layer that enables engineers to safely do development and QA work directly in production. Say goodbye to maintaining multiple environments and hello to faster, more efficient development workflows.

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

Kardinal main components are the Kardinal CLI and the Kardinal Manager.  The Kardinal CLI allows the user to manage the development flows.  The Kardinal Manager retrieves the latest configuration from the Kardinal Cloud and applies changes to the K8S user services topology.

![kardinal-dev-overview](./img/kardinal-dev-overview.png?raw=true)

The Kardinal Cloud code is not open-source.

### Kardinal CLI

The Kardinal CLI is a standalone tool interacting with the Kardinal Cloud to manage the dev flows.

### Kardinal Manager

The Kardinal Manager retrieves the latest user services topology from the Kardinal Cloud and applies the changes by interacting with the Istio client and K8S client. The Manager manages traffic using Istio objects such as virtual services and destination rules. The Manager also updates the K8S services and deployments.

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
     value: "http://localhost:8080/tenant/{36e22127-3c9e-4110-aa83-af552cd94b88}/cluster-resources"
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
