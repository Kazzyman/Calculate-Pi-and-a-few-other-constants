package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {

	for {
		mainMenu()
	}

} // end func main



func mainMenu() {
	for {
	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
			fmt.Printf("  Main Menu: \n\n")
			fmt.Printf("  Enter 1 to practice recognizing Katakana prompts \n")
			fmt.Printf("  Enter 2 to practice recognizing Romaji-Katakana prompts \n\n\n")

			fmt.Scan(&menu)

		if menu == 1 {
				fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")

			fmt.Println("\n\n\n\n")
			fmt.Println("         'practicing recognizing Katakana' ")

			fmt.Println("      Using Hiragana input mode on your system, Type the Hiragana that corresponds to the Katakana prompt")
			fmt.Println("      Enter 'menu' to return to the the main menu (-change input mode to US or Alpha-numeric to enter 'menu' ) ")
			fmt.Println("      Enter '?' for context-sensitive help ) \n\n")

			kata()
		}
		if menu == 2 {
				fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")

			fmt.Println("\n\n\n\n")
			fmt.Println("         'practicing converting Romaji-Katakana prompts to Hiragana' ")
			fmt.Println("      Using Hiragana input mode on your system, Type the Hiragana that corresponds to the Romaji-Katakana prompt")
			fmt.Println("      Enter 'menu' to return to the the main menu (-change input mode to US or Alpha-numeric to enter 'menu' ) ")
			fmt.Println("      Enter '?' for context-sensitive help ) \n\n")

			hira()
		}
	}
}


func kata() {

	rand.Seed(time.Now().UnixNano())

/*
// sa group: 
//  sa       shi      su       se        so * 1   ||   za       ji        zu       ze       zo * 1         //  naked || ゛marks
	"サ, さ", "シ, し", "ス, す", "セ, せ", "ソ, そ",      "ザ, ざ", "ジ, じ", "ズ, ず", "ゼ, ぜ", "ゾ, ぞ",      //-* One key "zu", has two values づ and ず *-//
//           ^----^-v                                           ^^---^
	          "シャ, しゃ",  "シュ, しゅ",  "ショ, しょ",  // * 1    ^^---^-v                                          // shi,  ya yu yo (sha, shu, sho)
	                                                             "ジャ, じゃ",   "ジュ, じゅ",   "ジョ, じょ",  // * 1 // shi゛, ya yu yo (ja, ju, jo)
//                  ^            ^             ^                       ^             ^              ^    

// ta group: 　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　vv ji never gi
//  ta       chi      tsu      te        to * 1   ||   da        ji       zu       de       do * 1        //  naked || ゛marks
	"タ, た", "チ, ち", "ツ, つ", "テ, て", "ト, と",      "ダ, だ",  "ヂ, ぢ", "ヅ, づ", "デ, で", "ド, ど",    //-* One key "zu", has two values づ and ず *-//
//            ^--^--v                                            ^^--^
             "チャ, ちゃ", "チュ, ちゅ", "チョ, ちょ",   // * 1      ^^--^-v                                           // chi,  ya yu yo (cha, chu, cho)
	                                                              "ヂャ, ぢゃ",  "ヂュ, ぢゅ",  "ヂョ, ぢょ",    // * 1 // chi゛, ya yu yo (ja, ju, jo)
//                  ^           ^            ^                          ^            ^             ^

// na group:  'cause they are na-ty special -- they like to go naked 
//  na       ni       nu        ne       no * 1                                                       //  naked, all of them (there are no ゛s on na's)
	"ナ, な", "ニ, に", "ヌ, ぬ", "ネ, ね", "ノ, の",    
//            ^---^-v                                
             "ニャ, にゃ", "ニュ, にゅ", "ニョ, にょ",   "ニャ, にゃ", "ニュ, にゅ", "ニョ, にょ",                   // * 2 // ni, ya yu yo (nya, nyu, nyo)
//                  ^           ^            ^              ^           ^            ^ 

// ha group: -- the fully-clothed group (゛marks, and poo゜marks too)
//  ha       hi       hu       he        ho   * 1   ||    ba       bi       bu        be       bo          //  naked || ゛marks
   "ハ, は", "ヒ, ひ", "フ, ふ", "ヘ, へ", "ホ, ほ",         "バ, ば", "ビ, び", "ブ, ぶ", "ベ, べ", "ボ, ぼ",    //  ゛marks 
//            ^--^--v                                      pa        pi       pu       pe        po        //  poo゜ marks (drag)
                                                           "パ, ぱ", "ピ, ぴ", "プ, ぷ", "ペ, ぺ", "ポ, ぽ",  //   ^ ^  ^ ^ ^  ^ ^ ^
//            ^-----v                                                ^^--^
	          "ヒャ, ひゃ", "ヒュ, ひゅ", "ヒョ, ひょ",               //  ^^--^--v           v            v       // hi , ya yu yo (hi) hya, hyu, hyo
	                                                                  "ビャ, びゃ", "ビュ, びゅ", "ビョ, びょ",   // hi゛, ya yu yo (bi) bya, byu, byo
	                                                                  "ピャ, ぴゃ", "ピュ, ぴゅ", "ピョ, ぴょ",   // hi゜,  ya yu yo (pi) pya, pyu, pyo
//                  ^           ^            ^                                            ^           ^            ^   

// ma group:  'cause they too are special -- they too also like to go naked
//  ma       mi       mu        me       mo * 1                   //  naked, all of them (there are no ゛s on mama's)
   "マ, ま", "ミ, み", "ム, む", "メ, め", "モ, も",     
//            ^--^-v
             "ミャ, みゃ", "ミュ, みゅ", "ミョ, みょ"                // * 1 // mi , ya yu yo (mya, myu, myo)
//                  ^           ^            ^          

}  // go.lang requires that the "}" must be on this line rather than on a line of its own 
*/


doingContextSH = false
ContextSHskip = false 
// meat of the Katakana practice algorithm: ---------------------------------------------------------------------------
	for {
		// if we got here not via an interim handling of ? Context Sensitive Help
		//  v v v v v          v v
		if doingContextSH == false && ContextSHskip == false {

		randIndex := rand.Intn(len(fileOfCards))
		cardSelected := fileOfCards[randIndex]

		parts = strings.Split(cardSelected, ", ")

			if len(parts) != 2 {
				fmt.Println("Invalid key-value pair:", cardSelected)
				continue
				// or throw an error of some kind ?????????????????????????????????
			}

			fmt.Printf("%s", KeyK)
			fmt.Printf("%s", string(colorCyan))
			fmt.Printf(" ? ")
			fmt.Printf("%s", string(colorReset))

			fmt.Scan(&in)

			if in == "menu" {
				return 
				// returns to main menu
			}

			if in == "?" {
				doingContextSH = true
				// will do a kataMeatB() below in the parent else if 
			} else {
				doingContextSH = false 
				ContextSHskip = false
				kataMeatB() // probably no need to capture the return val here because ??????????????????????????
			}
		
		// else if -- if this is not our first rodeo 
		} else if ContextSHskip == false {
				giveAhint(value)
				fmt.Printf("%s", KeyK)
				fmt.Printf("%s", string(colorCyan))
				fmt.Printf(" ? ")
				fmt.Printf("%s", string(colorReset))
				doingContextSH = false
				ContextSHskip = true
				//continue
				fmt.Scan(&in)

				redoMenuKMB := kataMeatB() 
				if redoMenuKMB == true {
					return 
					// returns to main menu
				}

				// restart once in a func that does what the above if does and then sets ContextSHskip = false so that a future for iteration will run as if no ? had occured
				//             v v v v v v 
				redoMenuRS := restartKata() 
				if redoMenuRS == true {
					return 
					// returns to main menu
				}
		}

} // end for
} // end func kata 







