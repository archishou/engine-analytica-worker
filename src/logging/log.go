package logging

import "fmt"

func LogInfo(info ...any) {
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 33, "[INFO]")
	fmt.Println(colored, info)
}

func LogError(err ...any) {
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, "[ERROR]")
	fmt.Println(colored, err)
}
