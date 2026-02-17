// Package gotabulate provides table formatting and rendering capabilities.
// It supports multiple output formats including text grids, HTML, CSV, and more.
package gotabulate

import (
	"bytes"
	"fmt"
	"math"

	"github.com/mattn/go-runewidth"
)

// TableFormat defines the structure and styling of table output.
type TableFormat struct {
	LineTop         Line
	LineBelowHeader Line
	LineBetweenRows Line
	LineBottom      Line
	HeaderRow       Row
	DataRow         Row
	Padding         int
	HeaderHide      bool
	FitScreen       bool
}

// Line represents a horizontal line element in a table.
type Line struct {
	begin string
	hline string
	sep   string
	end   string
}

// Row represents a row element in a table.
type Row struct {
	begin string
	sep   string
	end   string
}

const htmlTemplate = `<!DOCTYPE html>
<html>
<div class="tabulate">
<style type="text/css">
#tabulate {
    font-family: "Trebuchet MS", Arial, Helvetica, sans-serif;
    width: 100%;
    border-collapse: collapse;
}
#tabulate td, #tabulate th {
    font-size: 1em;
    border: 1px solid #000000;
    padding: 3px 7px 2px 7px;
}
#tabulate th {
    text-align: left;
    padding-top: 5px;
    padding-bottom: 4px;
    background-color: #C0C0C0;
    color: #000000;
}
#tabulate tr.alt td {
    color: #000000;
    background-color: #EAF2D3;
}
</style>
</div>`

// ShowFormats prints all available table formats to standard output.
func ShowFormats() {
	availableFormats := []string{
		"Basic Text Formats:",
		"  simple      - Minimal ASCII table with dashes",
		"  plain       - No borders, space-separated columns",
		"  tab         - Tab-separated values (TSV)",
		"  text        - Space-separated text with minimal padding",
		"  csv         - Comma-separated values",
		"",
		"MySQL-style Formats:",
		"  mysql       - ASCII art grid styled like MySQL output",
		"  mysqlg      - Unicode grid with MySQL-like styling",
		"",
		"Unicode Box Drawing:",
		"  grid        - Full Unicode box drawing with row lines",
		"  gridt       - Text-based grid with ASCII characters",
		"  bingo       - Special format with box drawing",
		"",
		"Documentation Formats:",
		"  markdown    - Markdown table format",
		"  rst         - reStructuredText table format",
		"",
		"No-Header Variants (append -nohead):",
		"  simple-nohead, plain-nohead, tab-nohead, csv-nohead,",
		"  text-nohead, mysql-nohead, mysqlg-nohead, grid-nohead,",
		"  gridt-nohead, markdown-nohead, rst-nohead",
	}
	fmt.Println("Available Output Formats:")
	for _, format := range availableFormats {
		fmt.Println(format)
	}
}

