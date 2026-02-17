# Fix Summary: Removed Blank Lines Between Rows

## Problem
When rendering tables in formats that don't have `LineBetweenRows` defined (like "mysql", "simple", etc.), blank lines were appearing between data rows.

## Root Cause
In the `Render()` method around line 519 of `tabulate.go`, when `LineBetweenRows` was not defined (zero-value `Line{}`), the code would still call `buildLine()` which returned an empty string, and that empty string was being added to the lines array. When the output was assembled, this empty line would get a newline character appended, creating a blank line in the output.

## Solution
Modified the code to check if the line returned by `buildLine()` is empty before adding it to the lines array:

```go
// Old code (line 519):
lines = append(lines, t.buildLine(paddedWidths, cols, t.TableFormat.LineBetweenRows))

// New code (lines 519-522):
line := t.buildLine(paddedWidths, cols, t.TableFormat.LineBetweenRows)
if line != "" {
    lines = append(lines, line)
}
```

## Result
- Formats WITHOUT `LineBetweenRows` defined (mysql, simple, plain, etc.) now correctly show NO blank lines between rows
- Formats WITH `LineBetweenRows` defined (grid, gridt, etc.) continue to show separator lines between rows as expected
- No need to pipe output through `sed '/^$/d'` anymore!

## Testing
Run your original test:
```bash
echo -e "Name,Age,City\nAlice,30,NYC\nBob,25,LA\nCharlie,35,Chicago" | /tmp/csvtable
```

The output should now have NO blank lines between the data rows.

