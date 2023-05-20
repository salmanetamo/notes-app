package menu

import (
	"bufio"
	"fmt"
	"notes/utils"
	"os"
	"strconv"
	"strings"
)

type menuItem struct {
	name string
}

var in = bufio.NewReader(os.Stdin)

type menu []menuItem

func (m menu) print() {
	fmt.Println("What do you want to do? (Select an option)")
	for index, item := range m {
		fmt.Printf("\t%v - %v\n", index+1, item.name)
	}
	fmt.Println("\tq - Quit")
}

func (m menu) getUserChoice() int {
	var validOptions []int

	for index := range menuItems {
		validOptions = append(validOptions, index+1)
	}

	choice, _ := in.ReadString('\n')
	choiceTrimmed := strings.TrimSpace(choice)

	if utils.Contains([]string{"q", "Q"}, choiceTrimmed) {
		return 0
	}

	choiceAsInt, err := strconv.Atoi(choiceTrimmed)
	if err != nil {
		return -1
	}
	if utils.Contains(validOptions, choiceAsInt) {
		return choiceAsInt
	}
	return -1
}

func Print() {
	menuItems.print()
}

func GetUserChoice() int {
	return menuItems.getUserChoice()
}
