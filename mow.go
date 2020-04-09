package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	printMow()
	mowGame()
}

func printMow() {
	// YUM!!!!
	fmt.Println("yum")

	// SLOW WAY TO DECLARE AND SET A VARIABLE
	// declare variabubble
	var m string

	// set variabubble m's value to "Mow"
	m = "Mow"

	// printing/using my variabubble
	fmt.Println(m)

	m = "MowMow"

	// FAST WAY TO DECLARE AND SET A VARIABLE
	mmm := "MowMowMow"

	fmt.Println(mmm)
	// printing/using my variabubble
	fmt.Println(m)
}

func mowGame() {
	// declare a variabubble reader, which takes input from the
	// operating system "standard input" aka the terminal input
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user with a question
	fmt.Println("What is your name?")

	// the reader reads to you because you're an idiot and computers are all powerful
	// read the string from the standard input until an "end of line character is submitted"
	// \n is an end of line character in unix/linux based systems, which this Mac runs on
	name, _ := reader.ReadString('\n')

	// remove the end of line character to just get the name
	name = strings.Replace(name, "\n", "", -1)

	// loop
	for {
		fmt.Printf("%s, a pleasure to make your acquaintance - would you like me to eat your soul?\n", name)
		fmt.Println("Type y for yes or n for no")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("y", text) == 0 {
			fmt.Println("Your soul has been eaten - delicious!")
			os.Exit(0)
		} else {
			fmt.Println("aw man, but it looked so good!!!!")
		}

	}
}
