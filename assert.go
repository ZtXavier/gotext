package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
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

func (p Phone) Call (){
	fmt.Println("Dididi call phone ...")
}

func (cp Computer) Working (usb Usb) {
	usb.Start()
	usb.Stop()
	// 在这里使用断言的好处,如果一个类其中没有某个方法
	// 可以通过判断来解决,否则会让程序退出
	if ph, ok := usb.(Phone); ok {
		ph.Call()
	}
}

func main() {
	var usbarr [3]Usb
	usbarr[0] = Phone{11,"iphone"}
	usbarr[1] = Camera{10,"sony"}
	usbarr[2] = Phone{2,"anzhuo"}
	var cmp Computer
	
	for _, v := range usbarr {
		cmp.Working(v)
	}

	
}