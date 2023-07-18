package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)
// 51 15 25 21 = 112 if we include the suffix combinations 
var in string

func main() {
	rand.Seed(time.Now().UnixNano())

	keyValuePairs := []string{

	//  a * 2             i*1        u e o * 2
	"a, あ", "a, あ",    "i, い",    "u, う", "e, え", "o, お", "u, う", "e, え", "o, お",      
	
	//  ka ki ku ke ko * 1                              ga       gi        gu       ge       go * 1
	"ka, か", "ki, き", "ku, く", "ke, け", "ko, こ",     "ga, が", "gi, ぎ", "gu, ぐ", "ge, げ", "go, ご",
	          "kiya, きゃ", "kiyu, きゅ", "kiyo, きょ",             "giya, ぎゃ", "giyu, ぎゅ", "giyo, ぎょ", 

	//  sa shi su se so * 1                             za       gi       zu        ze       zo * 1
	"sa, さ", "shi, し", "su, す", "se, せ", "so, そ",    "za, ざ", "gi, じ", "zu, ず", "ze, ぜ", "zo, ぞ", // one key "zu", has two values づ and ず 
	         "shiya, しゃ", "shiyu, しゅ", "shiyo, しょ",            "giya, じゃ", "giyu, じゅ", "giyo, じょ", // can i do this, yes 

	//  ta chi tsu se so * 1                            da        gi        zu       de       do * 1
	"ta, た", "chi, ち", "tsu, つ", "te, て", "to, と",   "da, だ",  "gi, ぢ", "zu, づ", "de, で", "do, ど", // one key "zu", has two values づ and ず
	         "chiya, ちゃ", "chiyu, ちゅ", "chiyo, ちょ",            "giya, ぢゃ", "giyu, ぢゅ", "giyo, ぢょ", 

	//  na ni nu ne no * 1                              niya        niyu         niyo 
	"na, な", "ni, に", "nu, ぬ", "ne, ね", "no, の",    "niya, にゃ", "niyu, にゅ", "niyo, にょ", 

	//  ha hi hu he ho * 1                          * 2                    ba       bi        bu       be        bo          pa        pi       pu        pe       po
	"ha, は", "hi, ひ", "hu, ふ", "he, へ", "ho, ほ", "he, へ", "ho, ほ",    "ba, ば", "bi, び", "bu, ぶ", "be, べ", "bo, ぼ",    "pa, ぱ", "pi, ぴ", "pu, ぷ", "pe, ぺ", "po, ぽ", 
   // then just the ひs below:
	"hiya, ひゃ", "hiyu, ひゅ", "hiyo, ひょ",     "biya, びゃ", "biyu, びゅ", "biyo, びょ",     "piya, ぴゃ", "piyu, ぴゅ", "piyo, ぴょ", // it did not used to like this line ????



	"ma, ま", "mi, み", "mu, む", "me, め",
	//"mo, も", "ya, や", "yu, ゆ", "yo, よ", "ra, ら", "ri, り", "ru, る", "re, れ", "ro, ろ", "wa, わ", "wo, を", "nh, ん" 
	}

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
		fmt.Printf("%s ? ", key)
		fmt.Scan(&in)

		if in == value {
			//fmt.Println("You guessed right! You typed", in, "... the key was", key, "and the value was indeed", value)
			fmt.Printf("\nRight! \n")
		} else {
			//fmt.Println("You typed", in, "which is not", value, "so you guessed wrongly")
			fmt.Printf("\n Oops! it was %s", value)

			// hints: 

			if value == "ほ" {
				fmt.Printf(", hint: ring--->, \"-\" \n")
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
				fmt.Printf(", hint: ring<---, 1 char \n")
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
				fmt.Printf(", hint: _ char \n")
			}


			if value == "や" {
				fmt.Printf(", hint: 7 char \n")
			}

			if value == "ゆ" {
				fmt.Printf(", hint: index->, 8 \n")
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
