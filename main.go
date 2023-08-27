package main

import (
	"fmt"
	rabbitconnect "myapp/rabbitConnect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeUI(connect *rabbitconnect.Connection) []fyne.CanvasObject {
	url := widget.NewEntry()
	url.SetPlaceHolder("Введите адрес rabbit")
	url.OnChanged = func(s string) {
		connect.URL = s
	}

	name := widget.NewEntry()
	name.SetPlaceHolder("Введите имя очереди")

	name.OnChanged = func(s string) {
		connect.Name = s
	}

	button := widget.NewButton("print", func() {
		fmt.Println(*connect)
	})

	return []fyne.CanvasObject{url, name, button}
}

func main() {

	connectData := rabbitconnect.NewConnection()

	myApp := app.New()
	myWindow := myApp.NewWindow("RABBIT WTF")

	myWindow.SetContent(container.New(connectData, makeUI(connectData)...))

	myWindow.Resize(fyne.NewSize(1000, 562))
	myWindow.ShowAndRun()
}
