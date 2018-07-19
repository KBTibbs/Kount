package main

import (
	"fmt"
	"github.com/fatih/color"
	"bufio"
	"os"
	"strings"
	"strconv"
	"github.com/kbtibbs/happy_numbers"
)

func main() {
	// Introduction to the application
	fmt.Fprintf(color.Output, "Welcome to the %s \n", color.YellowString("Happy Number Tester"))
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter a positive integer you wish to test: ")
	text, err := reader.ReadString('\n')

	if err == nil {
		// Always test input for safety
		num, err := parseInput(text)
		if err == nil {
			happy, err := happynumber.TestNumber(num)

			if err == nil {
				// report if the number is happy or not.
				if happy == true {
					fmt.Fprintf(color.Output, "%d is a %s \n", num, color.YellowString("Happy Number!"))
				} else {
					fmt.Fprintf(color.Output, "%d is %s a %s \n", num, color.RedString("not"), 
						color.YellowString("Happy Number!"))
				}
			} else {
				fmt.Println(err)
			}
		}
	}

	if err != nil {
		fmt.Println(err)
	}
}

func parseInput(text string) (int, error) {
	// remove whitespace
	text = strings.TrimSpace(text)
	// if the text cannot convert to an Int, reject it.
	num, convErr := strconv.Atoi(text)
	if convErr != nil {
		fmt.Println("Text is unparsable: '", text, "'.")
		fmt.Println("This is not a positive integer or is too large to work with.")
	}
	return num, convErr
}