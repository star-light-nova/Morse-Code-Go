package dictionary

import (
	"github.com/star-light-nova/morse-code-go/pkg/mapassembler"
)

func All() map[rune]string {
    alphabet := Alphabet()
    digits := Digits()
    p35s := P35S()

    all := mapassembler.Assemble(alphabet, digits, p35s)

    return all
}
