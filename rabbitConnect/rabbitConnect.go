package rabbitconnect

import (
	"fyne.io/fyne/v2"
)

type Connection struct {
	URL  string
	Name string
}

func (d *Connection) MinSize(objects []fyne.CanvasObject) fyne.Size {

	return fyne.NewSize(200, 130)
}

func (c *Connection) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {

	pos := fyne.NewPos(5, 10)

	objects[0].Resize(fyne.Size{Width: 200, Height: 35})
	objects[0].Move(pos)

	pos = pos.Add(fyne.NewPos(0, 42))

	objects[1].Resize(fyne.Size{Width: 200, Height: 35})
	objects[1].Move(pos)

	pos = pos.Add(fyne.NewPos(0, 42))

	objects[2].Resize(fyne.Size{Width: 50, Height: 40})
	objects[2].Move(pos)
	pos = pos.Add(fyne.NewPos(0, 42))

}

func NewConnection() *Connection {
	return &Connection{}
}
