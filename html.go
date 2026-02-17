package gotabulate

import (
	"bytes"
	"fmt"
	"strings"
)

// Bootstrap5MinimalCSS is a minimal but functional Bootstrap 5 CSS embedded for offline use
const bootstrapCSS = `
/* Bootstrap 5 Minimal CSS - Embedded for offline use */
*{margin:0;padding:0;box-sizing:border-box}
html{font-family:sans-serif;line-height:1.15;-webkit-text-size-adjust:100%}
body{margin:0;font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif;font-size:1rem;font-weight:400;line-height:1.5;color:#212529;background-color:#fff}
table{border-collapse:collapse}
thead th{text-align:left;background-color:#e9ecef;border-bottom:2px solid #dee2e6;padding:.75rem}
tbody td{padding:.75rem;border-bottom:1px solid #dee2e6}
.table{width:100%;margin-bottom:1rem;color:#212529;border-collapse:collapse}
.table>:not(caption)>*>*{padding:.5rem .5rem;background-color:rgba(0,0,0,0);border-bottom-width:1px}
.table>tbody{border-color:inherit}
.table-striped>tbody>tr:nth-of-type(odd)>*{background-color:rgba(0,0,0,.05)}
.table-striped>tbody>tr:nth-of-type(odd){background-color:rgba(0,0,0,.025)}
.table-hover>tbody>tr:hover>*{background-color:rgba(0,0,0,.075)}
.table-bordered{border:1px solid #dee2e6}
.table-bordered>:not(caption)>*>*{border-width:1px 0}
.table-bordered>:not(caption)>tr>th{border-bottom-width:2px}
.table-sm>:not(caption)>*>*{padding:.25rem}
.table-primary{background-color:#cfe2ff;border-color:#b6d4fe}
.table-secondary{background-color:#e2e3e5;border-color:#d3d6db}
.table-success{background-color:#d1e7dd;border-color:#badbcc}
.table-danger{background-color:#f8d7da;border-color:#f5c2c7}
.table-active{background-color:rgba(0,0,0,.075)}
.table-dark{color:#fff;background-color:#212529;border-color:#373b3f}
.table-dark>:not(caption)>*>*{background-color:#212529}
.container-fluid{--bs-gutter-x:1.5rem;--bs-gutter-y:0;width:100%;padding-right:calc(var(--bs-gutter-x)*.5);padding-left:calc(var(--bs-gutter-x)*.5);margin-right:auto;margin-left:auto}
.mt-4{margin-top:1.5rem!important}
h1,h2,h3,h4,h5,h6{margin-top:0;margin-bottom:.5rem;font-weight:500;line-height:1.2}
h2{font-size:calc(1.325rem + .9vw)}
`

// HTMLConfig provides customizable options for HTML table rendering.
type HTMLConfig struct {
	// TableClass is the CSS class(es) applied to the table element
	TableClass string
	// HeaderClass is the CSS class(es) applied to header cells
	HeaderClass string
	// RowClass is the CSS class(es) applied to data row <tr> elements
	RowClass string
	// CellClass is the CSS class(es) applied to data cells
	CellClass string
	// StripedRows applies alternating row colors when true
	StripedRows bool
	// Bordered adds borders to all cells when true
	Bordered bool
	// Condensed reduces padding for a more compact table
	Condensed bool
	// Hover highlights rows on hover
	Hover bool
	// ResponsiveWidth makes the table responsive to screen width
	ResponsiveWidth bool
	// CustomCSS contains custom CSS rules to include in the table
	CustomCSS string
	// BootstrapVersion specifies Bootstrap version (4, 5, or empty for no framework)
	BootstrapVersion string
	// UseBootstrapCDN determines whether to use CDN link (true) or embed CSS (false)
	// When false, Bootstrap CSS is embedded in the HTML for offline use
	UseBootstrapCDN bool
	// Title is an optional title displayed above the table
	Title string
}

