package morse_test

import (
	"testing"
	"github.com/StarLightNova/Morse-Code-Go/pkg/morse"
)

func TestTranslate(t *testing.T) {
    originalText := "SALEM, ALEM"

    morseText := morse.Translate(originalText)

    if morseText != "... .- .-.. . -- --..-- / .- .-.. . --" {
        t.Fatal("The function \"morse.Translate\" does not return the right value.")
    }
}
