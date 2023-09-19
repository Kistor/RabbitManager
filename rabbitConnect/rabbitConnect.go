package rabbitconnect

import (
	"fyne.io/fyne/v2"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Connection struct {
	URL    string
	Name   string
	conn   *amqp.Connection
	chanel *amqp.Channel
	queue  amqp.Queue
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

func (c *Connection) Clear() {

}

func (c *Connection) ConnectServe() (err error) {
	/// коннектимся к серверу
	c.conn, err = amqp.Dial(c.Name)
	if err != nil {
		return
	}

	// Конект к каналу (wtf)
	c.chanel, err = c.conn.Channel()
	if err != nil {
		return
	}
	return
}

func (c *Connection) SetQueue() {
	if c.chanel != nil {
		c.queue, _ = c.chanel.QueueDeclare(
			c.Name, // name
			false,  // durable
			false,  // delete when unused
			false,  // exclusive
			false,  // no-wait
			nil,    // arguments
		)
	}
}
