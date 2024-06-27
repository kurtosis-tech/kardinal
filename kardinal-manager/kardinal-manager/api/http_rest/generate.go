package http_rest

//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=./specs/types_cfg.yaml ./specs/kardinal_manager_api.yaml
//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=./specs/server_cfg.yaml ./specs/kardinal_manager_api.yaml
//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=./specs/client_cfg.yaml ./specs/kardinal_manager_api.yaml
