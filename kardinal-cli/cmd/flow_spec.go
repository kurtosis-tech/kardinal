package cmd

const (
	flowSpecFilePerms = 0644
)

//
//// FlowSpec represents a map of service names to their corresponding PodSpec
//type FlowSpec map[string]corev1.PodSpec
//
//// DeserializeConfig deserializes a YAML file into a ServiceConfig
//func DeserializeFlowSpec(filePath string) (*FlowSpec, error) {
//	data, err := os.ReadFile(filePath)
//	if err != nil {
//		return nil, fmt.Errorf("error reading YAML file: %w", err)
//	}
//
//	var config FlowSpec
//	err = yaml.Unmarshal(data, &config)
//	if err != nil {
//		return nil, fmt.Errorf("error unmarshaling YAML: %w", err)
//	}
//
//	return &config, nil
//}
