package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("飞机大战")
	w.Resize(fyne.Size{
		Width:  500,
		Height: 700,
	})

	image := canvas.NewImageFromFile("./img/bg1.jpg")
	image2 := canvas.NewImageFromFile("./img/hero.png")
	image2.Resize(fyne.Size{
		Width:  10,
		Height: 10,
	})
	//image2.FillMode = canvas.ImageFillOriginal

	//lay := container.New(layout.NewGridLayout(2), image, image2)
	w.SetContent(image)
	w.SetContent(image2)
	//w.SetContent(lay)

	//pan, err := panel.NewPanel()
	//if err != nil {
	//	log.Println("create panel error: ", err)
	//	return
	//}
	//
	//w.SetContent(pan)
	//panel.NewPanel()
	w.ShowAndRun()
}