// TableFormats contains all available table format configurations.
// Users can define custom formats by adding entries to this map and calling Render.
//
// Formats:
//   - simple: Minimal ASCII table with dashes
//   - plain: No borders, space-separated columns
//   - tab: Tab-separated values (TSV)
//   - text: Space-separated text (minimal padding)
//   - csv: Comma-separated values
//   - mysql: ASCII art grid styled like MySQL output
//   - mysqlg: Unicode grid styled like MySQL with box drawing characters
//   - grid: Full Unicode box drawing with lines between rows
//   - gridt: Text-based grid with ASCII characters
//   - markdown: Markdown table format
//   - rst: reStructuredText table format
//   - json: JSON array format
//   - plus suffix "-nohead" variants that suppress headers
var TableFormats = map[string]TableFormat{
	// ==================== Basic Text Formats ====================
	"simple": {
		LineTop:         Line{"", "-", "  ", ""},
		LineBelowHeader: Line{"", "-", "  ", ""},
		LineBottom:      Line{"", "-", "  ", ""},
		HeaderRow:       Row{"", "  ", ""},
		DataRow:         Row{"", "  ", ""},
		Padding:         1,
	},
	"plain": {
		HeaderRow: Row{"", "  ", ""},
		DataRow:   Row{"", "  ", ""},
		Padding:   1,
	},
	"tab": {
		HeaderRow: Row{"", "\t", ""},
		DataRow:   Row{"", "\t", ""},
		Padding:   0,
	},
	"text": {
		HeaderRow: Row{"", " ", ""},
		DataRow:   Row{"", " ", ""},
		Padding:   0,
	},
	"csv": {
		HeaderRow: Row{"", ",", ""},
		DataRow:   Row{"", ",", ""},
		Padding:   0,
	},

	// ==================== MySQL-style Formats ====================
	"mysql": {
		LineTop:         Line{"+", "-", "+", "+"},
		LineBelowHeader: Line{"+", "=", "+", "+"},
		LineBottom:      Line{"+", "-", "+", "+"},
		HeaderRow:       Row{"|", "|", "|"},
		DataRow:         Row{"|", "|", "|"},
		Padding:         1,
	},
	"mysqlg": {
		LineTop:         Line{"╒", "═", "╤", "╕"},
		LineBelowHeader: Line{"╞", "═", "╪", "╡"},
		LineBottom:      Line{"└", "─", "┴", "┘"},
		HeaderRow:       Row{"│", "│", "│"},
		DataRow:         Row{"│", "│", "│"},
		Padding:         1,
	},

	// ==================== Unicode Box Drawing ====================
	"grid": {
		LineTop:         Line{"╒", "═", "╤", "╕"},
		LineBelowHeader: Line{"╞", "═", "╪", "╡"},
		LineBetweenRows: Line{"│", "─", "┼", "│"},
		LineBottom:      Line{"└", "─", "┴", "┘"},
		HeaderRow:       Row{"│", "│", "│"},
		DataRow:         Row{"│", "│", "│"},
		Padding:         1,
	},
	"gridt": {
		LineTop:         Line{"+", "-", "+", "+"},
		LineBelowHeader: Line{"+", "=", "+", "+"},
		LineBetweenRows: Line{"+", "-", "+", "+"},
		LineBottom:      Line{"+", "-", "+", "+"},
		HeaderRow:       Row{"|", "|", "|"},
		DataRow:         Row{"|", "|", "|"},
		Padding:         1,
	},
	"bingo": {
		LineTop:         Line{"╒", "═", "╤", "╕"},
		LineBelowHeader: Line{"╞", "═", "╪", "╡"},
		LineBetweenRows: Line{"│", "─", "┼", "│"},
		LineBottom:      Line{"└", "─", "┴", "┘"},
		HeaderRow:       Row{"│", "│", "│"},
		DataRow:         Row{"│", "│", "│"},
		Padding:         1,
	},

	// ==================== Documentation Formats ====================
	"markdown": {
		LineBelowHeader: Line{"|", "-", "|", "|"},
		HeaderRow:       Row{"| ", " | ", " |"},
		DataRow:         Row{"| ", " | ", " |"},
		Padding:         0,
	},
	"rst": {
		LineTop:         Line{"", "=", " ", ""},
		LineBelowHeader: Line{"", "=", " ", ""},
		LineBottom:      Line{"", "=", " ", ""},
		HeaderRow:       Row{"", "  ", ""},
		DataRow:         Row{"", "  ", ""},
		Padding:         1,
	},

	// ==================== No Header Variants ====================
	"simple-nohead": {
		LineTop:    Line{"", "-", "  ", ""},
		LineBottom: Line{"", "-", "  ", ""},
		HeaderRow:  Row{"", "  ", ""},
		DataRow:    Row{"", "  ", ""},
		Padding:    1,
	},
	"plain-nohead": {
		HeaderRow: Row{"", "  ", ""},
		DataRow:   Row{"", "  ", ""},
		Padding:   0,
	},
	"tab-nohead": {
		HeaderRow: Row{"", "\t", ""},
		DataRow:   Row{"", "\t", ""},
		Padding:   0,
	},
	"csv-nohead": {
		HeaderRow: Row{"", ",", ""},
		DataRow:   Row{"", ",", ""},
		Padding:   0,
	},
	"text-nohead": {
		HeaderRow: Row{"", " ", ""},
		DataRow:   Row{"", " ", ""},
		Padding:   0,
	},
	"mysql-nohead": {
		LineTop:    Line{"+", "-", "+", "+"},
		LineBottom: Line{"+", "-", "+", "+"},
		HeaderRow:  Row{"|", "|", "|"},
		DataRow:    Row{"|", "|", "|"},
		Padding:    1,
	},
	"mysqlg-nohead": {
		LineTop:    Line{"╒", "─", "╤", "╕"},
		LineBottom: Line{"└", "─", "┴", "┘"},
		HeaderRow:  Row{"│", "│", "│"},
		DataRow:    Row{"│", "│", "│"},
		Padding:    1,
	},
	"grid-nohead": {
		LineTop:         Line{"╒", "─", "╤", "╕"},
		LineBetweenRows: Line{"│", "─", "┼", "│"},
		LineBottom:      Line{"└", "─", "┴", "┘"},
		HeaderRow:       Row{"│", "│", "│"},
		DataRow:         Row{"│", "│", "│"},
		Padding:         1,
	},
	"gridt-nohead": {
		LineTop:         Line{"+", "-", "+", "+"},
		LineBetweenRows: Line{"+", "-", "+", "+"},
		LineBottom:      Line{"+", "-", "+", "+"},
		HeaderRow:       Row{"|", "|", "|"},
		DataRow:         Row{"|", "|", "|"},
		Padding:         1,
	},
	"markdown-nohead": {
		HeaderRow: Row{"| ", " | ", " |"},
		DataRow:   Row{"| ", " | ", " |"},
		Padding:   0,
	},
	"rst-nohead": {
		LineTop:    Line{"", "=", " ", ""},
		LineBottom: Line{"", "=", " ", ""},
		HeaderRow:  Row{"", "  ", ""},
		DataRow:    Row{"", "  ", ""},
		Padding:    1,
	},
}

