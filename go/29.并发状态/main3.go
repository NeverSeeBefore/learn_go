package main

import (
	"fmt"
	"image"
	"time"
)

// 事件循环
// event loop

// 中心循环
// centrall loop

// 可以把任何工作进程 goroutine 看到独立的事件循环
func main() {
	r := NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
	r.Stop()
	time.Sleep(3 * time.Second)
	r.Start()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)

}

type command int

const (
	left  = command(1)
	right = command(2)
	start = command(3)
	stop  = command(4)
)

type RoverDriver struct {
	cmd chan command
}

func (rd RoverDriver) drive() {
	pos := image.Point{0, 0}
	direction := image.Point{1, 0}
	updateInterval := 500 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case cmd := <-rd.cmd:
			var cmdDesc string
			switch cmd {
			case left:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
				cmdDesc = "change direction to left " + direction.String()
			case right:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
				cmdDesc = "change direction to right " + direction.String()
			case start:
				nextMove = time.After(updateInterval)
				cmdDesc = "moving"
			case stop:
				nextMove = nil
				cmdDesc = "stoped"
			}
			fmt.Println(cmdDesc)
		case <-nextMove:
			pos = pos.Add(direction)
			fmt.Println("moveTo", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

func NewRoverDriver() *RoverDriver {
	rd := &RoverDriver{
		cmd: make(chan command),
	}
	go rd.drive()
	return rd
}

func (rd *RoverDriver) Left() {
	rd.cmd <- left
}
func (rd *RoverDriver) Right() {
	rd.cmd <- left
}
func (rd *RoverDriver) Start() {
	rd.cmd <- start
}
func (rd *RoverDriver) Stop() {
	rd.cmd <- stop
}
