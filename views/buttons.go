package views

import "fyne.io/fyne/v2/widget"

func NewStartButton(startGameFunc func()) *widget.Button {
	return widget.NewButton("Play", startGameFunc)
}

func NewStopButton(stopGameFunc func()) *widget.Button {
	return widget.NewButton("Stop", stopGameFunc)
}
