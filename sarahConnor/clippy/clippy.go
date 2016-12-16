package clippy

import (
	"fmt"
	"github.com/kr/text"
	"strings"
)

const maxLineLimit = 39
const lineFormat = "\n| %s%s |"

var clippy = [2]string{`
 _________________________________________
/                                         \`, `
\_________________________________________/
 \
  \
     __
    /  \
    |  |
    @  @
    |  |
    || |/
    || ||
    |\_/|
    \___/
`,
}

func SayGoodbye(note string) {
	note = text.Wrap(note, maxLineLimit)
	splitNote := strings.Split(note, "\n")

	fmt.Print(clippy[0])
	for _, line := range splitNote {
		fmt.Printf(lineFormat, line, strings.Repeat(" ", maxLineLimit-len(line)))
	}
	fmt.Print(clippy[1])
}

func DanceAway() {
	fmt.Print(danceCat)
}

const danceCat = `
(_＼ヽ
　 ＼＼ .Λ＿Λ.
　　 ＼(　ˇωˇ)　
　　　 >　⌒ヽ
　　　/ 　 へ＼
　　 /　　/　＼＼
　　 ﾚ　ノ　　 ヽ_つ
　　/　/
　 /　/|
　(　(ヽ
　|　|、＼
　| 丿 ＼ ⌒)
　| |　　) /
 ノ ) 　 Lﾉ
(_／
`
