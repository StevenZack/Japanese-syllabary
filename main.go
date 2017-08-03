package main

import (
	"fmt"
	"github.com/gizak/termui"
	"strconv"
)

func main() {
	err := termui.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer termui.Close()
	data := make(map[string]string)
	data["あ"] = "a"
	data["い"] = "i"
	data["う"] = "u"
	data["え"] = "e"
	data["お"] = "o"

	data["か"] = "ka"
	data["き"] = "ki"
	data["く"] = "ku"
	data["け"] = "ke"
	data["こ"] = "ko"
	data["さ"] = "sa"
	data["し"] = "si"
	data["す"] = "su"
	data["せ"] = "se"
	data["そ"] = "so"
	data["た"] = "ta"
	data["ち"] = "chi"
	data["つ"] = "tsu"
	data["て"] = "te"
	data["と"] = "to"
	data["な"] = "na"
	data["に"] = "ni"
	data["ぬ"] = "nu"
	data["ね"] = "ne"
	data["の"] = "no"
	data["は"] = "ha"
	data["ひ"] = "hi"
	data["ふ"] = "hu"
	data["へ"] = "he"
	data["ほ"] = "ho"
	data["ま"] = "ma"
	data["み"] = "mi"
	data["む"] = "mu"
	data["め"] = "me"
	data["も"] = "mo"

	data["や"] = "ya"
	data["ゆ"] = "yu"
	data["よ"] = "yo"

	data["ら"] = "ra"
	data["り"] = "ri"
	data["る"] = "ru"
	data["れ"] = "re"
	data["ろ"] = "ro"

	data["わ"] = "wa"
	data["を"] = "wo"
	data["ん"] = "n"
	//片假名
	data["ア"] = "a"
	data["イ"] = "i"
	data["ウ"] = "u"
	data["エ"] = "e"
	data["オ"] = "o"

	data["カ"] = "ka"
	data["キ"] = "ki"
	data["ク"] = "ku"
	data["ケ"] = "ke"
	data["コ"] = "ko"
	data["サ"] = "sa"
	data["シ"] = "si"
	data["ス"] = "su"
	data["セ"] = "se"
	data["ソ"] = "so"
	data["タ"] = "ta"
	data["チ"] = "chi"
	data["ツ"] = "tsu"
	data["テ"] = "te"
	data["ト"] = "to"
	data["ナ"] = "na"
	data["ニ"] = "ni"
	data["ヌ"] = "nu"
	data["ネ"] = "ne"
	data["ノ"] = "no"
	data["ハ"] = "ha"
	data["ヒ"] = "hi"
	data["フ"] = "hu"
	data["ヘ"] = "he"
	data["ホ"] = "ho"
	data["マ"] = "ma"
	data["ミ"] = "mi"
	data["ム"] = "mu"
	data["メ"] = "me"
	data["モ"] = "mo"

	data["ヤ"] = "ya"
	data["ユ"] = "yu"
	data["ヨ"] = "yo"

	data["ラ"] = "ra"
	data["リ"] = "ri"
	data["ル"] = "ru"
	data["レ"] = "re"
	data["ロ"] = "ro"

	data["ワ"] = "wa"
	data["ヲ"] = "wo"
	data["ン"] = "n"

	p := termui.NewPar("ン")
	p.Height = 3
	p.Width = 40
	c := make(chan bool, 1)
	go circle(data, 0, p, c)

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})
	termui.Handle("/sys/kbd/<space>", func(termui.Event) {
		c <- true
	})
	termui.Handle("/sys/kbd/m", func(termui.Event) {
		c <- false
	})
	termui.Render(p)
	termui.Loop()
}
func circle(da map[string]string, round int, p *termui.Par, c chan bool) {
	if len(da) < 1 {
		p.Text = "恭喜你记忆完成，共计" + strconv.FormatInt(int64(round), 10) + "次循环"
		termui.Render(p)
		return
	}
	submap := make(map[string]string)
	for k, v := range da {
		p.Text = k
		termui.Render(p)
		c1 := <-c
		p.Text = k + "  " + v
		termui.Render(p)
		c2 := <-c
		if !c1 || !c2 {
			submap[k] = v
		}
	}
	p.Text = "第" + strconv.FormatInt(int64(round+1), 10) + "轮循环结束"
	termui.Render(p)
	<-c
	circle(submap, round+1, p, c)
	<-c
	termui.StopLoop()
}