func giveAhint(value string) {

	// fix all these hints so that they only refer to the value one needs to type that corresponds to the katakana prompt (to the key) 

	if value == "あ" { // a ア, maybe a grotesque A
		fmt.Printf(", hint: middle<- to the 3 char, looks nothing-like the hiragana a, but a lot like a te ア, あ, fuck mae! \n")
	} else 
	if value == "い" { // i イ, shift the two lines of the hiragana 
		fmt.Printf(", hint: middle < to the E char, イ looks more like a hiragana te, but at least it is still two mostly-vertical lines イ　\n")
	} else 
	if value == "う" { // u ウ, um-kay
		fmt.Printf(", hint: middle> to the 4 char, u ウ　is ok, having had to look for angles \n")
	} else 
	if value == "え" { // e エ, eye see it as a ... an, eye 
		fmt.Printf(", hint: index > to the 5 char, it エ　does have a vague, angular resemblance  \n")
	} else 
	if value == "お" { // o オ, on-the-go maybe
		fmt.Printf(", hint: index--> to the 6 char, オ　has only a vague resemblance, albeit with less curves \n")
	} else 


	if value == "か" { // ka カ
		fmt.Printf(", hint: index --> to the T char \n")
	} else 
	if value == "き" { // ki キ
		fmt.Printf(", hint: L-index > to the G char \n")
	} else 
	if value == "く" { // ku ク, compare to ta タ, and ke ケ
		fmt.Printf(", hint: R-index <- to the H char, ク　no, just no. Starting with one angle, they settled for this? \n")
	} else 
	if value == "け" { // ke ケ, compare to ku ク, and ta タ
		fmt.Printf(", hint: pinky -> one over to the :* chars, ケ, bits of it are there, just as many curves though \n")
	} else 
	if value == "こ" { // ko コ, compare to ni ニ
		fmt.Printf(", hint: index <--- to the B char, alright, コ　it makes sense, 'cause angles \n")
	} else 


	if value == "さ" { // sa サ, if you 'se' so sa 
		fmt.Printf(", hint: ring >v to the X char, at least the sa goes to the left, and it looks a lot like se, I'd say \n")
	} else 
	if value == "し" { // shi シ, she sleeps and snores, compare to tsu ツ, and so ソ, 
		fmt.Printf(", hint: middle on the D char, シ she is sleeping/snoring here, but has no excuse to2 look like this (no angles here, less curve though) \n")
	} else 
	if value == "す" { // su ス, sue is running 
		fmt.Printf(", hint: index < to the R char, they were looking for angles (sue is running ス) \n")
	} else 
	if value == "せ" { // se セ
		fmt.Printf(", hint: pinky to the P char \n")
	} else 
	if value == "そ" { // so ソ, compare to ん ン which lays down more and has a more laid-down back-tick 
		fmt.Printf(", hint: index <-- to the C char, ソ Looks like しシ or no ノ, or ン starting with all angles - `had to `backtrack?  \n")
	} else 


	if value == "た" { // ta タ, compare to ku ク and ke ケ
		fmt.Printf(", hint: pinky < to the Q char, the top left looks like ta, more at least than ku ク, or he ケ \n")
	} else 
	if value == "ち" { // chi チ, compare to te テ
		fmt.Printf(", hint: pinky on the A char, 'some' resemblance if we see the middle line as the top of the backwards c \n")
	} else 
	if value == "つ" { // tsu ツ, compare to shi シ, she who sleeps low and snores loudly 
		fmt.Printf(", hint: pinky to the Z char, understandable?, if we are generous. Very generous..゛as in adding two ツ  shi シ said \n")
	} else 
	if value == "て" { // te テ, compare to chi チ
		fmt.Printf(", hint: ring < to the W char, te went the wrong way, and gained a flat hat \n")
	} else 
	if value == "と" { // to ト
		fmt.Printf(", hint: ring on the S char, toe is flipped-out. Kicked in the balls, on its head \n")
	} else 


	if value == "な" { // na ナ, compare to me メ, and nu ヌ
		fmt.Printf(", hint: index < to the U char, is simple, very simple \n")
	} else 
	if value == "に" { // ni ニ
		fmt.Printf(", hint: middle < to the I char \n")
	} else 
	if value == "ぬ" { // nu ヌ, compare to me メ, and na ナ
		fmt.Printf(", hint: R-ring<--- to the 1 char, at least two lines cross, nu flew  \n")
	} else 
	if value == "ね" { // ne ネ
		fmt.Printf(", hint: L-ring <v to the ,< chars, a hoe that got nailed \n")
	} else 
	if value == "の" { // no ノ, no extra marks needed on this classic no symbol 
		fmt.Printf(", hint: middle on the k char, drop the curve, and I'll take it \n")
	} else 


	if value == "ま" { // ma マ, mama is unique, see the breast pump? 
		fmt.Printf(", hint: index on the J char, (a breast pump maybe?) \n")
	} else 
	if value == "み" { // mi ミ, and two makes three 
		fmt.Printf(", hint: index <v to the N char, me thinks 3 \n")
	} else 
	if value == "む" { // mu ム, an even-fatter-jawed moo-cow
		fmt.Printf(", hint: pinky two over ---> to the ]} chars, now it's an even-fatter-jawed moo-cow  \n")
	} else 
	if value == "め" { // me メ
		fmt.Printf(", hint: pinky -> slide down to the /? chars, cross off the mess. But that's nothing nu ヌ \n")
	} else 
	if value == "も" { // mo モ
		fmt.Printf(", hint: middle <v to the M char, way to hang to the right; you go Joe! \n")
	} else 


	if value == "ら" { // ra ラ
		fmt.Printf(", hint: ring ^ to the o char, Similar ラ, ら \n")
	} else 
	if value == "り" { // ri リ
		fmt.Printf(", hint: ring on the L char, longer on the right ring \n")
	} else 
	if value == "る" { // ru ル, is two
		fmt.Printf(", hint: pinky < to the .> chars, ル, る, ru is now two? \n")
	} else 
	if value == "れ" { // re レ, more, or less, a ru ル - しshe-said
		fmt.Printf(", hint: pinky on the ;+ chars, レ, れ; しshe-it \n")
	} else 
	if value == "ろ" { // ro ロ
		fmt.Printf(", hint: pinky ---> long slide to the _ and backslash chars, maybe it had always been the way to go to do a square O for ro? \n")
	} else 


	if value == "ほ" { // ha ハ 
		fmt.Printf(", hint: L-ring---> to the \"-\" char\n")
	} else 
	if value == "ひ" { // hi ヒ
		fmt.Printf(", hint: index > to the V char \n")
	} else 
	if value == "ふ" { // hu フ, squinting it is a ふフ
		fmt.Printf(", hint: ring<, to the 2 char, if we squint? フ, ふ\n")
	} else 
	if value == "へ" { // he ヘ
		fmt.Printf(", hint: ;+ char \n")
	} else 
	if value == "ほ" { // ho ホ, now with a dress and all
		fmt.Printf(", hint: ring ---> to the - char, now looks like a hoe \n")
	} else 


	if value == "や" { // ya ヤ
		fmt.Printf(", hint: index <-- to the 7 char \n")
	} else 
	if value == "ゆ" { // yu ユ
		fmt.Printf(", hint: index->, to the 8 char, yu looks like ユ \n")
	} else 
	if value == "よ" { // yo ヨ
		fmt.Printf(", hint: middle to the 9 char, tripple yo ヨ \n")
	} else 


	if value == "わ" { // wa ワ, compare to wo ヲ
		fmt.Printf(", hint: ring > to the 0 char, so now it's a water fall for wa? \n")
	} else 
	if value == "を" { // wo ヲ, compare to wa ワ
		fmt.Printf(", hint: ring > shifted 0 char, NOW, it looks like wa. Now! For fucks sake, Now? \n")
	} else 

	
	if value == "ん" { // nh ン, compare to so ソ
		fmt.Printf(", hint: Y char, pointing at the one last remaining bent stroke of the hiragana char \n")
	} else {
		fmt.Printf("\n")
	}
} // end of func giveAhint()