// minPadding is the minimum padding applied to table cells.
const minPadding = 5

// Tabulate represents a data table with formatting options.
type Tabulate struct {
	Data          []*TabulateRow
	Headers       []string
	FloatFormat   byte
	TableFormat   TableFormat
	Align         string
	EmptyVar      string
	HideLines     []string
	MaxSize       int
	WrapStrings   bool
	RemEmptyLines bool
	NoHeader      bool
	Index         bool
	HTMLConfig    HTMLConfig
}

// TabulateRow represents a normalized row in a table.
type TabulateRow struct {
	Elements   []string
	Continuous bool
}

// writeBuffer is a utility for efficient string building.
type writeBuffer struct {
	Buffer bytes.Buffer
}

func newWriteBuffer() *writeBuffer {
	return &writeBuffer{}
}

// Write repeats str count times in the buffer.
func (b *writeBuffer) Write(str string, count int) *writeBuffer {
	for i := 0; i < count; i++ {
		b.Buffer.WriteString(str)
	}
	return b
}

// String returns the buffer contents as a string.
func (b *writeBuffer) String() string {
	return b.Buffer.String()
}

// padRow adds padding to each element in the array.
func (t *Tabulate) padRow(arr []string, padding int) []string {
	if len(arr) < 1 {
		return arr
	}
	padded := make([]string, len(arr))
	for i, el := range arr {
		b := newWriteBuffer()
		b.Write(" ", padding)
		b.Write(el, 1)
		b.Write(" ", padding)
		padded[i] = b.String()
	}
	return padded
}

// padLeft right-aligns text by adding padding on the left.
func (t *Tabulate) padLeft(width int, str string) string {
	b := newWriteBuffer()
	b.Write(" ", (width - runewidth.StringWidth(str)))
	b.Write(str, 1)
	return b.String()
}

// padRight left-aligns text by adding padding on the right.
func (t *Tabulate) padRight(width int, str string) string {
	b := newWriteBuffer()
	b.Write(str, 1)
	b.Write(" ", (width - runewidth.StringWidth(str)))
	return b.String()
}

// padCenter centers text in a cell by adding padding on both sides.
func (t *Tabulate) padCenter(width int, str string) string {
	b := newWriteBuffer()
	strWidth := runewidth.StringWidth(str)
	padding := int(math.Ceil(float64((width - strWidth)) / 2.0))
	b.Write(" ", padding)
	b.Write(str, 1)
	b.Write(" ", (width - runewidth.StringWidth(b.String())))
	return b.String()
}

