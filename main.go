package main

import (
	"fmt"
	"math/rand"
	"time"
)

type charSetStruct struct {
	KeyK  string
	KeyH  string
	KeyR  string
	KeyRK string

	Value string

	Hint1k string
	Hint2k string

	Hint1h string
	Hint2h string
}

/*
func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		var skip = false // unless a guess has been handled by a special section, we will skip the normal handeler

		randIndex := rand.Intn(len(fileOfCards))
		aCard := fileOfCards[randIndex]
		key := aCard.KeyRK
		value := aCard.Value

		fmt.Printf("%s ? ", key)
		fmt.Scan(&in)
*/

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
	doingContextSH = false
	ContextSHskip = false
	// meat of the Katakana practice algorithm: ---------------------------------------------------------------------------
	for {
		// if we got here not via an interim handling of ? Context Sensitive Help
		//  v v v v v          v v
		if doingContextSH == false && ContextSHskip == false {
			randIndex := rand.Intn(len(fileOfCards))
			aCard := fileOfCards[randIndex]
			key := aCard.KeyRK
			value := aCard.Value
			if value == "rick" { // kluge ??????????????????????
				fmt.Printf("woo")
			}
			fmt.Printf("%s", key)
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
			giveAhintKata(value)
			fmt.Printf("%s", key)
			fmt.Printf("%s", string(colorCyan))
			fmt.Printf(" ? ")
			fmt.Printf("%s", string(colorReset))
			doingContextSH = false
			ContextSHskip = true
			//continue // ?????????????????????????????????????
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

func giveAhintKata(value string) {

	// fix all these hints so that they only refer to the value one needs to type that corresponds to the katakana prompt (to the key)

	if value == "あ" { // a ア, maybe a grotesque A
		fmt.Printf(", hint: middle<- to the 3 char, looks nothing-like the hiragana a, but a lot like a te ア, あ, fuck mae! \n")
	} else if value == "い" { // i イ, shift the two lines of the hiragana
		fmt.Printf(", hint: middle < to the E char, イ looks more like a hiragana te, but at least it is still two mostly-vertical lines イ　\n")
	} else if value == "う" { // u ウ, um-kay
		fmt.Printf(", hint: middle> to the 4 char, u う ウ　is ok, having had to look for angles \n")
	} else if value == "え" { // e エ, eye see it as a ... an, eye
		fmt.Printf(", hint: index > to the 5 char, it エ　does have a vague, angular resemblance  \n")
	} else if value == "お" { // o オ, on-the-go maybe
		fmt.Printf(", hint: index--> to the 6 char, オ　has only a vague resemblance, albeit with less curves \n")
	} else if value == "か" { // ka カ
		fmt.Printf(", hint: index --> to the T char \n")
	} else if value == "き" { // ki キ
		fmt.Printf(", hint: L-index > to the G char \n")
	} else if value == "く" { // ku ク, compare to ta タ, and ke ケ
		fmt.Printf(", hint: R-index <- to the H char, ク　no, just no. Starting with one angle, they settled for this? \n")
	} else if value == "け" { // ke ケ, compare to ku ク, and ta タ
		fmt.Printf(", hint: pinky -> one over to the :* chars, ケ, bits of it are there, just as many curves though \n")
	} else if value == "こ" { // ko コ, compare to ni ニ
		fmt.Printf(", hint: index <--- to the B char, alright, コ　it makes sense, 'cause angles \n")
	} else if value == "さ" { // sa サ, if you 'se' so sa
		fmt.Printf(", hint: ring >v to the X char, at least the sa goes to the left, and it looks a lot like se, I'd say \n")
	} else if value == "し" { // shi シ, she sleeps and snores, compare to tsu ツ, and so ソ,
		fmt.Printf(", hint: middle on the D char, シ she is sleeping/snoring here, but has no excuse to2 look like this (no angles here, less curve though) \n")
	} else if value == "す" { // su ス, sue is running
		fmt.Printf(", hint: index < to the R char, they were looking for angles (sue is running ス) \n")
	} else if value == "せ" { // se セ
		fmt.Printf(", hint: pinky to the P char \n")
	} else if value == "そ" { // so ソ, compare to ん ン which lays down more and has a more laid-down back-tick
		fmt.Printf(", hint: index <-- to the C char, ソ Looks like しシ or no ノ, or ン starting with all angles - `had to `backtrack?  \n")
	} else if value == "た" { // ta タ, compare to ku ク and ke ケ
		fmt.Printf(", hint: pinky < to the Q char, the top left looks like ta, more at least than ku ク, or he ケ \n")
	} else if value == "ち" { // chi チ, compare to te テ
		fmt.Printf(", hint: pinky on the A char, 'some' resemblance if we see the middle line as the top of the backwards c \n")
	} else if value == "つ" { // tsu ツ, compare to shi シ, she who sleeps low and snores loudly
		fmt.Printf(", hint: pinky to the Z char, understandable?, if we are generous. Very generous..゛as in adding two ツ  shi シ said \n")
	} else if value == "て" { // te テ, compare to chi チ
		fmt.Printf(", hint: ring < to the W char, te went the wrong way, and gained a flat hat \n")
	} else if value == "と" { // to ト
		fmt.Printf(", hint: ring on the S char, toe is flipped-out. Kicked in the balls, on its head \n")
	} else if value == "な" { // na ナ, compare to me メ, and nu ヌ
		fmt.Printf(", hint: index < to the U char, is simple, very simple \n")
	} else if value == "に" { // ni ニ
		fmt.Printf(", hint: middle < to the I char \n")
	} else if value == "ぬ" { // nu ヌ, compare to me メ, and na ナ
		fmt.Printf(", hint: R-ring<--- to the 1 char, at least two lines cross, nu flew  \n")
	} else if value == "ね" { // ne ネ
		fmt.Printf(", hint: L-ring <v to the ,< chars, a hoe that got nailed \n")
	} else if value == "の" { // no ノ, no extra marks needed on this classic no symbol
		fmt.Printf(", hint: middle on the k char, drop the curve, and I'll take it \n")
	} else if value == "ま" { // ma マ, mama is unique, see the breast pump?
		fmt.Printf(", hint: index on the J char, (a breast pump maybe?) \n")
	} else if value == "み" { // mi ミ, and two makes three
		fmt.Printf(", hint: index <v to the N char, me thinks 3 \n")
	} else if value == "む" { // mu ム, an even-fatter-jawed moo-cow
		fmt.Printf(", hint: pinky two over ---> to the ]} chars, now it's an even-fatter-jawed moo-cow  \n")
	} else if value == "め" { // me メ
		fmt.Printf(", hint: pinky -> slide down to the /? chars, cross off the mess. But that's nothing nu ヌ \n")
	} else if value == "も" { // mo モ
		fmt.Printf(", hint: middle <v to the M char, way to hang to the right; you go Joe! \n")
	} else if value == "ら" { // ra ラ
		fmt.Printf(", hint: ring ^ to the o char, Similar ラ, ら \n")
	} else if value == "り" { // ri リ
		fmt.Printf(", hint: ring on the L char, longer on the right ring \n")
	} else if value == "る" { // ru ル, is two
		fmt.Printf(", hint: pinky < to the .> chars, ル, る, ru is now two? \n")
	} else if value == "れ" { // re レ, more, or less, a ru ル - しshe-said
		fmt.Printf(", hint: pinky on the ;+ chars, レ, れ; しshe-it \n")
	} else if value == "ろ" { // ro ロ
		fmt.Printf(", hint: pinky ---> long slide to the _ and backslash chars, maybe it had always been the way to go to do a square O for ro? \n")
	} else if value == "ほ" { // ha ハ
		fmt.Printf(", hint: L-ring---> to the \"-\" char\n")
	} else if value == "ひ" { // hi ヒ
		fmt.Printf(", hint: index > to the V char \n")
	} else if value == "ふ" { // hu フ, squinting it is a ふフ
		fmt.Printf(", hint: ring<, to the 2 char, if we squint? フ, ふ\n")
	} else if value == "へ" { // he ヘ
		fmt.Printf(", hint: ;+ char \n")
	} else if value == "ほ" { // ho ホ, now with a dress and all
		fmt.Printf(", hint: ring ---> to the - char, now looks like a hoe \n")
	} else if value == "や" { // ya ヤ
		fmt.Printf(", hint: index <-- to the 7 char \n")
	} else if value == "ゆ" { // yu ユ
		fmt.Printf(", hint: index->, to the 8 char, yu looks like ユ \n")
	} else if value == "よ" { // yo ヨ
		fmt.Printf(", hint: middle to the 9 char, tripple yo ヨ \n")
	} else if value == "わ" { // wa ワ, compare to wo ヲ
		fmt.Printf(", hint: ring > to the 0 char, so now it's a water fall for wa? \n")
	} else if value == "を" { // wo ヲ, compare to wa ワ
		fmt.Printf(", hint: ring > shifted 0 char, NOW, it looks like wa. Now! For fucks sake, Now? \n")
	} else if value == "ん" { // nh ン, compare to so ソ
		fmt.Printf(", hint: Y char, pointing at the one last remaining bent stroke of the hiragana char \n")
	} else {
		fmt.Printf("\n")
	}
} // end of func giveAhintKata()

func restartKata() (returnToMenuRS bool) {

	rand.Seed(time.Now().UnixNano())

	randIndex := rand.Intn(len(fileOfCards))
	aCard := fileOfCards[randIndex]

	key := aCard.KeyRK
	value := aCard.Value
	if value == "rick" {
		fmt.Printf("woo")
	}

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
		} else if value == "い" { // i イ, shift the two lines of the hiragana
			fmt.Printf(", hint: middle < to the E char, イ looks more like a hiragana te, but at least it is still two mostly-vertical lines イ　\n")
		} else if value == "う" { // u ウ, um-kay
			fmt.Printf(", hint: middle> to the 4 char, u ウ　is ok, having had to look for angles \n")
		} else if value == "え" { // e エ, eye see it as a ... an, eye
			fmt.Printf(", hint: index > to the 5 char, it エ　does have a vague, angular resemblance  \n")
		} else if value == "お" { // o オ, on-the-go maybe
			fmt.Printf(", hint: index--> to the 6 char, オ　has only a vague resemblance, albeit with less curves \n")
		} else if value == "か" { // ka カ
			fmt.Printf(", hint: index --> to the T char \n")
		} else if value == "き" { // ki キ
			fmt.Printf(", hint: L-index > to the G char \n")
		} else if value == "く" { // ku ク, compare to ta タ, and ke ケ
			fmt.Printf(", hint: R-index <- to the H char, ク　no, just no. Starting with one angle, they settled for this? \n")
		} else if value == "け" { // ke ケ, compare to ku ク, and ta タ
			fmt.Printf(", hint: pinky -> one over to the :* chars, ケ, bits of it are there, just as many curves though \n")
		} else if value == "こ" { // ko コ, compare to ni ニ
			fmt.Printf(", hint: index <--- to the B char, alright, コ　it makes sense, 'cause angles \n")
		} else if value == "さ" { // sa サ, if you 'se' so sa
			fmt.Printf(", hint: ring >v to the X char, at least the sa goes to the left, and it looks a lot like se, I'd say \n")
		} else if value == "し" { // shi シ, she sleeps and snores, compare to tsu ツ, and so ソ,
			fmt.Printf(", hint: middle on the D char, シ she is sleeping/snoring here, but has no excuse to2 look like this (no angles here, less curve though) \n")
		} else if value == "す" { // su ス, sue is running
			fmt.Printf(", hint: index < to the R char, they were looking for angles (sue is running ス) \n")
		} else if value == "せ" { // se セ
			fmt.Printf(", hint: pinky to the P char \n")
		} else if value == "そ" { // so ソ, compare to ん ン which lays down more and has a more laid-down back-tick
			fmt.Printf(", hint: index <-- to the C char, ソ Looks like しシ or no ノ, or ン starting with all angles - `had to `backtrack?  \n")
		} else if value == "た" { // ta タ, compare to ku ク and ke ケ
			fmt.Printf(", hint: pinky < to the Q char, the top left looks like ta, more at least than ku ク, or he ケ \n")
		} else if value == "ち" { // chi チ, compare to te テ
			fmt.Printf(", hint: pinky on the A char, 'some' resemblance if we see the middle line as the top of the backwards c \n")
		} else if value == "つ" { // tsu ツ, compare to shi シ, she who sleeps low and snores loudly
			fmt.Printf(", hint: pinky to the Z char, understandable?, if we are generous. Very generous..゛as in adding two ツ  shi シ said \n")
		} else if value == "て" { // te テ, compare to chi チ
			fmt.Printf(", hint: ring < to the W char, te went the wrong way, and gained a flat hat \n")
		} else if value == "と" { // to ト
			fmt.Printf(", hint: ring on the S char, toe is flipped-out. Kicked in the balls, on its head \n")
		} else if value == "な" { // na ナ, compare to me メ, and nu ヌ
			fmt.Printf(", hint: index < to the U char, is simple, very simple \n")
		} else if value == "に" { // ni ニ
			fmt.Printf(", hint: middle < to the I char \n")
		} else if value == "ぬ" { // nu ヌ, compare to me メ, and na ナ
			fmt.Printf(", hint: R-ring<--- to the 1 char, at least two lines cross, nu flew  \n")
		} else if value == "ね" { // ne ネ
			fmt.Printf(", hint: L-ring <v to the ,< chars, a hoe that got nailed \n")
		} else if value == "の" { // no ノ, no extra marks needed on this classic no symbol
			fmt.Printf(", hint: middle on the k char, drop the curve, and I'll take it \n")
		} else if value == "ま" { // ma マ, mama is unique, see the breast pump?
			fmt.Printf(", hint: index on the J char, (a breast pump maybe?) \n")
		} else if value == "み" { // mi ミ, and two makes three
			fmt.Printf(", hint: index <v to the N char, me thinks 3 \n")
		} else if value == "む" { // mu ム, an even-fatter-jawed moo-cow
			fmt.Printf(", hint: pinky two over ---> to the ]} chars, now it's an even-fatter-jawed moo-cow  \n")
		} else if value == "め" { // me メ
			fmt.Printf(", hint: pinky -> slide down to the /? chars, cross off the mess. But that's nothing nu ヌ \n")
		} else if value == "も" { // mo モ
			fmt.Printf(", hint: middle <v to the M char, way to hang to the right; you go Joe! \n")
		} else if value == "ら" { // ra ラ
			fmt.Printf(", hint: ring ^ to the o char, Similar ラ, ら \n")
		} else if value == "り" { // ri リ
			fmt.Printf(", hint: ring on the L char, longer on the right ring \n")
		} else if value == "る" { // ru ル, is two
			fmt.Printf(", hint: pinky < to the .> chars, ル, る, ru is now two? \n")
		} else if value == "れ" { // re レ, more, or less, a ru ル - しshe-said
			fmt.Printf(", hint: pinky on the ;+ chars, レ, れ; しshe-it \n")
		} else if value == "ろ" { // ro ロ
			fmt.Printf(", hint: pinky ---> long slide to the _ and backslash chars, maybe it had always been the way to go to do a square O for ro? \n")
		} else if value == "ほ" { // ha ハ
			fmt.Printf(", hint: L-ring---> to the \"-\" char\n")
		} else if value == "ひ" { // hi ヒ
			fmt.Printf(", hint: index > to the V char \n")
		} else if value == "ふ" { // hu フ, squinting it is a ふフ
			fmt.Printf(", hint: ring<, to the 2 char, if we squint? フ, ふ\n")
		} else if value == "へ" { // he ヘ
			fmt.Printf(", hint: ;+ char \n")
		} else if value == "ほ" { // ho ホ, now with a dress and all
			fmt.Printf(", hint: ring ---> to the - char, now looks like a hoe \n")
		} else if value == "や" { // ya ヤ
			fmt.Printf(", hint: index <-- to the 7 char \n")
		} else if value == "ゆ" { // yu ユ
			fmt.Printf(", hint: index->, to the 8 char, yu looks like ユ \n")
		} else if value == "よ" { // yo ヨ
			fmt.Printf(", hint: middle to the 9 char, tripple yo ヨ \n")
		} else if value == "わ" { // wa ワ, compare to wo ヲ
			fmt.Printf(", hint: ring > to the 0 char, so now it's a water fall for wa? \n")
		} else if value == "を" { // wo ヲ, compare to wa ワ
			fmt.Printf(", hint: ring > shifted 0 char, NOW, it looks like wa. Now! For fucks sake, Now? \n")
		} else if value == "ん" { // nh ン, compare to so ソ
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
	rand.Seed(time.Now().UnixNano())

	doingContextSH = false
	ContextSHskip = false
	// meat of the Hiragana practice algorithm: ---------------------------------------------------------------------------
	for {
		// if we got here not via an interim handling of ? Context Sensitive Help
		//  v v v v v          v v
		if doingContextSH == false && ContextSHskip == false {

			randIndex := rand.Intn(len(fileOfCards))
			aCard := fileOfCards[randIndex]

			key := aCard.KeyRK
			value := aCard.Value

			fmt.Printf("%s", key)
			fmt.Printf("%s", string(colorCyan))
			fmt.Printf(" ? ")
			fmt.Printf("%s", string(colorReset))

			fmt.Scan(&in)

			if in == "menu" {
				return
				// returns to main menu
			}
			if value == "rick" {
				fmt.Printf("woo")
			}
			if in == "?" {
				doingContextSH = true
				// will do a hiraMeatB() below in the parent if's , else if , section
			} else {
				doingContextSH = false
				ContextSHskip = false
				kataMeatB() // probably no need to capture the return val here because ??????????????????????????
			}

			// else if -- if this is not our first rodeo -- re-prompt after the handling of ?
		} else if ContextSHskip == false {
			giveAhintHira(value)
			fmt.Printf("%s", key)
			fmt.Printf("%s", string(colorCyan))
			fmt.Printf(" ? ")
			fmt.Printf("%s", string(colorReset))
			doingContextSH = false
			ContextSHskip = true
			//continue // ????????????????????????????????????
			fmt.Scan(&in)

			redoMenuKMB := hiraMeatB()
			if redoMenuKMB == true {
				return
				// returns to main menu
			}

			// restart once in a func that does what the above if does and then sets ContextSHskip = false so that a future for iteration will run as if no ? had occured
			//             v v v v v v
			redoMenuRS := restartHira()
			if redoMenuRS == true {
				return
				// returns to main menu
			}
		}

	} // end for
} // end func hira

func giveAhintHira(value string) {
	/*
		ぎ:gi ギ
		じ:ji ジ

		ji:じ or ジ In the traditional Hepburn romanization system, the sound じ in hiragana is romanized as "ji," and the katakana ジ is also romanized as "ji"
		zi:じ However, in some other romanization systems like the Nihon-shiki and Kunrei-shiki, the sound じ is romanized as "zi" instead of "ji."
		gi:ぎ or ギ The sound gi:ぎ in hiragana is romanized as "gi," and the katakana ギ is also romanized as "gi."
	*/
	// proper use of the ji sound from the sa group (sometimes spelled zi) : sa->za,ji , special key handlers
	if key == "jiジ" {
		if in == "じ" {
			fmt.Printf("%s", string(colorGreen))
			fmt.Printf("\n      ^Right! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("ジ is always from shi し and NEVER from chi ち ,it is the sound ji, NEVER gi ( that being ぎ:gi ギ )  \n")
			// ContextSHskip = needs to replace all skip s as in skip = true
			ContextSHskip = true
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
			ContextSHskip = true
		}
	}
	// proper use of the ji sound from the ta group : ta->da,ji , a special key handeler
	if key == "jiヂ" {
		if in == "ぢ" {
			fmt.Printf("%s", string(colorGreen))
			fmt.Printf("\n      ^Right! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("ヂ is always from chi ち ,it is the sound ji, NEVER gi ( that being ぎ:gi ギ )  \n")
			ContextSHskip = true
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
			ContextSHskip = true
		}
	}

	// proper use of the gi sound from the ka group : ka->ga , a special key handeler
	if key == "giギ" {
		if in == "ぎ" {
			fmt.Printf("%s", string(colorGreen))
			fmt.Printf("\n      ^Right! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it is always from ki き and NEVER from shi し ,and it is the sound gi, NEVER ji ( that being じ:ji ジ )  \n")
			ContextSHskip = true
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
			ContextSHskip = true
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
			ContextSHskip = true
		} else {
			fmt.Printf("%s", string(colorRed))
			fmt.Printf("      ^Oops! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it was ")
			fmt.Printf("%s", string(colorCyan))
			fmt.Printf("%s", value)
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("and remember that it is nearly-never from chi ち but is nearly-always from shi し ")
			ContextSHskip = true
		}
	}
	if key == "juジュ" {
		if in == "じゅ" {
			fmt.Printf("%s", string(colorGreen))
			fmt.Printf("\n      ^Right! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it is nearly always from shi し and really never from chi ち  \n")
			ContextSHskip = true
		} else {
			fmt.Printf("%s", string(colorRed))
			fmt.Printf("      ^Oops! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it was ")
			fmt.Printf("%s", string(colorCyan))
			fmt.Printf("%s", value)
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("and remember that it is nearly-never from chi ち but is nearly-always from shi し ")
			ContextSHskip = true
		}
	}
	if key == "joジョ" {
		if in == "じょ" {
			fmt.Printf("%s", string(colorGreen))
			fmt.Printf("\n      ^Right! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it is nearly always from shi し and really never from chi ち  \n")
			ContextSHskip = true
		} else {
			fmt.Printf("%s", string(colorRed))
			fmt.Printf("      ^Oops! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it was ")
			fmt.Printf("%s", string(colorCyan))
			fmt.Printf("%s", value)
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("and remember that it is nearly-never from chi ち but is nearly-always from shi し ")
			ContextSHskip = true
		}
	}

	// /* obsolete forms of ja, ju and jo , special key handelers
	if key == "ja obsヂャ" {
		if in == "ぢゃ" {
			fmt.Printf("%s", string(colorGreen))
			fmt.Printf("\n      ^Right! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it is nearly always from shi し and really never from chi ち  \n")
			ContextSHskip = true
		} else {
			fmt.Printf("%s", string(colorRed))
			fmt.Printf("      ^Oops! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it was ")
			fmt.Printf("%s", string(colorCyan))
			fmt.Printf("%s", value)
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("and remember that it is nearly-never from chi ち but is nearly-always from shi し ")
			ContextSHskip = true
		}
	}
	if key == "ju obsヂュ" {
		if in == "ぢゅ" {
			fmt.Printf("%s", string(colorGreen))
			fmt.Printf("\n      ^Right! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it is nearly always from shi し and really never from chi ち  \n")
			ContextSHskip = true
		} else {
			fmt.Printf("%s", string(colorRed))
			fmt.Printf("      ^Oops! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it was ")
			fmt.Printf("%s", string(colorCyan))
			fmt.Printf("%s", value)
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("and remember that it is nearly-never from chi ち but is nearly-always from shi し ")
			ContextSHskip = true
		}
	}
	if key == "jo obsヂョ" {
		if in == "ぢょ" {
			fmt.Printf("%s", string(colorGreen))
			fmt.Printf("\n      ^Right! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it is nearly always from shi し and really never from chi ち  \n")
			ContextSHskip = true
		} else {
			fmt.Printf("%s", string(colorRed))
			fmt.Printf("      ^Oops! ")
			fmt.Printf("%s", string(colorReset))

			fmt.Printf("it was ")
			fmt.Printf("%s", string(colorCyan))
			fmt.Printf("%s", value)
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("and remember that it is nearly-never from chi ち but is nearly-always from shi し ")
			ContextSHskip = true
		}
	}
	//

	// then we check for additional very-special prompt(key)/value events or conditions,  , a very-special key handeler

	if key == "zu ズ" || key == "zu ヅ" || key == "zu ズ from つ or す" || key == "zu ヅ from つ or す" {
		if in == "づ" || in == "ず" {
			fmt.Printf("%s", string(colorGreen))
			fmt.Printf("\n      ^Right! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it could have been either ず or づ as they are the same sound \n")
			ContextSHskip = true
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

		//
		//		} else if in == value && skip == false {
		//			fmt.Printf("\n      ^Right! \n")
		//		} else {
		//			if skip == false {
		//				fmt.Printf("      ^Oops! actually it was %s", value)
		//			}
		//

	} else if in == value && ContextSHskip == false {
		fmt.Printf("%s", string(colorGreen))
		fmt.Printf("      ^Right! \n")
		fmt.Printf("%s", string(colorReset))
	} else {
		if ContextSHskip == false {
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
		} else if value == "い" {
			fmt.Printf(", hint: middle < to the E char \n")
		} else if value == "う" {
			fmt.Printf(", hint: middle> to the 4 char \n")
		} else if value == "え" {
			fmt.Printf(", hint: index > to the 5 char \n")
		} else if value == "お" {
			fmt.Printf(", hint: index--> to the 6 char \n")
		} else
		// 5

		if value == "か" {
			fmt.Printf(", hint: index --> to the T char \n")
		} else if value == "き" {
			fmt.Printf(", hint: G char, key|ki of G \n")
		} else if value == "く" {
			fmt.Printf(", hint: H char \n")
		} else if value == "け" {
			fmt.Printf(", hint: :* chars \n")
		} else if value == "こ" {
			fmt.Printf(", hint: B char, ko way below \n")
		} else
		// 10

		if value == "さ" {
			fmt.Printf(", hint: X char \n")
		} else if value == "し" {
			fmt.Printf(", hint: D char \n")
		} else if value == "す" {
			fmt.Printf(", hint: R char \n")
		} else if value == "せ" {
			fmt.Printf(", hint: P char \n")
		} else if value == "そ" {
			fmt.Printf(", hint: C char \n")
		} else
		// 15

		if value == "た" {
			fmt.Printf(", hint: Q char, pinky tag \n")
		} else if value == "ち" {
			fmt.Printf(", hint: A char, A cheater \n")
		} else if value == "つ" {
			fmt.Printf(", hint: Z char, pinky mowed down by Tsunami \n")
		} else if value == "て" {
			fmt.Printf(", hint: W char, next to ta \n")
		} else if value == "と" {
			fmt.Printf(", hint: S char, wedding-ring instead of toe \n")
		} else
		// 20

		if value == "な" {
			fmt.Printf(", hint: U char \n")
		} else if value == "に" {
			fmt.Printf(", hint: I char \n")
		} else if value == "ぬ" {
			fmt.Printf(", hint: ring<---, 1 char \n")
		} else if value == "ね" {
			fmt.Printf(", hint: ring< to the , char \n")
		} else if value == "の" {
			fmt.Printf(", hint: K char \n")
		} else
		// 25

		if value == "ま" {
			fmt.Printf(", hint: J char \n")
		} else if value == "み" {
			fmt.Printf(", hint: N char \n")
		} else if value == "む" {
			fmt.Printf(", hint: pinky---> one-over \n")
		} else if value == "め" {
			fmt.Printf(", hint: slide down to the /? chars \n")
		} else if value == "も" {
			fmt.Printf(", hint: M char \n")
		} else
		// 30

		if value == "ら" {
			fmt.Printf(", hint: ra:raised-ring-finger \n")
		} else if value == "り" {
			fmt.Printf(", hint: ri:ring-stays-put \n")
		} else if value == "る" {
			fmt.Printf(", hint: ru:loop-left \n")
		} else if value == "れ" {
			fmt.Printf(", hint: re:lay-little-finger \n")
		} else if value == "ろ" {
			fmt.Printf(", hint: ro:roll-way-over \n")
		} else
		// 35

		if value == "ほ" {
			fmt.Printf(", hint: F \n")
		}
		if value == "ひ" {
			fmt.Printf(", hint: V \n")
		} else if value == "ふ" {
			fmt.Printf(", hint: 2 \n")
		} else if value == "へ" {
			fmt.Printf(", hint: far-upper-right-corner \n")
		} else if value == "ほ" {
			fmt.Printf(", hint: ring-->long-reach, dash char\n")
		} else
		// 40

		if value == "や" {
			fmt.Printf(", hint: index long-reach<--, to the 7 char \n")
		} else if value == "ゆ" {
			fmt.Printf(", hint: index->, to the 8 char \n")
		} else if value == "よ" {
			fmt.Printf(", hint: middle>, to the 9 char \n")
		} else
		// 43

		if value == "わ" {
			fmt.Printf(", hint: ring>, to the zero char \n")
		} else if value == "を" {
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

} // end of func giveAhintHira() (value string)

func hiraMeatB() (returnToMenuKMB bool) {

	if in == "menu" {
		returnToMenuKMB = true
		return
	}

	// good, but we still get an unwanted Ooops instead of the wanted hint
	/*

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
	   		}
	   		// this } added just to try it ?????????????????????????????????

	   // hira hints: (copied from kataMeatB hints) may want to add additional special hints as special cases ??????????????-above-???????????????????


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
			ContextSHskip = true
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
			ContextSHskip = true
		}
	}
	// proper use of the ji sound from the ta group : ta->da,ji , a special key handeler
	if key == "jiヂ" {
		if in == "ぢ" {
			fmt.Printf("%s", string(colorGreen))
			fmt.Printf("\n      ^Right! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("ヂ is always from chi ち ,it is the sound ji, NEVER gi ( that being ぎ:gi ギ )  \n")
			ContextSHskip = true
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
			ContextSHskip = true
		}
	}

	// proper use of the gi sound from the ka group : ka->ga , a special key handeler
	if key == "giギ" {
		if in == "ぎ" {
			fmt.Printf("%s", string(colorGreen))
			fmt.Printf("\n      ^Right! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it is always from ki き and NEVER from shi し ,and it is the sound gi, NEVER ji ( that being じ:ji ジ )  \n")
			ContextSHskip = true
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
			ContextSHskip = true
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
			ContextSHskip = true
		} else {
			fmt.Printf("%s", string(colorRed))
			fmt.Printf("      ^Oops! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it was ")
			fmt.Printf("%s", string(colorCyan))
			fmt.Printf("%s", value)
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("and remember that it is nearly-never from chi ち but is nearly-always from shi し ")
			ContextSHskip = true
		}
	}
	if key == "juジュ" {
		if in == "じゅ" {
			fmt.Printf("%s", string(colorGreen))
			fmt.Printf("\n      ^Right! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it is nearly always from shi し and really never from chi ち  \n")
			ContextSHskip = true
		} else {
			fmt.Printf("%s", string(colorRed))
			fmt.Printf("      ^Oops! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it was ")
			fmt.Printf("%s", string(colorCyan))
			fmt.Printf("%s", value)
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("and remember that it is nearly-never from chi ち but is nearly-always from shi し ")
			ContextSHskip = true
		}
	}
	if key == "joジョ" {
		if in == "じょ" {
			fmt.Printf("%s", string(colorGreen))
			fmt.Printf("\n      ^Right! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it is nearly always from shi し and really never from chi ち  \n")
			ContextSHskip = true
		} else {
			fmt.Printf("%s", string(colorRed))
			fmt.Printf("      ^Oops! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it was ")
			fmt.Printf("%s", string(colorCyan))
			fmt.Printf("%s", value)
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("and remember that it is nearly-never from chi ち but is nearly-always from shi し ")
			ContextSHskip = true
		}
	}

	// /* obsolete forms of ja, ju and jo , special key handelers
	if key == "ja obsヂャ" {
		if in == "ぢゃ" {
			fmt.Printf("%s", string(colorGreen))
			fmt.Printf("\n      ^Right! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it is nearly always from shi し and really never from chi ち  \n")
			ContextSHskip = true
		} else {
			fmt.Printf("%s", string(colorRed))
			fmt.Printf("      ^Oops! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it was ")
			fmt.Printf("%s", string(colorCyan))
			fmt.Printf("%s", value)
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("and remember that it is nearly-never from chi ち but is nearly-always from shi し ")
			ContextSHskip = true
		}
	}
	if key == "ju obsヂュ" {
		if in == "ぢゅ" {
			fmt.Printf("%s", string(colorGreen))
			fmt.Printf("\n      ^Right! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it is nearly always from shi し and really never from chi ち  \n")
			ContextSHskip = true
		} else {
			fmt.Printf("%s", string(colorRed))
			fmt.Printf("      ^Oops! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it was ")
			fmt.Printf("%s", string(colorCyan))
			fmt.Printf("%s", value)
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("and remember that it is nearly-never from chi ち but is nearly-always from shi し ")
			ContextSHskip = true
		}
	}
	if key == "jo obsヂョ" {
		if in == "ぢょ" {
			fmt.Printf("%s", string(colorGreen))
			fmt.Printf("\n      ^Right! ")
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("it is nearly always from shi し and really never from chi ち  \n")
			ContextSHskip = true
		} else {
			fmt.Printf("%s", string(colorRed))
			fmt.Printf("      ^Oops! ")
			fmt.Printf("%s", string(colorReset))

			fmt.Printf("it was ")
			fmt.Printf("%s", string(colorCyan))
			fmt.Printf("%s", value)
			fmt.Printf("%s", string(colorReset))
			fmt.Printf("and remember that it is nearly-never from chi ち but is nearly-always from shi し ")
			ContextSHskip = true
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
			ContextSHskip = true
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

	} else {
		// this } also added just to try it , bad idea ????  at least without a } else  to follow it  ???????????????????????????

		// the next two following conditions are for all normal (not special) prompt(key)/value events or conditions

		/*
		   		} else if in == value && skip == false {
		   			fmt.Printf("\n      ^Right! \n")
		   		} else {
		   			if skip == false {
		   				fmt.Printf("      ^Oops! actually it was %s", value)
		   			}
		   ///////////////////////

		   		} else if in == value &&ContextSHskip == false {
		   			fmt.Printf("%s", string(colorGreen))
		   			fmt.Printf("      ^Right! \n")
		   			fmt.Printf("%s", string(colorReset))
		   		} else {
		   			ifContextSHskip == false {
		   				fmt.Printf("%s", string(colorRed))
		   				fmt.Printf("      ^Oops! ")
		   				fmt.Printf("%s", string(colorReset))

		   				fmt.Printf("it was: ")
		   				fmt.Printf("%s", string(colorCyan))
		   				fmt.Printf("%s", value)
		   				fmt.Printf("%s", string(colorReset))  // ???????????????????????? above 1160 and 61 70 ????????????? and below ????????????????????????
		   //			} else

		*/

		// seeing if removeing this section will eliminate the unwanted Ooops
		/*
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
		*/

		// hints for katakana (key) prompts: (values are in hiragana)
		// hints only are given for values that are naked (lack ゛゜or suffixes of や, ゆ, or よ)

		if value == "あ" { // a ア, maybe a grotesque A
			fmt.Printf(", hint: middle<- to the 3 char, looks nothing-like the hiragana a, but a lot like a te ア, あ, fuck mae! \n")
		} else if value == "い" { // i イ, shift the two lines of the hiragana
			fmt.Printf(", hint: middle < to the E char, イ looks more like a hiragana te, but at least it is still two mostly-vertical lines イ　\n")
		} else if value == "う" { // u ウ, um-kay
			fmt.Printf(", hint: middle> to the 4 char, u ウ　is ok, having had to look for angles \n")
		} else if value == "え" { // e エ, eye see it as a ... an, eye
			fmt.Printf(", hint: index > to the 5 char, it エ　does have a vague, angular resemblance  \n")
		} else if value == "お" { // o オ, on-the-go maybe
			fmt.Printf(", hint: index--> to the 6 char, オ　has only a vague resemblance, albeit with less curves \n")
		} else if value == "か" { // ka カ
			fmt.Printf(", hint: index --> to the T char \n")
		} else if value == "き" { // ki キ
			fmt.Printf(", hint: L-index > to the G char \n")
		} else if value == "く" { // ku ク, compare to ta タ, and ke ケ
			fmt.Printf(", hint: R-index <- to the H char, ク　no, just no. Starting with one angle, they settled for this? \n")
		} else if value == "け" { // ke ケ, compare to ku ク, and ta タ
			fmt.Printf(", hint: pinky -> one over to the :* chars, ケ, bits of it are there, just as many curves though \n")
		} else if value == "こ" { // ko コ, compare to ni ニ
			fmt.Printf(", hint: index <--- to the B char, alright, コ　it makes sense, 'cause angles \n")
		} else if value == "さ" { // sa サ, if you 'se' so sa
			fmt.Printf(", hint: ring >v to the X char, at least the sa goes to the left, and it looks a lot like se, I'd say \n")
		} else if value == "し" { // shi シ, she sleeps and snores, compare to tsu ツ, and so ソ,
			fmt.Printf(", hint: middle on the D char, シ she is sleeping/snoring here, but has no excuse to2 look like this (no angles here, less curve though) \n")
		} else if value == "す" { // su ス, sue is running
			fmt.Printf(", hint: index < to the R char, they were looking for angles (sue is running ス) \n")
		} else if value == "せ" { // se セ
			fmt.Printf(", hint: pinky to the P char \n")
		} else if value == "そ" { // so ソ, compare to ん ン which lays down more and has a more laid-down back-tick
			fmt.Printf(", hint: index <-- to the C char, ソ Looks like しシ or no ノ, or ン starting with all angles - `had to `backtrack?  \n")
		} else if value == "た" { // ta タ, compare to ku ク and ke ケ
			fmt.Printf(", hint: pinky < to the Q char, the top left looks like ta, more at least than ku ク, or he ケ \n")
		} else if value == "ち" { // chi チ, compare to te テ
			fmt.Printf(", hint: pinky on the A char, 'some' resemblance if we see the middle line as the top of the backwards c \n")
		} else if value == "つ" { // tsu ツ, compare to shi シ, she who sleeps low and snores loudly
			fmt.Printf(", hint: pinky to the Z char, understandable?, if we are generous. Very generous..゛as in adding two ツ  shi シ said \n")
		} else if value == "て" { // te テ, compare to chi チ
			fmt.Printf(", hint: ring < to the W char, te went the wrong way, and gained a flat hat \n")
		} else if value == "と" { // to ト
			fmt.Printf(", hint: ring on the S char, toe is flipped-out. Kicked in the balls, on its head \n")
		} else if value == "な" { // na ナ, compare to me メ, and nu ヌ
			fmt.Printf(", hint: index < to the U char, is simple, very simple \n")
		} else if value == "に" { // ni ニ
			fmt.Printf(", hint: middle < to the I char \n")
		} else if value == "ぬ" { // nu ヌ, compare to me メ, and na ナ
			fmt.Printf(", hint: R-ring<--- to the 1 char, at least two lines cross, nu flew  \n")
		} else if value == "ね" { // ne ネ
			fmt.Printf(", hint: L-ring <v to the ,< chars, a hoe that got nailed \n")
		} else if value == "の" { // no ノ, no extra marks needed on this classic no symbol
			fmt.Printf(", hint: middle on the k char, drop the curve, and I'll take it \n")
		} else if value == "ま" { // ma マ, mama is unique, see the breast pump?
			fmt.Printf(", hint: index on the J char, (a breast pump maybe?) \n")
		} else if value == "み" { // mi ミ, and two makes three
			fmt.Printf(", hint: index <v to the N char, me thinks 3 \n")
		} else if value == "む" { // mu ム, an even-fatter-jawed moo-cow
			fmt.Printf(", hint: pinky two over ---> to the ]} chars, now it's an even-fatter-jawed moo-cow  \n")
		} else if value == "め" { // me メ
			fmt.Printf(", hint: pinky -> slide down to the /? chars, cross off the mess. But that's nothing nu ヌ \n")
		} else if value == "も" { // mo モ
			fmt.Printf(", hint: middle <v to the M char, way to hang to the right; you go Joe! \n")
		} else if value == "ら" { // ra ラ
			fmt.Printf(", hint: ring ^ to the o char, Similar ラ, ら \n")
		} else if value == "り" { // ri リ
			fmt.Printf(", hint: ring on the L char, longer on the right ring \n")
		} else if value == "る" { // ru ル, is two
			fmt.Printf(", hint: pinky < to the .> chars, ル, る, ru is now two? \n")
		} else if value == "れ" { // re レ, more, or less, a ru ル - しshe-said
			fmt.Printf(", hint: pinky on the ;+ chars, レ, れ; しshe-it \n")
		} else if value == "ろ" { // ro ロ
			fmt.Printf(", hint: pinky ---> long slide to the _ and backslash chars, maybe it had always been the way to go to do a square O for ro? \n")
		} else if value == "ほ" { // ha ハ
			fmt.Printf(", hint: L-ring---> to the \"-\" char\n")
		} else if value == "ひ" { // hi ヒ
			fmt.Printf(", hint: index > to the V char \n")
		} else if value == "ふ" { // hu フ, squinting it is a ふフ
			fmt.Printf(", hint: ring<, to the 2 char, if we squint? フ, ふ\n")
		} else if value == "へ" { // he ヘ
			fmt.Printf(", hint: ;+ char \n")
		} else if value == "ほ" { // ho ホ, now with a dress and all
			fmt.Printf(", hint: ring ---> to the - char, now looks like a hoe \n")
		} else if value == "や" { // ya ヤ
			fmt.Printf(", hint: index <-- to the 7 char \n")
		} else if value == "ゆ" { // yu ユ
			fmt.Printf(", hint: index->, to the 8 char, yu looks like ユ \n")
		} else if value == "よ" { // yo ヨ
			fmt.Printf(", hint: middle to the 9 char, tripple yo ヨ \n")
		} else if value == "わ" { // wa ワ, compare to wo ヲ
			fmt.Printf(", hint: ring > to the 0 char, so now it's a water fall for wa? \n")
		} else if value == "を" { // wo ヲ, compare to wa ワ
			fmt.Printf(", hint: ring > shifted 0 char, NOW, it looks like wa. Now! For fucks sake, Now? \n")
		} else if value == "ん" { // nh ン, compare to so ソ
			fmt.Printf(", hint: Y char, pointing at the one last remaining bent stroke of the hiragana char \n")
		} else {
			fmt.Printf("\n")
		}
	}

	fmt.Println()
	returnToMenuKMB = false
	return
} // end of func hiraMeatB

func restartHira() (returnToMenuRS bool) {

	rand.Seed(time.Now().UnixNano())

	//keyValuePairs := []string{
	randIndex := rand.Intn(len(fileOfCards))
	aCard := fileOfCards[randIndex]

	key := aCard.KeyRK
	value := aCard.Value
	if value == "rick" {
		fmt.Printf("woo")
	}

	fmt.Printf("%s", key)
	fmt.Printf("%s", string(colorCyan))
	fmt.Printf(" ? ")
	fmt.Printf("%s", string(colorReset))

	fmt.Scan(&in)

	if in == "menu" {
		returnToMenuRS = true // the return val
		return
	}

	hiraMeatB()
	ContextSHskip = false

	// this is the return val:
	returnToMenuRS = false
	return
} // end of func restartHira()

var menu int
var in string
var doingContextSH bool
var ContextSHskip bool
var value string
var key string

var colorReset = "\033[0m"
var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorCyan = "\033[36m"

// kinda broken, and probably also it is far from being totally refactored

// and the latest 'fix' produces this strange set of mixed behaviors when running 2
/*

roロ ? ?
      ^Oops! it was: ろ
// - this ? elicits an ^Oops! instead of the expected help comment


roロ ? ろ
, hint: pinky ---> long slide to the _ and backslash chars, maybe it had always been the way to go to do a square O for ro?
// - this correct responce offers a hint instead of the expected ^Right!


ワ ? わ
, hint: ring > to the 0 char, so now it's a water fall for wa?
// - this correct responce offers a hint instead of the expected ^Right!


hyuヒュ ? ひゅ
      ^Right!
// - this correct responce elicits the expected comment


buブ ? ?
      ^Oops! it was: ぶ
// - this ? elicits an ^Oops! instead of the expected help comment


buブ ? ぶ
// - this particular correct responce elicits nothing at all and instead proceeds to simply generate a new random prompt ?????


シ ? し
, hint: middle on the D char, シ she is sleeping/snoring here, but has no excuse to2 look like this (no angles here, less curve though)
// - this correct responce offers a hint instead of the expected ^Right!


ン ? ?
, hint: Y char, pointing at the one last remaining bent stroke of the hiragana char
// - this ? gives the expected hint, but then just generates another random prompt instead of re-issuing the original prompt as it should

byoビョ ?


*/

var fileOfCards = []charSetStruct{
	//// done to 203
	// We are never prompted with Hiragana chars, never!  We always see a Katakana prompt, so we always need a Hiragana hint

	// vowels:
	//
	// Kp   Hp    R    RK    Value:
	{" ア", "あ", "a", "aア", "あ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" a:あ looks nothing-like the hiragana a, but a lot like a hiragana te:て, あ fuck mae! a:ア is maybe a grotesque A:ア ?",
		" a:あ ア is maybe a grotesque A:ア ?",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" middle<- to the 3 char    あ ",
		" a"},
	//
	//
	// Kp   Hp    R    RK    Value:
	{" イ", "い", "i", "iイ", "い",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" i:い イ looks more like a hiragana te:て, still two mostly-vertical lines イ -- shift the two lines of the hiragana",
		" i:い maybe shift the two lines of the hiragana い",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" middle < to the E char    い ",
		" i"},
	//
	//
	// Kp   Hp    R    RK    Value:
	{" ウ", "う", "u", "uウ", "う",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" u:う is ok, having had to look for angles. But the u:ウ looks like the wa:ワ albeit with a mohawk ",
		" u:う more angles and with a tick for its top line : the u:ウ looks like the wa:ワ albeit with a mohawk ",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" middle> to the 4 char    う ",
		" u"},
	//
	//
	// Kp   Hp    R    RK    Value:
	{" エ", "え", "e", "eエ", "え",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" e:え エ does have a vague, angular resemblance. e:エ, eye see it as a ... an, eye; may-bey, may-b-eh e ",
		" e:え eye see it as a ... an, eye; may-bey, may-b-eh e ",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" index > to the 5 char    え ",
		" e"},
	//
	//
	// Kp   Hp    R    RK    Value:
	{" オ", "お", "o", "oオ", "お",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" o:お オ has only a vague resemblance, albeit with less curves. o オ, is someone on-the-go g-o o maybe ",
		" o:お is someone on-the-go maybe",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" index--> to the 6 char    お ",
		" o"},

	// ka group: (ka-ga) ============================================================================================
	//
	// Kp   Hp    R     RK     Value:
	{" カ", "か", "ka", "kaカ", "か",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ka:か カ is an easy one",
		" ka same カ albeit more angular and one less line to draw than か",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" index --> to the T char    か ",
		" ka"},
	//
	//
	// Kp   Hp    R     RK     Value:
	{" キ", "き", "ki", "kiキ", "き",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ki:き is an easy one, キ has the same top",
		" ki:き キ has the same top",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" G char, ki of G    き, compare to sa:さ ki:き さき",
		" ki"},
	//
	//
	// Kp   Hp    R     RK     Value:
	{" ク", "く", "ku", "kuク", "く",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ku:く Starting with one angle, they settled for this?  Compare to ta タ, and ke ケ",
		" ku:く ク compare ta タ, ke ケ ",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" H char    く ",
		" ku"},
	//
	//
	// Kp   Hp    R     RK     Value:
	{" ケ", "け", "ke", "keケ", "け",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ke:け bits of it are there, just as many curves though. Compare to ku ク, and ta タ",
		" ke:け compare ku ク, ta タ",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" :* chars    け ",
		" ke"},
	//
	//
	// Kp   Hp    R     RK     Value:
	{" コ", "こ", "ko", "koコ", "こ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ko:こ it makes sense, 'cause angles",
		" ko:こ Compare ni ニ",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" B char, ko way below    こ ",
		" ko"},

	// ka becomes ga group: -----------------------------------------------------------------------------------------
	//
	// Kp   Hp    R     RK     Value:
	{" ガ", "が", "ga", "gaガ", "が",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ga:が same ガ albeit more angular with one-less line to draw than が",
		" ga:が is an easy one",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" ga index --> to the T char    が ",
		" ga"},
	//
	//
	// Kp   Hp    R     RK     Value:
	{" ギ", "ぎ", "gi", "giギ", "ぎ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" gi:ぎ is an easy one, ギ has the same top",
		" gi:ぎ ギ has the same top",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" gi G char,    ぎ, compare to za:ざ ",
		" gi"},
	//
	//
	// Kp   Hp    R     RK     Value:
	{" グ", "ぐ", "gu", "guグ", "ぐ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" gu:ぐ Starting with one angle, they settled for this? Compare to ta:za:ダ, and ke:ge:ゲ",
		" gu:ぐ Compare za ダ, ge ゲ",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" H char    ぐ ",
		" gu"},
	//
	//
	// Kp   Hp    R     RK     Value:
	{" ゲ", "げ", "ge", "geゲ", "げ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ge:げ bits of it are there, just as many curves though. Compare to ku:gu:グ, and ta:za:ダ",
		" ge:げ Compare to ku:gu:グ, and ta:za:ダ",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" :* chars    げ",
		" ge"},
	//
	//
	// Kp   Hp    R     RK     Value:
	{" ゴ", "ご", "go", "goゴ", "ご",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" go:ご it makes sense, 'cause angles' ",
		" go:",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" index <--- to the B char    ご",
		" go:"},

	// ya, yu, yo's of ki:き ---------------------------------------------------------------------------
	//
	// Kp     Hp     R      RK        Value:
	{" キャ", "きゃ", "kya", "kyaキャ", "きゃ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" kya:きゃ is an easy one, キャ has the same top and the ya is similar",
		" kya:きゃ キャ has the same top and the ya is similar",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" G char,    きゃ, compare to sa:さ ",
		" kya:"},
	//
	//
	// Kp     Hp     R      RK        Value:
	{" キュ", "きゅ", "kyu", "kyuキュ", "きゅ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" kyu",
		" kyu",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" kyu",
		" kyu"},
	//
	//
	// Kp     Hp     R      RK        Value:
	{" キョ", "きょ", "kyo", "kyoキョ", "きょ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" kyo",
		" kyo",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" kyo",
		" kyo"},

	// ya, yu, yo's of gi:ぎ ---------------------------------------------------------------------------
	//
	// Kp     Hp     R      RK        Value:
	{" ギャ", "ぎゃ", "gya", "gyaギャ", "ぎゃ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" gya",
		" gya",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" gya",
		" gya"},
	//
	//
	// Kp     Hp     R      RK        Value:
	{" ギュ", "ぎゅ", "gyu", "gyuギュ", "ぎゅ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" gyu",
		" gyu",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" gyu",
		" gyu"},
	//
	//
	// Kp     Hp     R      Rk        Value;
	{" ギョ", "ぎょ", "gyo", "gyoギョ", "ぎょ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" gyo",
		" gyo",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" gyo",
		" gyo"},

	// sa group: (sa-za,ji,zu,za,zo) ================================================================================
	//
	// Kp   Hp    R     Rk     Value;
	{" サ", "さ", "sa", "saサ", "さ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" sa",
		" sa",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" sa",
		" sa"},
	//
	//
	// Kp   Hp    R      Rk      Value;
	{" シ", "し", "shi", "shiシ", "し",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" shi",
		" shi",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" shi",
		" shi"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ス", "す", "su", "suス", "す",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" su",
		" su",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" su",
		" su"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" セ", "せ", "se", "seセ", "せ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" se",
		" se",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" se",
		" se"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ソ", "そ", "so", "soソ", "そ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" so",
		" so",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" so",
		" so"},

	// sa becomes za, ji, zu, ze, zo --------------------------------------------------------------------------------
	//
	// Kp   Hp    R     Rk     Value;
	{" ザ", "ざ", "za", "zaザ", "ざ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" za",
		" za",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" za",
		" za"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ジ", "じ", "ji", "jiジ", "じ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ji",
		" ji",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" ji",
		" ji"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ズ", "ず", "zu", "zuズ", "ず",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" zu",
		" zu",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" zu",
		" zu"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ゼ", "ぜ", "ze", "zeゼ", "ぜ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ze",
		" ze",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" ze",
		" ze"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ゾ", "ぞ", "zo", "zoゾ", "ぞ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" zo",
		" zo",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" zo",
		" zo"},

	// ya, yu, yo's of shi:し:シ ------------------------------------------------------------------------
	//
	// Kp     Hp      R      Rk       Value;
	{" シャ", "しゃ", "sha", "shaシャ", "しゃ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" sha",
		" sha",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" sha",
		" sha"},
	//
	//
	// Kp     Hp      R      Rk       Value;
	{" シュ", "しゅ", "shu", "shuシュ", "しゅ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" shu",
		" shu",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" shu",
		" shu"},
	//
	//
	// Kp     Hp      R      Rk       Value;
	{" ショ", "しょ", "sho", "shoショ", "しょ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" sho",
		" sho",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" sho",
		" sho"},

	// ya, yu, yo's of ji:じ:ジ -------------------------------------------------------------------------
	//
	// Kp     Hp      R     Rk      Value;
	{" ジャ", "じゃ", "ja", "jaジャ", "じゃ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ja:",
		" ja:",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" ja:",
		" ja:"},
	//
	//
	// Kp     Hp      R     Rk      Value;
	{" ジュ", "じゅ", "ju", "juジュ", "じゅ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ju",
		" ju",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" ju",
		" ju"},
	//
	//
	// Kp     Hp      R     Rk      Value;
	{" ジョ", "じょ", "jo", "joジョ", "じょ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" jo",
		" jo",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" jo",
		" jo"},

	// ta group: (ta-da,ji,zu,de,do) ================================================================================
	//
	// Kp   Hp    R      Rk    Value;
	{" タ", "た", "ta", "taタ", "た",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ta",
		" ta",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" ta",
		" ta"},
	//
	//
	// Kp   Hp     R      Rk     Value;
	{" チ", "ち", "chi", "chiチ", "ち",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" chi",
		" chi",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" chi",
		" chi"},
	//
	//
	// Kp   Hp     R      Rk     Value;
	{" ツ", "つ", "tsu", "tsuツ", "つ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" tsu",
		" tsu",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" tsu",
		" tsu"},
	//
	//
	// Kp   Hp     R     Rk    Value;
	{" テ", "て", "te", "teテ", "て",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" te",
		" te",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" te",
		" te"},
	//
	//
	// Kp   Hp     R     Rk    Value;
	{" ト", "と", "to", "toト", "と",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" to",
		" to",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" to",
		" to"},

	// ta becomes da, ji, zu, de, do --------------------------------------------------------------------------------
	//
	// Kp   Hp    R     Rk     Value;
	{" ダ", "だ", "da", "daダ", "だ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" da",
		" da",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" da",
		" da"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ヂ", "ぢ", "ji", "jiヂ", "ぢ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ji",
		" ji",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" ji",
		" ji"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ヅ", "づ", "zu", "zuヅ", "づ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" zu",
		" zu",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" zu",
		" zu"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" デ", "で", "de", "deデ", "で",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" de",
		" de",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" de",
		" de"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ド", "ど", "do", "doド", "ど",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" do",
		" do",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" do",
		" do"},

	// ya, yu, yo's of chi:ち:チ ------------------------------------------------------------------------
	//
	// Kp     Hp      R       Rk      Value;
	{" チャ", "ちゃ", "cha", "chaチャ", "ちゃ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" cha",
		" cha",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" cha",
		" cha"},
	//
	//
	// Kp     Hp      R       Rk      Value;
	{" チュ", "ちゅ", "chu", "chuチュ", "ちゅ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" chu",
		" chu",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" chu",
		" chu"},
	//
	//
	// Kp     Hp      R      Rk       Value;
	{" チョ", "ちょ", "cho", "choチョ", "ちょ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" cho",
		" cho",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" cho",
		" cho"},

	// ya, yu, yo's of ji:ぢ:ヂ -------------------------------------------------------------------------
	//
	// Kp     Hp      R      Rk     Value;
	{" ヂャ", "ぢゃ", "ja", "jaヂャ", "ぢゃ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ja",
		" ja",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" ja",
		" ja"},
	//
	//
	// Kp      Hp     R      Rk     Value;
	{" ヂュ", "ぢゅ", "ju", "juヂュ", "ぢゅ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ju",
		" ju",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" ju",
		" ju"},
	//
	//
	// Kp     Hp      R      Rk     Value;
	{" ヂョ", "ぢょ", "jo", "joヂョ", "ぢょ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" jo",
		" jo",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" jo",
		" jo"},

	// na group: (always a naked group) =============================================================================
	//
	// Kp   Hp    R     Rk     Value;
	{" ナ", "な", "na", "naナ", "な",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" na",
		" na",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" na",
		" na"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ニ", "に", "ni", "niニ", "に",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ni",
		" ni",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" ni",
		" ni"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ヌ", "ぬ", "nu", "nuヌ", "ぬ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" nu",
		" nu",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" nu",
		" nu"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ネ", "ね", "ne", "neネ", "ね",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ne",
		" ne",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" ne",
		" ne"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ノ", "の", "no", "noノ", "の",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" no",
		" no",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" no",
		" no"},

	// ya, yu, yo's of ni:に:ニ -------------------------------------------------------------------------
	//
	// Kp     Hp      R       Rk      Value;
	{" ニャ", "にゃ", "nya", "nyaニャ", "にゃ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" nya",
		" nya",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" nya",
		" nya"},
	//
	//
	// Kp     Hp      R       Rk      Value;
	{" ニュ", "にゅ", "nyu", "nyuニュ", "にゅ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" nyu",
		" nyu",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" nyu",
		" nyu"},
	//
	//
	// Kp     Hp      R       Rk      Value;
	{" ニョ", "にょ", "nyo", "nyoニョ", "にょ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" nyo",
		" nyo",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" nyo",
		" nyo"},

	// ha group: (ha-fu ba:pa) ha:hi:fu:he:ho =======================================================================
	//
	// Kp   Hp    R     Rk     Value;
	{" ハ", "は", "ha", "haハ", "は",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ha",
		" ha",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" ha",
		" ha"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ヒ", "ひ", "hi", "hiヒ", "ひ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" hi",
		" hi",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" hi",
		" hi"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" フ", "ふ", "fu", "fuフ", "ふ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" fu",
		" fu",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" fu",
		" fu"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ヘ", "へ", "he", "heヘ", "へ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" he",
		" he",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" he",
		" he"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ホ", "ほ", "ho", "hoホ", "ほ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ho",
		" ho",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" ho",
		" ho"},

	// ba, bi, bu, be, bo -------------------------------------------------------------------------------------------
	//
	// Kp   Hp    R     Rk     Value;
	{" バ", "ば", "ba", "baバ", "ば",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ba",
		" ba",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" ba",
		" ba"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ビ", "び", "bi", "biビ", "び",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" bi",
		" bi",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" bi",
		" bi"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ブ", "ぶ", "bu", "buブ", "ぶ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" bu",
		" bu",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" bu",
		" bu"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ベ", "べ", "be", "beベ", "べ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" be",
		" be",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" be",
		" be"},
	//
	//
	// Kp   Hp    R      Rk    Value;
	{" ボ", "ぼ", "bo", "boボ", "ぼ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" bo",
		" bo",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" bo",
		" bo"},

	// pa, pi, pu, pe, po -------------------------------------------------------------------------------------------
	//
	// Kp   Hp    R      Rk    Value;
	{" パ", "ぱ", "pa", "paパ", "ぱ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" pa",
		" pa",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" pa",
		" pa"},
	//
	//
	// Kp   Hp    R      Rk    Value;
	{" ピ", "ぴ", "pi", "piピ", "ぴ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" pi",
		" pi",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" pi",
		" pi"},
	//
	//
	// Kp   Hp    R      Rk    Value;
	{" プ", "ぷ", "pu", "puプ", "ぷ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" pu",
		" pu",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" pu",
		" pu"},
	//
	//
	// Kp   Hp    R      Rk    Value;
	{" ペ", "ぺ", "pe", "peペ", "ぺ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" pe",
		" pe",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" pe",
		" pe"},
	//
	//
	// Kp   Hp    R      Rk    Value;
	{" ポ", "ぽ", "po", "poポ", "ぽ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" po",
		" po",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" po",
		" po"},

	// ya, yu, yo's of hi:ひ:ヒ -------------------------------------------------------------------------
	//
	// Kp     Hp      R       Rk      Value;
	{" ヒャ", "ひゃ", "hya", "hyaヒャ", "ひゃ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" hya",
		" hya",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" hya",
		" hya"},
	//
	//
	// Kp      Hp     R        Rk     Value;
	{" ヒュ", "ひゅ", "hyu", "hyuヒュ", "ひゅ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" hyu",
		" hyu",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" hyu",
		" hyu"},
	//
	//
	// Kp     Hp      R       Rk      Value;
	{" ヒョ", "ひょ", "hyo", "hyoヒョ", "ひょ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" hyo",
		" hyo",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" hyo",
		" hyo"},

	// ya, yu, yo's of bi:び:ビ -------------------------------------------------------------------------
	//
	// Kp      Hp     R      Rk       Value;
	{" ビャ", "びゃ", "bya", "byaビャ", "びゃ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" bya",
		" bya",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" bya",
		" bya"},
	//
	//
	// Kp     Hp      R      Rk        Value;
	{" ビュ", "びゅ", "byu", "byuビュ", "びゅ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" byu",
		" byu",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" byu",
		" byu"},
	//
	//
	// Kp     Hp      R       Rk      Value;
	{" ビョ", "びょ", "byo", "byoビョ", "びょ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" byo",
		" byo",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" byo",
		" byo"},

	//
	//
	// Kp     Hp      R       Rk      Value;
	{" ピャ", "ぴゃ", "pya", "pyaピャ", "ぴゃ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" pya",
		" pya",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" pya",
		" pya"},

	// ya, yu, yo's of pi:ぴ:ピ -------------------------------------------------------------------------
	//
	// Kp      Hp     R       Rk      Value;
	{" ピュ", "ぴゅ", "pyu", "pyuピュ", "ぴゅ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" pyu",
		" pyu",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" pyu",
		" pyu"},
	//
	//
	// Kp     Hp      R       Rk      Value;
	{" ピョ", "ぴょ", "pyo", "pyoピョ", "ぴょ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" pyo",
		" pyo",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" pyo",
		" pyo"},

	// ma group: ====================================================================================================
	//
	// Kp   Hp    R     Rk     Value;
	{" マ", "ま", "ma", "maマ", "ま",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ma",
		" ma",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" ma",
		" ma"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ミ", "み", "mi", "miミ", "み",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" mi",
		" mi",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" mi",
		" mi"},
	//
	//
	// Kp   Hp    R      Rk    Value;
	{" ム", "む", "mu", "muム", "む",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" mu",
		" mu",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" mu",
		" mu"},
	//
	//
	// Kp   Hp    R      Rk    Value;
	{" メ", "め", "me", "meメ", "め",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" me",
		" me",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" me",
		" me"},
	//
	//
	// Kp   Hp    R      Rk    Value;
	{" モ", "も", "mo", "moモ", "も",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" mo",
		" mo",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" mo",
		" mo"},

	// ya, yu, yo's of mi:み:ミ -------------------------------------------------------------------------
	//
	// Kp     Hp      R      Rk       Value;
	{" ミャ", "みゃ", "mya", "myaミャ", "みゃ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" mya",
		" mya",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" mya",
		" mya"},
	//
	//
	// Kp     Hp      R       Rk      Value;
	{" ミュ", "みゅ", "myu", "myuミュ", "みゅ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" myu",
		" myu",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" myu",
		" myu"},
	//
	//
	// Kp      Hp     R       Rk      Value;
	{" ミョ", "みょ", "myo", "myoミョ", "みょ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" myo",
		" myo",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" myo",
		" myo"},

	// group ya yu yo ===============================================================================================
	//
	// Kp   Hp    R      Rk    Value;
	{" ヤ", "や", "ya", "yaヤ", "や",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ya",
		" ya",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" ya",
		" ya"},
	//
	//
	// Kp   Hp    R      Rk    Value;
	{" ユ", "ゆ", "yu", "yuユ", "ゆ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" yu",
		" yu",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" yu",
		" yu"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ヨ", "よ", "yo", "yoヨ", "よ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" yo",
		" yo",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" yo",
		" yo"},

	// ra group: (pronounced more like la) ==========================================================================
	//
	// Kp   Hp    R     Rk     Value;
	{" ラ", "ら", "ra", "raラ", "ら",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ra",
		" ra",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" ra",
		" ra"},
	//
	//
	// Kp   Hp    R      Rk    Value;
	{" リ", "り", "ri", "riリ", "り",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ri",
		" ri",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" ri",
		" ri"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ル", "る", "ru", "ruル", "る",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ru",
		" ru",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" ru",
		" ru"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" レ", "れ", "re", "reレ", "れ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" re",
		" re",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" re",
		" re"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ロ", "ろ", "ro", "roロ", "ろ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ro",
		" ro",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" ro",
		" ro"},

	// ya, yu, yo's of ri:り:リ ---------------------------------------------------------------------------
	//
	// Kp     Hp      R       Rk      Value;
	{" リャ", "りゃ", "rya", "ryaリャ", "りゃ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" rya",
		" rya",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" rya",
		" rya"},
	//
	//
	// Kp     Hp      R       Rk      Value;
	{" リュ", "りゅ", "ryu", "ryuリュ", "りゅ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ryu",
		" ryu",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" ryu",
		" ryu"},
	//
	//
	// Kp     Hp      R       Rk      Value;
	{" リョ", "りょ", "ryo", "ryoリョ", "りょ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" ryo",
		" ryo",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" ryo",
		" ryo"},

	// wa and wo ====================================================================================================
	//
	// Kp   Hp    R     Rk     Value;
	{" ワ", "わ", "wa", "waワ", "わ",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" wa",
		" wa",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" wa",
		" wa"},
	//
	//
	// Kp   Hp    R     Rk     Value;
	{" ヲ", "を", "wo", "woヲ", "を",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" wo",
		" wo",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" wo",
		" wo"},

	// n ============================================================================================================
	//
	// Kp   Hp    R    Rk    Value;
	{" ン", "ん", "n", "nン", "ん",

		// Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
		" n",
		" n",

		// Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
		" n",
		" n"},
}
