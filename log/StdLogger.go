package log

import "fmt"

type StdLogger struct{}

func (sl *StdLogger) Log(s string) {
	fmt.Print(s)
}
