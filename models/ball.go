package models

import (
    "fmt"
    "time"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
)

type Pelota struct {
    posX   float32
    posY   float32
    upperLimit float32
    lowerLimit float32
    status bool
    pel    *canvas.Image
    topo   *Topos // Agrega una referencia al topo
}

func NewPelota(posx float32, posy float32, upper float32, lower float32, img *canvas.Image, topo *Topos) *Pelota {
    return &Pelota{
        posX:   posx,
        posY:   posy,
        upperLimit: upper,
        lowerLimit: lower,
        status: true,
        pel:    img,
        topo:   topo, // Inicializa la referencia al topo
    }
}

func (p *Pelota) Run() {
    var incY float32 = 2 // Velocidad vertical

    p.status = true
    for p.status {
        if p.posY > p.upperLimit {
            incY = -2 // Cambia la dirección cuando alcanza el límite superior
        } else if p.posY < p.lowerLimit {
            incY = 2 // Cambia la dirección cuando alcanza el límite inferior
        }

        p.posY += incY

        // Verifica si la pelota ha chocado con el topo y terminar el proceso del topo


        fmt.Println(p.posY)
        p.pel.Move(fyne.NewPos(p.posX, p.posY)) // Usa la posición X de la estructura
        time.Sleep(25 * time.Millisecond)    // Controla la velocidad de movimiento
    }
}

func (p *Pelota) SetStatus(status bool) {
    p.status = status
}
