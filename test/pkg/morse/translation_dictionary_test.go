package morse_test

import (
    "github.com/star-light-nova/Morse-Code-Go/pkg/reversemap"
)

func TRANSLATE_DICTIONARY() map[string]string{
    return map[string]string{
        "SALEM, ALEM": "... .- .-.. . -- --..-- / .- .-.. . --",
        "TO BE OR NOT TOBE,. . . . . . .": "- --- / -... . / --- .-. / -. --- - / - --- -... . --..-- .-.-.- / .-.-.- / .-.-.- / .-.-.- / .-.-.- / .-.-.- / .-.-.-",
        "@AUDIAM (ULLAMCORPER) :PRI: = AD,, EU' NOVUM TANTAS CUM, EI+US DELENITI -PLACERAT AN QUI!": ".--.-. .- ..- -.. .. .- -- / -.--. ..- .-.. .-.. .- -- -.-. --- .-. .--. . .-. -.--.- / ---... .--. .-. .. ---... / -...- / .- -.. --..-- --..-- / . ..- .----. / -. --- ...- ..- -- / - .- -. - .- ... / -.-. ..- -- --..-- / . .. .-.-. ..- ... / -.. . .-.. . -. .. - .. / -....- .--. .-.. .- -.-. . .-. .- - / .- -. / --.- ..- .. -.-.--",
        "5-380940655A 0": "..... -....- ...-- ---.. ----- ----. ....- ----- -.... ..... ..... .- / -----",
        "2B, 9S, A2": "..--- -... --..-- / ----. ... --..-- / .- ..---",
    }
}

func REVERSE_TRANSLATION_DICTIONARY() map[string]string {
    return reversemap.Reverse(TRANSLATE_DICTIONARY())
}