// DefaultHTMLConfig returns a sensible default HTML configuration using Bootstrap 5.
//
// By default, Bootstrap CSS is loaded from a CDN (Content Delivery Network)
// for smaller file sizes (~2KB). This assumes your users have internet access,
// which is reasonable for most modern web applications.
//
// Configuration defaults:
//   - Framework: Bootstrap 5
//   - CSS Source: CDN links (jsdelivr)
//   - File Size: ~2KB (CSS loaded externally)
//   - Striped Rows: Enabled
//   - Hover Effects: Enabled
//   - Borders: Enabled
//   - Theme: Primary header color
//
// If users encounter rendering issues due to lack of internet, you can
// enable offline mode using SetHTMLOfflineMode() or OfflineHTMLConfig()
// to embed the CSS directly.
//
// Example with default CDN mode:
//
//	table := gotabulate.Create(data)
//	table.SetHeaders([]string{"Name", "Value"})
//	html := table.Render("html")  // Small file, uses CDN
//
// Example switching to offline mode:
//
//	table := gotabulate.Create(data)
//	table.SetHeaders([]string{"Name", "Value"})
//	table.SetHTMLOfflineMode()  // If needed for offline use
//	html := table.Render("html")  // Large file, self-contained
func DefaultHTMLConfig() HTMLConfig {
	return HTMLConfig{
		TableClass:       "table table-striped table-hover",
		HeaderClass:      "table-primary",
		StripedRows:      true,
		Bordered:         true,
		Hover:            true,
		BootstrapVersion: "5",
		UseBootstrapCDN:  true, // Use CDN by default for smaller files
		Condensed:        false,
	}
}

// OfflineHTMLConfig returns an HTML configuration with fully embedded Bootstrap CSS.
//
// This configuration embeds the complete Bootstrap 5 CSS directly in the HTML file,
// making it completely self-contained and usable without internet access.
// The resulting file is larger (~100KB) but requires no external resources.
//
// Use this function or SetHTMLOfflineMode when generating HTML on systems
// that don't have internet access, such as backend servers, data processing
// systems, or air-gapped environments.
//
// The HTML produced with this configuration can be:
//   - Emailed as an attachment and opened anywhere
//   - Saved to a file and opened later
//   - Served by internal web applications
//   - Archived for long-term storage
//   - Used in offline or disconnected environments
//
// Example:
//
//	config := gotabulate.OfflineHTMLConfig()
//	config.Title = "Offline Report"
//	table.SetHTMLConfig(config)
//	html := table.Render("html")
func OfflineHTMLConfig() HTMLConfig {
	config := DefaultHTMLConfig()
	config.UseBootstrapCDN = false
	return config
}

// MinimalHTMLConfig returns a minimal HTML configuration with no framework dependencies.
// This is completely self-contained and works offline without any external resources.
// Useful if you don't want to depend on Bootstrap or CDN.
func MinimalHTMLConfig() HTMLConfig {
	return HTMLConfig{
		TableClass:      "tabulate-table",
		UseBootstrapCDN: false,
		CustomCSS: `
.tabulate-table {
    font-family: Arial, sans-serif;
    border-collapse: collapse;
    width: 100%;
}
.tabulate-table thead th {
    background-color: #f8f9fa;
    border: 1px solid #dee2e6;
    padding: 12px;
    text-align: left;
    font-weight: 600;
}
.tabulate-table tbody td {
    border: 1px solid #dee2e6;
    padding: 10px 12px;
}
.tabulate-table tbody tr:nth-child(even) {
    background-color: #f8f9fa;
}
.tabulate-table tbody tr:hover {
    background-color: #e9ecef;
}`,
	}
}

