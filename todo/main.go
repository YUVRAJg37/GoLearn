package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TodoItem struct {
	Id          int    `json:"id"`
	IsDone      bool   `json:"isDone"`
	Description string `json:"description"`
}

type mode int

const (
	DELETE mode = iota
	UPDATE
)

const FILE_NAME = "data.json"

var reader = bufio.NewReader(os.Stdin)

func main() {
	for {
		println("Choose the Options.")
		println("1. Add Item")
		println("2. Remove Item")
		println("3. Mark Item")
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

		switch userInput {
		//-----Creation-----
		case "1":
			input, err := takeInputFromUser("Enter your task: ")
			if err != nil {
				fmt.Println(err)
				return
			}

			if strings.TrimSpace(input) == "" {
				fmt.Println("Input not valid")
				return
			}

			err = handleItemAddition(input)

			if err != nil {
				fmt.Println(err)
				return
			}
		//-----Deletion-----
		case "2":
			input, err := takeInputFromUser("Enter item id to remove: ")
			if err != nil {
				fmt.Println(err)
				return
			}
			id, err := strconv.Atoi(input)

			if err != nil {
				fmt.Println(err)
				return
			}

			err = handleUpdation(id, DELETE)

			if err != nil {
				fmt.Println(err)
				return
			}
		//-----Mark-----
		case "3":
			input, err := takeInputFromUser("Enter item id to mark: ")
			if err != nil {
				fmt.Println(err)
				return
			}
			id, err := strconv.Atoi(input)

			if err != nil {
				fmt.Println(err)
				return
			}

			err = handleUpdation(id, UPDATE)

			if err != nil {
				fmt.Println(err)
				return
			}
		default:
			fmt.Println("Invalid Input")
			continue
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
}

func takeInputFromUser(prompt string) (string, error) {
	if prompt != "" {
		fmt.Println(prompt)
	}
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

func handleItemAddition(description string) error {
	jsonData, err := ReadJson[TodoItem](FILE_NAME)

	if err != nil {
		return err
	}

	data := TodoItem{
		Id:          len(jsonData),
		IsDone:      false,
		Description: description,
	}

	jsonData = append(jsonData, data)
	fmt.Println(jsonData)
	return WriteJson(string(FILE_NAME), jsonData)
}

func handleUpdation(id int, m mode) error {
	jsonData, err := ReadJson[TodoItem](FILE_NAME)

	if err != nil {
		return err
	}

	if len(jsonData) == 0 {
		return errors.New("item doesnt exist")
	}

	if m == UPDATE {
		itemIndex := -1
		for index, item := range jsonData {
			if item.Id == id {
				itemIndex = index
				break
			}
		}

		if itemIndex == -1 {
			return errors.New("item doesnt exist")
		}
		jsonData[itemIndex].IsDone = !jsonData[itemIndex].IsDone
		return WriteJson(string(FILE_NAME), jsonData)
	} else if m == DELETE {
		var data []TodoItem = make([]TodoItem, 0)
		for _, item := range jsonData {
			if item.Id != id {
				data = append(data, item)
			}
		}

		if len(data) == 0 {
			return errors.New("item doesnt exist")
		}

		return WriteJson(string(FILE_NAME), data)
	}

	return nil
}
