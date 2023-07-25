package main

import "fmt"

func logInfo(info string) {
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 33, "[INFO]")
	fmt.Println(colored, info)
}

func logError(err string) {
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, "[INFO]")
	fmt.Println(colored, err)
}