// buildLine builds a horizontal line based on column widths.
func (t *Tabulate) buildLine(paddedWidths []int, padding []int, l Line) string {
	cells := make([]string, len(paddedWidths))

	for i := range cells {
		b := newWriteBuffer()
		b.Write(l.hline, padding[i]+minPadding)
		cells[i] = b.String()
	}

	var buffer bytes.Buffer
	buffer.WriteString(l.begin)

	// Print contents
	for i := 0; i < len(cells); i++ {
		buffer.WriteString(cells[i])
		if i != len(cells)-1 {
			buffer.WriteString(l.sep)
		}
	}

	buffer.WriteString(l.end)
	return buffer.String()
}

// set Index to automatically add a column for line number
func (t *Tabulate) SetIndex(index bool) {
	t.Index = index
}

// buildRow builds a formatted row based on column widths.
func (t *Tabulate) buildRow(elements []string, paddedWidths []int, paddings []int, d Row) string {

	var buffer bytes.Buffer
	buffer.WriteString(d.begin)
	padFunc := t.getAlignFunc()
	// Print contents
	for i := 0; i < len(paddedWidths); i++ {
		output := ""
		if len(elements) <= i || (len(elements) > i && elements[i] == " nil ") {
			output = padFunc(paddedWidths[i], t.EmptyVar)
		} else if len(elements) > i {
			output = padFunc(paddedWidths[i], elements[i])
		}
		buffer.WriteString(output)
		if i != len(paddedWidths)-1 {
			buffer.WriteString(d.sep)
		}
	}

	buffer.WriteString(d.end)
	return buffer.String()
}

// Render formats the data table using the specified format.
// If no format is provided, uses the TableFormat set in the Tabulate struct.
// For HTML format, uses the HTMLConfig for customization.
func (t *Tabulate) Render(format ...interface{}) string {
	var lines []string

	// Determine the format
	formatStr := ""
	if len(format) > 0 {
		if fs, ok := format[0].(string); ok {
			formatStr = fs
		}
	}

	// Special handling for HTML format - use modern renderHTMLTable
	if formatStr == "html" || formatStr == "html-nohead" {
		if formatStr == "html-nohead" {
			t.NoHeader = true
		}
		return t.renderHTMLTable(t.HTMLConfig)
	}

	// If headers are set use them, otherwise pop the first row
	if len(t.Headers) < 1 {
		if t.NoHeader {
			t.Headers, t.Data = t.Data[0].Elements, t.Data[0:]
		} else {
			t.Headers, t.Data = t.Data[0].Elements, t.Data[1:]
		}
	}

	// Use the format that was passed as parameter, otherwise
	// use the format defined in the struct
	if formatStr != "" {
		if tableFormat, exists := TableFormats[formatStr]; exists {
			t.TableFormat = tableFormat
		}
	}

	// If Wrap Strings is set to True,then break up the string to multiple cells
	if t.WrapStrings {
		t.Data = t.wrapCellData()
	}

	// Check if Data is present
	if len(t.Data) < 1 {
		return fmt.Sprintln("go tabulate render - no data specified")
	}

	// Get Column widths for all columns
	cols := t.getWidths(t.Headers, t.Data)

	paddedWidths := make([]int, len(cols))
	for i := range paddedWidths {
		paddedWidths[i] = cols[i] + minPadding*t.TableFormat.Padding
	}

	// Start appending lines

	// Append top line if not hidden
	if !inSlice("top", t.HideLines) {
		lines = append(lines, t.buildLine(paddedWidths, cols, t.TableFormat.LineTop))
	}

	if !t.NoHeader {
		// Add Header
		if len(t.Headers) < len(t.Data[0].Elements) {
			diff := len(t.Data[0].Elements) - len(t.Headers)
			paddedHeader := make([]string, diff)
			for _, e := range t.Headers {
				paddedHeader = append(paddedHeader, e)
			}
			t.Headers = paddedHeader
		}
		lines = append(lines, t.buildRow(t.padRow(t.Headers, t.TableFormat.Padding), paddedWidths, cols, t.TableFormat.HeaderRow))

		// Add Line Below Header if not hidden
		if !inSlice("belowheader", t.HideLines) {
			lines = append(lines, t.buildLine(paddedWidths, cols, t.TableFormat.LineBelowHeader))
		}
	}

	// Add Data Rows
	for index, element := range t.Data {
		lines = append(lines, t.buildRow(t.padRow(element.Elements, t.TableFormat.Padding), paddedWidths, cols, t.TableFormat.DataRow))
		if index < len(t.Data)-1 {
			if !element.Continuous {
				if !t.RemEmptyLines {
					lines = append(lines, t.buildLine(paddedWidths, cols, t.TableFormat.LineBetweenRows))
				}
			}
		}
	}

	if !inSlice("bottomLine", t.HideLines) {
		lines = append(lines, t.buildLine(paddedWidths, cols, t.TableFormat.LineBottom))
	}

	// Join lines
	var buffer bytes.Buffer
	for _, line := range lines {
		buffer.WriteString(line + "\n")
	}

	return buffer.String()
}

