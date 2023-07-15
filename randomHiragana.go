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

	keyValuePairs := []string{"a, あ", "i, い", "u, う", "e, え", "o, お", "ka, か", "ki, き", "ku, く", "ke, け", "ko, こ",
	 "sa, さ", "shi, し", "su, す", "se, せ", "so, そ", "ta, た", "chi, ち", "tsu, つ", "te, て", "to, と", "na, な", "ni, に",
	  "nu, ぬ", "ne, ね", "no, の", "ha, は", "hi, ひ", "hu, ふ", "he, へ", "ho, ほ", "ma, ま", "mi, み", "mu, む", "me, め",
	   "mo, も", "ya, や", "yu, ゆ", "yo, よ", "ra, ら", "ri, り", "ru, る", "re, れ", "ro, ろ", "wa, わ", "wo, を", "nh, ん"}

	for {
		randIndex := rand.Intn(len(keyValuePairs))
		keyValuePairSelected := keyValuePairs[randIndex]

		parts := strings.Split(keyValuePairSelected, ", ")
		if len(parts) != 2 {
			fmt.Println("Invalid key-value pair:", keyValuePairSelected)
			continue
		}

		key := parts[0]
		value := parts[1]

		//fmt.Println("Type the key for this sound:", key)
		fmt.Printf("%s?", key)
		fmt.Scan(&in)

		if in == value {
			//fmt.Println("You guessed right! You typed", in, "... the key was", key, "and the value was indeed", value)
			fmt.Printf("\nRight! \n")
		} else {
			//fmt.Println("You typed", in, "which is not", value, "so you guessed wrongly")
			fmt.Printf("%s, make a note and try again", value)

			// hints: 

			if value == "あ" {
				fmt.Printf(", hint: 3 char \n\n")
			}
			if value == "い" {
				fmt.Printf(", hint: E char \n\n")
			}
			if value == "う" {
				fmt.Printf(", hint: 4 char \n\n")
			}
			if value == "え" {
				fmt.Printf(", hint: 5 char \n\n")
			}
			if value == "お" {
				fmt.Printf(", hint: 6 char \n\n")
			}


			if value == "か" {
				fmt.Printf(", hint: T char \n\n")
			}
			if value == "き" {
				fmt.Printf(", hint: G char \n\n")
			}
			if value == "く" {
				fmt.Printf(", hint: H char \n\n")
			}
			if value == "け" {
				fmt.Printf(", hint: :* char \n\n")
			}
			if value == "こ" {
				fmt.Printf(", hint: B char \n\n")
			}


			if value == "さ" {
				fmt.Printf(", hint: X char \n\n")
			}
			if value == "し" {
				fmt.Printf(", hint: D char \n\n")
			}
			if value == "す" {
				fmt.Printf(", hint: R char \n\n")
			}
			if value == "せ" {
				fmt.Printf(", hint: P char \n\n")
			}
			if value == "そ" {
				fmt.Printf(", hint: C char \n\n")
			}


			if value == "た" {
				fmt.Printf(", hint: Q char \n\n")
			}
			if value == "ち" {
				fmt.Printf(", hint: A char \n\n")
			}
			if value == "つ" {
				fmt.Printf(", hint: Z char \n\n")
			}
			if value == "て" {
				fmt.Printf(", hint: W char \n\n")
			}
			if value == "と" {
				fmt.Printf(", hint: S char \n\n")
			}


			if value == "な" {
				fmt.Printf(", hint: U char \n\n")
			}
			if value == "に" {
				fmt.Printf(", hint: I char \n\n")
			}
			if value == "ぬ" {
				fmt.Printf(", hint: 1 char \n\n")
			}
			if value == "ね" {
				fmt.Printf(", hint: ,< char \n\n")
			}
			if value == "の" {
				fmt.Printf(", hint: k char \n\n")
			}


			if value == "ま" {
				fmt.Printf(", hint: J char \n\n")
			}
			if value == "ひ" {
				fmt.Printf(", hint: V char \n\n")
			}
			if value == "む" {
				fmt.Printf(", hint: ]} char \n\n")
			}
			if value == "め" {
				fmt.Printf(", hint: /? char \n\n")
			}
			if value == "も" {
				fmt.Printf(", hint: M char \n\n")
			}


			if value == "ら" {
				fmt.Printf(", hint: o char \n\n")
			}
			if value == "り" {
				fmt.Printf(", hint: L char \n\n")
			}
			if value == "る" {
				fmt.Printf(", hint: .> char \n\n")
			}
			if value == "れ" {
				fmt.Printf(", hint: ;+ char \n\n")
			}
			if value == "ろ" {
				fmt.Printf(", hint: _ char \n\n")
			}


			if value == "や" {
				fmt.Printf(", hint: 7 char \n\n")
			}
			if value == "ゆ" {
				fmt.Printf(", hint: 8 char \n\n")
			}
			if value == "よ" {
				fmt.Printf(", hint: 9 char \n\n")
			}


			if value == "わ" {
				fmt.Printf(", hint: 0 char \n\n")
			}
			if value == "を" {
				fmt.Printf(", hint: 0^ char \n\n")
			}

			
			if value == "ん" {
				fmt.Printf(", hint: Y char \n\n")
			}
		}
		fmt.Println()
	}
}
