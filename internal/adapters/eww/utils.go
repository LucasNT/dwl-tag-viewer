package eww

import "os"

func ChangeStdoutToStderr() *os.File {
	ret := os.Stdout
	os.Stdout = os.Stderr
	return ret
}
