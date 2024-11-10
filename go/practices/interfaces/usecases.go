package main

import (
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
	err := os.WriteFile(fileName+fileType, []byte(info), 0644)
	if err != nil {
		return err
	}
	return nil
}
func (file *LogWriter) logWriter(writer FileWriter) error {
	err := writer.write(file.fileName, file.fileType)
	if err != nil {
		err.Error()
	}
	return nil
}
func (file *LogWriter) execute() error {

	if file.schema == "car" && file.fileType == ".txt" {
		car := Car{name: "Audi", color: "Black"}
		file.logWriter(&car)

	} else {
		err := InvalidSchemaFileTypeError{"Invalid schema and file type"}
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
