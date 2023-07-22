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

// 2 lines of the "vowels" (inclusive)
//  a * 2 :                i * 1       u         e         o * 2 :
	"ア, あ", "ア, あ",    "イ, い",    "ウ, う", "エ, え", "オ, お",     "ウ, う", "エ, え", "オ, お",      

// 6 lines of the ka group: (inclusive) キ　
//   ka         ki         ku        ke         ko * 1       /  ga         gi         gu        ge         go * 1          //  naked / ゛
	"カ, か", "ギ, き", "ク, く", "ケ, け", "コ, こ",       "ガ, が", "ギ, ぎ", "グ, ぐ", "ゲ, げ", "ゴ, ご",   //-* ???????????????????????????????????? *-//
//               ^----^--v                                                 ^^----^
	           "キャ, きゃ",  "キュ, きゅ",  "キョ, きょ",  // * 1      ^^----^-v                                                // ki, ya yu yo
	                                                                      "ギャ, ぎゃ",  "ギュ, ぎゅ",  "ギョ, ぎょ",     // * 1 // ki゛, ya yu yo
//                       ^               ^               ^                          ^               ^               ^        

// 6 lines of the sa group: (inclusive)
//   sa         shi         su        se         so * 1      /   za         gi         zu         ze        zo * 1         //  naked / ゛ 
	"サ, さ", "シ, し", "ス, す", "セ, せ", "ソ, そ",      "ザ, ざ", "ジ, じ", "ズ, ず", "ゼ, ぜ", "ゾ, ぞ",  //-* One key "zu", has two values づ and ず *-//
//               ^----^---v                                                 ^^----^
	           "シャ, しゃ",  "シュ, しゅ",  "ショ, しょ",  // * 1    ^^----^-v                                               // shi, ya yu yo
	                                                                       "ジャ, じゃ",   "ジュ, じゅ",   "ジョ, じょ",  // * 1 // shi゛, ya yu yo
//                        ^                ^                ^                        ^                ^                ^    

// 6 lines of the ta group: (inclusive)
//   ta         chi        tsu         te         to * 1     /   da         gi         zu         de         do * 1        //  naked / ゛
	"タ, た", "チ, ち", "ツ, つ", "テ, て", "ト, と",     "ダ, だ",  "ザ, ぢ", "ヅ, づ", "デ, で", "ド, ど",  //-* One key "zu", has two values づ and ず *-//
//               ^----^---v                                                 ^^----^
	           "チャ, ちゃ", "チュ, ちゅ", "チョ, ちょ",   // * 1     ^^----^-v                                               // chi, ya yu yo
	                                                                       "ヂャ, ぢゃ",  "ヂュ, ぢゅ",  "ヂョ, ぢょ",    // * 1 // chi゛, ya yu yo
//                        ^               ^               ^                          ^               ^               ^

// 5 lines of the na group: (inclusive) 'cause they are special -- they like to go naked 
//   na         ni         nu        ne         no * 1                                                                     //  naked, all of them (there are no ゛s on na's)
	"ナ, な", "ニ, に", "ヌ, ぬ", "ネ, ね", "ノ, の",    
//               ^---^---v                                
               "ニャ, にゃ", "ニュ, にゅ", "ニョ, にょ",   "ニャ, にゃ", "ニュ, にゅ", "ニョ, にょ",                   // * 2 // ni, ya yu yo
//                       ^              ^              ^                ^               ^              ^ 

// 9 lines of the ha group: (inclusive) -- the fully-clothed group 
//  ha         hi         hu * 1     he        ho          he        ho * 2      /  ba         bi         bu         be         bo             //  naked / ゛
   "ハ, は", "ヒ, ひ", "フ, ふ", "ヘ, へ", "ホ, ほ", "ヘ, へ", "ホ, ほ",     "バ, ば", "ビ, び", "ブ, ぶ", "ベ, べ", "ボ, ぼ",  //-* ????? *-// 
//              ^---^--v                                                             pa         pi         pu        pe         po            //  poo゜ marks (drag)
                                                                                    "パ, ぱ", "ピ, ぴ", "プ, ぷ", "ペ, ぺ", "ポ, ぽ", //-* ????? *-//
//              ^------v
	          "ヒャ, ひゃ", "ヒュ, ひゅ", "ヒョ, ひょ",                                                                         // hi, ya yu yo
	                                                                                "ビャ, びゃ", "ビュ, びゅ", "ビョ, びょ",   // hi゛, ya yu yo (bi)
	                                                                                "ピャ, ぴゃ", "ビュ, ぴゅ", "ピョ, ぴょ",   // hi゜,  yu yo (pi)
//                     ^               ^              ^                                       ^              ^              ^   

// 5 lines of the ma group: (inclusive) 'cause they too are special -- and also like to go naked
//  ma         mi        mu         me         mo * 1                                                                     //  naked, all of them (there are no ゛s on ma's)
   "マ, ま", "ミ, み", "ム, む", "メ, め", "モ, も",     
//              ^--^-v
              "ミャ, みゃ", "ミュ, みゅ", "ミョ, みょ",     "ミャ, みゃ", "ミュ, みゅ", "ミョ, みょ",                  // * 2 // mi, ya yu yo 
//                      ^              ^              ^                  ^              ^              ^

// 3 lines of the ya, yu, yo set: (inclusive)
//  ya         yu         yo * 3
   "ヤ, や", "ユ, ゆ", "ヨ, よ","ヤ, や", "ユ, ゆ", "ヨ, よ", "ヤ, や", "ユ, ゆ", "ヨ, よ",    // * 3 (also used as suffixes)
// 

// 5 lines of the ra/la group: (inclusive) actually spelled with the "r" but pronounced with the "L" sounds -- they too are special -- and also like to go naked
//  ra        ri         ru         re         ro * 1    
   "ラ, ら", "リ, り", "ル, る", "レ, れ", "ロ, ろ", 
//             ^----^---v
              "リャ, りゃ", "リュ, りゅ", "リョ, りょ",    "ャ, りゃ", "リュ, りゅ", "リョ, りょ",    // * 2 // ri, ya yu yo 
//                      ^              ^              ^             ^               ^             ^ 

// 2 lines (inclusive) of the wa set, plus the nh consonant -- always fully naked AND never have suffixes of ya yu or yo 
//  wa wo nh * 2
   "ワ, わ", "ヲ, を",   "ン, ん",       "ワ, わ", "ヲ, を",   "ン, ん" }  // * 2 //  and the "}" must be on this line rather than on a line of its own 

   

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
/*
ChatGPT 3.5 proofread all of the katakana, hiragana pairs used in this code; and my Japanese wife verified them also. 
*/