package main

import (
	"fmt"
	"os"
)

type BaseExtractUseCase interface {
	extract() string
}

type LogWriter struct {
	schema   string
	fileType string
	fileName string
}

type FileWriter interface {
	write(fileName string, fileType string) error
}
type Car struct {
	name  string
	color string
}

func (car *Car) extract() string {
	return car.name + " " + car.color
}

func (car *Car) write(fileName string, fileType string) error {
	info := car.extract()
	file, err := os.Create(fileName + fileType)
	if err != nil {
		return err
	}
	defer func() {
		closeFileError := file.Close()
		if closeFileError != nil {
			panic(closeFileError)
		}
	}()
	file.WriteString(info)
	return nil
}
func (file *LogWriter) logWriter(writer FileWriter) error {
	err := writer.write(file.fileName, file.fileType)
	if err != nil {
		err.Error()
	}
	return nil
}
func infoLog(value interface{}) {
	val, ok := value.(*LogWriter)
	if ok {
		fmt.Printf("[Info] Operation for %v interface successful.\n", val)
	}
}
func (file *LogWriter) execute() error {

	if file.schema == "car" && file.fileType == ".txt" {
		car := Car{name: "Audi", color: "Black"}
		file.logWriter(&car)
		infoLog(file)

	} else {
		err := InvalidSchemaFileTypeError{"[Error] Invalid schema and file type."}
		return &err
	}
	return nil
}

type InvalidSchemaFileTypeError struct {
	errorMessage string
}

func (message *InvalidSchemaFileTypeError) Error() string {
	return message.errorMessage
}
