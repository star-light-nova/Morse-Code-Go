package morse

import (
	"strings"
	"github.com/StarLightNova/Morse-Code-Go/pkg/dictionary"
)

func Decode(morseText string) string {
    splitMorse := strings.Split(morseText, LETTER_SEPARATOR)
    last := len(splitMorse) - 1

    // The buffer reader adds `\n` to the end of the last string added after scan.
    splitMorse[last] = strings.Trim(splitMorse[last], "\n")

    result := ""

    for _, morseCharacter := range splitMorse {
        if morseCharacter == WORD_SEPARATOR {
            result += LETTER_SEPARATOR
        } else if morseCharacter != LETTER_SEPARATOR {
            result += string(dictionary.ReversedAll()[morseCharacter])
        }
    }

    return result
}
