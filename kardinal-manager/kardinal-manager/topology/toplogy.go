package topology

import (
	"github.com/sirupsen/logrus"
)

type RawKialiGraph struct {
	Elements struct {
		Nodes []struct {
			Data struct {
				ID       string `json:"id"`
				NodeType string `json:"nodeType"`
				Service  string `json:"service"`
				App      string `json:"app"`
				Version  string `json:"version"`
			} `json:"data"`
		} `json:"nodes"`
		Edges []struct {
			Data struct {
				Source string `json:"source"`
				Target string `json:"target"`
			} `json:"data"`
		} `json:"edges"`
	} `json:"elements"`
}

type Node struct {
	RawKialiGraphID string
	ID              string // serviceName_version
	ServiceName     string
	ServiceVersion  string
	TalksTo         []string // List of IDs (serviceName_version)
}

func graphToNodesMap(graph *RawKialiGraph) map[string]*Node {
	nodesMap := make(map[string]*Node)
	idMap := make(map[string]string) // Map from raw graph ID to readable ID
	nodesByUsefulID := make(map[string]*Node)

	// Populate nodes
	for _, n := range graph.Elements.Nodes {
		serviceName := n.Data.Service
		if serviceName == "" {
			serviceName = n.Data.App // Use app name if service name is not specified
		}
		serviceVersion := n.Data.Version
		if serviceVersion == "" {
			serviceVersion = "latest" // Default to 'latest' if no version is specified
		}
		readableID := serviceName + "_" + serviceVersion

		node := &Node{
			RawKialiGraphID: n.Data.ID,
			ID:              readableID,
			ServiceName:     serviceName,
			ServiceVersion:  serviceVersion,
			TalksTo:         make([]string, 0),
		}
		nodesMap[n.Data.ID] = node
		idMap[n.Data.ID] = readableID
		nodesByUsefulID[readableID] = node
	}

	// Populate connections using readable IDs
	for _, e := range graph.Elements.Edges {
		if sourceNode, ok := nodesMap[e.Data.Source]; ok {
			if targetID, ok := idMap[e.Data.Target]; ok {
				sourceNode.TalksTo = append(sourceNode.TalksTo, targetID)
			}
		}
	}

	// Print the nodes and their connections
	for _, node := range nodesMap {
		logrus.Debugf("Node ID: %s (RawKialiGraphID: %s)", node.ID, node.RawKialiGraphID)
		logrus.Debugf("  Service: %s Version: %s", node.ServiceName, node.ServiceVersion)
		logrus.Debugf("  Talks To: %v", node.TalksTo)
	}

	return nodesByUsefulID
}
