package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"kardinal.cli/tenant"
	"log"
	"net/http"

	"github.com/compose-spec/compose-go/cli"
	"github.com/compose-spec/compose-go/types"
	"github.com/spf13/cobra"

	api "github.com/kurtosis-tech/kardinal/libs/cli-kontrol-api/api/golang/client"
	api_types "github.com/kurtosis-tech/kardinal/libs/cli-kontrol-api/api/golang/types"
)

const (
	projectName          = "kardinal"
	devMode              = true
	kontrolServiceApiUrl = "ad718d90d54d54dd084dea50a9f011af-1140086995.us-east-1.elb.amazonaws.com"
	kontrolServicePort   = 8080
)

var composeFile string

var rootCmd = &cobra.Command{
	Use:   "kardinal",
	Short: "Kardinal CLI to manage deployment flows",
}

var flowCmd = &cobra.Command{
	Use:   "flow",
	Short: "Manage deployment flows",
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy services",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		services, err := parseComposeFile(composeFile)
		if err != nil {
			log.Fatalf("Error loading compose file: %v", err)
		}
		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}
		deploy(tenantUuid, services)
	},
}

var createCmd = &cobra.Command{
	Use:   "create [service name] [image name]",
	Short: "Create a new service in development mode",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		serviceName, imageName := args[0], args[1]
		services, err := parseComposeFile(composeFile)
		if err != nil {
			log.Fatalf("Error loading compose file: %v", err)
		}

		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}

		fmt.Printf("Creating service %s with image %s in development mode...\n", serviceName, imageName)
		createDevFlow(tenantUuid, services, imageName, serviceName)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete services",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		services, err := parseComposeFile(composeFile)
		if err != nil {
			log.Fatalf("Error loading compose file: %v", err)
		}

		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}
		deleteFlow(tenantUuid, services)

		fmt.Print("Deleting dev flow")
	},
}

func init() {
	rootCmd.AddCommand(flowCmd)
	rootCmd.AddCommand(deployCmd)
	flowCmd.AddCommand(createCmd, deleteCmd)

	flowCmd.PersistentFlags().StringVarP(&composeFile, "docker-compose", "d", "", "Path to the Docker Compose file")
	flowCmd.MarkPersistentFlagRequired("docker-compose")
	deployCmd.PersistentFlags().StringVarP(&composeFile, "docker-compose", "d", "", "Path to the Docker Compose file")
	deployCmd.MarkPersistentFlagRequired("docker-compose")
}

func Execute() error {
	return rootCmd.Execute()
}

func loadComposeFile(filename string) (*types.Project, error) {
	opts, err := cli.NewProjectOptions([]string{filename},
		cli.WithOsEnv,
		cli.WithDotEnv,
		cli.WithName(projectName),
	)
	if err != nil {
		return nil, err
	}

	project, err := cli.ProjectFromOptions(opts)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func parseComposeFile(composeFile string) ([]types.ServiceConfig, error) {
	project, err := loadComposeFile(composeFile)
	if err != nil {
		log.Fatalf("Error loading compose file: %v", err)
		return nil, err
	}

	fmt.Println("Services in the Docker Compose file:")
	for _, service := range project.Services {
		fmt.Println(service.Name)
	}

	projectYAML, err := project.MarshalJSON()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var dockerCompose map[string]interface{}
	err = json.Unmarshal(projectYAML, &dockerCompose)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}

	return project.Services, nil
}

func createDevFlow(tenantUuid uuid.UUID, services []types.ServiceConfig, imageLocator, serviceName string) {
	ctx := context.Background()

	params := &api_types.PostFlowCreateParams{
		Tenant: tenantUuid.String(),
	}

	// fmt.Printf("Services:\n%v", services)
	// fmt.Printf("%v", serviceName)
	// fmt.Printf("%v", imageLocator)
	body := api_types.PostFlowCreateJSONRequestBody{
		DockerCompose: &services,
		ServiceName:   &serviceName,
		ImageLocator:  &imageLocator,
	}
	client := getKontrolServiceClient()

	resp, err := client.PostFlowCreateWithResponse(ctx, params, body)
	if err != nil {
		log.Fatalf("Failed to create dev flow: %v", err)
	}

	fmt.Printf("Response: %s\n", string(resp.Body))
	fmt.Printf("Response: %s\n", resp)
}

func deploy(tenantUuid uuid.UUID, services []types.ServiceConfig) {
	ctx := context.Background()

	params := &api_types.PostDeployParams{
		Tenant: tenantUuid.String(),
	}

	body := api_types.PostDeployJSONRequestBody{
		DockerCompose: &services,
	}
	client := getKontrolServiceClient()

	resp, err := client.PostDeployWithResponse(ctx, params, body)
	if err != nil {
		log.Fatalf("Failed to deploy: %v", err)
	}

	fmt.Printf("Response: %s\n", string(resp.Body))
}

func deleteFlow(tenantUuid uuid.UUID, services []types.ServiceConfig) {
	ctx := context.Background()

	params := &api_types.PostFlowDeleteParams{
		Tenant: tenantUuid.String(),
	}

	body := api_types.PostFlowDeleteJSONRequestBody{
		DockerCompose: &services,
	}
	client := getKontrolServiceClient()

	resp, err := client.PostFlowDeleteWithResponse(ctx, params, body)
	if err != nil {
		log.Fatalf("Failed to delete flow: %v", err)
	}

	fmt.Printf("Response: %s\n", string(resp.Body))
}

func getKontrolServiceClient() *api.ClientWithResponses {
	if devMode {
		client, err := api.NewClientWithResponses("http://localhost:8080", api.WithHTTPClient(http.DefaultClient))
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}
		return client
	} else {
		client, err := api.NewClientWithResponses(fmt.Sprintf("http://%s:%v", kontrolServiceApiUrl, kontrolServicePort))
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}
		return client
	}
}
