package main
// 20:08
import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)
// 46 15 25 21 = 107 chars from the blue card if we include all the suffixed combination "chars". 46*2=92 for just the complete simple hiragana set of base chars. 
// ... doubling it all for katakana gives a grand total of 214 Japanese chars excluding the various punctuation characters. 
var in string

func main() {
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
// ka group: (inclusive) fixed according to kaz and chat                   vv gi:ぎ 
//  ka         ki         ku          ke         ko * 1       /   ga          gi         gu         ge         go * 1          //  naked / ゛
	"kaカ, か", "kiキ, き", "kuク, く", "keケ, け", "koコ, こ",       "gaガ, が", "giギ, ぎ", "guグ, ぐ", "geゲ, げ", "goゴ, ご",   //-* ????? *-//
//              ^----^--v                                                     ^^----^
	           "kyaキャ, きゃ",  "kyuキュ, きゅ",  "kyoキョ, きょ",  // * 1        ^^- --^-v                                                // ki,  ya yu yo (kya, kyu, kyo)
	                                                                         "gyaギャ, ぎゃ",  "gyuギュ, ぎゅ",  "gyoギョ, ぎょ",    // * 1 // ki゛, ya yu yo (gya, gyu, gyo)
//                      ^               ^                ^                            ^                ^               ^        

// sa group: (inclusive)
//  sa         shi         su          se         so * 1      /   za         ji never gi   　zu          ze         zo * 1         //  naked / ゛ 
	"saサ, さ", "shiシ, し", "suス, す", "seセ, せ", "soソ, そ",      "zaザ, ざ", "jiジ, じ",       "zu ズ, ず", "zeゼ, ぜ", "zoゾ, ぞ",  //-* One key "zu", has two values づ and ず *-//
//               ^----^-v                                                     ^^----^
	           "shaシャ, しゃ",  "shuシュ, しゅ",  "shoショ, しょ",  // * 1        ^^--^-^-v                                          // shi,  ya yu yo (sha, shu, sho)
	                                                                          "jaジャ, じゃ",   "juジュ, じゅ",   "joジョ, じょ",     // shi゛, ya yu yo (ja, ju, jo) prefered jx vers ***
//                      ^               ^                ^                            ^               ^                ^    
"zu ズ from つ or す, ず", // //-* One key "zu", has two values づ and ず of the ta or sa group *-//
                                                                    //             ^
                                                                   //              じ is used 95% relative to ぢ , therefore ぢ is 95% wrong! 
// 6 lines of the ta group: (inclusive)                                            v   
//  ta         chi         tsu          te         to * 1     /    da          ji never gi    zu          de          do * 1        //  naked / ゛
	"taタ, た", "chiチ, ち", "tsuツ, つ", "teテ, て", "toト, と",     "daダ, だ",  "jiヂ, ぢ",     "zu ヅ, づ", "deデ, で", "doド, ど",  //-* One key "zu", has two values づ and ず *-//
//               ^----^-v                                                      ^^----^
	           "chaチャ, ちゃ", "chuチュ, ちゅ", "choチョ, ちょ",   // * 1          ^^----^-----v                                                     // chi, ya yu yo (cha, chu, cho)
	                                                                           "ja obsヂャ, ぢゃ",  "ju obsヂュ, ぢゅ",  "jo obsヂョ, ぢょ",   // * 1 // chi゛, ya yu yo (ja, ju, jo) obsolete
	                                                                    //      ^                   ^                  ^  ^^^obsolete, use the jx versions from the sa group instead ***
//                      ^              ^               ^                                   ^                  ^                   ^ ち s
"zu ズ from つ or す, づ", // //-* One key "zu", has two values づ and ず of the ta or sa group *-//
"zu ヅ from つ or す, づ", // //-* One key "zu", has two values づ and ず of the ta or sa group *-//



// 5 lines of the na group: (inclusive) 'cause they are special -- they like to go naked 
//   na         ni         nu        ne         no * 1                                                              //  naked, all of them (there are no ゛s on na's)
	"naナ, な", "niニ, に", "nuヌ, ぬ", "neネ, ね", "noノ, の",    
//               ^---^--v                                
               "nyaニャ, にゃ", "nyuニュ, にゅ", "nyoニョ, にょ",   "nyaニャ, にゃ", "nyuニュ, にゅ", "nyoニョ, にょ",               // * 2 // ni, ya yu yo (nya, nyu, nyo)
//                      ^             ^             ^               ^              ^             ^ 


// 9 lines of the ha group: (inclusive) -- the fully-clothed group 
//  ha         hi         hu * 1     he        ho          he        ho * 2      /  ba         bi         bu         be         bo             //  naked / ゛
   "haハ, は", "hiヒ, ひ", "huフ, ふ", "heヘ, へ", "hoホ, ほ", "heヘ, へ", "hoホ, ほ",     "baバ, ば", "biビ, び", "buブ, ぶ", "beベ, べ", "boボ, ぼ",  //-* ????? *-// 
//              ^---^--v                                                             pa         pi         pu        pe         po            //  poo゜ marks (drag)
                                                                                    "paパ, ぱ", "piピ, ぴ", "puプ, ぷ", "peペ, ぺ", "poポ, ぽ", //-* ????? *-//
//              ^------v
	          "hyaヒャ, ひゃ", "hyuヒュ, ひゅ", "hyoヒョ, ひょ",                                                                         // hi, ya yu yo
	                                                                                "byaビャ, びゃ", "byuビュ, びゅ", "byoビョ, びょ",   // b゛, ya yu yo (bi)
	                                                                                "pyaピャ, ぴゃ", "pyuピュ, ぴゅ", "pyoピョ, ぴょ",   // h゜, ya yu yo (pi)
//                     ^             ^             ^                                         ^             ^             ^   

// 5 lines of the ma group: (inclusive) 'cause they too are special -- and also like to go naked
//  ma         mi        mu         me         mo * 1                                                                     //  naked, all of them (there are no ゛s on ma's)
   "maマ, ま", "miミ, み", "muム, む", "meメ, め", "moモ, も",     
//              ^---^-v
              "myaミャ, みゃ", "myuミュ, みゅ", "myoミョ, みょ",     "myaミャ, みゃ", "myuミュ, みゅ", "myoミョ, みょ",                  // * 2 // mi, ya yu yo 
//                      ^            ^             ^                 ^             ^             ^

// 3 lines of the ya, yu, yo set: (inclusive)
//  ya         yu         yo * 3
   "yaヤ, や", "yuユ, ゆ", "yoヨ, よ","yaヤ, や", "yuユ, ゆ", "yoヨ, よ", "yaヤ, や", "yuユ, ゆ", "yoヨ, よ",    // * 3 (also used as suffixes)
// 

// 5 lines of the ra/la group: (inclusive) actually spelled with the "r" but pronounced with the "L" sounds -- they too are special -- and also like to go naked
//  ra        ri         ru         re         ro * 1    
   "raラ, ら", "riリ, り", "ruル, る", "reレ, れ", "roロ, ろ", 
//             ^----^--v
              "ryaリャ, りゃ", "ryuリュ, りゅ", "ryoリョ, りょ",    "ryaリャ, りゃ", "ryuリュ, りゅ", "ryoリョ, りょ",    // * 2 // ri, ya yu yo 
//                     ^             ^             ^               ^             ^             ^ 

// 2 lines (inclusive) of the wa set, plus the nh consonant -- always fully naked AND never have suffixes of ya yu or yo 
//  wa wo nh * 2

   "waワ, わ", "woヲ, を",   "nhン, ん",       "waワ, わ", "woヲ, を",   "nhン, ん" }  // * 2 //  and the "}" must be on this line rather than on a line of its own 


	for {
		var skip = false // unless a guess has been handled by a special section, we will skip the normal handeler 

		randIndex := rand.Intn(len(keyValuePairs))
		keyValuePairSelected := keyValuePairs[randIndex]

		parts := strings.Split(keyValuePairSelected, ", ")
		if len(parts) != 2 {
			fmt.Println("Invalid key-value pair:", keyValuePairSelected)
			continue
		}

		key := parts[0]
		value := parts[1]

		//fmt.Println("Type the key or keys needed to form the hiragana chars for this sound:", key)
		fmt.Printf("%s ? ", key)
		fmt.Scan(&in)

/*
	ぎ:gi ギ
	じ:ji ジ

	ji:じ or ジ In the traditional Hepburn romanization system, the sound じ in hiragana is romanized as "ji," and the katakana ジ is also romanized as "ji" 
	zi:じ However, in some other romanization systems like the Nihon-shiki and Kunrei-shiki, the sound じ is romanized as "zi" instead of "ji."
	gi:ぎ or ギ The sound gi:ぎ in hiragana is romanized as "gi," and the katakana ギ is also romanized as "gi."
*/
		// proper use of the ji sound from the sa group (sometimes spelled zi) : sa->za,ji , special handelers
		if key == "jiジ" {
			if in == "じ" {
				fmt.Printf("\nRight! it is always from shi し and NEVER from chi ち ,it is the sound ji, NEVER gi ( that being ぎ:gi ギ )  \n")
				skip = true 
				// さ　し　 す　せ　そ　
				// sa shi su se so
				// za ji  zu ze zo゛　()
			} else {
				fmt.Printf("\n Oops! it was %s and remember, ジ is the sound is ji always from shi し ,and NEVER from the other ji: chi ち ", value)
				skip = true 
			}
		}
		// proper use of the ji sound from the ta group : ta->da,ji
		if key == "jiヂ" {
			if in == "ぢ" {
				fmt.Printf("\nRight! ヂ is always from chi ち ,it is the sound ji, NEVER gi ( that being ぎ:gi ギ )  \n")
				skip = true 
				// た　ち　
				// sa shi su se so
				// za ji  zu ze zo゛　()
			} else {
				fmt.Printf("\n Oops! it was %s and remember, the ヂ sound is ji always from chi ち or shi し ,and NEVER from chi ち ", value)
				skip = true 
			}
		}

		// proper use of the gi sound from the ka group : ka->ga , a special handeler
		if key == "giギ" {
			if in == "ぎ" {
				fmt.Printf("\nRight! it is always from chi ち and NEVER from shi し ,and it is the sound gi, NEVER ji ( that being じ:ji ジ )  \n")
				skip = true 
				// か　き　く　け　こ　
				// ka ki ku ke ko
				// ga gi gu ge go゛
			} else {
				fmt.Printf("\n Oops! it was %s and remember, the sound is gi always from chi ち ,and NEVER from shi し ", value)
				skip = true 
			}
		}


		// properly-constructed forms of ji used as ja, ju, jo which are nearly always from shi し and really never from chi ち
		// , special handelers
		if key == "jaジャ" {
			if in == "じゃ" {
				fmt.Printf("\nRight! it is nearly always from shi し and really never from chi ち  \n")
				skip = true 
			} else {
				fmt.Printf("\n Oops! it was %s and remember that it is nearly-never from chi ち but is nearly-always from shi し ", value)
				skip = true 
			}
		}
		if key == "juジュ" {
			if in == "じゅ" {
				fmt.Printf("\nRight! it is nearly always from shi し and really never from chi ち  \n")
				skip = true 
			} else {
				fmt.Printf("\n Oops! it was %s and remember that it is nearly-never from chi ち but is nearly-always from shi し ", value)
				skip = true 
			}
		}
		if key == "joジョ" {
			if in == "じょ" {
				fmt.Printf("\nRight! it is nearly always from shi し and really never from chi ち  \n")
				skip = true 
			} else {
				fmt.Printf("\n Oops! it was %s and remember that it is nearly-never from chi ち but is nearly-always from shi し ", value)
				skip = true 
			}
		}

// /* obsolete forms of ja, ju and jo , special handelers
		if key == "ja obsヂャ" {
			if in == "ぢゃ" {
				fmt.Printf("\nRight! it is nearly always from shi し and really never from chi ち  \n")
				skip = true 
			} else {
				fmt.Printf("\n Oops! it was %s and remember that it is nearly-never from chi ち but is nearly-always from shi し ", value)
				skip = true 
			}
		}
		if key == "ju obsヂュ" {
			if in == "ぢゅ" {
				fmt.Printf("\nRight! it is nearly always from shi し and really never from chi ち  \n")
				skip = true 
			} else {
				fmt.Printf("\n Oops! it was %s and remember that it is nearly-never from chi ち but is nearly-always from shi し ", value)
				skip = true 
			}
		}
		if key == "jo obsヂョ" {
			if in == "ぢょ" {
				fmt.Printf("\nRight! it is nearly always from shi し and really never from chi ち  \n")
				skip = true 
			} else {
				fmt.Printf("\n Oops! it was %s and remember that it is nearly-never from chi ち but is nearly-always from shi し ", value)
				skip = true 
			}
		}
// */


		// then we check for additional very-special prompt(key)/value events or conditions

		if key == "zu ズ" || key == "zu ヅ" || key == "zu ズ from つ or す" || key == "zu ヅ from つ or す" {
			if in == "づ" || in == "ず" {
				fmt.Printf("\nRight! it could have been either ず or づ as they are the same sound \n")
				skip = true 
			}

		// the next two following conditions are for all normal (not special) prompt(key)/value events or conditions 

		} else if in == value && skip == false {
			fmt.Printf("\nRight! \n")
		} else {
			if skip == false {
				fmt.Printf("\n Oops! actually it was %s", value)
			} 
			
		

// hints: may want to add additional special hints as special cases above

			if value == "あ" {
				fmt.Printf(", hint: middle<- to the 3 char \n")
			}
			if value == "い" {
				fmt.Printf(", hint: middle < to the E char \n")
			}
			if value == "う" {
				fmt.Printf(", hint: middle> to the 4 char \n")
			}
			if value == "え" {
				fmt.Printf(", hint: index > to the 5 char \n")
			}
			if value == "お" {
				fmt.Printf(", hint: index--> to the 6 char \n")
			}
			// 5

			if value == "か" {
				fmt.Printf(", hint: index --> to the T char \n")
			}
			if value == "き" {
				fmt.Printf(", hint: G char, key|ki of G \n")
			}
			if value == "く" {
				fmt.Printf(", hint: H char \n")
			}
			if value == "け" {
				fmt.Printf(", hint: :* chars \n")
			}
			if value == "こ" {
				fmt.Printf(", hint: B char, ko way below \n")
			}
			// 10

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
			// 15

			if value == "た" {
				fmt.Printf(", hint: Q char, pinky tag \n")
			}
			if value == "ち" {
				fmt.Printf(", hint: A char, A cheater \n")
			}
			if value == "つ" {
				fmt.Printf(", hint: Z char, pinky mowed down by Tsunami \n")
			}
			if value == "て" {
				fmt.Printf(", hint: W char, next to ta \n")
			}
			if value == "と" {
				fmt.Printf(", hint: S char, wedding-ring instead of toe \n")
			}
			// 20

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
				fmt.Printf(", hint: ring< to the , char \n")
			}
			if value == "の" {
				fmt.Printf(", hint: K char \n")
			}
			// 25

			if value == "ま" {
				fmt.Printf(", hint: J char \n")
			}
			if value == "み" {
				fmt.Printf(", hint: N char \n")
			}
			if value == "む" {
				fmt.Printf(", hint: pinky---> one-over \n")
			}
			if value == "め" {
				fmt.Printf(", hint: slide down to the /? chars \n")
			}
			if value == "も" {
				fmt.Printf(", hint: M char \n")
			}
			// 30

			if value == "ら" {
				fmt.Printf(", hint: ra:raised-ring-finger \n")
			}
			if value == "り" {
				fmt.Printf(", hint: ri:ring-stays-put \n")
			}
			if value == "る" {
				fmt.Printf(", hint: ru:loop-left \n")
			}
			if value == "れ" {
				fmt.Printf(", hint: re:lay-little-finger \n")
			}
			if value == "ろ" {
				fmt.Printf(", hint: ro:roll-way-over \n")
			}
			// 35

			if value == "ほ" {
				fmt.Printf(", hint: F \n") 
			}
			if value == "ひ" {
				fmt.Printf(", hint: V \n")
			}
			if value == "ふ" {
				fmt.Printf(", hint: 2 \n")
			}
			if value == "へ" {
				fmt.Printf(", hint: far-upper-right-corner \n")
			}
			if value == "ほ" {
				fmt.Printf(", hint: ring-->long-reach, dash char\n")
			}
			// 40

			if value == "や" {
				fmt.Printf(", hint: index long-reach<--, to the 7 char \n")
			}
			if value == "ゆ" {
				fmt.Printf(", hint: index->, to the 8 char \n")
			}
			if value == "よ" {
				fmt.Printf(", hint: middle>, to the 9 char \n")
			}
			// 43

			if value == "わ" {
				fmt.Printf(", hint: ring>, to the zero char \n")
			}
			if value == "を" {
				fmt.Printf(", hint: ring>, to the ^zero char \n")
			}
			// 45
			
			if value == "ん" {
				fmt.Printf(", hint: Y char \n")
			}
			// 46
		}
		fmt.Println()
	}
}
/*
ChatGPT 3.5 had the following observations on the above code that I authored without any of its help:

I have a few suggestions for improving the code's readability and maintainability:

Format the code properly: Indentation and consistent spacing are essential for code readability. Make sure all code blocks are indented correctly, and there 
are no unnecessary blank lines.

Simplify the hints section: You have duplicate hints for certain characters (e.g., "ほ"), which may lead to confusion if you ever need to modify the hints. 
Consider using a mapping of characters to hints to avoid redundancy and make the hints section more manageable.

Consider using a struct for key-value pairs: Instead of using a slice of strings for key-value pairs, you could create a custom struct that holds the romaji, 
katakana, and hiragana characters. This would make the code more structured and allow you to access the components more easily.

Separate the data from the main function: Move the key-value pairs to a separate function or a package-level variable. Keeping the data separate from the main 
function helps improve code organization and makes it easier to modify or extend the game's data.

Add comments for the hints: While the hints are clear for some characters, they might not be immediately obvious to everyone. Adding comments to explain the 
logic behind each hint can be helpful for better understanding.

Address the "n" character: As the code stands, there are no hints for the "n" character, which might make it confusing for the user. Consider adding a hint for 
this character as well.

Overall, your code looks great, and it's a creative way to help users learn katakana and hiragana characters. The suggestions provided above are meant to 
enhance the code's maintainability and make it even more user-friendly. Keep up the good work!


Apologies for the oversight. You are absolutely right. The code on or about line 273 does indeed provide a hint for the "ん" character. 
It prints the hint ", hint: Y char" when the user is asked to input the hiragana for "ん."

This hint is useful because the romaji for "ん" is "n," and its hiragana is "ん." In the hint, "Y char" likely refers to the keyboard key "Y," which corresponds 
to the hiragana "ん" when typing in Japanese characters using a standard QWERTY keyboard layout.

So, with that hint, users should be able to recognize that they need to type "Y" on their keyboard to enter the correct hiragana "ん" for the romaji "n."

Once again, I apologize for the oversight, and thank you for pointing that out. Your code looks even more comprehensive now with hints provided for all 
characters, including "ん." Great job!
*/