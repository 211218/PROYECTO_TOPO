package models

import (
    "fmt"
    "time"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/storage"
)

type Topos struct {
    posX   float32
    posY   float32
    status bool
    topo   *canvas.Image
    pelota1 *Pelota // Agrega una referencia a la primera pelota
    pelota2 *Pelota // Agrega una referencia a la segunda pelota
    pelota3 *Pelota // Agrega una referencia a la tercera pelota
    topoWin *TopoWin
}

func NewTopos(posX float32, posY float32, img *canvas.Image, pelota1 *Pelota, pelota2 *Pelota, pelota3 * Pelota, topoWin *TopoWin) *Topos {
    return &Topos{
        posX:   posX,
        posY:   posY,
        status: true,
        topo:   img,
        pelota1: pelota1, // Inicializa la referencia a la primera pelota
        pelota2: pelota2, // Inicializa la referencia a la segunda pelota
        pelota3: pelota3, // Inicializa la referencia a la tercera pelota
        topoWin: topoWin,

    }
}

// Función para verificar si dos rectángulos se superponen
func Intersects(r1, r2, r3, r4 fyne.Position, s1, s2, s3, s4 fyne.Size) bool {
    // ajustar los valores de los rectángulos para que coincidan con los tamaños de las imágenes
    r1.X += 20
    r1.Y += 20
    r2.X += 20
    r2.Y += 20
    r3.X += 20
    r3.Y += 20
    r4.X += 20
    r4.Y += 20
    s1.Width -= 60
    s1.Height -= 60
    s2.Width -= 60
    s2.Height -= 60
    s3.Width -= 60
    s3.Height -= 60
    s4.Width -= 60
    s4.Height -= 60

    return r1.X < r2.X+s2.Width &&
           r1.X+s1.Width > r2.X &&
           r1.Y < r2.Y+s2.Height &&
           r1.Y+s1.Height > r2.Y ||
              r1.X < r3.X+s3.Width &&
                r1.X+s1.Width > r3.X &&
                r1.Y < r3.Y+s3.Height &&
                r1.Y+s1.Height > r3.Y ||
                     r1.X < r4.X+s4.Width &&
                          r1.X+s1.Width > r4.X &&
                            r1.Y < r4.Y+s4.Height &&
                                r1.Y+s1.Height > r4.Y

                
}

func (t *Topos) Run() {
        // mover el topo con teclado A W S D
        fyne.CurrentApp().Driver().CanvasForObject(t.topo).SetOnTypedKey(func(key *fyne.KeyEvent) {
            switch key.Name {
            case fyne.KeyA:
                t.posX -= 10
            case fyne.KeyD:
                t.posX += 10
            case fyne.KeyW:
                t.posY -= 10
            case fyne.KeyS:
                t.posY += 10
            }
        })

    t.status = true
    for t.status {
        // Verifica si el topo ha chocado con alguna de las pelotas
        if Intersects(t.topo.Position(), t.pelota1.pel.Position(), fyne.NewPos(0, 0), fyne.NewPos(0, 0), t.topo.Size(), t.pelota1.pel.Size(), fyne.NewSize(0, 0), fyne.NewSize(0, 0)) ||
        Intersects(t.topo.Position(), t.pelota2.pel.Position(), fyne.NewPos(0, 0), fyne.NewPos(0, 0), t.topo.Size(), t.pelota2.pel.Size(), fyne.NewSize(0, 0), fyne.NewSize(0, 0)) ||
           Intersects(t.topo.Position(), t.pelota3.pel.Position(), t.pelota2.pel.Position(), fyne.NewPos(0, 0), t.topo.Size(), t.pelota3.pel.Size(), t.pelota2.pel.Size(), fyne.NewSize(0, 0))  {
           
           // mostrar una imagen de perdedor
            lose := canvas.NewImageFromURI(storage.NewFileURI("./assets/gameover.jpeg"))
            lose.Resize(fyne.NewSize(50,50))
            lose.Move(fyne.NewPos(0,0))

            // Añade la imagen a la ventana de tu aplicación
            fyne.CurrentApp().Driver().CanvasForObject(t.topo).SetContent(lose)
            fyne.CurrentApp().Driver().CanvasForObject(t.topo).Refresh(lose)
           
            t.pelota1.SetStatus(false) // Detiene la primera pelota
            t.pelota2.SetStatus(false) // Detiene la segunda pelota
            t.pelota3.SetStatus(false) // Detiene la tercera pelota
            t.SetStatus(false) // Detiene el topo 
            break
        }
// Verifica si el topo ha chocado con el topoWin
if Intersects(t.topo.Position(), t.topoWin.topoWin.Position(), fyne.NewPos(0, 0), fyne.NewPos(0, 0), t.topo.Size(), t.topoWin.topoWin.Size(), fyne.NewSize(0, 0), fyne.NewSize(0, 0)) {
    // mostrar una imagen de ganador
    win := canvas.NewImageFromURI(storage.NewFileURI("./assets/winn.jpg"))
    win.Resize(fyne.NewSize(50,50))
    win.Move(fyne.NewPos(0,0))

    // Añade la imagen a la ventana de tu aplicación
    fyne.CurrentApp().Driver().CanvasForObject(t.topo).SetContent(win)
    fyne.CurrentApp().Driver().CanvasForObject(t.topo).Refresh(win)

    // cambiar el tamaño de ventana para que se vea la imagen


    t.pelota1.SetStatus(false) // Detiene la primera pelota
    t.pelota2.SetStatus(false) // Detiene la segunda pelota
    t.pelota3.SetStatus(false) // Detiene la tercera pelota
    t.SetStatus(false) // Detiene el topo
    t.topoWin.SetStatus(false) // Detiene el topoWin
    break
}


        fmt.Println(t.posX)
        t.topo.Move(fyne.NewPos(t.posX, t.posY)) // Cambia el valor "100" por la posición Y deseada
        time.Sleep(16 * time.Millisecond)    // Controla la velocidad de movimiento
    }
}

func (t *Topos) SetStatus(status bool) {
    t.status = status
}

// Agrega este método para obtener la posición actual del topo
func (t *Topos) GetPos() fyne.Position {
    return t.topo.Position()
}