// Calculate the max column width for each element
func (t *Tabulate) getWidths(headers []string, data []*TabulateRow) []int {
	widths := make([]int, len(headers))
	current_max := len(t.EmptyVar)
	for i := 0; i < len(headers); i++ {
		current_max = runewidth.StringWidth(headers[i])
		for _, item := range data {
			if len(item.Elements) > i && len(widths) > i {
				element := item.Elements[i]
				strLength := runewidth.StringWidth(element)
				if strLength > current_max {
					widths[i] = strLength
					current_max = strLength
				} else {
					widths[i] = current_max
				}
			}
		}
	}

	return widths
}

func (t *Tabulate) SetNoHeader() {
	t.NoHeader = true
}

// Set Headers of the table
// If Headers count is less than the data row count, the headers will be padded to the right
func (t *Tabulate) SetHeaders(headers []string) *Tabulate {
	t.Headers = headers
	return t
}

// Set Float Formatting
// will be used in strconv.FormatFloat(element, format, -1, 64)
func (t *Tabulate) SetFloatFormat(format byte) *Tabulate {
	t.FloatFormat = format
	return t
}

// Set Align Type, Available options: left, right, center
func (t *Tabulate) SetAlign(align string) {
	t.Align = align
}

// Select the padding function based on the align type
func (t *Tabulate) getAlignFunc() func(int, string) string {
	if len(t.Align) < 1 || t.Align == "right" {
		return t.padLeft
	} else if t.Align == "left" {
		return t.padRight
	} else {
		return t.padCenter
	}
}

// Set how an empty cell will be represented
func (t *Tabulate) SetEmptyString(empty string) {
	t.EmptyVar = empty + " "
}

// Set which lines to hide.
// Can be:
// top - Top line of the table,
// belowheader - Line below the header,
// bottom - Bottom line of the table
func (t *Tabulate) SetHideLines(hide []string) {
	t.HideLines = hide
}

func (t *Tabulate) SetWrapStrings(wrap bool) {
	t.WrapStrings = wrap
}

func (t *Tabulate) SetRemEmptyLines(remEmptyLines bool) {
	t.RemEmptyLines = remEmptyLines
}

// Sets the maximum size of cell
// If WrapStrings is set to true, then the string inside
// the cell will be split up into multiple cell
func (t *Tabulate) SetMaxCellSize(max int) {
	t.MaxSize = max
}

// If string size is larger than t.MaxSize, then split it to multiple cells (downwards)
func (t *Tabulate) wrapCellData() []*TabulateRow {
	var arr []*TabulateRow
	next := t.Data[0]
	for index := 0; index <= len(t.Data); index++ {
		elements := next.Elements
		newElements := make([]string, len(elements))

		for i, e := range elements {
			if runewidth.StringWidth(e) > t.MaxSize {
				elements[i] = runewidth.Truncate(e, t.MaxSize, "")
				newElements[i] = e[len(elements[i]):]
				next.Continuous = true
			}
		}

		if next.Continuous {
			arr = append(arr, next)
			next = &TabulateRow{Elements: newElements}
			index--
		} else if index+1 < len(t.Data) {
			arr = append(arr, next)
			next = t.Data[index+1]
		} else if index >= len(t.Data) {
			arr = append(arr, next)
		}

	}
	return arr
}

