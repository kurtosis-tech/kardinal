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
      /** @description Create a dev flow */
      requestBody: {
        content: {
          "application/json": components["schemas"]["FlowSpec"];
        };
      };
      responses: {
        /** @description Template creation status */
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
  "/tenant/{uuid}/flow/{flow-id}": {
    delete: {
      parameters: {
        path: {
          uuid: components["parameters"]["uuid"];
          "flow-id": components["parameters"]["flow-id"];
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
        /** @description Dev flow creation status */
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
}

export type webhooks = Record<string, never>;

export interface components {
  schemas: {
    MainClusterConfig: {
      "service-configs"?: components["schemas"]["ServiceConfig"][];
    };
    Flow: {
      "flow-id": string;
      "flow-urls": string[];
    };
    FlowSpec: {
        /** @example backend-a:latest */
        "image-locator": string;
        /** @example backend-service-a */
        "service-name": string;
      }[];
    TemplateSpec: {
      /** @description name of the template */
      "template-id": string;
      arguments?: {
        [key: string]: unknown;
      };
    };
    Node: {
      /** @description Unique identifier for the node. */
      id: string;
      /** @description Label for the node. */
      label?: string;
      /**
       * @description Type of the node
       * @enum {string}
       */
      type: "gateway" | "service" | "service-version" | "redis";
      /** @description Parent node */
      parent?: string;
      /** @description Node versions */
      versions?: string[];
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
      name: unknown;
      /** @description The description of the template */
      description?: unknown;
    };
    Template: {
      "template-id": string;
      name: string;
      description?: string;
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
          "resource-type": string;
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
    "flow-id": string;
  };
  requestBodies: never;
  headers: never;
  pathItems: never;
}

export type $defs = Record<string, never>;

export type external = Record<string, never>;

export type operations = Record<string, never>;
