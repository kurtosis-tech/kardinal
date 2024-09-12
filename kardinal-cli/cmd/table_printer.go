package cmd

import (
	"fmt"
	"strings"

	api_types "github.com/kurtosis-tech/kardinal/libs/cli-kontrol-api/api/golang/types"
	"github.com/samber/lo"
)

func printTemplateTable(templates []api_types.Template) {
	if len(templates) == 0 {
		fmt.Println("No templates found.")
		return
	}

	data := lo.Map(templates, func(template api_types.Template, _ int) []string {
		description := "N/A"
		if template.Description != nil {
			description = *template.Description
		}
		return []string{
			template.TemplateId,
			template.Name,
			description,
		}
	})

	header := []string{"Template ID", "Name", "Description"}
	printGenericTable(header, data)
}

// Update the existing printFlowTable function to use the new generic function
func printFlowTable(flows []api_types.Flow) {
	data := lo.Map(flows, func(flow api_types.Flow, _ int) []string {
		var baselineStr string
		if *flow.IsBaseline {
			baselineStr = "✅"
		}
		return []string{
			flow.FlowId,
			strings.Join(lo.Map(flow.FlowUrls, func(item string, _ int) string { return fmt.Sprintf("http://%s", item) }), ", "),
			baselineStr,
		}
	})

	header := []string{"Flow ID", "Flow URL", "Baseline"}
	printGenericTable(header, data)
}

func printGenericTable(header []string, data [][]string) {
	// Prepend header to data
	allRows := append([][]string{header}, data...)

	// Calculate column widths
	colWidths := make([]int, len(header))
	for _, row := range allRows {
		for i, cell := range row {
			if len(cell) > colWidths[i] {
				colWidths[i] = len(cell)
			}
		}
	}

	// Print top border
	printTableBorder(colWidths)

	// Print rows
	for rowIndex, row := range allRows {
		for colIndex, cell := range row {
			fmt.Printf("| %-*s ", colWidths[colIndex], cell)
		}
		fmt.Println("|")

		// Print separator after header
		if rowIndex == 0 {
			printTableBorder(colWidths)
		}
	}

	// Print bottom border
	printTableBorder(colWidths)
}

func printTableBorder(colWidths []int) {
	for _, width := range colWidths {
		fmt.Print("+", strings.Repeat("-", width+2))
	}
	fmt.Println("+")
}
