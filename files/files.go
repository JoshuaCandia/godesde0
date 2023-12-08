package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/JoshuaCandia/godesde0/exercises"
)

var fileName string = "./files/txt/table.txt"

func TableRecorder() {
	var text string = exercises.Multiply()
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("error creating the file" + err.Error())
		return
	}
	fmt.Fprintln(file, text)
	file.Close()
}

func PlusTable() {
	var text string = exercises.Multiply()
	if !Append(text) {
		fmt.Println("Error appending content")
	}
}

func Append(text string) bool {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error during append" + err.Error())
		return false
	}

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error during WriteString" + err.Error())
		return false
	}
	file.Close()
	return true
}

func ReadFile() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error reading the file" + err.Error())
		return
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		record := scanner.Text()
		fmt.Println("> " + record)
	}
	file.Close()
}
