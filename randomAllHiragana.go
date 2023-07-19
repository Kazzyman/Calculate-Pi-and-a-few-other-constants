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
	"aア, あ", "aア, あ",    "iイ, い",    "uウ, う", "eエ, え", "oオ, お",     "uウ, う", "eエ, え", "oオ, お",      

// 6 lines of the ka group: (inclusive)
//   ka         ki         ku        ke         ko * 1       /  ga         gi         gu        ge         go * 1          //  naked / ゛
	"kaカ, か", "kiチ, き", "kuク, く", "keケ, け", "koコ, こ",       "gaガ, が", "giヂ, ぎ", "guグ, ぐ", "geゲ, げ", "goゴ, ご",   //-* ???????????????????????????????????? *-//
//               ^----^--v                                                 ^^----^
	           "kiyaチャ, きゃ",  "kiyuチュ, きゅ",  "kiyoチョ, きょ",  // * 1      ^^----^-v                                                // ki, ya yu yo
	                                                                      "giyaヂャ, ぎゃ",  "giyuヂュ, ぎゅ",  "giyoヂョ, ぎょ",     // * 1 // ki゛, ya yu yo
//                       ^               ^               ^                          ^               ^               ^        

// 6 lines of the sa group: (inclusive)
//   sa         shi         su        se         so * 1      /   za         gi         zu         ze        zo * 1         //  naked / ゛ 
	"saサ, さ", "shiシ, し", "suス, す", "seセ, せ", "soソ, そ",      "zaザ, ざ", "giジ, じ", "zuズ, ず", "zeゼ, ぜ", "zoゾ, ぞ",  //-* One key "zu", has two values づ and ず *-//
//               ^----^---v                                                 ^^----^
	           "shiyaシャ, しゃ",  "shiyuシュ, しゅ",  "shiyoショ, しょ",  // * 1    ^^----^-v                                               // shi, ya yu yo
	                                                                       "giyaジャ, じゃ",   "giyuジュ, じゅ",   "giyoジョ, じょ",  // * 1 // shi゛, ya yu yo
//                        ^                ^                ^                        ^                ^                ^    

// 6 lines of the ta group: (inclusive)
//   ta         chi        tsu         te         to * 1     /   da         gi         zu         de         do * 1        //  naked / ゛
	"taタ, た", "chiチ, ち", "tsuツ, つ", "teテ, て", "toト, と",     "daダ, だ",  "giザ, ぢ", "zuヅ, づ", "deデ, で", "doド, ど",  //-* One key "zu", has two values づ and ず *-//
//               ^----^---v                                                 ^^----^
	           "chiyaチャ, ちゃ", "chiyuチュ, ちゅ", "chiyoチョ, ちょ",   // * 1     ^^----^-v                                               // chi, ya yu yo
	                                                                       "giyaヂャ, ぢゃ",  "giyuヂュ, ぢゅ",  "giyoヂョ, ぢょ",    // * 1 // chi゛, ya yu yo
//                        ^               ^               ^                          ^               ^               ^

// 5 lines of the na group: (inclusive) 'cause they are special -- they like to go naked 
//   na         ni         nu        ne         no * 1                                                                     //  naked, all of them (there are no ゛s on na's)
	"naナ, な", "niニ, に", "nuヌ, ぬ", "neネ, ね", "noノ, の",    
//               ^---^---v                                
               "niyaニャ, にゃ", "niyuニュ, にゅ", "niyoニョ, にょ",   "niyaニャ, にゃ", "niyuニュ, にゅ", "niyoニョ, にょ",                   // * 2 // ni, ya yu yo
//                       ^              ^              ^                ^               ^              ^ 

// 9 lines of the ha group: (inclusive) -- the fully-clothed group 
//  ha         hi         hu * 1     he        ho          he        ho * 2      /  ba         bi         bu         be         bo             //  naked / ゛
   "haハ, は", "hiヒ, ひ", "huフ, ふ", "heヘ, へ", "hoホ, ほ", "heヘ, へ", "hoホ, ほ",     "baバ, ば", "biビ, び", "buブ, ぶ", "beベ, べ", "boボ, ぼ",  //-* ????? *-// 
//              ^---^--v                                                             pa         pi         pu        pe         po            //  poo゜ marks (drag)
                                                                                    "paパ, ぱ", "piピ, ぴ", "puプ, ぷ", "peペ, ぺ", "poポ, ぽ", //-* ????? *-//
//              ^------v
	          "hiyaヒャ, ひゃ", "hiyuヒュ, ひゅ", "hiyoヒョ, ひょ",                                                                         // hi, ya yu yo
	                                                                                "biyaビャ, びゃ", "biyuビュ, びゅ", "biyoビョ, びょ",   // hi゛, ya yu yo (bi)
	                                                                                "piyaピャ, ぴゃ", "piyuビュ, ぴゅ", "piyoピョ, ぴょ",   // hi゜, ya yu yo (pi)
//                     ^               ^              ^                                       ^              ^              ^   

// 5 lines of the ma group: (inclusive) 'cause they too are special -- and also like to go naked
//  ma         mi        mu         me         mo * 1                                                                     //  naked, all of them (there are no ゛s on ma's)
   "maマ, ま", "miミ, み", "muム, む", "meメ, め", "moモ, も",     
//              ^--^-v
              "miyaミャ, みゃ", "miyuミュ, みゅ", "miyoミョ, みょ",     "miyaミャ, みゃ", "miyuミュ, みゅ", "miyoミョ, みょ",                  // * 2 // mi, ya yu yo 
//                      ^              ^              ^                  ^              ^              ^

// 3 lines of the ya, yu, yo set: (inclusive)
//  ya         yu         yo * 3
   "yaヤ, や", "yuユ, ゆ", "yoヨ, よ","yaヤ, や", "yuユ, ゆ", "yoヨ, よ", "yaヤ, や", "yuユ, ゆ", "yoヨ, よ",    // * 3 (also used as suffixes)
// 

// 5 lines of the ra/la group: (inclusive) actually spelled with the "r" but pronounced with the "L" sounds -- they too are special -- and also like to go naked
//  ra        ri         ru         re         ro * 1    
   "raラ, ら", "riリ, り", "ruル, る", "reレ, れ", "roロ, ろ", 
//             ^----^---v
              "riyaリャ, りゃ", "riyuリュ, りゅ", "riyoリョ, りょ",    "riyaリャ, りゃ", "riyuリュ, りゅ", "riyoリョ, りょ",    // * 2 // ri, ya yu yo 
//                      ^              ^              ^             ^               ^             ^ 

// 2 lines (inclusive) of the wa set, plus the nh consonant -- always fully naked AND never have suffixes of ya yu or yo 
//  wa wo nh * 2
   "waワ, わ", "woヲ, を",   "nhン, ん",       "waワ, わ", "woヲ, を",   "nhン, ん" }  // * 2 //  and the "}" must be on this line rather than on a line of its own 

   

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
