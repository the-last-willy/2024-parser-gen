package log

import "os"

type FileLogger struct {
	file *os.File
}

func OpenFileLogger(filepath string) *FileLogger {
	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	return &FileLogger{
		file: file,
	}
}

func (fl *FileLogger) Close() {
	_ = fl.file.Close()
}

func (fl *FileLogger) Log(s string) {
	if fl.file == nil {
		panic("")
	}
	_, err := fl.file.WriteString(s)
	if err != nil {
		panic(err)
	}
}
