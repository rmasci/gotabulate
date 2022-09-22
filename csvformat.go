package gotabulate

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

/*-
 * ============LICENSE_START=======================================================
 * Author: Richard Masci
 * ================================================================================
 * Copyright (C) 2017 - 2020 AT&T Intellectual Property. All rights reserved.
 * ================================================================================
 * The MIT License (MIT)
 *
 * Copyright <year> AT&T Intellectual Property. All other rights reserved.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software
 * and associated documentation files (the "Software"), to deal in the Software without
 * restriction, including without limitation the rights to use, copy, modify, merge, publish,
 * distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or
 * substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN
 * AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
 * WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 * ============LICENSE_END=========================================================
 */

type CSVOut struct {
	Text          string
	Render        string
	Columns       string
	RemEmptyLines bool
	NoHead        bool
	Align         string
	Wrap          bool
	Indent        string
	Delimeter     string
}

// NewCSVOut Create struct for csvout table.
// Text: Raw csv text to make into a table.
// Columns: custom columns -- will use first line if this is blank. CSV Formatted one,two,three...
// LineBetweenRow -- Output has a line between the rows
// Wrap -- Wrap long lines
// NoHead -- Do not print a header
// Indent -- String, defaults to "\t"
// Render: This is how you want to render the  output: can be one of:
//     "simple" -- Simple Underline for header and underling at end, contents tabed in.
//     "plain" -- Plain text
//     "tab" -- Tab based output
//     "html" -- Outputs in html table format
//     "mysql" -- Looks like mysql shell output
//     "grid" -- Spreadsheet like grid
//     "gridt" -- Text based grid, like MySQL but with rowlines

func NewCSVOut() *CSVOut {
	return &CSVOut{
		Render:        "simple",
		Indent:        "\t",
		Align:         "left",
		Delimeter:     ",",
		RemEmptyLines: true,
	}
}

// Table takes csv text and turns it into a table output
func (csvout *CSVOut) Table() string {
	var columns string
	lines := strings.Split(strings.TrimSpace(csvout.Text), "\n")
	var minLength int
	masterStr := [][]string{}
	if csvout.Columns != "" {
		var t []string
		t = strings.Split(csvout.Columns, ",")
		masterStr = append(masterStr, t)
		csvout.NoHead = false
	}
	minLength = 1
	for _, line := range lines {
		if line != "" {
			var lineArray []string
			lineArray = PrintCol(line, columns, false)
			masterStr = append(masterStr, lineArray)
		}
	}
	if len(masterStr) <= minLength {
		csvout.NoHead = true
	}
	// Send to gridulate
	gridulate := Create(masterStr)
	gridulate.NoHeader = csvout.NoHead
	//gridulate.SetHeaders(c1)
	gridulate.SetAlign(csvout.Align)
	gridulate.SetWrapStrings(csvout.Wrap)
	gridulate.SetRemEmptyLines(csvout.RemEmptyLines)
	/*
		scanner := bufio.NewScanner(strings.NewReader(gridulate.Render(csvout.Render)))
		for scanner.Scan() {
			fmt.Println(csvout.Indent + scanner.Text())
		}
		return
	*/
	return gridulate.Render(csvout.Render)
}

func PrintCol(line, columns string, space bool) []string {
	var delimeter = ","
	numcolumns := false
	var l []string
	var out []string
	if space {
		l = strings.Fields(line)
		if columns == "" {
			return l
		}
	} else {
		l = strings.Split(line, delimeter)
		if columns == "" {
			return l
		}
	}
	for _, idx := range strings.Split(columns, ",") {
		i, err := strconv.Atoi(idx)
		if err != nil {
			fmt.Printf("ERROR: in columns passed: %v\nColumns need to be a comma separated list of numbers ex '-c 1,4,5,6,7'.\n", err)
			os.Exit(1)
		}

		if i <= len(l) && i >= 0 {
			if numcolumns {
				tstr := fmt.Sprintf("%v %v", i, l[i])
				out = append(out, tstr)

			} else {
				out = append(out, l[i])
			}
		} else {
			continue
		}
	}
	return out
}

// Take CSV formatted text (csvIn String) and print it out as an XLS file.
func ExcelOut(csvIn, sheetname, excelFile string) error {
	lines := strings.Split(csvIn, "\n")
	xlsxFile := xlsx.NewFile()
	if sheetname == "" {
		sheetname = "Sheet1"
	}
	sheet, err := xlsxFile.AddSheet(sheetname)
	if err != nil {
		return err
	}

	for _, l := range lines {
		lArr :=strings.Split(l, ",")
		row := sheet.AddRow()
		for _, field := range lArr {
			cell := row.AddCell()
			cell.Value = field
		}
	}
	if err != nil {
		fmt.Printf(err.Error())
	}
	return xlsxFile.Save(excelFile)
}
