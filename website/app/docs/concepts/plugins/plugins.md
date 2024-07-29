# Kardinal Plugin System

## Overview

The Plugin System allow developers to write Python-based plugins that can dynamically alter Kubernetes deployment specs, providing a way to handle stateful services and managed services in a safer manner.

## How It Works

1. **Plugin Execution**: Kardinal executes the specified plugins when creating or deleting a flow.
2. **Deployment Spec Modification**: Plugins can modify the deployment specification before it's applied to the cluster.
3. **Config Map Generation**: Plugins can generate a config map to store information for later use, particularly during flow deletion.

## Designing Plugins

Plugins are Python scripts hosted on GitHub. Each plugin should have two main functions:

1. **create_flow**: Called when creating a new flow.
2. **delete_flow**: Called when deleting a flow.

### Basic Plugin Structure

```python
# main.py

def create_flow(service_spec, deployment_spec, flow_uuid, optional_argument):
    # Modify the deployment spec
    # Generate a config map if needed
    # service_spec - the Kubernetes service spec json
    # deployment_spec - Deployment spec of the service json
    # flow_uuid - the uuid of the flow string
    # optional_argument - you can have any number of these, passed via annotations
    return {
        "deployment_spec": modified_deployment_spec,
        "config_map": config_map
    }

def delete_flow(config_map, flow_uuid):
    # Perform any necessary cleanup
    pass
```

### Plugin Guidelines

- You need a `main.py` in the root of your repository with the above structure
- Modify the `deployment_spec` as needed in the `create_flow` function.
- Use the `config_map` to store information that might be needed during flow deletion.
- If your plugin has external dependencies, include a `requirements.txt` file in the root of your repository.

### Example Plugin

Here's an example of a simple plugin that replaces text in various parts of the deployment spec:

```python
REPLACED = "the-text-has-been-replaced"

def create_flow(service_spec, deployment_spec, flow_uuid, text_to_replace):
    deployment_spec['template']['metadata']['labels']['app'] = deployment_spec['template']['metadata']['labels']['app'].replace(text_to_replace, REPLACED)
    deployment_spec['selector']['matchLabels']['app'] = deployment_spec['selector']['matchLabels']['app'].replace(text_to_replace, REPLACED)
    deployment_spec['template']['spec']['containers'][0]['name'] = deployment_spec['template']['spec']['containers'][0]['name'].replace(text_to_replace, REPLACED)
    
    config_map = {
        "original_text": text_to_replace
    }
    
    return {
        "deployment_spec": deployment_spec,
        "config_map": config_map
    }

def delete_flow(config_map, flow_uuid):
    print(config_map["original_text"])
```

### Plugin Design Best Practices

1. **Documentation**: Clearly document your plugin's purpose, required arguments, and effects on the deployment spec.
1. **Dependency Management**: If your plugin requires external libraries, list them in a `requirements.txt` file in the root of your repository.
1. **Error Handling**: Implement proper error handling; raise an error and a non zero exit code if the plugin fails
1. **Config Map Usage**: Use the config map to store any information that might be needed during the delete_flow operation.


## Associating Plugins with Kubernetes Services

In the Kardinal system, plugins are associated with specific services in your Kubernetes cluster using annotations. This allows you to specify which plugins should be applied to each service, providing fine-grained control over your deployment modifications.

### How to Tag a Service Spec

To use a plugin with a particular service, you need to add special annotations to your Kubernetes service specification. Here's how you can do it:

1. Open your Kubernetes service specification YAML file.
2. Add an annotation under the `metadata` section of your service.
3. Use the `kardinal.dev/plugins` key to specify the plugins you want to use.

Here's an example of how your service spec might look:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-awesome-service
  annotations:
    kardinal.dev/plugins: |
      - name: github.com/kurtosis-tech/redis-plugin
        args:
          text_to_replace: "original-text"
spec:
  selector:
    app: my-awesome-service
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
```

### Annotation Structure

The `kardinal.dev/plugins` annotation uses a YAML-formatted list of plugins. Each plugin in the list has two main parts:

1. `name`: This is the GitHub repository URL of the plugin.
2. `args`: These are the arguments that will be passed to the plugin's `create_flow` function.

You can specify multiple plugins for a single service by adding more items to the list:

```yaml
annotations:
  kardinal.dev/plugins: |
    - name: github.com/username/repo1
      args:
        arg1: value1
    - name: github.com/username/repo2
      args:
        arg2: value2
```

### Plugin Execution

When Kardinal processes your service, it will:

1. Read the `kardinal.dev/plugins` annotation.
2. For each plugin listed:
   - Fetch the plugin code from the specified GitHub repository.
   - Execute the plugin's `create_flow` function, passing in the service spec, deployment spec, a generated flow UUID, and any arguments specified in the `args` section.
   - Apply the modifications returned by the plugin to the deployment spec.

### Plugin Annotation Best Practices

1. **Argument Naming**: Use clear, descriptive names for your plugin arguments.
1. **Plugin Order**: If using multiple plugins, consider their order as they will be applied sequentially.

By using annotations in your Kubernetes service specs, you can easily associate Kardinal plugins with specific services. This allows for powerful, targeted modifications to your deployments, enhancing the flexibility and manageability of your Kubernetes applications.

## Existing plugins

1. [Redis Sidecar Plugin](https://github.com/kurtosis-tech/redis-sidecar-plugin) - Adds a  thin layer over redis that  allows for shared reads and isolated writes
1. [Neon DB Plugin](https://github.com/kurtosis-tech/neondb-plugin) - Creates a new branch for the database you are using
1. [Postgres Seed Plugin](https://github.com/kurtosis-tech/postgres-seed-plugin) - Allows you to spin up a postgres database with seeded data
