package main

import (
	"fmt"
	"os"

	"github.com/rmasci/gotabulate"
)

func main() {
	// Sample data
	data := [][]string{
		{"John Doe", "30", "Engineer", "$120,000"},
		{"Jane Smith", "28", "Designer", "$110,000"},
		{"Bob Johnson", "35", "Manager", "$130,000"},
		{"Alice Brown", "26", "Developer", "$105,000"},
	}

	// Example 1: Default Bootstrap 5 HTML
	fmt.Println("Example 1: Default Bootstrap 5 HTML")
	table := gotabulate.Create(data)
	table.SetHeaders([]string{"Name", "Age", "Position", "Salary"})
	table.SetHTMLTitle("Employee Directory")
	html := table.Render("html")
	writeFile("example_bootstrap5.html", html)
	fmt.Println("✓ Created example_bootstrap5.html")

	// Example 2: Bootstrap 4 HTML
	fmt.Println("\nExample 2: Bootstrap 4 HTML")
	table2 := gotabulate.Create(data)
	table2.SetHeaders([]string{"Name", "Age", "Position", "Salary"})
	table2.SetHTMLTitle("Employee Directory")
	table2.SetHTMLBootstrap("4")
	html2 := table2.Render("html")
	writeFile("example_bootstrap4.html", html2)
	fmt.Println("✓ Created example_bootstrap4.html")

	// Example 3: Minimal HTML (no Bootstrap)
	fmt.Println("\nExample 3: Minimal HTML (no Bootstrap)")
	table3 := gotabulate.Create(data)
	table3.SetHeaders([]string{"Name", "Age", "Position", "Salary"})
	table3.SetHTMLTitle("Employee Directory")
	table3.SetHTMLMinimal()
	html3 := table3.Render("html")
	writeFile("example_minimal.html", html3)
	fmt.Println("✓ Created example_minimal.html")

	// Example 4: Custom styled HTML
	fmt.Println("\nExample 4: Custom styled HTML")
	table4 := gotabulate.Create(data)
	table4.SetHeaders([]string{"Name", "Age", "Position", "Salary"})
	table4.SetHTMLTitle("Employee Directory")

	customConfig := gotabulate.DefaultHTMLConfig()
	customConfig.BootstrapVersion = "5"
	customConfig.TableClass = "table table-dark table-striped table-hover"
	customConfig.Title = "Employee Directory - Dark Mode"
	customConfig.CustomCSS = `
		.tabulate-dark { background-color: #1a1a1a; color: #fff; }
		.table-dark th { background-color: #2d3748; }
	`

	table4.SetHTMLConfig(customConfig)
	html4 := table4.Render("html")
	writeFile("example_custom.html", html4)
	fmt.Println("✓ Created example_custom.html")

	fmt.Println("\n✅ All examples created successfully!")
}

func writeFile(filename, content string) {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error writing %s: %v\n", filename, err)
	}
}
