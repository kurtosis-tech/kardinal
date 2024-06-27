#!/usr/bin/env bash

set -euo pipefail # Bash "strict mode"
script_dirpath="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
openapi_root_dirpath="$(dirname "${script_dirpath}")"
api_root_dirpath="$(dirname "${openapi_root_dirpath}")"

echo "Generating data models for REST API "
oapi-codegen --config="$openapi_root_dirpath/generators/api_types.cfg.yaml" "$openapi_root_dirpath/specs/api.yaml"

echo "Generating server code for REST API "
oapi-codegen --config="$openapi_root_dirpath/generators/go_server.cfg.yaml" "$openapi_root_dirpath/specs/api.yaml"

echo "Generating Go client code for REST API "
oapi-codegen --config="$openapi_root_dirpath/generators/go_client.cfg.yaml" "$openapi_root_dirpath/specs/api.yaml"

echo "Generating Typescript client code for REST API "
openapi-typescript "$openapi_root_dirpath/specs/api.yaml" -o "$openapi_root_dirpath/api/typescript/client/types.d.ts"