func restartKata() (returnToMenuRS bool) {

	rand.Seed(time.Now().UnixNano())

	keyValuePairs := []string{

/*
ChatGPT 3.5 proofread all of the katakana, hiragana pairs used in this code; and my Japanese wife verified them also. But I did catch them both having gotten one wrong:
they made one mistake in saying that "ビュ, ぴゅ" was correct when it should have been "ピュ, ぴゅ"
*/

// "vowels" :
//  a         * 2 :     　i  * 1      u       　e        o             u        e        o  * 2 :
	"ア, あ", "ア, あ",    "イ, い",    "ウ, う", "エ, え", "オ, お",     "ウ, う", "エ, え", "オ, お",      

// ka group: 
//  ka       ki       ku      　ke       ko * 1    ||  　ga       gi       gu      　ge       go * 1         //  naked || ゛marks
	"カ, か", "キ, き", "ク, く", "ケ, け", "コ, こ",       "ガ, が", "ギ, ぎ", "グ, ぐ", "ゲ, げ", "ゴ, ご",  
//            ^---^--v                                            ^^--^
	           "キャ, きゃ",  "キュ, きゅ",  "キョ, きょ",  // * 1     ^^--^-v                                           // ki,  ya yu yo (kya, kyu, kyo)
	                                                              "ギャ, ぎゃ",  "ギュ, ぎゅ",  "ギョ, ぎょ",     // * 1 // ki゛, ya yu yo (gya, gyu, gyo)
//                   ^            ^             ^                       ^             ^            ^        

// sa group: 
//  sa       shi      su       se        so * 1   ||   za       ji        zu       ze       zo * 1         //  naked || ゛marks
	"サ, さ", "シ, し", "ス, す", "セ, せ", "ソ, そ",      "ザ, ざ", "ジ, じ", "ズ, ず", "ゼ, ぜ", "ゾ, ぞ",      //-* One key "zu", has two values づ and ず *-//
//           ^----^-v                                           ^^---^
	          "シャ, しゃ",  "シュ, しゅ",  "ショ, しょ",  // * 1    ^^---^-v                                          // shi,  ya yu yo (sha, shu, sho)
	                                                             "ジャ, じゃ",   "ジュ, じゅ",   "ジョ, じょ",  // * 1 // shi゛, ya yu yo (ja, ju, jo)
//                  ^            ^             ^                       ^             ^              ^    

// ta group: 　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　vv ji never gi
//  ta       chi      tsu      te        to * 1   ||   da        ji       zu       de       do * 1        //  naked || ゛marks
	"タ, た", "チ, ち", "ツ, つ", "テ, て", "ト, と",      "ダ, だ",  "ヂ, ぢ", "ヅ, づ", "デ, で", "ド, ど",    //-* One key "zu", has two values づ and ず *-//
//            ^--^--v                                            ^^--^
             "チャ, ちゃ", "チュ, ちゅ", "チョ, ちょ",   // * 1      ^^--^-v                                           // chi,  ya yu yo (cha, chu, cho)
	                                                              "ヂャ, ぢゃ",  "ヂュ, ぢゅ",  "ヂョ, ぢょ",    // * 1 // chi゛, ya yu yo (ja, ju, jo)
//                  ^           ^            ^                          ^            ^             ^

// na group:  'cause they are na-ty special -- they like to go naked 
//  na       ni       nu        ne       no * 1                                                       //  naked, all of them (there are no ゛s on na's)
	"ナ, な", "ニ, に", "ヌ, ぬ", "ネ, ね", "ノ, の",    
//            ^---^-v                                
             "ニャ, にゃ", "ニュ, にゅ", "ニョ, にょ",   "ニャ, にゃ", "ニュ, にゅ", "ニョ, にょ",                   // * 2 // ni, ya yu yo (nya, nyu, nyo)
//                  ^           ^            ^              ^           ^            ^ 

// ha group: -- the fully-clothed group (゛marks, and poo゜marks too)
//  ha       hi       hu       he        ho   * 1   ||    ba       bi       bu        be       bo          //  naked || ゛marks
   "ハ, は", "ヒ, ひ", "フ, ふ", "ヘ, へ", "ホ, ほ",         "バ, ば", "ビ, び", "ブ, ぶ", "ベ, べ", "ボ, ぼ",    //  ゛marks 
//            ^--^--v                                      pa        pi       pu       pe        po        //  poo゜ marks (drag)
                                                           "パ, ぱ", "ピ, ぴ", "プ, ぷ", "ペ, ぺ", "ポ, ぽ",  //   ^ ^  ^ ^ ^  ^ ^ ^
//            ^-----v                                                ^^--^
	          "ヒャ, ひゃ", "ヒュ, ひゅ", "ヒョ, ひょ",               //  ^^--^--v           v            v       // hi , ya yu yo (hi) hya, hyu, hyo
	                                                                  "ビャ, びゃ", "ビュ, びゅ", "ビョ, びょ",   // hi゛, ya yu yo (bi) bya, byu, byo
	                                                                  "ピャ, ぴゃ", "ピュ, ぴゅ", "ピョ, ぴょ",   // hi゜,  ya yu yo (pi) pya, pyu, pyo
//                  ^           ^            ^                                            ^           ^            ^   

// ma group:  'cause they too are special -- they too also like to go naked
//  ma       mi       mu        me       mo * 1                   //  naked, all of them (there are no ゛s on mama's)
   "マ, ま", "ミ, み", "ム, む", "メ, め", "モ, も",     
//            ^--^-v
             "ミャ, みゃ", "ミュ, みゅ", "ミョ, みょ",                // * 1 // mi , ya yu yo (mya, myu, myo)
//                  ^           ^            ^          

// ya, yu, yo set: 
//  ya       yu       yo  * 1
   "ヤ, や", "ユ, ゆ", "ヨ, よ",    // * 1 (also used as suffixes)
// 

// ra/la group: spelled with the "r" but pronounced with the "L" sounds -- they too are special -- and also like to go naked
//  ra       ri       ru       re        ro  * 1    
   "ラ, ら", "リ, り", "ル, る", "レ, れ", "ロ, ろ", 
//            ^--^--v
             "リャ, りゃ", "リュ, りゅ", "リョ, りょ",       // * 1 // ri, ya yu yo (ryo)
//                  ^           ^            ^       

// wa set, plus the nh consonant -- always fully naked AND never even wear suffixes of ya yu, or yo 
//     wa       wo         nh                                   * 2
   "ワ, わ", "ヲ, を",   "ン, ん",       "ワ, わ", "ヲ, を",   "ン, ん" } 

		randIndex := rand.Intn(len(keyValuePairs))
		cardSelected := keyValuePairs[randIndex]

		parts = strings.Split(cardSelected, ", ")
		if len(parts) != 2 {
			fmt.Println("Invalid key-value pair:", cardSelected)
			//continue
		}

		key = parts[0]
		value = parts[1]

		fmt.Printf("%s", key)
		fmt.Printf("%s", string(colorCyan))
		fmt.Printf(" ? ")
		fmt.Printf("%s", string(colorReset))

		fmt.Scan(&in)

 
		if in == "menu" {
			returnToMenuRS = true // the return val
			return 
		}

		kataMeatB()
		ContextSHskip = false

		// this is the return val:
			returnToMenuRS = false
		return 
} // end of func restartKata()







