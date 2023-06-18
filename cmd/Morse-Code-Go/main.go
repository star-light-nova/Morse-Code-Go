package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/StarLightNova/Morse-Code-Go/pkg/morse"
)

func main() {
    var userChoice string

    fmt.Println("Welcome to the Morse Code Encoder/Decoder.\nWhat do you want to do? (1: encode, 2: decode)")

    fmt.Scanln(&userChoice)

    fmt.Println("\nEnter your text: ")

    userInput := inputReader()

    if userChoice == "1" {
        fmt.Println("\nYour encoded text:")
        fmt.Println(morse.Encode(userInput))
    } else if userChoice == "2" {
        fmt.Println("\nYour decoded text:")
        fmt.Println(morse.Decode(userInput))
    }
}

func inputReader() string {
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')

    return text
}
