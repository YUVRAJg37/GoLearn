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
		println("4. Show All")
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
			handleShowAll()
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
			handleShowAll()
		//-----Mark-----
		case "3":
			handleShowAll()
			input, err := takeInputFromUser("Enter item id to mark: ")
			if err != nil {
				fmt.Println(err)
				return
			}
			index, err := strconv.Atoi(input)

			if err != nil {
				fmt.Println(err)
				return
			}

			err = handleUpdation(index-1, UPDATE)

			if err != nil {
				fmt.Println(err)
				return
			}
			handleShowAll()
		case "4":
			handleShowAll()
		default:
			fmt.Println("Invalid Input")
			continue
		}

		userInput, err = takeInputFromUser("Press Enter to continue....?")

		if err != nil {
			fmt.Println("Error reading input: ", err)
			break
		}

		if strings.ToLower(userInput) == "\n" {
			continue
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
	return WriteJson(string(FILE_NAME), jsonData)
}

func handleUpdation(index int, m mode) error {
	jsonData, err := ReadJson[TodoItem](FILE_NAME)

	if err != nil {
		return err
	}

	if len(jsonData) == 0 || index > len(jsonData) {
		return errors.New("item doesnt exist")
	}

	switch m {
	case UPDATE:
		jsonData[index].IsDone = !jsonData[index].IsDone
		return WriteJson(string(FILE_NAME), jsonData)
	case DELETE:
		var data []TodoItem = make([]TodoItem, 0)
		count := 0
		for id, item := range jsonData {
			count++
			jsonData[id].Id = count
			if item.Id != index {
				data = append(data, item)
			}
		}
		fmt.Println(count)

		if len(data) == 0 {
			return errors.New("item doesnt exist")
		}

		return WriteJson(string(FILE_NAME), data)
	}

	return nil
}

func handleShowAll() error {
	jsonData, err := ReadJson[TodoItem](FILE_NAME)

	if err != nil {
		return err
	}

	for index, item := range jsonData {
		tick := " "
		if item.IsDone {
			tick = "✓"
		}
		fmt.Printf("%d.[%s] %s\n", index+1, tick, item.Description)
	}

	return nil
}