func kataMeatB() (returnToMenuKMB bool) {

		if in == "menu" {
			returnToMenuKMB = true
			return 
		}
		if in == value {
			fmt.Printf("%s", string(colorGreen))
			fmt.Printf("      ^Right! \n")
			fmt.Printf("%s", string(colorReset))
		} else {
			fmt.Printf("%s", string(colorRed))
			fmt.Printf("      ^Oops! ")
			fmt.Printf("%s", string(colorReset))

			fmt.Printf("it was: ")
			fmt.Printf("%s", string(colorCyan))
			fmt.Printf("%s ", value)
			fmt.Printf("%s", string(colorReset))


			// hints for katakana (key) prompts: (values are in hiragana)
			// hints only are given for values that are naked (lack ゛゜or suffixes of や, ゆ, or よ)

			if value == "あ" { // a ア, maybe a grotesque A
				fmt.Printf(", hint: middle<- to the 3 char, looks nothing-like the hiragana a, but a lot like a te ア, あ, fuck mae! \n")
			} else 
			if value == "い" { // i イ, shift the two lines of the hiragana 
				fmt.Printf(", hint: middle < to the E char, イ looks more like a hiragana te, but at least it is still two mostly-vertical lines イ　\n")
			} else 
			if value == "う" { // u ウ, um-kay
				fmt.Printf(", hint: middle> to the 4 char, u ウ　is ok, having had to look for angles \n")
			} else 
			if value == "え" { // e エ, eye see it as a ... an, eye 
				fmt.Printf(", hint: index > to the 5 char, it エ　does have a vague, angular resemblance  \n")
			} else 
			if value == "お" { // o オ, on-the-go maybe
				fmt.Printf(", hint: index--> to the 6 char, オ　has only a vague resemblance, albeit with less curves \n")
			} else 


			if value == "か" { // ka カ
				fmt.Printf(", hint: index --> to the T char \n")
			} else 
			if value == "き" { // ki キ
				fmt.Printf(", hint: L-index > to the G char \n")
			} else 
			if value == "く" { // ku ク, compare to ta タ, and ke ケ
				fmt.Printf(", hint: R-index <- to the H char, ク　no, just no. Starting with one angle, they settled for this? \n")
			} else 
			if value == "け" { // ke ケ, compare to ku ク, and ta タ
				fmt.Printf(", hint: pinky -> one over to the :* chars, ケ, bits of it are there, just as many curves though \n")
			} else 
			if value == "こ" { // ko コ, compare to ni ニ
				fmt.Printf(", hint: index <--- to the B char, alright, コ　it makes sense, 'cause angles \n")
			} else 


			if value == "さ" { // sa サ, if you 'se' so sa 
				fmt.Printf(", hint: ring >v to the X char, at least the sa goes to the left, and it looks a lot like se, I'd say \n")
			} else 
			if value == "し" { // shi シ, she sleeps and snores, compare to tsu ツ, and so ソ, 
				fmt.Printf(", hint: middle on the D char, シ she is sleeping/snoring here, but has no excuse to2 look like this (no angles here, less curve though) \n")
			} else 
			if value == "す" { // su ス, sue is running 
				fmt.Printf(", hint: index < to the R char, they were looking for angles (sue is running ス) \n")
			} else 
			if value == "せ" { // se セ
				fmt.Printf(", hint: pinky to the P char \n")
			} else 
			if value == "そ" { // so ソ, compare to ん ン which lays down more and has a more laid-down back-tick 
				fmt.Printf(", hint: index <-- to the C char, ソ Looks like しシ or no ノ, or ン starting with all angles - `had to `backtrack?  \n")
			} else 


			if value == "た" { // ta タ, compare to ku ク and ke ケ
				fmt.Printf(", hint: pinky < to the Q char, the top left looks like ta, more at least than ku ク, or he ケ \n")
			} else 
			if value == "ち" { // chi チ, compare to te テ
				fmt.Printf(", hint: pinky on the A char, 'some' resemblance if we see the middle line as the top of the backwards c \n")
			} else 
			if value == "つ" { // tsu ツ, compare to shi シ, she who sleeps low and snores loudly 
				fmt.Printf(", hint: pinky to the Z char, understandable?, if we are generous. Very generous..゛as in adding two ツ  shi シ said \n")
			} else 
			if value == "て" { // te テ, compare to chi チ
				fmt.Printf(", hint: ring < to the W char, te went the wrong way, and gained a flat hat \n")
			} else 
			if value == "と" { // to ト
				fmt.Printf(", hint: ring on the S char, toe is flipped-out. Kicked in the balls, on its head \n")
			} else 


			if value == "な" { // na ナ, compare to me メ, and nu ヌ
				fmt.Printf(", hint: index < to the U char, is simple, very simple \n")
			} else 
			if value == "に" { // ni ニ
				fmt.Printf(", hint: middle < to the I char \n")
			} else 
			if value == "ぬ" { // nu ヌ, compare to me メ, and na ナ
				fmt.Printf(", hint: R-ring<--- to the 1 char, at least two lines cross, nu flew  \n")
			} else 
			if value == "ね" { // ne ネ
				fmt.Printf(", hint: L-ring <v to the ,< chars, a hoe that got nailed \n")
			} else 
			if value == "の" { // no ノ, no extra marks needed on this classic no symbol 
				fmt.Printf(", hint: middle on the k char, drop the curve, and I'll take it \n")
			} else 


			if value == "ま" { // ma マ, mama is unique, see the breast pump? 
				fmt.Printf(", hint: index on the J char, (a breast pump maybe?) \n")
			} else 
			if value == "み" { // mi ミ, and two makes three 
				fmt.Printf(", hint: index <v to the N char, me thinks 3 \n")
			} else 
			if value == "む" { // mu ム, an even-fatter-jawed moo-cow
				fmt.Printf(", hint: pinky two over ---> to the ]} chars, now it's an even-fatter-jawed moo-cow  \n")
			} else 
			if value == "め" { // me メ
				fmt.Printf(", hint: pinky -> slide down to the /? chars, cross off the mess. But that's nothing nu ヌ \n")
			} else 
			if value == "も" { // mo モ
				fmt.Printf(", hint: middle <v to the M char, way to hang to the right; you go Joe! \n")
			} else 


			if value == "ら" { // ra ラ
				fmt.Printf(", hint: ring ^ to the o char, Similar ラ, ら \n")
			} else 
			if value == "り" { // ri リ
				fmt.Printf(", hint: ring on the L char, longer on the right ring \n")
			} else 
			if value == "る" { // ru ル, is two
				fmt.Printf(", hint: pinky < to the .> chars, ル, る, ru is now two? \n")
			} else 
			if value == "れ" { // re レ, more, or less, a ru ル - しshe-said
				fmt.Printf(", hint: pinky on the ;+ chars, レ, れ; しshe-it \n")
			} else 
			if value == "ろ" { // ro ロ
				fmt.Printf(", hint: pinky ---> long slide to the _ and backslash chars, maybe it had always been the way to go to do a square O for ro? \n")
			} else 


			if value == "ほ" { // ha ハ 
				fmt.Printf(", hint: L-ring---> to the \"-\" char\n")
			} else 
			if value == "ひ" { // hi ヒ
				fmt.Printf(", hint: index > to the V char \n")
			} else 
			if value == "ふ" { // hu フ, squinting it is a ふフ
				fmt.Printf(", hint: ring<, to the 2 char, if we squint? フ, ふ\n")
			} else 
			if value == "へ" { // he ヘ
				fmt.Printf(", hint: ;+ char \n")
			} else 
			if value == "ほ" { // ho ホ, now with a dress and all
				fmt.Printf(", hint: ring ---> to the - char, now looks like a hoe \n")
			} else 


			if value == "や" { // ya ヤ
				fmt.Printf(", hint: index <-- to the 7 char \n")
			} else 
			if value == "ゆ" { // yu ユ
				fmt.Printf(", hint: index->, to the 8 char, yu looks like ユ \n")
			} else 
			if value == "よ" { // yo ヨ
				fmt.Printf(", hint: middle to the 9 char, tripple yo ヨ \n")
			} else 


			if value == "わ" { // wa ワ, compare to wo ヲ
				fmt.Printf(", hint: ring > to the 0 char, so now it's a water fall for wa? \n")
			} else 
			if value == "を" { // wo ヲ, compare to wa ワ
				fmt.Printf(", hint: ring > shifted 0 char, NOW, it looks like wa. Now! For fucks sake, Now? \n")
			} else 

			
			if value == "ん" { // nh ン, compare to so ソ
				fmt.Printf(", hint: Y char, pointing at the one last remaining bent stroke of the hiragana char \n")
			} else {
				fmt.Printf("\n")
			}
		}

		fmt.Println()
	returnToMenuKMB = false
	return 
} // end of func kataMeatB 
















