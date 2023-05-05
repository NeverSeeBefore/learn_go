package main

import "fmt"

// usb接口
type Usber interface {
	start()
	stop()
}

type Compuer struct {
}

// 计算机的work方法会操作usb接口
func (c Compuer) work(usb Usber) {
	usb.start()
	usb.stop()
}

// Phone 和 Camera 都实现了usb的所有方法
type Phone struct {
	name string
}

func (p Phone) start() {
	fmt.Println(p.name, "start")
}
func (p Phone) stop() {
	fmt.Println(p.name, "stop")
}

type Camera struct {
	Name string
}

func (c Camera) start() {
	fmt.Println(c.Name, "started")
}

func (c Camera) stop() {
	fmt.Println(c.Name, "stopped")
}

func main() {
	computer := Compuer{}

	phone := Phone{"vivo"}

	camera := Camera{"leica"}

	computer.work(phone)

	computer.work(camera)
}
