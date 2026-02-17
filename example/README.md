# Gotabulate Examples

This directory contains example programs demonstrating different features of gotabulate.

## Example Programs

### usage.go
Demonstrates basic table rendering with multiple output formats. Shows how to create a table from interface data, set headers, configure alignment, and render to different formats.

**Run:** `go run usage.go`

### usage2.go
Simple example showing how to render data in MySQL format. Minimal setup for quick table generation.

**Run:** `go run usage2.go`

### csvOut.go
Demonstrates parsing CSV text and rendering it as a formatted table. Shows the CSVOut feature for converting raw CSV data.

**Run:** `go run csvOut.go`

### html_example.go
Comprehensive HTML table generation examples including Bootstrap 5 (CDN), Bootstrap 4, minimal CSS, and custom styled tables. Shows when to use offline mode and CDN mode.

**Run:** `go run html_example.go`

---

## Format Samples

All samples below use this data:
```go
data := [][]string{
    {"John Doe", "30", "Engineer", "$120,000"},
    {"Jane Smith", "28", "Designer", "$110,000"},
    {"Bob Johnson", "35", "Manager", "$130,000"},
    {"Alice Brown", "26", "Developer", "$105,000"},
}
table := gotabulate.Create(data)
table.SetHeaders([]string{"Name", "Age", "Position", "Salary"})
```

### Table of Contents
- **Text Grids:** simple, grid, gridt, mysql, mysqlg, bingo
- **Data Formats:** csv, tab, plain, text
- **Markup:** markdown, rst
- **HTML:** CDN (default) or offline mode

---

## simple

```
----------------  --------  --------------  -------------
           Name       Age        Position         Salary 
----------------  --------  --------------  -------------
       John Doe        30        Engineer       $120,000 

     Jane Smith        28        Designer       $110,000 

    Bob Johnson        35         Manager       $130,000 

    Alice Brown        26       Developer       $105,000 
----------------  --------  --------------  -------------
```

## grid

```
╒════════════════╤════════╤══════════════╤═════════════╕
│           Name │    Age │     Position │      Salary │
╞════════════════╪════════╪══════════════╪═════════════╡
│       John Doe │     30 │     Engineer │    $120,000 │
│────────────────┼────────┼──────────────┼─────────────│
│     Jane Smith │     28 │     Designer │    $110,000 │
│────────────────┼────────┼──────────────┼─────────────│
│    Bob Johnson │     35 │      Manager │    $130,000 │
│────────────────┼────────┼──────────────┼─────────────│
│    Alice Brown │     26 │    Developer │    $105,000 │
└────────────────┴────────┴──────────────┴─────────────┘
```

## gridt

```
+----------------+--------+--------------+-------------+
           Name     Age      Position       Salary 
+================+========+==============+=============+
       John Doe      30      Engineer     $120,000 
+----------------+--------+--------------+-------------+
     Jane Smith      28      Designer     $110,000 
+----------------+--------+--------------+-------------+
    Bob Johnson      35       Manager     $130,000 
+----------------+--------+--------------+-------------+
    Alice Brown      26     Developer     $105,000 
+----------------+--------+--------------+-------------+
```

## mysql

```
+----------------+--------+--------------+-------------+
           Name     Age      Position       Salary 
+================+========+==============+=============+
       John Doe      30      Engineer     $120,000 

     Jane Smith      28      Designer     $110,000 

    Bob Johnson      35       Manager     $130,000 

    Alice Brown      26     Developer     $105,000 
+----------------+--------+--------------+-------------+
```

## mysqlg

```
╒════════════════╤════════╤══════════════╤═════════════╕
│           Name │    Age │     Position │      Salary │
╞════════════════╪════════╪══════════════╪═════════════╡
│       John Doe │     30 │     Engineer │    $120,000 │

│     Jane Smith │     28 │     Designer │    $110,000 │

│    Bob Johnson │     35 │      Manager │    $130,000 │

│    Alice Brown │     26 │    Developer │    $105,000 │
└────────────────┴────────┴──────────────┴─────────────┘
```

## bingo

```
╒════════════════╤════════╤══════════════╤═════════════╕
│           Name │    Age │     Position │      Salary │
╞════════════════╪════════╪══════════════╪═════════════╡
│       John Doe │     30 │     Engineer │    $120,000 │
│────────────────┼────────┼──────────────┼─────────────│
│     Jane Smith │     28 │     Designer │    $110,000 │
│────────────────┼────────┼──────────────┼─────────────│
│    Bob Johnson │     35 │      Manager │    $130,000 │
│────────────────┼────────┼──────────────┼─────────────│
│    Alice Brown │     26 │    Developer │    $105,000 │
└────────────────┴────────┴──────────────┴─────────────┘
```

## csv

```
Name,Age,Position,Salary
John Doe,30,Engineer,$120,000
Jane Smith,28,Designer,$110,000
Bob Johnson,35,Manager,$130,000
Alice Brown,26,Developer,$105,000
```

## tab

```
Name	Age	Position	Salary
John Doe	30	Engineer	$120,000
Jane Smith	28	Designer	$110,000
Bob Johnson	35	Manager	$130,000
Alice Brown	26	Developer	$105,000
```

## plain

```
           Name       Age        Position         Salary 
       John Doe        30        Engineer       $120,000 
     Jane Smith        28        Designer       $110,000 
    Bob Johnson        35         Manager       $130,000 
    Alice Brown        26       Developer       $105,000 
```

## text

```
Name Age Position Salary
John Doe 30 Engineer $120,000
Jane Smith 28 Designer $110,000
Bob Johnson 35 Manager $130,000
Alice Brown 26 Developer $105,000
```

## markdown

```
|Name|Age|Position|Salary|
|---|---|---|---|
|John Doe|30|Engineer|$120,000|
|Jane Smith|28|Designer|$110,000|
|Bob Johnson|35|Manager|$130,000|
|Alice Brown|26|Developer|$105,000|
```

## rst

```
================ ======== ============== =============
           Name       Age        Position         Salary 
================ ======== ============== =============
       John Doe        30        Engineer       $120,000 
     Jane Smith        28        Designer       $110,000 
    Bob Johnson        35         Manager       $130,000 
    Alice Brown        26       Developer       $105,000 
================ ======== ============== =============
```

## HTML (Default: CDN)

```go
table.Render("html")  // Bootstrap 5 via CDN (~2KB)
```

Produces a modern, professional HTML table with Bootstrap 5 styling. CSS is loaded from CDN.

## HTML (Offline Mode)

```go
table.SetHTMLOfflineMode()
table.Render("html")  // Bootstrap 5 embedded (~100KB)
```

Generates a self-contained HTML file with all CSS embedded. Works offline on any system.

---

## No-Header Variants

Append `-nohead` to any format to suppress the header row:

```go
table.Render("simple-nohead")
table.Render("grid-nohead")
table.Render("csv-nohead")
// ... etc for all formats
```

