// views/buttons.go
package views

import (
    "fyne.io/fyne/v2/widget"
    "pelota/models"
)

func StartGameButton(t *models.Topos) *widget.Button {
    return widget.NewButton("Start Game", func() {
        go t.Run()
    })
}

func StopGameButton(p1, p2 *models.Pelota, t *models.Topos) *widget.Button {
    return widget.NewButton("Stop Game", func() {
        p1.SetStatus(false)
        p2.SetStatus(false)
        t.SetStatus(false)
    })
}

func StartPelota1Button(p1 *models.Pelota) *widget.Button {
    return widget.NewButton("Ball 1", func() {
        go p1.Run()
    })
}

func StartPelota2Button(p2 *models.Pelota) *widget.Button {
    return widget.NewButton("Ball 2", func() {
        go p2.Run()
    })
}
