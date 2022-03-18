package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var CHARACTERS = map[string]string{
	"a":   "あ",
	"e":   "え",
	"u":   "う",
	"i":   "い",
	"o":   "お",
	"ka":  "か",
	"ke":  "け",
	"ku":  "く",
	"ki":  "き",
	"ko":  "こ",
	"sa":  "さ",
	"se":  "せ",
	"su":  "す",
	"shi": "し",
	"so":  "そ",
	"ta":  "た",
	"te":  "て",
	"tsu": "つ",
	"chi": "ち",
	"to":  "と",
	"na":  "な",
	"ne":  "ね",
	"nu":  "ぬ",
	"ni":  "に",
	"no":  "の",
	"ha":  "は",
	"he":  "へ",
	"fu":  "ふ",
	"hi":  "ひ",
	"ho":  "ほ",
	"ma":  "ま",
	"me":  "め",
	"mu":  "む",
	"mi":  "み",
	"mo":  "も",
	"ra":  "ら",
	"re":  "れ",
	"ru":  "る",
	"ri":  "り",
	"ro":  "ろ",
	"ya":  "や",
	"yu":  "ゆ",
	"yo":  "よ",
	"n":   "ん",
	"wa":  "わ",
	"wo":  "を",
	"ga":  "が",
	"ge":  "げ",
	"gu":  "ぐ",
	"gi":  "ぎ",
	"go":  "ご",
	"za":  "ざ",
	"ze":  "ぜ",
	"zu":  "ず",
	"ji":  "じ",
	"zo":  "ぞ",
	"da":  "だ",
	"de":  "で",
	"do":  "ど",
	"ba":  "ば",
	"be":  "べ",
	"bu":  "ぶ",
	"bi":  "び",
	"bo":  "ぼ",
	"pa":  "ぱ",
	"pe":  "ぺ",
	"pu":  "ぷ",
	"pi":  "ぴ",
	"po":  "ぽ",
}

const (
	ModeMainMenu = iota
	ModePractice
)

var RUNNING bool = true
var CURRENT int32 = ModeMainMenu

var PREVIOUS_WORDS [5]string = [5]string{
	"one", "two", "three", "four", "five",
}

func GetLine(msg string) string {

	fmt.Print(msg)

	var line string
	fmt.Scanln(&line)

	return line
}

func PushBackPreviousWord(value string) {
	for k := range PREVIOUS_WORDS {
		if k < len(PREVIOUS_WORDS)-1 {
			PREVIOUS_WORDS[k] = PREVIOUS_WORDS[k+1]
		} else {
			PREVIOUS_WORDS[4] = value
		}
	}
}

func PickKana() string {
	for {
		num := rand.Intn(len(CHARACTERS))
		index := 0
		valid := true

		for k := range CHARACTERS {
			if index == num {
				for _, v := range PREVIOUS_WORDS {
					if v == CHARACTERS[k] {
						valid = false
						break
					}
				}

				if !valid {
					break
				}

				PushBackPreviousWord(CHARACTERS[k])
				return k
			} else {
				index += 1
			}
		}
	}
}

func main() {

	fmt.Println("------------ GO HIRAGANA PRACTICE ------------")
	fmt.Println("- Type 'Practice' to enter the practice mode.")
	fmt.Println("- Type 'Words' to enter the word mode.")

	for RUNNING {

		var action string = ""

		if CURRENT == ModeMainMenu {
			action = GetLine(">> ")

			if strings.ToLower(action) == "practice" {
				CURRENT = ModePractice
				continue
			} else if strings.ToLower(action) == "words" {
				fmt.Println("Not implemented yet.")
				continue
			}
		} else if CURRENT == ModePractice {

			kana := PickKana()
			fmt.Printf("Current : %v\n", CHARACTERS[kana])

			action = GetLine(">> ")

			if strings.ToLower(action) == kana {
				fmt.Println("Correct!")
				continue
			} else if strings.ToLower(action) != "quit" {
				fmt.Printf("False, it was '%v'.\n", kana)
				continue
			}
		}

		if strings.ToLower(action) == "quit" {
			RUNNING = false
			continue
		} else {
			fmt.Println("Unknown command.")
			continue
		}
	}
}
