package menu

import (
	"bufio"
	"fmt"
	"os"
)

import "github.com/shrekie/phonegok/phone"

var reader = bufio.NewReader(os.Stdin)

func readInput(t string) string {
	fmt.Print(t)
	text, _ := reader.ReadString('\n')
	return text[:len(text)-1]
}

func StartMenu() {
	for {
		choice := readInput("'search' or 'create'")
		fmt.Println(choice)
		switch choice {
		case "search":
			search()
		case "create":
			create()
		default:
			fmt.Println("Too far away.")
		}
	}
}

func search() {
	choice := readInput("enter name")
	result := phone.ContactByName(choice)
	fmt.Println(result)
}

func create() {
	number := readInput("please type number")
	name := readInput("please type name")
	contact := phone.PhoneContact{Number: number, Name: name}
	phone.AddContact(contact)
}