func hira() {

	//fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")

	rand.Seed(time.Now().UnixNano())

	keyValuePairs := []string{

// the "vowels" (inclusive)
//  a * 2 :                 i * 1        u         e         o              u         e         o  * 2 :
	"aア, あ", "aア, あ",    "iイ, い",    "uウ, う", "eエ, え", "oオ, お",     "uウ, う", "eエ, え", "oオ, お",      

/*
   "ji" is for じ ,and "gi" is for ぎ .
ぎ:gi ギ
じ:ji ジ

ji:じ or ジ In the traditional Hepburn romanization system, the sound じ in hiragana is romanized as "ji," and the katakana ジ is also romanized as "ji" 
zi:じ However, in some other romanization systems like the Nihon-shiki and Kunrei-shiki, the sound じ is romanized as "zi" instead of "ji."
gi:ぎ or ギ The sound gi:ぎ in hiragana is romanized as "gi," and the katakana ギ is also romanized as "gi."

So, both "gi" and "ji" are used as romanizations for different contexts, but the standard in most systems is "ji" for じ in hiragana and "gi" for ぎ in hiragana.

*/
// ka group: fixed according to kaz and chat                                  vv gi:ぎ 
//  ka         ki         ku          ke         ko * 1      ||   ga          gi         gu         ge         go * 1          //  naked || ゛marks
	"kaカ, か", "kiキ, き", "kuク, く", "keケ, け", "koコ, こ",       "gaガ, が", "giギ, ぎ", "guグ, ぐ", "geゲ, げ", "goゴ, ご",   
//              ^----^--v                                                     ^^----^
	           "kyaキャ, きゃ",  "kyuキュ, きゅ",  "kyoキョ, きょ",  // * 1        ^^- --^-v                                                // ki,  ya yu yo (kya, kyu, kyo)
	                                                                         "gyaギャ, ぎゃ",  "gyuギュ, ぎゅ",  "gyoギョ, ぎょ",    // * 1 // ki゛, ya yu yo (gya, gyu, gyo)
//                      ^               ^                ^                            ^                ^               ^        

// sa group: 
//  sa         shi         su          se         so * 1     ||    za         ji never gi   　zu          ze         zo * 1         //  naked || ゛marks
	"saサ, さ", "shiシ, し", "suス, す", "seセ, せ", "soソ, そ",      "zaザ, ざ", "jiジ, じ",       "zu ズ, ず", "zeゼ, ぜ", "zoゾ, ぞ",  //-* One key "zu", has two values づ and ず *-//
//               ^----^-v                                                     ^^----^
	           "shaシャ, しゃ",  "shuシュ, しゅ",  "shoショ, しょ",  // * 1        ^^--^-^-v                                          // shi,  ya yu yo (sha, shu, sho)
	                                                                          "jaジャ, じゃ",   "juジュ, じゅ",   "joジョ, じょ",     // shi゛, ya yu yo (ja, ju, jo) prefered jx vers ***
//                      ^               ^                ^                            ^               ^                ^    
"zu ズ from つ or す, ず", // //-* One key "zu", has two values づ and ず of the ta or sa group *-//
                                                                    //             ^
                                                                   //              じ is used 95% relative to ぢ , therefore ぢ is 95% wrong! 
// ta group:   
//  ta         chi         tsu          te         to * 1     /    da          ji never gi    zu          de          do * 1        //  naked / ゛
	"taタ, た", "chiチ, ち", "tsuツ, つ", "teテ, て", "toト, と",     "daダ, だ",  "jiヂ, ぢ",     "zu ヅ, づ", "deデ, で", "doド, ど",  //-* One key "zu", has two values づ and ず *-//
//               ^----^-v                                                      ^^----^
	           "chaチャ, ちゃ", "chuチュ, ちゅ", "choチョ, ちょ",   // * 1          ^^----^-----v                                                     // chi, ya yu yo (cha, chu, cho)
	                                                                           "ja obsヂャ, ぢゃ",  "ju obsヂュ, ぢゅ",  "jo obsヂョ, ぢょ",   // * 1 // chi゛, ya yu yo (ja, ju, jo) obsolete
	                                                                    //      ^                   ^                  ^  ^^^obsolete, use the jx versions from the sa group instead ***
//                      ^              ^               ^                                   ^                  ^                   ^ ち s
"zu ズ from つ or す, づ", // //-* One key "zu", has two values づ and ず of the ta or sa group *-//
"zu ヅ from つ or す, づ", // //-* One key "zu", has two values づ and ず of the ta or sa group *-//



// na group: 'cause they are special -- they really prefer to go-naked 
//   na         ni         nu        ne         no * 1       //  naked, all of them (there are no ゛s on na's)
	"naナ, な", "niニ, に", "nuヌ, ぬ", "neネ, ね", "noノ, の",    
//              ^----^---v                                
               "nyaニャ, にゃ", "nyuニュ, にゅ", "nyoニョ, にょ",   // ni, ya yu yo (nya, nyu, nyo)
//                       ^              ^               ^       


// ha group: -- the fully-clothed group 
//  ha         hi         hu          he         ho    * 1      ||    ba         bi         bu         be          bo            //  naked || ゛marks
   "haハ, は", "hiヒ, ひ", "huフ, ふ", "heヘ, へ", "hoホ, ほ",          "baバ, ば", "biビ, び", "buブ, ぶ", "beベ, べ", "boボ, ぼ",  
//              ^-^--v                                                pa         pi         pu          pe         po            //  poo゜ marks (drag)
                                                                     "paパ, ぱ", "piピ, ぴ", "puプ, ぷ", "peペ, ぺ", "poポ, ぽ", 
//              ^-------v                                                         v----v
	            "hyaヒャ, ひゃ", "hyuヒュ, ひゅ", "hyoヒョ, ひょ",              //     v---------v                                      // hi, ya yu yo (hya, hyu, hyo)
	                                                                               "byaビャ, びゃ", "byuビュ, びゅ", "byoビョ, びょ",   // bi, ya yu yo (bya, byu, byo)
	                                                                               "pyaピャ, ぴゃ", "pyuピュ, ぴゅ", "pyoピョ, ぴょ",   // pi, ya yu yo (pya, pyu, pyo)
//                   -  ^       --      ^    --         ^                                    ^             ^               ^   

// ma group: 'cause they too are special -- they too also like to go naked
//  ma         mi         mu          me         mo * 1          //  naked, all of them (there are no ゛s on mama's)
   "maマ, ま", "miミ, み", "muム, む", "meメ, め", "moモ, も",     
//              ^----^--v
               "myaミャ, みゃ", "myuミュ, みゅ", "myoミョ, みょ",    // mi, ya yu yo (mya, myu, myo)
//                       ^              ^               ^   

// ya, yu, yo set: 
//  ya         yu         yo 
   "yaヤ, や", "yuユ, ゆ", "yoヨ, よ",    // (also used as suffixes)
// 

// ra/la group: spelled with the "r" but pronounced with the "L" sounds -- they too are special -- and also like to go naked
//  ra         ri          ru         re         ro     
   "raラ, ら", "riリ, り", "ruル, る", "reレ, れ", "roロ, ろ", 
//               ^---^--v
               "ryaリャ, りゃ", "ryuリュ, りゅ", "ryoリョ, りょ",   // ri, ya yu yo (rya, ryu, ryo)
//                      ^               ^               ^  

// wa set, plus the nh consonant -- always fully naked AND never have suffixes of ya yu, or yo 
//  wa         wo           nh   * 2
   "waワ, わ", "woヲ, を",   "nhン, ん",   "waワ, わ", "woヲ, を",   "nhン, ん" }  //  go.lang requies that the "}" must be on this line rather than on a line of its own 

doingContextSH = false
	for {
		var skip = false // unless a guess has been handled by a special section, we will skip the normal handeler 

if doingContextSH == false {


		randIndex := rand.Intn(len(keyValuePairs))
		cardSelected := keyValuePairs[randIndex]

		parts = strings.Split(cardSelected, ", ")
		if len(parts) != 2 {
			fmt.Println("Invalid key-value pair:", cardSelected)
			continue
		}

		key = parts[0]
		value = parts[1]

/*
		fmt.Printf("%s ? ", key)
*/

		fmt.Printf("%s", key)
		fmt.Printf("%s", string(colorCyan))
		fmt.Printf(" ? ")
		fmt.Printf("%s", string(colorReset))

		fmt.Scan(&in)

		if in == "menu" {
			return 
		}
		if in == "?" {
			fmt.Printf("how is this")
			doingContextSH = true
		}

} else {
		//key := parts[0]
		//value := parts[1]
		fmt.Printf("doingCSH")
		fmt.Printf("%s", key)
		fmt.Printf("%s", string(colorCyan))
		fmt.Printf(" ? ")
		fmt.Printf("%s", string(colorReset))
		doingContextSH = false
}




/*
	ぎ:gi ギ
	じ:ji ジ

	ji:じ or ジ In the traditional Hepburn romanization system, the sound じ in hiragana is romanized as "ji," and the katakana ジ is also romanized as "ji" 
	zi:じ However, in some other romanization systems like the Nihon-shiki and Kunrei-shiki, the sound じ is romanized as "zi" instead of "ji."
	gi:ぎ or ギ The sound gi:ぎ in hiragana is romanized as "gi," and the katakana ギ is also romanized as "gi."
*/
		// proper use of the ji sound from the sa group (sometimes spelled zi) : sa->za,ji , special key handelers
		if key == "jiジ" {
			if in == "じ" {
				fmt.Printf("%s", string(colorGreen))
				fmt.Printf("\n      ^Right! ")
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("ジ is always from shi し and NEVER from chi ち ,it is the sound ji, NEVER gi ( that being ぎ:gi ギ )  \n")
				skip = true 
				// さ　し　 す　せ　そ　
				// sa shi su se so
				// za ji  zu ze zo゛
			} else {
				fmt.Printf("%s", string(colorRed))
				fmt.Printf("      ^Oops! ")
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("it was ")
				fmt.Printf("%s", string(colorCyan))
				fmt.Printf("%s", value)
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("and remember, ジ is the sound ji always from shi し ,and NEVER from the other ji: chi ぢ ")
				skip = true 
			}
		}
		// proper use of the ji sound from the ta group : ta->da,ji , a special key handeler
		if key == "jiヂ" {
			if in == "ぢ" {
				fmt.Printf("%s", string(colorGreen))
				fmt.Printf("\n      ^Right! ")
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("ヂ is always from chi ち ,it is the sound ji, NEVER gi ( that being ぎ:gi ギ )  \n")
				skip = true 
				// た　ち   つ　て　と　
				// ta chi tsu te to
				// da ji  zu  de do゛
			} else {
				fmt.Printf("%s", string(colorRed))
				fmt.Printf("      ^Oops! ")
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("it was ")
				fmt.Printf("%s", string(colorCyan))
				fmt.Printf("%s", value)
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("and remember, the ヂ sound ji is always from either chi ち or shi し ")
				skip = true 
			}
		}

		// proper use of the gi sound from the ka group : ka->ga , a special key handeler
		if key == "giギ" {
			if in == "ぎ" {
				fmt.Printf("%s", string(colorGreen))
				fmt.Printf("\n      ^Right! ")
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("it is always from ki き and NEVER from shi し ,and it is the sound gi, NEVER ji ( that being じ:ji ジ )  \n")
				skip = true 
				// か　き　く　け　こ　
				// ka ki ku ke ko
				// ga gi gu ge go゛
			} else {
				fmt.Printf("%s", string(colorRed))
				fmt.Printf("      ^Oops! ")
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("it was ")
				fmt.Printf("%s", string(colorCyan))
				fmt.Printf("%s", value)
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("and remember, the sound gi is always from ki き ,and NEVER from shi し or chi ち ")
				skip = true 
			}
		}


		// properly-constructed forms of ji used as ja, ju, jo which are nearly always from shi し and really never from chi ち
		// , special key handelers
		if key == "jaジャ" {
			if in == "じゃ" {
				fmt.Printf("%s", string(colorGreen))
				fmt.Printf("\n      ^Right! ")
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("it is nearly always from shi し and really never from chi ち  \n")
				skip = true 
			} else {
				fmt.Printf("%s", string(colorRed))
				fmt.Printf("      ^Oops! ")
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("it was ")
				fmt.Printf("%s", string(colorCyan))
				fmt.Printf("%s", value)
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("and remember that it is nearly-never from chi ち but is nearly-always from shi し ")
				skip = true 
			}
		}
		if key == "juジュ" {
			if in == "じゅ" {
				fmt.Printf("%s", string(colorGreen))
				fmt.Printf("\n      ^Right! ")
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("it is nearly always from shi し and really never from chi ち  \n")
				skip = true 
			} else {
				fmt.Printf("%s", string(colorRed))
				fmt.Printf("      ^Oops! ")
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("it was ")
				fmt.Printf("%s", string(colorCyan))
				fmt.Printf("%s", value)
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("and remember that it is nearly-never from chi ち but is nearly-always from shi し ")
				skip = true 
			}
		}
		if key == "joジョ" {
			if in == "じょ" {
				fmt.Printf("%s", string(colorGreen))
				fmt.Printf("\n      ^Right! ")
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("it is nearly always from shi し and really never from chi ち  \n")
				skip = true 
			} else {
				fmt.Printf("%s", string(colorRed))
				fmt.Printf("      ^Oops! ")
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("it was ")
				fmt.Printf("%s", string(colorCyan))
				fmt.Printf("%s", value)
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("and remember that it is nearly-never from chi ち but is nearly-always from shi し ")
				skip = true 
			}
		}

// /* obsolete forms of ja, ju and jo , special key handelers
		if key == "ja obsヂャ" {
			if in == "ぢゃ" {
				fmt.Printf("%s", string(colorGreen))
				fmt.Printf("\n      ^Right! ")
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("it is nearly always from shi し and really never from chi ち  \n")
				skip = true 
			} else {
				fmt.Printf("%s", string(colorRed))
				fmt.Printf("      ^Oops! ")
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("it was ")
				fmt.Printf("%s", string(colorCyan))
				fmt.Printf("%s", value)
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("and remember that it is nearly-never from chi ち but is nearly-always from shi し ")
				skip = true 
			}
		}
		if key == "ju obsヂュ" {
			if in == "ぢゅ" {
				fmt.Printf("%s", string(colorGreen))
				fmt.Printf("\n      ^Right! ")
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("it is nearly always from shi し and really never from chi ち  \n")
				skip = true 
			} else {
				fmt.Printf("%s", string(colorRed))
				fmt.Printf("      ^Oops! ")
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("it was ")
				fmt.Printf("%s", string(colorCyan))
				fmt.Printf("%s", value)
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("and remember that it is nearly-never from chi ち but is nearly-always from shi し ")
				skip = true 
			}
		}
		if key == "jo obsヂョ" {
			if in == "ぢょ" {
				fmt.Printf("%s", string(colorGreen))
				fmt.Printf("\n      ^Right! ")
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("it is nearly always from shi し and really never from chi ち  \n")
				skip = true 
			} else {
				fmt.Printf("%s", string(colorRed))
				fmt.Printf("      ^Oops! ")
				fmt.Printf("%s", string(colorReset))


				fmt.Printf("it was ")
				fmt.Printf("%s", string(colorCyan))
				fmt.Printf("%s", value)
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("and remember that it is nearly-never from chi ち but is nearly-always from shi し ")
				skip = true 
			}
		}
// */


		// then we check for additional very-special prompt(key)/value events or conditions,  , a very-special key handeler

		if key == "zu ズ" || key == "zu ヅ" || key == "zu ズ from つ or す" || key == "zu ヅ from つ or す" {
			if in == "づ" || in == "ず" {
				fmt.Printf("%s", string(colorGreen))
				fmt.Printf("\n      ^Right! ")
				fmt.Printf("%s", string(colorReset))
				fmt.Printf("it could have been either ず or づ as they are the same sound \n")
				skip = true 
			} else { //just added :
				fmt.Printf("%s", string(colorRed))
				fmt.Printf("      ^Oops! ")
				fmt.Printf("%s", string(colorReset))

				//fmt.Printf("it was: %s \n", value)

				fmt.Printf("it was: ")
				fmt.Printf("%s", string(colorCyan))
				fmt.Printf("%s \n", value)
				fmt.Printf("%s", string(colorReset))
			}
			// the above needs a CR as in: ("it was: %s \n", value)





		// the next two following conditions are for all normal (not special) prompt(key)/value events or conditions 

/*
		} else if in == value && skip == false {
			fmt.Printf("\n      ^Right! \n")
		} else {
			if skip == false {
				fmt.Printf("      ^Oops! actually it was %s", value)
			} 
*/

		} else if in == value && skip == false {
			fmt.Printf("%s", string(colorGreen))
			fmt.Printf("      ^Right! \n")
			fmt.Printf("%s", string(colorReset))
		} else {
			if skip == false {
				fmt.Printf("%s", string(colorRed))
				fmt.Printf("      ^Oops! ")
				fmt.Printf("%s", string(colorReset))

				fmt.Printf("it was: ")
				fmt.Printf("%s", string(colorCyan))
				fmt.Printf("%s", value)
				fmt.Printf("%s", string(colorReset))
			} else 


// hints: may want to add additional special hints as special cases above

			if value == "あ" {
				fmt.Printf(", hint: middle<- to the 3 char \n")
			} else 
			if value == "い" {
				fmt.Printf(", hint: middle < to the E char \n")
			} else 
			if value == "う" {
				fmt.Printf(", hint: middle> to the 4 char \n")
			} else 
			if value == "え" {
				fmt.Printf(", hint: index > to the 5 char \n")
			} else 
			if value == "お" {
				fmt.Printf(", hint: index--> to the 6 char \n")
			} else 
			// 5

			if value == "か" {
				fmt.Printf(", hint: index --> to the T char \n")
			} else 
			if value == "き" {
				fmt.Printf(", hint: G char, key|ki of G \n")
			} else 
			if value == "く" {
				fmt.Printf(", hint: H char \n")
			} else 
			if value == "け" {
				fmt.Printf(", hint: :* chars \n")
			} else 
			if value == "こ" {
				fmt.Printf(", hint: B char, ko way below \n")
			} else 
			// 10

			if value == "さ" {
				fmt.Printf(", hint: X char \n")
			} else 
			if value == "し" {
				fmt.Printf(", hint: D char \n")
			} else 
			if value == "す" {
				fmt.Printf(", hint: R char \n")
			} else 
			if value == "せ" {
				fmt.Printf(", hint: P char \n")
			} else 
			if value == "そ" {
				fmt.Printf(", hint: C char \n")
			} else 
			// 15

			if value == "た" {
				fmt.Printf(", hint: Q char, pinky tag \n")
			} else 
			if value == "ち" {
				fmt.Printf(", hint: A char, A cheater \n")
			} else 
			if value == "つ" {
				fmt.Printf(", hint: Z char, pinky mowed down by Tsunami \n")
			} else 
			if value == "て" {
				fmt.Printf(", hint: W char, next to ta \n")
			} else 
			if value == "と" {
				fmt.Printf(", hint: S char, wedding-ring instead of toe \n")
			} else 
			// 20

			if value == "な" {
				fmt.Printf(", hint: U char \n")
			} else 
			if value == "に" {
				fmt.Printf(", hint: I char \n")
			} else 
			if value == "ぬ" {
				fmt.Printf(", hint: ring<---, 1 char \n")
			} else 
			if value == "ね" {
				fmt.Printf(", hint: ring< to the , char \n")
			} else 
			if value == "の" {
				fmt.Printf(", hint: K char \n")
			} else 
			// 25

			if value == "ま" {
				fmt.Printf(", hint: J char \n")
			} else 
			if value == "み" {
				fmt.Printf(", hint: N char \n")
			} else 
			if value == "む" {
				fmt.Printf(", hint: pinky---> one-over \n")
			} else 
			if value == "め" {
				fmt.Printf(", hint: slide down to the /? chars \n")
			} else 
			if value == "も" {
				fmt.Printf(", hint: M char \n")
			} else 
			// 30

			if value == "ら" {
				fmt.Printf(", hint: ra:raised-ring-finger \n")
			} else 
			if value == "り" {
				fmt.Printf(", hint: ri:ring-stays-put \n")
			} else 
			if value == "る" {
				fmt.Printf(", hint: ru:loop-left \n")
			} else 
			if value == "れ" {
				fmt.Printf(", hint: re:lay-little-finger \n")
			} else 
			if value == "ろ" {
				fmt.Printf(", hint: ro:roll-way-over \n")
			} else 
			// 35

			if value == "ほ" {
				fmt.Printf(", hint: F \n") 
			}
			if value == "ひ" {
				fmt.Printf(", hint: V \n")
			} else 
			if value == "ふ" {
				fmt.Printf(", hint: 2 \n")
			} else 
			if value == "へ" {
				fmt.Printf(", hint: far-upper-right-corner \n")
			} else 
			if value == "ほ" {
				fmt.Printf(", hint: ring-->long-reach, dash char\n")
			} else 
			// 40

			if value == "や" {
				fmt.Printf(", hint: index long-reach<--, to the 7 char \n")
			} else 
			if value == "ゆ" {
				fmt.Printf(", hint: index->, to the 8 char \n")
			} else 
			if value == "よ" {
				fmt.Printf(", hint: middle>, to the 9 char \n")
			} else 
			// 43

			if value == "わ" {
				fmt.Printf(", hint: ring>, to the zero char \n")
			} else 
			if value == "を" {
				fmt.Printf(", hint: ring>, to the ^zero char \n")
			} else 
			// 45
			
			if value == "ん" {
				fmt.Printf(", hint: Y char \n")
			} else {
				fmt.Printf("\n")
			}
			// 46
		}
		fmt.Println()
	}
}

