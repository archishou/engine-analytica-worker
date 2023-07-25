package logging

import "fmt"

func LogInfo(info string) {
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 33, "[INFO]")
	fmt.Println(colored, info)
}

func LogError(err string) {
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, "[INFO]")
	fmt.Println(colored, err)
}
