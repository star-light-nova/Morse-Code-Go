package morse_test

import (
	"fmt"
	"testing"
	"github.com/StarLightNova/Morse-Code-Go/pkg/morse"
)

func TestEncode(t *testing.T) {
    result := map[string]string{}

    for text := range TRANSLATE_DICTIONARY() {
        result[text] = morse.Encode(text)
    }

    for text := range TRANSLATE_DICTIONARY() {
        if TRANSLATE_DICTIONARY()[text] != result[text] {
            tableText := fmt.Sprintf("\n📑ORIGINAL TEXT: %s\n🧪CORRECT TRANSLATION: %s\n❌RETURNED TRANSLATION: %s", text, TRANSLATE_DICTIONARY()[text], result[text])
            t.Fatalf("The Encode function from text ➡️  morse has returned incorrect result(-s).%s", tableText)
        }
    }
}
