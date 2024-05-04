package output

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"
)

func PrintTidySlice[T any](slice []T, rowMapper func(any) []string) {
	last := len(slice) - 1
	i := 0
	for i < last {
		fmt.Println(slice[i])
		i++
	}
}

func PrintAsTable(data any, rowMapper func(any) []string) {
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
