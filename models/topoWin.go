package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"time"
)

type TopoWin struct {
	posX   float32
	posY   float32
	upperLimit float32
	lowerLimit float32
	status bool
	topoWin   *canvas.Image
	topo 	  *Topos
}

func NewTopoWin(posX float32, posY float32, upper float32, lower float32, img *canvas.Image, topo * Topos) *TopoWin {
	return &TopoWin{
		posX:   posX,
		posY:   posY,
		upperLimit: upper,
		lowerLimit: lower,
		status: true,
		topoWin:   img,
		topo: topo,
	}
}

func (t *TopoWin) Run() {
	var incY float32 = 2 // Velocidad vertical

	t.status = true
	for t.status {
		if t.posY > t.upperLimit {
			incY = -2 // Cambia la dirección cuando alcanza el límite superior
		} else if t.posY < t.lowerLimit {
			incY = 2 // Cambia la dirección cuando alcanza el límite inferior
		}

		t.posY += incY

		t.topoWin.Move(fyne.NewPos(t.posX, t.posY)) // Actualiza la posición Y de la estructura
		time.Sleep(1 * time.Millisecond)    // Controla la velocidad de movimiento
	}
}

func (t *TopoWin) SetStatus(status bool) {
	t.status = status
}
