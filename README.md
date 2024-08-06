[![Docker Hub](https://img.shields.io/badge/dockerhub-images-important.svg?logo=docker)](https://hub.docker.com/u/kurtosistech)

# Kardinal

![Kardi B](https://kardinal.dev/_next/static/media/kardinal-orange.65ea335b.png)

The lightest-weight Kubernetes development environments in the world. Stop duplicating everything across your dev, test, and QA Kubernetes clusters. Deploy the minimum resources necessary to develop and test directly in one production-like environment.

## Want to try it out without installing anything?

Visit the [Kardinal playground](https://github.com/kurtosis-tech/kardinal-playground/) to experience a simple demonstration of how Kardinal can enhance your dev flow.

## What is Kardinal?

Kardinal is a traffic routing and state isolation tool for Kubernetes that enables engineers to efficiently do development, testing, and QA work within a single stable cluster. Instead of implementing isolation at the cluster level, Kardinal implements isolation by deploying development versions of services side-by-side with their "staging" versions, and creating isolated traffic routes through the cluster.

These traffic routes connect development versions to their appropriate dependencies, and to development versions of any databases, queues, caches, and external APIs that you may need. They are effectively "logical environments" or "views" on top of a single cluster. In Kardinal, these are called "[flows](https://kardinal.dev/docs/concepts/flows)". Flows enable application-level isolation for the purpose of development, testing, and QA with the lowest possible resource footprint. It all happens in one cluster, with the absolute minimum duplication of resources necessary.

There are many ways to isolate different environments in the context of cloud/Kubernetes deployments. To get an idea of how Kardinal fits into other methods, see the table below:

| Isolation method | Level of Isolation | Cost | # of Duplicated Resources |
| :--- | :--- | :--- | :--- |
| Separate VPCs | Most coarse-grained | Highest Cost | Highest |
| Separate Kubernetes Clusters | Coarse-grained | High Cost | High |
| Separate Namespaces (vclusters) | Fine-grained | Low Cost | Low |
| Separate Traffic Routes (Kardinal) | Most fine-grained | Lowest Cost | Lowest |

## How it works

In Kardinal, an application deployment with multiple logical environments (”flows”) running in it may look like the following image. 

![infographic](https://github.com/user-attachments/assets/343a44bc-2119-4368-a338-f27dc2271d8f)

Some services have multiple distinct versions (5 for `order-service`, 8 for `analytics-service`), because there are many version requirements for that service across the set of logical environments. Others only have one version deployed (`entity-service` ). In this example, every single logical environment only depends on the stable version of `entity-service`.

Because isolation is implemented at the level of traffic routing, you can maximally reuse shared resources, and you can spin up a new logical environment by only deploying the absolute minimum number of changes resources necessary.

The same multi-version deploy and de-duplication mechanism works for stateful services like Postgres or state-bearing external services like Stripe. If an isolated version of Postgres is required in a logical environment, seeded with new test data or a snapshot of data in the stable environment, it will spin up automatically for the logical environments that need it. In the case of external service like Stripe, a proxy that switches requests between dev API keys will provide the necessary isolation between logical environments.

Stateful services like databases and managed services like external APIs are supported via an open-source plugin ecosystem, so its easy to add support for the set of dependencies you have in your architecture.

Kardinal is implemented as a set of sidecars that are deployed next to your stateless services, and proxies that sit on top of those stateful services or external, managed services (like Stripe or an external API hosted outside of Kubernetes).

It's easy to install and easy to uninstall - just deploy the the sidecars in your staging cluster, and use the Kardinal control plane to manage your development and test environments. If you want to uninstall Kardinal, just remove the sidecars.

## Resources

- Explore our [docs](https://kardinal.dev/docs) to learn more about how Kardinal works.
- Ask questions and get help in our community [forum](https://discuss.kardinal.dev).
- Read our [blog](https://blog.kardinal.dev/) for tips from developers and creators.

## Quick start with a demo application

### Prerequisites

If you want to test without deploying to your machine, check out our [playground](https://github.com/kurtosis-tech/kardinal-playground/). Or if you already have your own application in mind, check out our docs regarding [running on your own application](https://kardinal.dev/docs/getting-started/install).

Before getting started make sure you have the following installed:

- [Docker](https://docs.docker.com/get-docker/)
- [Minikube](https://minikube.sigs.k8s.io/docs/start/?arch=%2Fmacos%2Fx86-64%2Fstable%2Fbinary+download)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/)
- [Istio](https://istio.io/latest/docs/setup/getting-started/)

The last prerequisite is you'll need to run Minikube with Istio enabled. To do this, run the following:

```
minikube start --driver=docker --cpus=10 --memory 8192 --disk-size 32g;
minikube addons enable ingress;
minikube addons enable metrics-server;
istioctl install --set profile=default -y;
```

### Step 1: Install Kardinal
To install Kardinal, run the following command:

```curl get.kardinal.dev -sL | sh```

### Step 2: Deploy the Kardinal Manager to your cluster

`kardinal manager deploy kloud-kontrol`

### Step 3: Deploy the demo app
Since this guide is using minikube, you'll need to set up the minikube tunnel to access the frontend of the application you're about to deploy:

`minikube tunnel`

You can leave the tunnel running. In a new terminal window, deploy the demo app via Kardinal:

```
curl https://raw.githubusercontent.com/kurtosis-tech/new-obd/main/release/obd-kardinal.yaml > ./obd-kardinal.yaml;
kardinal deploy obd-kardinal.yaml
```

You can view the frontend of the demo app by going to:

`http://prod.app.localhost`

Feel free to click around, add items to your cart, and shop!

The Kardinal dashboard will show the architecture of your application, along with any logical environments (flows) you create on top of it. To view the dashboard, run:

`kardinal dashboard`

and click on the "Traffic configuration" sidebar item.

### Step 4: Create a lightweight development environment (dev flow)

Create a new flow by specifying a service name and a container image.

Here is an example of creating a dev flow for the frontend service, using an image we've prepared for this demo:

`kardinal flow create frontend leoporoli/newobd-frontend:0.0.6`

This command will output a URL that you can use to access the frontend of the development environment. You can view the frontend of the application by going to the URL provided.

Notice that there are already items in your cart in the development environment. We've configured the development "flow" in this demo to run with it's own database which is seeded with test data. This demonstrates how dev flows can be configured with the data that the development team needs to do their testing work.

To inspect the resources in your cluster, and see how Kardinal is reusing resources in your stable environment in the dev environment, go to the dashboard again:

`kardinal dashboard`

and click on the "Traffic configuration" sidebar item.

### Step 5: Clean up your development flow
When you're done with your development flow, you can delete it by running:

`kardinal flow delete <flow_id>`

The flow_id is in the output of the kardinal flow create command, but if you've lost it, you can get it again by running:

`kardinal flow ls`

Once you've deleted the flow, you can verify that the resources have been cleaned up by going to the dashboard again.

### Ready to test on your own application?
Check out [our docs](https://kardinal.dev/docs/getting-started/install) to learn how.