// Create creates a new Tabulate object from various data types.
// Accepts:
//   - 2D String Array ([][]string)
//   - 2D Int Array ([][]int)
//   - 2D Int32 Array ([][]int32)
//   - 2D Int64 Array ([][]int64)
//   - 2D Bool Array ([][]bool)
//   - 2D Float64 Array ([][]float64)
//   - 2D Interface Array ([][]interface{})
//   - 1D String Array ([]string)
//   - 1D Interface Array ([]interface{})
//   - Map[string][]interface{}
//   - Map[string][]string
func Create(data interface{}) *Tabulate {
	t := &Tabulate{
		FloatFormat: 'f',
		MaxSize:     30,
		TableFormat: TableFormats["simple"],
		Align:       "right",
		HTMLConfig:  DefaultHTMLConfig(),
	}

	switch v := data.(type) {
	case [][]string:
		t.Data = createFromString(v)
	case [][]int32:
		t.Data = createFromInt32(v)
	case [][]int64:
		t.Data = createFromInt64(v)
	case [][]int:
		t.Data = createFromInt(v)
	case [][]bool:
		t.Data = createFromBool(v)
	case [][]float64:
		t.Data = createFromFloat64(v, t.FloatFormat)
	case [][]interface{}:
		t.Data = createFromMixed(v, t.FloatFormat)
	case []string:
		t.Data = createFromString([][]string{v})
	case []interface{}:
		t.Data = createFromMixed([][]interface{}{v}, t.FloatFormat)
	case map[string][]interface{}:
		t.Headers, t.Data = createFromMapMixed(v, t.FloatFormat)
	case map[string][]string:
		t.Headers, t.Data = createFromMapString(v)
	}

	return t
}

// SetHTMLConfig sets a custom HTML configuration for table rendering.
func (t *Tabulate) SetHTMLConfig(config HTMLConfig) {
	t.HTMLConfig = config
}

// SetHTMLStyle is a convenience method to set the HTML table class for styling.
func (t *Tabulate) SetHTMLStyle(tableClass string) {
	t.HTMLConfig.TableClass = tableClass
}

// SetHTMLBootstrap sets the Bootstrap version (4, 5, or empty for none).
func (t *Tabulate) SetHTMLBootstrap(version string) {
	t.HTMLConfig.BootstrapVersion = version
	if version == "5" {
		t.HTMLConfig.TableClass = "table table-striped table-hover"
	} else if version == "4" {
		t.HTMLConfig.TableClass = "table table-striped table-hover"
	}
}

// SetHTMLTitle sets an optional title for the HTML table.
func (t *Tabulate) SetHTMLTitle(title string) {
	t.HTMLConfig.Title = title
}

// SetHTMLCustomCSS adds custom CSS to the HTML output.
func (t *Tabulate) SetHTMLCustomCSS(css string) {
	t.HTMLConfig.CustomCSS = css
}

// SetHTMLMinimal uses minimal CSS styling without Bootstrap.
func (t *Tabulate) SetHTMLMinimal() {
	t.HTMLConfig = MinimalHTMLConfig()
}

// SetHTMLOfflineMode enables offline mode with fully embedded Bootstrap CSS.
//
// By default, HTML generation uses CDN links for smaller file sizes (~2KB).
// Call SetHTMLOfflineMode to embed the complete Bootstrap 5 CSS directly
// in the HTML file for use on systems without internet access.
//
// The resulting HTML file will be self-contained and larger (~100KB),
// but can be opened and viewed in any browser regardless of internet connectivity.
//
// This is useful for:
//   - Generating reports on backend servers without internet access
//   - Creating portable HTML documents for email or file sharing
//   - Archiving HTML reports that must work indefinitely
//   - Serving tables in offline or air-gapped environments
//
// Example:
//
//	table := gotabulate.Create(data)
//	table.SetHeaders([]string{"Name", "Value"})
//	table.SetHTMLOfflineMode()
//	html := table.Render("html")
func (t *Tabulate) SetHTMLOfflineMode() {
	t.HTMLConfig = OfflineHTMLConfig()
}

// SetHTMLUseCDN configures whether Bootstrap CSS is loaded from CDN or embedded.
//
// By default (true), Bootstrap is loaded from a CDN for smaller file sizes (~2KB).
// When set to false, the complete Bootstrap 5 CSS is embedded in the HTML (~100KB),
// making the file completely self-contained and usable without internet.
//
// Parameters:
//   - useCDN: true to use CDN (default, smaller files, requires internet)
//   - useCDN: false to embed CSS (larger files, works offline)
//
// See SetHTMLOfflineMode for a convenience method to enable offline mode.
func (t *Tabulate) SetHTMLUseCDN(useCDN bool) {
	t.HTMLConfig.UseBootstrapCDN = useCDN
}
