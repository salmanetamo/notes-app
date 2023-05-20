package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Contains[V comparable](list []V, lookup V) bool {
	for _, v := range list {
		if v == lookup {
			return true
		}
	}

	return false
}

func GetNumberUserInput() int {
	var in = bufio.NewReader(os.Stdin)
	choice, _ := in.ReadString('\n')
	choiceTrimmed := strings.TrimSpace(choice)
	choiceAsInt, err := strconv.Atoi(choiceTrimmed)
	if err != nil {
		return -1
	}
	return choiceAsInt
}

func GetTextUserInput() (string, bool) {
	var in = bufio.NewReader(os.Stdin)
	choice, err := in.ReadString('\n')
	if err != nil {
		return "", false
	}
	choiceTrimmed := strings.TrimSpace(choice)
	if choiceTrimmed != "" {
		return choiceTrimmed, true
	}
	return "", false
}

func GetBooleanUserInput() (bool, bool) {
	var in = bufio.NewReader(os.Stdin)
	choice, err := in.ReadString('\n')
	if err != nil {
		return false, false
	}
	choiceTrimmed := strings.TrimSpace(choice)
	choiceLowercase := strings.ToLower(choiceTrimmed)
	containsTrue := Contains([]string{"yes", "y"}, choiceLowercase)
	containsFalse := Contains([]string{"no", "n"}, choiceLowercase)
	if containsTrue {
		return true, true
	}
	if containsFalse {
		return false, true
	}
	return false, false
}
