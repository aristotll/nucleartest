package panel

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"plane-war/plane"
)

func NewPanel() (*fyne.Container, error) {
	background := canvas.NewImageFromFile("/Users/zz/GolandProjects/plane-war/img/bg1.jpg")
	redPlane, err := plane.NewRedPlane()
	if err != nil {
		return nil, err
	}
	planeImg := canvas.NewImageFromFile(redPlane.Image())
	//box := container.NewHBox(background, planeImg)

	//lay := container.NewHSplit(background, planeImg)

	lay := container.NewWithoutLayout(background, planeImg)
	return lay, nil
}
