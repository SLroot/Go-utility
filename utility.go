package goutility

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// take an integer input from user(terminal, console)
func IntInputFromUser() int {
	fmt.Print("enter your number: ")
	iinput := bufio.NewReader(os.Stdin)
	in, err := iinput.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Replace(in, "\n", "", 1)
	intInput, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return intInput
}
