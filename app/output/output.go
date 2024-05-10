package output

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"

	"github.com/pterm/pterm"
)

type OutputPrinter interface {
	Println(args ...any)
	PrintAsTable(data any, header []string, rowMapper func(any) []string)
	Errorf(format string, a ...any)
	Infof(format string, a ...any)
}

var out OutputPrinter

func init() {
	out = PTermOutputPrinter{}
}

func Println(a ...any) {
	out.Println(a...)
}

func PrintAsTable(data any, rowMapper func(any) []string) {
	out.PrintAsTable(data, make([]string, 0), rowMapper)
}

func PrintAsTableH(data any, header []string, rowMapper func(any) []string) {
	out.PrintAsTable(data, header, rowMapper)
}

func Errorf(format string, a ...any) {
	out.Errorf(format, a...)
}

func Errorln(a any) {
	out.Errorf("%s\n", a)
}

func Infoln(a any) {
	out.Infof("%s\n", a)
}

func Infof(format string, a ...any) {
	out.Infof(format, a...)
}

/*
*
----- STDIO
*/
type StdioOutputPrinter struct {
}

func (self StdioOutputPrinter) Println(args ...any) {
	fmt.Println(args...)
}

func (self StdioOutputPrinter) Errorf(format string, a ...any) {
	fmt.Errorf(format, a)
}

func (self StdioOutputPrinter) Infof(format string, a ...any) {
	fmt.Printf(format, a)
}

func (self StdioOutputPrinter) PrintAsTable(data any, header []string, rowMapper func(any) []string) {
	fmt.Println("")
	w := tabwriter.NewWriter(os.Stdout, 1, 8, 4, ' ', 0)

	refData := reflect.ValueOf(data)
	if refData.Kind() == reflect.Map {
		for _, key := range refData.MapKeys() {
			row := refData.MapIndex(key)
			// now we have key.Interface(), strct.Interface()
			cols := rowMapper(row.Interface())
			// fmt.Println(key.Interface(), row.Interface())
			fmt.Fprintln(w, strings.Join(cols, "\t"))
		}
	} else if refData.Kind() == reflect.Slice {
		fmt.Errorf("unable to print. slices are not supported. ")
	} else {
		fmt.Errorf("unable to print. input data is not a map")
	}

	//	for _, row := range data.(map[interface{}]interface{}) {
	//		cols := rowMapper(row)
	//		fmt.Fprintln(w, strings.Join(cols, "\t"))
	//
	//	}
	w.Flush()
}

/**
----- PTerm
*/

type PTermOutputPrinter struct {
}

func (self PTermOutputPrinter) PrintAsTable(data any, header []string, rowMapper func(any) []string) {
	tableData := pterm.TableData{header}

	refData := reflect.ValueOf(data)
	if refData.Kind() == reflect.Map {
		for _, key := range refData.MapKeys() {
			row := refData.MapIndex(key)
			// now we have key.Interface(), strct.Interface()
			cols := rowMapper(row.Interface())
			// fmt.Println(key.Interface(), row.Interface())
			tableData = append(tableData, cols)
		}
	} else if refData.Kind() == reflect.Slice {
		fmt.Errorf("unable to print. slices are not supported. ")
	} else {
		fmt.Errorf("unable to print. input data is not a map")
	}

	// Create a table with a header and the defined data, then render it
	pterm.DefaultTable.WithHasHeader().WithSeparator("\t").WithData(tableData).Render()
}

func (self PTermOutputPrinter) Println(a ...any) {
	pterm.Println(a...)
}

func (self PTermOutputPrinter) Errorf(format string, a ...any) {
	pterm.Error.Printf(format, a...)
}

func (self PTermOutputPrinter) Infof(format string, a ...any) {
	pterm.Info.Printf(format, a...)
}

// --------------------------------------------------------------------------
func PrintTidySlice[T any](slice []T, header []string, rowMapper func(any) []string) {
	tableData := pterm.TableData{header}
	last := len(slice) - 1
	i := 0
	for i < last {
		cols := rowMapper(slice[i])
		tableData = append(tableData, cols)
		i++
	}
	pterm.DefaultTable.WithHasHeader().WithSeparator("\t").WithData(tableData).Render()
}
