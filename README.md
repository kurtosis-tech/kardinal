[![Docker Hub](https://img.shields.io/badge/dockerhub-images-important.svg?logo=docker)](https://hub.docker.com/u/kurtosistech)

# Kardinal

![Kardi B](https://kardinal.dev/_next/static/media/kardinal-orange.65ea335b.png)

The lightest-weight Kubernetes development environments in the world. Stop duplicating everything across your dev, test, and QA Kubernetes clusters. Deploy the minimum resources necessary to develop and test directly in one production-like environment.

## Can't wait to get started?

Visit the [Kardinal playground](https://github.com/kurtosis-tech/kardinal-playground/) to experience a simple demonstration of how Kardinal can enhance your development workflow.

## What is Kardinal?

Kardinal is a multi-tenancy tool for Kubernetes that enables engineers to efficiently do development, testing, and QA work within a single stable cluster. Instead of implementing isolation at the cluster level, Kardinal implements isolation by deploying development versions of services side-by-side with their "staging" versions, and creating isolated traffic routes through the cluster.

These traffic routes connect development versions to their appropriate dependencies, and to development versions of any databases, queues, caches, and external APIs that you may need. They are effectively "logical environments" or "views" on top of a single cluster. In Kardinal, these are called "flows". Flows enable application-level isolation for the purpose of development, testing, and QA with the lowest possible resource footprint. It all happens in one cluster, with the absolute minimum duplication of resources necessary.

There are many ways to isolate different environments in the context of cloud/Kubernetes deployments. To get an idea of how Kardinal fits into other methods, see the table below:

| Isolation method | Level of Isolation | Cost | # of Duplicated Resources |
| :--- | :--- | :--- | :--- |
| Separate VPCs | Most coarse-grained | Highest Cost | Highest |
| Separate Kubernetes Clusters | Coarse-grained | High Cost | High |
| Separate Namespaces (vclusters) | Fine-grained | Low Cost | Low |
| Separate Traffic Routes (Kardinal) | Most fine-grained | Lowest Cost | Lowest |

## How it Works

In Kardinal, an application deployment with multiple logical environments (”flows”) running in it may look like the following image. 

![infographic](https://github.com/user-attachments/assets/343a44bc-2119-4368-a338-f27dc2271d8f)

Some services have multiple distinct versions (5 for `order-service`, 8 for `analytics-service`), because there are many version requirements for that service across the set of logical environments. Others only have one version deployed (`entity-service` ). In this example, every single logical environment only depends on the stable version of `entity-service`.

Because isolation is implemented at the level of traffic routing, you can maximally reuse shared resources, and you can spin up a new logical environment by only deploying the absolute minimum number of changes resources necessary.

The same multi-version deploy and de-duplication mechanism works for stateful services like Postgres or state-bearing external services like Stripe. If an isolated version of Postgres is required in a logical environment, seeded with new test data or a snapshot of data in the stable environment, it will spin up automatically for the logical environments that need it. In the case of external service like Stripe, a proxy that switches requests between dev API keys will provide the necessary isolation between logical environments.

Stateful services like databases and managed services like external APIs are supported via an open-source plugin ecosystem, so its easy to add support for the set of dependencies you have in your architecture.

Kardinal is implemented as a set of sidecars that are deployed next to your stateless services, and proxies that sit on top of those stateful services or external, managed services (like Stripe or an external API hosted outside of Kubernetes).

It's easy to install and easy to uninstall - just deploy the the sidecars in your staging cluster, and use the Kardinal control plane to manage your development and test environments. If you want to uninstall Kardinal, just remove the sidecars.

## Quick start

Steps needed to get up and running with Kardinal will be published when we release a version that works on generic cluster topologies. Estimated release data is August 2nd, 2024.

## Resources

- Explore our [docs](https://kardinal.dev) to learn more about how Kardinal works
- Ask questions and get help in our community [forum](https://discuss.kardinal.dev).
- Read our [blog](https://blog.kardinal.dev/) for tips from developers and creators.
