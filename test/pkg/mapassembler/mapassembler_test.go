package mapassembler_test

import (
	"testing"
	"github.com/star-light-nova/Morse-Code-Go/pkg/mapassembler"
)

func TestAssembledMap(t *testing.T) {
    firstMap := map[rune]string {'a': "aa", 'b': "bb"}
    secondMap := map[rune]string {'c': "cc", 'd': "dd"}
    correctMap := map[rune]string {'a': "aa", 'b': "bb", 'c': "cc", 'd': "dd"}

    assembledMap := mapassembler.Assemble(firstMap, secondMap)

   for key := range correctMap {
        if correctMap[key] != assembledMap[key] {
            t.Fatal("ASSemble is not correct.")
        }
    }
}