var menu int 
var in string
var doingContextSH bool 
var ContextSHskip bool
//var parts []string 
//var value string
//var key string

  var colorReset = "\033[0m"          
  var colorRed = "\033[31m"
  var colorGreen = "\033[32m"
  var colorYellow = "\033[33m"
  var colorPurple = "\033[35m"
  var colorCyan = "\033[36m"
  var colorWhite = "\033[37m"

type charSetStruct struct {
	KeyK string
	KeyH string
	KeyR string 
	KeyRK string 

	Value string

	Hint1k string 
	Hint2k string

	Hint1h string
	Hint2h string
}
var fileOfCards = []charSetStruct{
{"ア", "あ", "a", "aア", "あ",
 " a あ ア looks nothing-like the hiragana a, but a lot like a te ア, あ fuck mae! a ア, maybe a grotesque A",
 " a ア あ ア maybe a grotesque A",
 " a",
 " a"},

{"イ", "い", "i", "iイ", "い",
 " i い イ looks more like a hiragana te, but at least it is still two mostly-vertical lines イ shift the two lines of the hiragana",
 " i イ い イ shift the two lines of the hiragana い",
 " i",
 " i"},

{"ウ", "う", "u", "uウ", "う",
 " u う ウ　is ok, having had to look for angles. u ウ, um-kay",
 " u ウ う ウ more angles and with a tick for its top line",
 " u",
 " u"},

{"エ", "え", "e", "eエ", "え",
 " e え エ　does have a vague, angular resemblance. e エ, eye see it as a ... an, eye; may-bey",
 " e エ え エ eye see it as a ... an, eye eh; well, may-bey",
 " e",
 " e"},

{"オ", "お", "o", "oオ", "お",
 " o お オ　has only a vague resemblance, albeit with less curves. o オ, is someone on-the-go maybe",
 " o オ お オ is someone on-the-go maybe",
 " o",
 " o"},

{"カ", "か", "ka", "kaカ", "か",
 " ka か カ is an easy one",
 " ka, same か albeit more angular and one less line to draw",
 " ka",
 " ka"},

{"キ", "き", "ki", "kiキ", "き",
 " ki き キ is an easy one",
 " ki き キ has the same top",
 " ki",
 " ki"},

{"ク", "く", "ku", "kuク", "く",
 " ku く ク　no, just no. Starting with one angle, they settled for this? ku ク, compare to ta タ, and ke ケ",
 " ku ク く ク compare ta タ, ke ケ",
 " ku",
 " ku"},

{"ケ", "け", "ke", "keケ", "け",
 " ke け ケ bits of it are there, just as many curves though. ke ケ, compare to ku ク, and ta タ",
 " ke ケ け ケ compare ku ク, ta タ",
 " ke",
 " ke"},

{"コ", "こ", "ko", "koコ", "こ",
 " ko こ コ it makes sense, 'cause angles",
 " ko コ　こ コ, compare ni ニ",
 " ko",
 " ko"},

{"ガ", "が", "ga", "gaガ", "が",
 " ga",
 " ga",
 " ga",
 " ga"},

{"ギ", "ぎ", "gi", "giギ", "ぎ",
 " gi",
 " gi",
 " gi",
 " gi"},

{"グ", "ぐ", "gu", "guグ", "ぐ",
 " gu",
 " gu",
 " gu",
 " gu"},

{"ゲ", "げ", "ge", "geゲ", "げ",
 " ge",
 " ge",
 " ge",
 " ge"},

{"ゴ", "ご", "go", "goゴ", "ご",
 " go",
 " go",
 " go",
 " go"},

{"キャ", "きゃ", "kya", "kyaキャ", "きゃ",
 " kya",
 " kya",
 " kya",
 " kya"},

{"キュ", "きゅ", "kyu", "kyuキュ", "きゅ",
 " kyu",
 " kyu",
 " kyu",
 " kyu"},

{"キョ", "きょ", "kyo", "kyoキョ", "きょ",
 " kyo",
 " kyo",
 " kyo",
 " kyo"},

{"ギャ", "ぎゃ", "gya", "gyaギャ", "ぎゃ",
 " gya",
 " gya",
 " gya",
 " gya"},

{"ギュ", "ぎゅ", "gyu", "gyuギュ", "ぎゅ",
 " gyu",
 " gyu",
 " gyu",
 " gyu"},

{"ギョ", "ぎょ", "gyo", "gyoギョ", "ぎょ",
 " gyo",
 " gyo",
 " gyo",
 " gyo"},

{"サ", "さ", "sa", "saサ", "さ",
 " sa",
 " sa",
 " sa",
 " sa"},

{"シ", "し", "shi", "shiシ", "し",
 " shi",
 " shi",
 " shi",
 " shi"},

{"ス", "す", "su", "suス", "す",
 " su",
 " su",
 " su",
 " su"},

{"セ", "せ", "se", "seセ", "せ",
 " se",
 " se",
 " se",
 " se"},

{"ソ", "そ", "so", "soソ", "そ",
 " so",
 " so",
 " so",
 " so"},

{"ザ", "ざ", "za", "zaザ", "ざ",
 " za",
 " za",
 " za",
 " za"},

{"ジ", "じ", "ji", "jiジ", "じ",
 " ji",
 " ji",
 " ji",
 " ji"},

{"ズ", "ず", "zu", "zuズ", "ず",
 " zu",
 " zu",
 " zu",
 " zu"},

{"ゼ", "ぜ", "ze", "zeゼ", "ぜ",
 " ze",
 " ze",
 " ze",
 " ze"},

{"ゾ", "ぞ", "zo", "zoゾ", "ぞ",
 " zo",
 " zo",
 " zo",
 " zo"},

{"シャ", "しゃ", "sha", "shaシャ", "しゃ",
 " sha",
 " sha",
 " sha",
 " sha"},

{"シュ", "しゅ", "shu", "shuシュ", "しゅ",
 " shu",
 " shu",
 " shu",
 " shu"},

{"ショ", "しょ", "sho", "shoショ", "しょ",
 " sho",
 " sho",
 " sho",
 " sho"},

{"ジャ", "じゃ", "ja", "jaジャ", "じゃ",
 " ja",
 " ja",
 " ja",
 " ja"},

{"ジュ", "じゅ", "ju", "juジュ", "じゅ",
 " ju",
 " ju",
 " ju",
 " ju"},

{"ジョ", "じょ", "jo", "joジョ", "じょ",
 " jo",
 " jo",
 " jo",
 " jo"},

{"タ", "た", "ta", "taタ", "た",
 " ta",
 " ta",
 " ta",
 " ta"},

{"チ", "ち", "chi", "chiチ", "ち",
 " chi",
 " chi",
 " chi",
 " chi"},

{"ツ", "つ", "tsu", "tsuツ", "つ",
 " tsu",
 " tsu",
 " tsu",
 " tsu"},

{"テ", "て", "te", "teテ", "て",
 " te",
 " te",
 " te",
 " te"},

{"ト", "と", "to", "toト", "と",
 " to",
 " to",
 " to",
 " to"},

{"ダ", "だ", "da", "daダ", "だ",
 " da",
 " da",
 " da",
 " da"},

{"ヂ", "ぢ", "ji", "jiヂ", "ぢ",
 " ji",
 " ji",
 " ji",
 " ji"},

{"ヅ", "づ", "zu", "zuヅ", "づ",
 " zu",
 " zu",
 " zu",
 " zu"},

{"デ", "で", "de", "deデ", "で",
 " de",
 " de",
 " de",
 " de"},

{"ド", "ど", "do", "doド", "ど",
 " do",
 " do",
 " do",
 " do"},

{"チャ", "ちゃ", "cha", "chaチャ", "ちゃ",
 " cha",
 " cha",
 " cha",
 " cha"},

{"チュ", "ちゅ", "chu", "chuチュ", "ちゅ",
 " chu",
 " chu",
 " chu",
 " chu"},

{"チョ", "ちょ", "cho", "choチョ", "ちょ",
 " cho",
 " cho",
 " cho",
 " cho"},

{"ヂャ", "ぢゃ", "ja", "jaヂャ", "ぢゃ",
 " ja",
 " ja",
 " ja",
 " ja"},

{"ヂュ", "ぢゅ", "ju", "juヂュ", "ぢゅ",
 " ju",
 " ju",
 " ju",
 " ju"},

{"ヂョ", "ぢょ", "jo", "joヂョ", "ぢょ",
 " jo",
 " jo",
 " jo",
 " jo"},

{"ナ", "な", "na", "naナ", "な",
 " na",
 " na",
 " na",
 " na"},

{"ニ", "に", "ni", "niニ", "に",
 " ni",
 " ni",
 " ni",
 " ni"},

{"ヌ", "ぬ", "nu", "nuヌ", "ぬ",
 " nu",
 " nu",
 " nu",
 " nu"},

{"ネ", "ね", "ne", "neネ", "ね",
 " ne",
 " ne",
 " ne",
 " ne"},

{"ノ", "の", "no", "noノ", "の",
 " no",
 " no",
 " no",
 " no"},

{"ニャ", "にゃ", "nya", "nyaニャ", "にゃ",
 " nya",
 " nya",
 " nya",
 " nya"},

{"ニュ", "にゅ", "nyu", "nyuニュ", "にゅ",
 " nyu",
 " nyu",
 " nyu",
 " nyu"},

{"ニョ", "にょ", "nyo", "nyoニョ", "にょ",
 " nyo",
 " nyo",
 " nyo",
 " nyo"},

{"ハ", "は", "ha", "haハ", "は",
 " ha",
 " ha",
 " ha",
 " ha"},

{"ヒ", "ひ", "hi", "hiヒ", "ひ",
 " hi",
 " hi",
 " hi",
 " hi"},

{"フ", "ふ", "fu", "fuフ", "ふ",
 " fu",
 " fu",
 " fu",
 " fu"},

{"ヘ", "へ", "he", "heヘ", "へ",
 " he",
 " he",
 " he",
 " he"},

{"ホ", "ほ", "ho", "hoホ", "ほ",
 " ho",
 " ho",
 " ho",
 " ho"},

{"バ", "ば", "ba", "baバ", "ば",
 " ba",
 " ba",
 " ba",
 " ba"},

{"ビ", "び", "bi", "biビ", "び",
 " bi",
 " bi",
 " bi",
 " bi"},

{"ブ", "ぶ", "bu", "buブ", "ぶ",
 " bu",
 " bu",
 " bu",
 " bu"},

{"ベ", "べ", "be", "beベ", "べ",
 " be",
 " be",
 " be",
 " be"},

{"ボ", "ぼ", "bo", "boボ", "ぼ",
 " bo",
 " bo",
 " bo",
 " bo"},

{"パ", "ぱ", "pa", "paパ", "ぱ",
 " pa",
 " pa",
 " pa",
 " pa"},

{"ピ", "ぴ", "pi", "piピ", "ぴ",
 " pi",
 " pi",
 " pi",
 " pi"},

{"プ", "ぷ", "pu", "puプ", "ぷ",
 " pu",
 " pu",
 " pu",
 " pu"},

{"ペ", "ぺ", "pe", "peペ", "ぺ",
 " pe",
 " pe",
 " pe",
 " pe"},

{"ポ", "ぽ", "po", "poポ", "ぽ",
 " po",
 " po",
 " po",
 " po"},

{"ヒャ", "ひゃ", "hya", "hyaヒャ", "ひゃ",
 " hya",
 " hya",
 " hya",
 " hya"},

{"ヒュ", "ひゅ", "hyu", "hyuヒュ", "ひゅ",
 " hyu",
 " hyu",
 " hyu",
 " hyu"},

{"ヒョ", "ひょ", "hyo", "hyoヒョ", "ひょ",
 " hyo",
 " hyo",
 " hyo",
 " hyo"},

{"ビャ", "びゃ", "bya", "byaビャ", "びゃ",
 " bya",
 " bya",
 " bya",
 " bya"},

{"ビュ", "びゅ", "byu", "byuビュ", "びゅ",
 " byu",
 " byu",
 " byu",
 " byu"},

{"ビョ", "びょ", "byo", "byoビョ", "びょ",
 " byo",
 " byo",
 " byo",
 " byo"},

{"ピャ", "ぴゃ", "pya", "pyaピャ", "ぴゃ",
 " pya",
 " pya",
 " pya",
 " pya"},

{"ピュ", "ぴゅ", "pyu", "pyuピュ", "ぴゅ",
 " pyu",
 " pyu",
 " pyu",
 " pyu"},

{"ピョ", "ぴょ", "pyo", "pyoピョ", "ぴょ",
 " pyo",
 " pyo",
 " pyo",
 " pyo"},

{"マ", "ま", "ma", "maマ", "ま",
 " ma",
 " ma",
 " ma",
 " ma"},

{"ミ", "み", "mi", "miミ", "み",
 " mi",
 " mi",
 " mi",
 " mi"},

{"ム", "む", "mu", "muム", "む",
 " mu",
 " mu",
 " mu",
 " mu"},

{"メ", "め", "me", "meメ", "め",
 " me",
 " me",
 " me",
 " me"},

{"モ", "も", "mo", "moモ", "も",
 " mo",
 " mo",
 " mo",
 " mo"},

{"ミャ", "みゃ", "mya", "myaミャ", "みゃ",
 " mya",
 " mya",
 " mya",
 " mya"},

{"ミュ", "みゅ", "myu", "myuミュ", "みゅ",
 " myu",
 " myu",
 " myu",
 " myu"},

{"ミョ", "みょ", "myo", "myoミョ", "みょ",
 " myo",
 " myo",
 " myo",
 " myo"},

{"ヤ", "や", "ya", "yaヤ", "や",
 " ya",
 " ya",
 " ya",
 " ya"},

{"ユ", "ゆ", "yu", "yuユ", "ゆ",
 " yu",
 " yu",
 " yu",
 " yu"},

{"ヨ", "よ", "yo", "yoヨ", "よ",
 " yo",
 " yo",
 " yo",
 " yo"},

{"ラ", "ら", "ra", "raラ", "ら",
 " ra",
 " ra",
 " ra",
 " ra"},

{"リ", "り", "ri", "riリ", "り",
 " ri",
 " ri",
 " ri",
 " ri"},

{"ル", "る", "ru", "ruル", "る",
 " ru",
 " ru",
 " ru",
 " ru"},

{"レ", "れ", "re", "reレ", "れ",
 " re",
 " re",
 " re",
 " re"},

{"ロ", "ろ", "ro", "roロ", "ろ",
 " ro",
 " ro",
 " ro",
 " ro"},

{"リャ", "りゃ", "rya", "ryaリャ", "りゃ",
 " rya",
 " rya",
 " rya",
 " rya"},

{"リュ", "りゅ", "ryu", "ryuリュ", "りゅ",
 " ryu",
 " ryu",
 " ryu",
 " ryu"},

{"リョ", "りょ", "ryo", "ryoリョ", "りょ",
 " ryo",
 " ryo",
 " ryo",
 " ryo"},

{"ワ", "わ", "wa", "waワ", "わ",
 " wa",
 " wa",
 " wa",
 " wa"},

{"ヲ", "を", "wo", "woヲ", "を",
 " wo",
 " wo",
 " wo",
 " wo"},

{"ン", "ん", "n", "nン", "ん",
 " n",
 " n",
 " n",
 " n"},
}
