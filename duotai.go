package main
import "fmt"


// 通过接口实现多态,感觉非常方便
type Usb interface {
	// Start()
	// Stop()
}


type Phone struct {
	ty int
	name string
}

type Camera struct {
	ty int
	name string
}

type Computer struct {

}

func (c Camera) Stop() {
	fmt.Println("stop camera ",c.ty,c.name)
}

func (p Phone) Stop() {
	fmt.Println("stop phone ",p.ty,p.name)
}

func (c Camera) Start() {
	fmt.Println("start camera",c.ty,c.name)
}

func (p Phone) Start() {
	fmt.Println("start phone ",p.ty,p.name)
}

// func (c Computer) Working(usb Usb) {
// 	usb.Start()
// 	usb.Stop()
// }

func main() {
	var p Phone
	p.ty = 1
	p.name = "Phone11"
	var c Camera
	c.ty = 2
	c.name = "sony"

	// var cm Computer

	// cm.Working(p)
	// cm.Working(c)

	// 定义一个多态数组,可以存放任意类型的结构体
	var usbarr [3]Usb
	usbarr[0] = p
	usbarr[1] = c
	usbarr[2] = Camera {}
	for i := 0; i < 3; i++ {
		fmt.Println(usbarr[i])
	}
}