package models

import (

	"fyne.io/fyne/v2/canvas"
)

// crear un topoWin

type TopoWin struct {
	posX   float32
	posY   float32
	status bool
	topoWin   *canvas.Image
}

func NewTopoWin(posX float32, posY float32, img *canvas.Image) *TopoWin {
	return &TopoWin{
		posX:   posX,
		posY:   posY,
		status: true,
		topoWin:   img,
	}
}

func (t *TopoWin) Run() {
	t.status = true
}

func (t *TopoWin) SetStatus(status bool) {
	t.status = status
}

