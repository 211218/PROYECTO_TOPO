package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"topos/scenes"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Game Topo")
	myWindow.CenterOnScreen()
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(800, 460))

	mainMenuScene := scenes.NewMainMenuScene(myWindow)
	mainMenuScene.Show()
	myWindow.ShowAndRun()
}
