package morse

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func prompt() {
    var userChoice string
    var userText string

    fmt.Println("Welcome to the Morse Code Encoder/Decoder 🙋‍♂️")

    // Ask till we get an expected prompt.
    for choice, error := encodeDecodeChoice(); ; choice, error = encodeDecodeChoice() {
        userChoice = choice

        if error == nil { break }
    }

    // Ask till we get an expected prompt.
    for text, error := textPrompt(); ; text, error = textPrompt() {
        userText = text

        if error == nil { break }
    }

    if userChoice == "1" {
        fmt.Println("\nYour encoded text:")
        fmt.Println(Encode(userText))
    } else if userChoice == "2" {
        fmt.Println("\nYour decoded text:")
        fmt.Println(Decode(userText))
    }
}

/*
 * User should enter either 1 or 2.
 * By this code know the encode and decode.
*/
func encodeDecodeChoice() (string, error) {
    var userChoice string

    fmt.Println("What do you want to do? (1: encode, 2: decode)")

    _, error := fmt.Scanln(&userChoice)

    if error != nil || (userChoice != "1" && userChoice != "2") {
        return "", errors.Join(errors.New("❌ The prompt for the encoder/decoder is incorrect.\n"), error)
    }

    return userChoice, nil
}

/*
 * Expecting to get the text, till user enters 'enter' (new line).
*/
func textPrompt() (string, error) {
    fmt.Println("\nEnter your text: ")

    userInput, error := inputReader()

    if error != nil {
        return "", errors.Join(errors.New("The given text can not be parsed/scaned."), error)
    }

    return userInput, nil
}

func inputReader() (string, error) {
    reader := bufio.NewReader(os.Stdin)
    text, error := reader.ReadString('\n')

    return strings.TrimRight(text, "\n"), error
}

