package scenes

import (
	_ "fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"topos/models"
	"topos/views" // Importa el paquete "views"
)

type MainMenuScene struct {
	window fyne.Window
}

var p1, p2, p3 *models.Ball
var t *models.Topos
var tw *models.TopoWin

func NewMainMenuScene(window fyne.Window) *MainMenuScene {
	return &MainMenuScene{window: window}
}

func (s *MainMenuScene) Show() {

	// Cargar la imagen de fondo
	background := canvas.NewImageFromURI(storage.NewFileURI("./assets/background.png"))
	background.Resize(fyne.NewSize(800, 600))
	background.Move(fyne.NewPos(-10, -100))

	pelota1 := canvas.NewImageFromURI(storage.NewFileURI("./assets/ball.png"))
	pelota1.Resize(fyne.NewSize(150, 90))
	pelota1.Move(fyne.NewPos(150, 300))
	//Creamos el modelo de la primera pelota
	p1 = models.NewBall(float32(150), float32(300), float32(300), float32(100), pelota1, t)

	pelota2 := canvas.NewImageFromURI(storage.NewFileURI("./assets/ball.png"))
	pelota2.Resize(fyne.NewSize(150, 90))
	pelota2.Move(fyne.NewPos(350, 100))
	//Creamos el modelo de la segunda pelota
	p2 = models.NewBall(float32(350), float32(100), float32(300), float32(100), pelota2, t)

	// creamos un tercer modelo de pelota
	pelota3 := canvas.NewImageFromURI(storage.NewFileURI("./assets/ball.png"))
	pelota3.Resize(fyne.NewSize(150, 90))
	pelota3.Move(fyne.NewPos(550, 300))
	//Creamos el modelo de la tercera pelota
	p3 = models.NewBall(float32(550), float32(300), float32(300), float32(100), pelota3, t)

	// Cargar la imagen del topo
	topo := canvas.NewImageFromURI(storage.NewFileURI("./assets/super_go.png"))
	topo.Resize(fyne.NewSize(100, 100))

	topo.Move(fyne.NewPos(0, 280))

	// creamos otro topo
	topoWin := canvas.NewImageFromURI(storage.NewFileURI("./assets/topo_elegant.png"))
	topoWin.Resize(fyne.NewSize(60, 100))
	topoWin.Move(fyne.NewPos(730, 180))
	tw = models.NewTopoWin(float32(720), float32(200), float32(300), float32(100), topoWin, t)

	// Creamos el modelo del topo
	t = models.NewTopos(float32(0), float32(280), topo, p1, p2, p3, tw)

	// Crea los botones utilizando las funciones del paquete "views"
	botonIniciar := views.NewStartButton(s.StartGame)
	botonIniciar.Resize(fyne.NewSize(100, 30))
	botonIniciar.Move(fyne.NewPos(280, 5))

	botonDetener := views.NewStopButton(s.StopGame)
	botonDetener.Resize(fyne.NewSize(100, 30))
	botonDetener.Move(fyne.NewPos(400, 5))

	// aqui se muestran los elementos en la ventana
	s.window.SetContent(container.NewWithoutLayout(background, pelota1, pelota2, pelota3, topo, topoWin, botonIniciar, botonDetener))
}


func (s *MainMenuScene) StartGame() {
	go t.Run()
	go p1.Run()
	go p2.Run()
	go p3.Run()
	go tw.Run()
}

func (s *MainMenuScene) StopGame() {
	p1.SetStatus(false)
	p2.SetStatus(false)
	p3.SetStatus(false)
	t.SetStatus(false)
	tw.SetStatus(false)
}
