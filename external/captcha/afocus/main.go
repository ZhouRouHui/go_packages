package main

import (
	"github.com/afocus/captcha"
	"image/color"
	"image/png"
	"net/http"
)

func main() {
	cpt := captcha.New()
	// 可以设置多个字体 或使用cap.AddFont("xx.ttf")追加
	_ = cpt.SetFont("comic.ttf")
	// 设置验证码大小
	cpt.SetSize(128, 64)
	// 设置干扰强度
	cpt.SetDisturbance(captcha.MEDIUM)
	// 设置前景色 可以多个 随机替换文字颜色 默认黑色
	cpt.SetFrontColor(color.RGBA{255, 255, 255, 255})
	// 设置背景色 可以多个 随机替换背景色 默认白色
	cpt.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})

	http.HandleFunc("/r", func(w http.ResponseWriter, r *http.Request) {
		img, str := cpt.Create(6, captcha.ALL)
		png.Encode(w, img)
		println(str)
	})

	_ = http.ListenAndServe(":8085", nil)
}
