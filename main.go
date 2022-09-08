package main

import (
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	goChrome "github.com/voyager-hang/go-chrome"
)

const indexURL = "index"

var gc *goChrome.GoChrome

func main() {
	gc = goChrome.Create(indexURL, goChrome.GoChromeOptions{
		AppName:              "go_chrome_demo",
		CliModule:            false,
		AppModule:            false,
		WindowWidth:          500,
		WindowHeight:         500,
		WindowPositionWidth:  400,
		WindowPositionHeight: 200,
		UseHttpServer:        true,
		HttpRoute:            nil,
		HttpPort:             0,
// 		AssetFile:            AssetFile(),
// 		RestoreAssets:        RestoreAssets,
	})
	// 绑定方法映射到js
	_ = gc.Bind("sss", Aaa)
	_ = gc.Bind("ffff", Bbb)
	gc.SetAction(goChrome.ActionTask{
		//chromedp.Click(`#aaa`, chromedp.NodeVisible),
		//chromedp.Click("#bbb", chromedp.NodeVisible),
		//chromedp.Click("#aaa", chromedp.NodeVisible),
		//chromedp.Click("#bbb", chromedp.NodeVisible),
		//chromedp.Click("#aaa", chromedp.NodeVisible),
		//chromedp.Click("#bbb", chromedp.NodeVisible),
	})
	gc.OpenAfter(func() {
		//打开完成后执行
		//fmt.Println("jsfunc")
		//chromedp.Evaluate(`function ddsss(){alert('ss')}`, false).Do(gc.ContextContext)
	})
	gc.ListenTarget(func(ev interface{}) {
		// 开启事件监听 事件列表：https://chromedevtools.github.io/devtools-protocol/
		switch ev := ev.(type) {
		case *network.EventLoadingFinished:
			// 页面加载完执行
			go func() {
				_ = chromedp.Evaluate(`function ddsss(){alert('ss')}`, false).Do(gc.ContextContext)
			}()
		case *string:
			// 意思意思
			fmt.Println(ev)
		default:
			//fmt.Printf("%T\n", ev)
		}

	})
	gc.Run()
}

func Aaa(a, b string) string {
	//fmt.Println(a)
	//fmt.Println(b)
	return "adsf"
	//return map[string]interface{}{
	//	"aa": 1,
	//	"bb": "dsa",
	//}
}

func Bbb(a, b float64) float64 {
	fmt.Println(a)
	fmt.Println(b)
	return a
	//go func() {
	//	// 调用js方法获取返回值
	//	time.Sleep(5 * time.Second)
	//	res := gc.JsFunc("aaa", "ddd", "ccc")
	//	fmt.Println(res)
	//}()
}
