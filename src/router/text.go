package router

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type Incoming struct {
	Text string
}

type Outgoing struct {
	Text  []string `json:"text"`
	Tags  []string `json:"tags"`
	Emoji []string `json:"emoji"`
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func TextFormatter(response http.ResponseWriter, request *http.Request) {

	// JSON
	decoder := json.NewDecoder(request.Body)
	var data Incoming
	err := decoder.Decode(&data)
	if err != nil {
		io.WriteString(response, "json data not valid")
		return
	}

	// Basic Modifications
	text := strings.ToLower(data.Text)
	emojiRegex := `😄|😃|😀|😊|☺️|😉|😍|😘|😚|😗|😙|😜|😝|😛|😳|😁|😔|😌|😒|😞|😣|😢|😂|😭|😪|😥|😰|😅|😓|😩|😫|😨|😱|😠|😡|😤|😖|😆|😋|😷|😎|😴|😵|😲|😟|😦|😧|😈|👿|😮|😬|😐|😕|😯|😶|😇|😏|😑|👲|👳|👮|👷|💂|👶|👦|👧|👨|👩|👴|👵|👱|👼|👸|😺|😸|😻|😽|😼|🙀|😿|😹|😾|👹|👺|🙈|🙉|🙊|💀|👽|💩|🔥|✨|🌟|💫|💥|💢|💦|💧|💤|💨|👂|👀|👃|👅|👄|👍|👎|👌|👊|✊|✌️|👋|✋|👐|👆|👇|👉|👈|🙌|🙏|☝️|👏|💪|🚶|🏃|💃|👫|👪|👬|👭|💏|💑|👯|🙆|🙅|💁|🙋|💆|💇|💅|👰|🙎|🙍|🙇|🎩|👑|👒|👟|👞|👡|👠|👢|👕|👔|👚|👗|🎽|👖|👘|👙|💼|👜|👝|👛|👓|🎀|🌂|💄|💛|💙|💜|💚|❤️|💔|💗|💓|💕|💖|💞|💘|💌|💋|💍|💎|👤|👥|💬|👣|💭|🐶|🐺|🐱|🐭|🐹|🐰|🐸|🐯|🐨|🐻|🐷|🐽|🐮|🐗|🐵|🐒|🐴|🐑|🐘|🐼|🐧|🐦|🐤|🐥|🐣|🐔|🐍|🐢|🐛|🐝|🐜|🐞|🐌|🐙|🐚|🐠|🐟|🐬|🐳|🐋|🐄|🐏|🐀|🐃|🐅|🐇|🐉|🐎|🐐|🐓|🐕|🐖|🐁|🐂|🐲|🐡|🐊|🐫|🐪|🐆|🐈|🐩|🐾|💐|🌸|🌷|🍀|🌹|🌻|🌺|🍁|🍃|🍂|🌿|🌾|🍄|🌵|🌴|🌲|🌳|🌰|🌱|🌼|🌐|🌞|🌝|🌚|🌑|🌒|🌓|🌔|🌕|🌖|🌗|🌘|🌜|🌛|🌙|🌍|🌎|🌏|🌋|🌌|🌠|⭐|☀️|⛅|☁️|⚡|☔|❄️|⛄|🌀|🌁|🌈|🌊|🎍|💝|🎎|🎒|🎓|🎏|🎆|🎇|🎐|🎑|🎃|👻|🎅|🎄|🎁|🎋|🎉|🎊|🎈|🎌|🔮|🎥|📷|📹|📼|💿|📀|💽|💾|💻|📱|☎️|📞|📟|📠|📡|📺|📻|🔊|🔉|🔈|🔇|🔔|🔕|📢|📣|⏳|⌛|⏰|⌚|🔓|🔒|🔏|🔐|🔑|🔎|💡|🔦|🔆|🔅|🔌|🔋|🔍|🛁|🛀|🚿|🚽|🔧|🔩|🔨|🚪|🚬|💣|🔫|🔪|💊|💉|💰|💴|💵|💷|💶|💳|💸|📲|📧|📥|📤|✉️|📩|📨|📯|📫|📪|📬|📭|📮|📦|📝|📄|📃|📑|📊|📈|📉|📜|📋|📅|📆|📇|📁|📂|✂️|📌|📎|✒️|✏️|📏|📐|📕|📗|📘|📙|📓|📔|📒|📚|📖|🔖|📛|🔬|🔭|📰|🎨|🎬|🎤|🎧|🎼|🎵|🎶|🎹|🎻|🎺|🎷|🎸|👾|🎮|🃏|🎴|🀄|🎲|🎯|🏈|🏀|⚽|⚾️|🎾|🎱|🏉|🎳|⛳|🚵|🚴|🏁|🏇|🏆|🎿|🏂|🏊|🏄|🎣|☕|🍵|🍶|🍼|🍺|🍻|🍸|🍹|🍷|🍴|🍕|🍔|🍟|🍗|🍖|🍝|🍛|🍤|🍱|🍣|🍥|🍙|🍘|🍚|🍜|🍲|🍢|🍡|🍳|🍞|🍩|🍮|🍦|🍨|🍧|🎂|🍰|🍪|🍫|🍬|🍭|🍯|🍎|🍏|🍊|🍋|🍒|🍇|🍉|🍓|🍑|🍈|🍌|🍐|🍍|🍠|🍆|🍅|🌽|🏠|🏡|🏫|🏢|🏣|🏥|🏦|🏪|🏩|🏨|💒|⛪|🏬|🏤|🌇|🌆|🏯|🏰|⛺|🏭|🗼|🗾|🗻|🌄|🌅|🌃|🗽|🌉|🎠|🎡|⛲|🎢|🚢|⛵|🚤|🚣|⚓|🚀|✈️|💺|🚁|🚂|🚊|🚉|🚞|🚆|🚄|🚅|🚈|🚇|🚝|🚋|🚃|🚎|🚌|🚍|🚙|🚘|🚗|🚕|🚖|🚛|🚚|🚨|🚓|🚔|🚒|🚑|🚐|🚲|🚡|🚟|🚠|🚜|💈|🚏|🎫|🚦|🚥|⚠️|🚧|🔰|⛽|🏮|🎰|♨️|🗿|🎪|🎭|📍|🚩|🇯🇵|🇰🇷|🇩🇪|🇨🇳|🇺🇸|🇫🇷|🇪🇸|🇮🇹|🇷🇺|🇬🇧|1️⃣|2️⃣|3️⃣|4️⃣|5️⃣|6️⃣|7️⃣|8️⃣|9️⃣|0️⃣|🔟|🔢|#️⃣|🔣|⬆️|⬇️|⬅️|➡️|🔠|🔡|🔤|↗️|↖️|↘️|↙️|↔️|↕️|🔄|◀️|▶️|🔼|🔽|↩️|↪️|ℹ️|⏪|⏩|⏫|⏬|⤵️|⤴️|🆗|🔀|🔁|🔂|🆕|🆙|🆒|🆓|🆖|📶|🎦|🈁|🈯|🈳|🈵|🈴|🈲|🉐|🈹|🈺|🈶|🈚|🚻|🚹|🚺|🚼|🚾|🚰|🚮|🅿️|♿|🚭|🈷️|🈸|🈂️|Ⓜ️|🛂|🛄|🛅|🛃|🉑|㊙️|㊗️|🆑|🆘|🆔|🚫|🔞|📵|🚯|🚱|🚳|🚷|🚸|⛔|✳️|❇️|❎|✅|✴️|💟|🆚|📳|📴|🅰️|🅱️|🆎|🅾️|💠|➿|♻️|♈|♉|♊|♋|♌|♍|♎|♏|♐|♑|♒|♓|⛎|🔯|🏧|💹|💲|💱|©️|®️|™️|❌|‼️|⁉️|❗|❓|❕|❔|⭕|🔝|🔚|🔙|🔛|🔜|🔃|🕛|🕧|🕐|🕜|🕑|🕝|🕒|🕞|🕓|🕟|🕔|🕠|🕕|🕖|🕗|🕘|🕙|🕚|🕡|🕢|🕣|🕤|🕥|🕦|✖️|➕|➖|➗|♠️|♥️|♣️|♦️|💮|💯|✔️|☑️|🔘|🔗|➰|〰️|〽️|🔱|◼️|◻️|◾|◽|▪️|▫️|🔺|🔲|🔳|⚫|⚪|🔴|🔵|🔻|⬜|⬛|🔶|🔷|🔸|🔹`
	textRegex := `([^А-ЯЁа-яёA-Za-z1-9# ])`
	hashTagRegex := `\S*#(?:\[[^\]]+\]|\S+)`

	// Emoji
	re, err := regexp.Compile(emojiRegex)
	emoji := re.FindAllString(text, -1)

	// Tags
	text = re.ReplaceAllString(text, "")
	text = strings.Replace(text, "#", " #", -1)
	re, err = regexp.Compile(textRegex)
	text = re.ReplaceAllString(text, " ")
	re, err = regexp.Compile(hashTagRegex)
	tags := re.FindAllString(text, -1)
	tags = Map(tags, func(v string) string {
		return strings.Replace(v, "#", "", -1)
	})

	// Words
	text = re.ReplaceAllString(text, "")
	words := strings.Split(text, " ")
	wordsData := Filter(words, func(v string) bool {
		if len(v) > 3 {
			return true
		} else {
			return false
		}
	})

	// Response
	result := Outgoing{wordsData, tags, emoji}
	res, err := json.Marshal(result)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(res)

}
