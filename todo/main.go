package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TodoItem struct {
	id          int
	isDone      bool
	description string
}

var reader = bufio.NewReader(os.Stdin)

func main() {

	file, err := os.OpenFile("data.csv", os.O_APPEND|os.O_CREATE, os.ModeAppend)

	if err != nil {
		fmt.Println("Error Reading CSV: ", err)
		file.Close()
	}

	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()

	if err != nil {
		fmt.Println("Error Reading CSV: ", err)
		file.Close()
	}

	writer := csv.NewWriter(file)

	var index int = len(records)

	for {
		println("Choose the Options.")
		println("1. Add Item")
		println("(Q/q) quit.")
		userInput, err := takeInputFromUser("")

		if err != nil {
			fmt.Println("Error reading input: ", err)
			break
		}

		if strings.ToLower(userInput) == "q" {
			fmt.Println("Exiting....")
			break
		}

		if userInput == "1" {
			input, err := takeInputFromUser("Enter your task: ")
			if err != nil {
				fmt.Println("Error reading task: ", err)
				break
			}

			index = index + 1
			data := []string{strconv.Itoa(index), "false", input}

			err = writer.Write(data)

			if err != nil {
				fmt.Println("Error Writing CSV: ", err)
				file.Close()
				writer.Flush()
				break
			}

			fmt.Println("Task Added: ", input)
		}

		userInput, err = takeInputFromUser("Continue....?[(Y/y)/(N/n)]")

		if err != nil {
			fmt.Println("Error reading input: ", err)
			break
		}

		if strings.ToLower(userInput) == "y" {
			continue
		} else if strings.ToLower(userInput) == "n" {
			break
		}

	}

	writer.Flush()
	err = writer.Error()
	if err != nil {
		fmt.Println(err)
	}
}

func takeInputFromUser(prompt string) (string, error) {
	if prompt != "" {
		fmt.Println(prompt)
	}
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}
