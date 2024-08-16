[![Docker Hub](https://img.shields.io/badge/dockerhub-images-important.svg?logo=docker)](https://hub.docker.com/u/kurtosistech)

# Kardinal

![Kardi B](https://kardinal.dev/_next/static/media/kardinal-orange.65ea335b.png)

## What is Kardinal?

Kardinal is a framework for creating extremely lightweight ephemeral development environments within a shared Kubernetes cluster. Read more about Kardinal in our [docs](https://kardinal.dev/docs).

### Why choose Kardinal?
- **Ephemeral Environments**: Spin up a new environment exactly when you need it, and just as quickly spin it down when you’re done.
- **Minimal Resource Usage**: Only deploy the services you’re actively working on. Kardinal takes care of the rest, so you don’t waste resources.
- **Flexible Flow Types**: Whether you need to test a single service or an entire application, Kardinal has you covered:
    - Single-Service Flows: Perfect for when you’re tweaking just one service.
    - Multi-Service Flows: Ideal for when your feature involves multiple services.
    - State-Isolated Flows: Great for features that need their own databases or caches.
    - Full Application Flows: For those times when you need end-to-end testing with full isolation.
- **Cost Savings**: Kardinal can help you save big by avoiding unnecessary resource duplication. It’s a game-changer for teams looking to cut costs.

## Installation

### **Step 1: Install Kardinal**
To install Kardinal, run the following command:
```
curl get.kardinal.dev -sL | sh
```
### **Step 2: Set up a development Kubernetes cluster**

All you need is a Kubernetes cluster with Istio enabled, and kubectl installed on your machine, pointing to your cluster. If you need help with this, read more [here](https://kardinal.dev/docs/getting-started/install)

### **Step 3: Deploy the Kardinal Manager to your cluster**
Make sure that kubectl is pointing to your cluster, and then run the following command:
```
kardinal manager deploy kloud-kontrol
```

## Try it out in a Playground

We have a playground that runs in Github Codespaces so you can try Kardinal right now without installing anything. Click below to open a Codespace with the playground. The default settings for the Codespace will work just fine.

[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://github.com/codespaces/new?hide_repo_select=true&ref=main&repo=818205437&skip_quickstart=true&machine=standardLinux32gb&devcontainer_path=.devcontainer%2Fdevcontainer.json)

## Quick start with a demo application

### Step 1: Deploy the demo app

Once you have Kardinal installed, you can run through the following demo. For step 1, since this guide is using minikube, you'll need to set up the minikube tunnel to access the frontend of the application you're about to deploy:

```bash
minikube tunnel
```

You can leave the tunnel running. In a new terminal window, deploy the demo app via Kardinal:

```bash
curl https://raw.githubusercontent.com/kurtosis-tech/new-obd/main/release/obd-kardinal.yaml > ./obd-kardinal.yaml
kardinal deploy -k ./obd-kardinal.yaml
```

You can view the frontend of the demo app by going to:

`http://prod.app.localhost`

Feel free to click around, add items to your cart, and shop!

The Kardinal dashboard will show the architecture of your application, along with any logical environments (flows) you create on top of it. To view the dashboard, run:

```bash
kardinal dashboard
```

and click on the "Traffic configuration" sidebar item.

### Step 2: Create a lightweight development environment (dev flow)

Create a new flow by specifying a service name and a container image.

Here is an example of creating a dev flow for the frontend service, using an image we've prepared for this demo:

```bash
kardinal flow create frontend kurtosistech/frontend:demo-frontend
```

This command will output a URL that you can use to access the frontend of the development environment. You can view the frontend of the application by going to the URL provided.

Notice that there are already items in your cart in the development environment. We've configured the development "flow" in this demo to run with it's own database which is seeded with test data. This demonstrates how dev flows can be configured with the data that the development team needs to do their testing work.

To inspect the resources in your cluster, and see how Kardinal is reusing resources in your stable environment in the dev environment, go to the dashboard again:

```bash
kardinal dashboard
```

and click on the "Traffic configuration" sidebar item.

### Step 3: Clean up your development flow

When you're done with your development flow, you can delete it by running:

```bash
kardinal flow delete <flow_id>
```

The flow_id is in the output of the kardinal flow create command, but if you've lost it, you can get it again by running:

```bash
kardinal flow ls
```

Once you've deleted the flow, you can verify that the resources have been cleaned up by going to the dashboard again.

## Helpful links

- Explore our [docs](https://kardinal.dev/docs) to learn more about how Kardinal works.
- Ask questions and get help in our community [forum](https://discuss.kardinal.dev).
- Read our [blog](https://blog.kardinal.dev/) for tips from developers and creators.