// renderHTMLTable generates an HTML table from the Tabulate data using the HTMLConfig.
func (t *Tabulate) renderHTMLTable(config HTMLConfig) string {
	var buf bytes.Buffer

	// Write DOCTYPE and opening tags
	buf.WriteString("<!DOCTYPE html>\n")
	buf.WriteString("<html lang=\"en\">\n")
	buf.WriteString("<head>\n")
	buf.WriteString("  <meta charset=\"UTF-8\">\n")
	buf.WriteString("  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n")
	buf.WriteString("  <title>Table</title>\n")

	// Add Bootstrap if specified
	if config.BootstrapVersion != "" {
		if config.UseBootstrapCDN {
			// Use CDN links
			if config.BootstrapVersion == "5" {
				buf.WriteString("  <link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css\" rel=\"stylesheet\">\n")
			} else if config.BootstrapVersion == "4" {
				buf.WriteString("  <link href=\"https://cdn.jsdelivr.net/npm/bootstrap@4.6.2/dist/css/bootstrap.min.css\" rel=\"stylesheet\">\n")
			}
		} else {
			// Embed Bootstrap CSS for offline use
			if config.BootstrapVersion == "5" || config.BootstrapVersion == "4" {
				buf.WriteString("  <style type=\"text/css\">\n")
				buf.WriteString(bootstrapCSS)
				buf.WriteString("  </style>\n")
			}
		}
	}

	// Write custom CSS if provided
	if config.CustomCSS != "" {
		buf.WriteString("  <style type=\"text/css\">\n")
		buf.WriteString(config.CustomCSS)
		buf.WriteString("  </style>\n")
	}

	buf.WriteString("</head>\n")
	buf.WriteString("<body>\n")

	// Add container if Bootstrap is enabled
	if config.BootstrapVersion != "" {
		buf.WriteString("  <div class=\"container-fluid mt-4\">\n")
	}

	// Add title if provided
	if config.Title != "" {
		if config.BootstrapVersion != "" {
			buf.WriteString(fmt.Sprintf("    <h2>%s</h2>\n", escapeHTML(config.Title)))
		} else {
			buf.WriteString(fmt.Sprintf("    <h2 style=\"margin-bottom: 20px;\">%s</h2>\n", escapeHTML(config.Title)))
		}
	}

	// Start table
	tableClass := config.TableClass
	if config.Bordered && !strings.Contains(tableClass, "bordered") {
		tableClass += " table-bordered"
	}
	if config.Condensed && !strings.Contains(tableClass, "condensed") {
		tableClass += " table-sm"
	}

	buf.WriteString(fmt.Sprintf("    <table class=\"%s\">\n", tableClass))

	// Write header if not hidden
	if !t.NoHeader {
		buf.WriteString("      <thead>\n")
		buf.WriteString("        <tr")
		if config.HeaderClass != "" {
			buf.WriteString(fmt.Sprintf(" class=\"%s\"", config.HeaderClass))
		}
		buf.WriteString(">\n")

		for _, header := range t.Headers {
			buf.WriteString(fmt.Sprintf("          <th>%s</th>\n", escapeHTML(header)))
		}

		buf.WriteString("        </tr>\n")
		buf.WriteString("      </thead>\n")
	}

	// Write data rows
	buf.WriteString("      <tbody>\n")
	for i, row := range t.Data {
		rowClass := ""
		if config.StripedRows && i%2 == 1 {
			rowClass = "table-active"
		}
		if config.RowClass != "" {
			if rowClass != "" {
				rowClass += " " + config.RowClass
			} else {
				rowClass = config.RowClass
			}
		}

		buf.WriteString("        <tr")
		if rowClass != "" {
			buf.WriteString(fmt.Sprintf(" class=\"%s\"", rowClass))
		}
		buf.WriteString(">\n")

		for _, cell := range row.Elements {
			cellClass := ""
			if config.CellClass != "" {
				cellClass = fmt.Sprintf(" class=\"%s\"", config.CellClass)
			}
			buf.WriteString(fmt.Sprintf("          <td%s>%s</td>\n", cellClass, escapeHTML(cell)))
		}

		buf.WriteString("        </tr>\n")
	}

	buf.WriteString("      </tbody>\n")
	buf.WriteString("    </table>\n")

	// Close container if Bootstrap is enabled
	if config.BootstrapVersion != "" {
		buf.WriteString("  </div>\n")
	}

	buf.WriteString("</body>\n")
	buf.WriteString("</html>\n")

	return buf.String()
}

// escapeHTML escapes special HTML characters in a string.
func escapeHTML(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "\"", "&quot;")
	s = strings.ReplaceAll(s, "'", "&#39;")
	return s
}
