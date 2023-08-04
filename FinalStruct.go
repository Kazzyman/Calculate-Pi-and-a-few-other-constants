var fileOfCards = []charSetStruct{
//// done to 175, header formatting to 268
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


// ka group: (ka-ga)
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


// ka becomes ga group:
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
     " ge",
     " ge"},
//
//
// Kp   Hp    R     RK     Value:
{" ゴ", "ご", "go", "goゴ", "ご",

    // Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
     " go",
     " go",

    // Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
     " go",
     " go"},


// ya, yu, yo's of ki:き
//
// Kp     Hp     R      RK        Value:
{" キャ", "きゃ", "kya", "kyaキャ", "きゃ",

    // Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
     " kya",
     " kya",
    
    // Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
     " kya",
     " kya"},
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


// ya, yu, yo's of gi:ぎ
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


// sa group: (sa-za)
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
//
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
//
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
//
//
// Kp     Hp      R     Rk      Value;
{" ジャ", "じゃ", "ja", "jaジャ", "じゃ",

    // Hints for Katakana prompts: (guessing the Hiragana from Katakana prompts):
     " ja",
     " ja",

    // Hints for Romaji|Katakana prompts: (guessing the Hiragana from Romaji|Katakana prompts):
     " ja",
     " ja"},
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
//
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
//
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
//
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
//
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
//
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
//
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
//
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
//
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
//
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
//
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
//
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
//
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
//
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
//
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
//
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
//
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
//
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
//
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
//
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