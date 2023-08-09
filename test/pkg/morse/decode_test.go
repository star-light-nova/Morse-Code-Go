package morse_test

import (
	"fmt"
	"testing"
	"github.com/star-light-nova/Morse-Code-Go/pkg/morse"
)

func TestDecode(t *testing.T) {
    result := map[string]string{}

    for morseText := range REVERSE_TRANSLATION_DICTIONARY() {
        result[morseText] = morse.Decode(morseText)
    }

    for morseText := range REVERSE_TRANSLATION_DICTIONARY() {
        if REVERSE_TRANSLATION_DICTIONARY()[morseText] != result[morseText] {
            tableText := fmt.Sprintf("\n📑ORIGINAL TEXT: %s\n🧪CORRECT TRANSLATION: %s\n❌RETURNED TRANSLATION: %s", morseText, REVERSE_TRANSLATION_DICTIONARY()[morseText], result[morseText])
            t.Fatalf("The Decode function from morse ➡️  text has returned incorrect result(-s).%s", tableText)
        }
    }
}
