[![Docker Hub](https://img.shields.io/badge/dockerhub-images-important.svg?logo=docker)](https://hub.docker.com/u/kurtosistech)

# Kardinal

![Kardi B](https://kardinal.dev/_next/static/media/kardinal-orange.65ea335b.png)

## Guide
1. [What is Kardinal?](https://github.com/kurtosis-tech/kardinal/?tab=readme-ov-file#what-is-kardinal)
2. [Playground](https://github.com/kurtosis-tech/kardinal/?tab=readme-ov-file#try-it-out-in-a-playground)
3. [Quick start](https://github.com/kurtosis-tech/kardinal/?tab=readme-ov-file#quick-start-with-a-demo-application)
4. [Helpful links](https://github.com/kurtosis-tech/kardinal/?tab=readme-ov-file#helpful-links)

## What is Kardinal?

Kardinal is a framework for creating extremely lightweight ephemeral development environments within a shared Kubernetes cluster.

In Kardinal, an environment is called a "flow" because it represents a path that a request takes through the cluster. Versions of services that are under development are deployed on-demand, and then shared across all development work that depends on that version. When you create a flow to test a feature, Kardinal deploys only the set of services that are changing for that feature. Then, any requests related to testing that feature are routed to those versions.

As you onboard deeper into Kardinal, you'll be able to set up isolated state for flows when desired (i.e. for testing database migrations or write-intensive workloads on shared state). Even with isolated state per flow, Kardinal will still deploy the absolute minimum resources necessary to test the changes. Isolation is done at the level of the request route, not by duplicating services in your cluster unnecessarily.

To see how Kardinal compares to other tools, and to get an idea of how would fit into your workflow, check out our [comparisons to alternatives](https://kardinal.dev/docs/references/comparisons).

Read more about Kardinal in our [docs](https://kardinal.dev/docs).

## Try it out in a Playground

We have a playground that runs in Github Codespaces so you can try Kardinal right now without installing anything. Click below to open a Codespace with the playground. The default settings for the Codespace will work just fine.

[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://github.com/codespaces/new?hide_repo_select=true&ref=main&repo=818205437&skip_quickstart=true&machine=standardLinux32gb&devcontainer_path=.devcontainer%2Fdevcontainer.json)


## Quick start with a demo application

If you want to get started with your own application, check out [our docs](https://kardinal.dev/docs/getting-started/install).

Otherwise, continue in this section to run Kardinal with a demo application to see how it works before trying it on your own.

### Prerequisites

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

## Helpful links

- Explore our [docs](https://kardinal.dev/docs) to learn more about how Kardinal works.
- Ask questions and get help in our community [forum](https://discuss.kardinal.dev).
- Read our [blog](https://blog.kardinal.dev/) for tips from developers and creators.
