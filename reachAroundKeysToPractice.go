//reachAroundKeysToPractice.go

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var in string

func main() {
	rand.Seed(time.Now().UnixNano())

	// only the most difficult to type chars are to be practiced in this version, and the hardest ones are duplicated in the slice for extra practice
	keyValuePairs := []string{ "u, う", "u, う",      "o, お", "o, お",      "nu, ぬ", "nu, ぬ",      "he, へ", "ho, ほ", "he, へ", "ho, ほ",      "hu, ふ", "hu, ふ",
		 "ya, や", "yu, ゆ", "yo, よ", "ya, や", "yu, ゆ", "yo, よ",     "ro, ろ",    "wa, わ",   "wo, を",    "nh, ん",    "mu, む",    "a, あ", "a, あ",    "e, え" }

	for {
		randIndex := rand.Intn(len(keyValuePairs))  // returns length of the slice to rand.Intn ... randIndex is set to 14, or 18, or some other random value
		//fmt.Println("this is what randIndex was set to: ", randIndex)  // ^^^^
		keyValuePairSelected := keyValuePairs[randIndex]  // gets one of the pairs selected at random 

		parts := strings.Split(keyValuePairSelected, ", ")  // parts is created as a new slice with two string values (key/value pair gets split on ,)
		//fmt.Println("parts is: ", parts)
		if len(parts) != 2 {
			fmt.Println("Invalid key-value pair:", keyValuePairSelected)  // just checking to ensure we have exactly two strings, as expected 
			continue
		}

		key := parts[0]
		value := parts[1]  // loading our randomly-selected pair 

		//fmt.Println("Type the key for this sound:", key)
		fmt.Printf("%s ? ", key)  // prompt with the sound

		//fmt.Printf(" ")s to show hints:


			if value == "は" {
				fmt.Printf(", hint: o char \n")
			}
			if value == "ひ" {
				fmt.Printf(", hint: L char \n")
			}
			if value == "ふ" {
				fmt.Printf(", hint: ring<, 2 char \n")
			}
			if value == "へ" {
				fmt.Printf(", hint: ;+ char \n")
			}


			if value == "ほ" {
				fmt.Printf(", hint: ring--->, \"-\" \n")
			}
			if value == "あ" {
				fmt.Printf(", hint: middle<-, 3 \n")
			}
			if value == "い" {
				fmt.Printf(", hint: E char \n")
			}
			if value == "う" {
				fmt.Printf(", hint: middle>, 4 \n")
			}
			if value == "え" {
				fmt.Printf(", hint: L-index> long-reach, 5 \n")
			}
			if value == "お" {
				fmt.Printf(", hint: L-index> long-reach-one-over, 6 \n")
			}
			if value == "か" {
				fmt.Printf(", hint: T char \n")
			}
			if value == "き" {
				fmt.Printf(", hint: G char \n")
			}
			if value == "く" {
				fmt.Printf(", hint: H char \n")
			}
			if value == "け" {
				fmt.Printf(", hint: :* char \n")
			}
			if value == "こ" {
				fmt.Printf(", hint: B char \n")
			}
			if value == "さ" {
				fmt.Printf(", hint: X char \n")
			}
			if value == "し" {
				fmt.Printf(", hint: D char \n")
			}
			if value == "す" {
				fmt.Printf(", hint: R char \n")
			}
			if value == "せ" {
				fmt.Printf(", hint: P char \n")
			}
			if value == "そ" {
				fmt.Printf(", hint: C char \n")
			}
			if value == "た" {
				fmt.Printf(", hint: Q char \n")
			}
			if value == "ち" {
				fmt.Printf(", hint: A char \n")
			}
			if value == "つ" {
				fmt.Printf(", hint: Z char \n")
			}
			if value == "て" {
				fmt.Printf(", hint: W char \n")
			}
			if value == "と" {
				fmt.Printf(", hint: S char \n")
			}
			if value == "な" {
				fmt.Printf(", hint: U char \n")
			}
			if value == "に" {
				fmt.Printf(", hint: I char \n")
			}
			if value == "ぬ" {
				fmt.Printf(", hint: ring<--- , 1  \n")
			}
			if value == "ね" {
				fmt.Printf(", hint: ,< char \n")
			}
			if value == "の" {
				fmt.Printf(", hint: k char \n")
			}
			if value == "ま" {
				fmt.Printf(", hint: J char \n")
			}
			if value == "み" {
				fmt.Printf(", hint: V char \n")
			}
			if value == "む" {
				fmt.Printf(", hint: one-over, ]}  \n")
			}
			if value == "め" {
				fmt.Printf(", hint: pinky-> slide-down, /? \n")
			}
			if value == "も" {
				fmt.Printf(", hint: M char \n")
			}
			if value == "ら" {
				fmt.Printf(", hint: o char \n")
			}
			if value == "り" {
				fmt.Printf(", hint: L char \n")
			}
			if value == "る" {
				fmt.Printf(", hint: .> char \n")
			}
			if value == "れ" {
				fmt.Printf(", hint: ;+ char \n")
			}
			if value == "ろ" {
				fmt.Printf(", hint: very-long-slide-down, _ \n")
			}
			if value == "や" {
				fmt.Printf(", hint: index<-, 7 \n")
			}
			if value == "ゆ" {
				fmt.Printf(", hint: index->, 8 \n")
			}
			if value == "よ" {
				fmt.Printf(", hint: middle>, 9 \n")
			}
			if value == "わ" {
				fmt.Printf(", hint: ring>, 0 \n")
			}
			if value == "を" {
				fmt.Printf(", hint: ring>, ^0^ \n")
			}
			if value == "ん" {
				fmt.Printf(", hint: Y char \n")
			}
		fmt.Scan(&in)  // user enters a hiragana char

		if in == value {
			//fmt.Println("You guessed right! You typed", in, "... the key was", key, "and the value was indeed", value)
			fmt.Printf("\nRight! \n")
		} else {
			//fmt.Println("You typed", in, "which is not", value, "so you guessed wrongly")
			fmt.Printf("\n Oops! it was %s, make a note and try again", value)

			// hints: (all naked hiragana are included in this and all versions)

			if value == "ほ" {
				fmt.Printf(", hint: ring--->, \"-\" \n")
			}
			if value == "ふ" {
				fmt.Printf(", hint: ring<-, 2 \n")
			}
			if value == "へ" {
				fmt.Printf(", hint: pinky--->, \"^\" \n")
			}

			if value == "あ" {
				fmt.Printf(", hint: middle<, 3 char \n")
			}
			if value == "い" {
				fmt.Printf(", hint: E char \n")
			}
			if value == "う" {
				fmt.Printf(", hint: middle->, 4, u \n")
			}
			if value == "え" {
				fmt.Printf(", hint: 5 char \n")
			}
			if value == "お" {
				fmt.Printf(", hint: index--->, 6 char \n")
			}


			if value == "か" {
				fmt.Printf(", hint: T char \n")
			}
			if value == "き" {
				fmt.Printf(", hint: G char \n")
			}
			if value == "く" {
				fmt.Printf(", hint: H char \n")
			}
			if value == "け" {
				fmt.Printf(", hint: :* char \n")
			}
			if value == "こ" {
				fmt.Printf(", hint: B char \n")
			}


			if value == "さ" {
				fmt.Printf(", hint: X char \n")
			}
			if value == "し" {
				fmt.Printf(", hint: D char \n")
			}
			if value == "す" {
				fmt.Printf(", hint: R char \n")
			}
			if value == "せ" {
				fmt.Printf(", hint: P char \n")
			}
			if value == "そ" {
				fmt.Printf(", hint: C char \n")
			}


			if value == "た" {
				fmt.Printf(", hint: Q char \n")
			}
			if value == "ち" {
				fmt.Printf(", hint: A char \n")
			}
			if value == "つ" {
				fmt.Printf(", hint: Z char \n")
			}
			if value == "て" {
				fmt.Printf(", hint: W char \n")
			}
			if value == "と" {
				fmt.Printf(", hint: S char \n")
			}


			if value == "な" {
				fmt.Printf(", hint: U char \n")
			}
			if value == "に" {
				fmt.Printf(", hint: I char \n")
			}
			if value == "ぬ" {
				fmt.Printf(", hint: 1 char \n")
			}
			if value == "ね" {
				fmt.Printf(", hint: ,< char \n")
			}
			if value == "の" {
				fmt.Printf(", hint: k char \n")
			}


			if value == "ま" {
				fmt.Printf(", hint: J char \n")
			}
			if value == "み" {
				fmt.Printf(", hint: V char \n")
			}
			if value == "む" {
				fmt.Printf(", hint: ]} char \n")
			}
			if value == "め" {
				fmt.Printf(", hint: /? char \n")
			}
			if value == "も" {
				fmt.Printf(", hint: M char \n")
			}


			if value == "ら" {
				fmt.Printf(", hint: o char \n")
			}
			if value == "り" {
				fmt.Printf(", hint: L char \n")
			}
			if value == "る" {
				fmt.Printf(", hint: .> char \n")
			}
			if value == "れ" {
				fmt.Printf(", hint: ;+ char \n")
			}
			if value == "ろ" {
				fmt.Printf(", hint: _ char \n")
			}


			if value == "や" {
				fmt.Printf(", hint: 7 char \n")
			}
			if value == "ゆ" {
				fmt.Printf(", hint: 8 char \n")
			}
			if value == "よ" {
				fmt.Printf(", hint: 9 char \n")
			}


			if value == "わ" {
				fmt.Printf(", hint: 0 char \n")
			}
			if value == "を" {
				fmt.Printf(", hint: 0^ char \n")
			}

			
			if value == "ん" {
				fmt.Printf(", hint: Y char \n")
			}
		}
		fmt.Println()
	}
}
