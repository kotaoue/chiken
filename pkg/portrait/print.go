package portrait

import "fmt"

var verbose bool

func vPrint(a ...interface{}) (int, error) {
	if verbose {
		return fmt.Print(a...)
	}
	return 0, nil
}

func vPrintln(a ...interface{}) (int, error) {
	if verbose {
		return fmt.Println(a...)
	}
	return 0, nil
}

func vPrintf(format string, a ...interface{}) (int, error) {
	if verbose {
		return fmt.Printf(format, a...)
	}
	return 0, nil
}
