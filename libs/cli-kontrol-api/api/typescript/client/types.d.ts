/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */


export interface paths {
  "/health": {
    get: {
      responses: {
        /** @description Successful response */
        200: {
          content: {
            "application/json": string;
          };
        };
        500: components["responses"]["Error"];
      };
    };
  };
  "/tenant/{uuid}/flow/create": {
    post: {
      parameters: {
        path: {
          uuid: components["parameters"]["uuid"];
        };
      };
      /** @description Create a dev flow using FlowSpec, optionally with a TemplateSpec */
      requestBody: {
        content: {
          "application/json": {
            flowSpec: components["schemas"]["FlowSpec"];
            templateSpec?: components["schemas"]["TemplateSpec"];
          };
        };
      };
      responses: {
        /** @description Flow creation status */
        200: {
          content: {
            "application/json": components["schemas"]["Flow"];
          };
        };
        404: components["responses"]["NotFound"];
        500: components["responses"]["Error"];
      };
    };
  };
  "/tenant/{uuid}/flows": {
    get: {
      parameters: {
        path: {
          uuid: components["parameters"]["uuid"];
        };
      };
      responses: {
        /** @description Dev flow creation status */
        200: {
          content: {
            "application/json": components["schemas"]["Flow"][];
          };
        };
        404: components["responses"]["NotFound"];
        500: components["responses"]["Error"];
      };
    };
  };
  "/tenant/{uuid}/flow/{flowId}": {
    delete: {
      parameters: {
        path: {
          uuid: components["parameters"]["uuid"];
          flowId: components["parameters"]["flowId"];
        };
      };
      responses: {
        404: components["responses"]["NotFound"];
        500: components["responses"]["Error"];
        /** @description Dev flow deletion status */
        "2xx": {
          content: never;
        };
      };
    };
  };
  "/tenant/{uuid}/deploy": {
    post: {
      parameters: {
        path: {
          uuid: components["parameters"]["uuid"];
        };
      };
      /** @description Deploy a prod only cluster */
      requestBody: {
        content: {
          "application/json": components["schemas"]["MainClusterConfig"];
        };
      };
      responses: {
        /** @description Dev flow creation status */
        200: {
          content: {
            "application/json": components["schemas"]["Flow"];
          };
        };
        404: components["responses"]["NotFound"];
        500: components["responses"]["Error"];
      };
    };
  };
  "/tenant/{uuid}/topology": {
    get: {
      parameters: {
        path: {
          uuid: components["parameters"]["uuid"];
        };
      };
      responses: {
        /** @description Topology information */
        200: {
          content: {
            "application/json": components["schemas"]["ClusterTopology"];
          };
        };
        404: components["responses"]["NotFound"];
        500: components["responses"]["Error"];
      };
    };
  };
  "/tenant/{uuid}/templates/create": {
    post: {
      parameters: {
        path: {
          uuid: components["parameters"]["uuid"];
        };
      };
      /** @description Create a new template */
      requestBody: {
        content: {
          "application/json": components["schemas"]["TemplateConfig"];
        };
      };
      responses: {
        /** @description Template creation status */
        200: {
          content: {
            "application/json": components["schemas"]["Template"];
          };
        };
        404: components["responses"]["NotFound"];
        500: components["responses"]["Error"];
      };
    };
  };
  "/tenant/{uuid}/templates": {
    get: {
      parameters: {
        path: {
          uuid: components["parameters"]["uuid"];
        };
      };
      responses: {
        /** @description List of templates for the given tenant */
        200: {
          content: {
            "application/json": components["schemas"]["Template"][];
          };
        };
        404: components["responses"]["NotFound"];
        500: components["responses"]["Error"];
      };
    };
  };
  "/tenant/{uuid}/templates/{templateName}": {
    delete: {
      parameters: {
        path: {
          uuid: components["parameters"]["uuid"];
          templateName: components["parameters"]["templateName"];
        };
      };
      responses: {
        404: components["responses"]["NotFound"];
        500: components["responses"]["Error"];
        /** @description Template deletion status */
        "2xx": {
          content: never;
        };
      };
    };
  };
  "/tenant/{uuid}/manifest": {
    /**
     * Cluster resource definition in a manifest YAML response
     * @description This endpoint returns all the Kubernetes resource for the cluster topology in one multi-resource manifest response
     */
    get: {
      parameters: {
        path: {
          uuid: components["parameters"]["uuid"];
        };
      };
      responses: {
        /** @description Successful response */
        200: {
          content: {
            "application/x-yaml": string;
          };
        };
        404: components["responses"]["NotFound"];
        500: components["responses"]["Error"];
      };
    };
  };
}

export type webhooks = Record<string, never>;

export interface components {
  schemas: {
    MainClusterConfig: {
      serviceConfigs?: components["schemas"]["ServiceConfig"][];
      ingressConfigs?: components["schemas"]["IngressConfig"][];
      namespace?: string;
    };
    Flow: {
      flowId: string;
      flowUrls: string[];
      isBaseline?: boolean;
    };
    FlowSpec: {
        /** @example backend-a:latest */
        imageLocator: string;
        /** @example backend-service-a */
        serviceName: string;
      }[];
    TemplateSpec: {
      /** @description name of the template */
      templateName: string;
      arguments?: {
        [key: string]: unknown;
      };
    };
    Node: {
      /** @description Unique identifier for the node. */
      id: string;
      /** @description Label for the node. */
      label: string;
      /**
       * @description Type of the node
       * @enum {string}
       */
      type: "gateway" | "service" | "external";
      /** @description Node versions */
      versions?: components["schemas"]["NodeVersion"][];
    };
    NodeVersion: {
      flowId: string;
      imageTag?: string;
      isBaseline: boolean;
    };
    Edge: {
      /** @description The identifier of the source node of the edge. */
      source: string;
      /** @description The identifier of the target node of the edge. */
      target: string;
      /** @description Label for the edge. */
      label?: string;
    };
    ClusterTopology: {
      nodes: components["schemas"]["Node"][];
      edges: components["schemas"]["Edge"][];
    };
    ServiceConfig: {
      service: unknown;
      deployment: unknown;
    };
    TemplateConfig: {
      service: unknown[];
      /** @description The name to give the template */
      name: string;
      /** @description The description of the template */
      description?: string;
    };
    Template: {
      templateId: string;
      name: string;
      description?: string;
    };
    IngressConfig: {
      ingress: unknown;
    };
  };
  responses: {
    /** @description Error */
    Error: {
      content: {
        "application/json": {
          /** @description Error type */
          error: string;
          /** @description Error message */
          msg?: string;
        };
      };
    };
    /** @description Resource not found */
    NotFound: {
      content: {
        "application/json": {
          /** @description Resource type */
          resourceType: string;
          /** @description Resource ID */
          id: string;
        };
      };
    };
  };
  parameters: {
    /** @description UUID of the resource */
    uuid: string;
    /** @description Flow identifier */
    flowId: string;
    /** @description name of the template */
    templateName: string;
  };
  requestBodies: never;
  headers: never;
  pathItems: never;
}

export type $defs = Record<string, never>;

export type external = Record<string, never>;

export type operations = Record<string, never>;
