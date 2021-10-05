package main

import "fmt"

type People struct {
	age int
	name string
}

type Speak interface {
	speak()
}

func (p People) speak(){
	fmt.Printf("我的名字叫%s,我%d岁了",p.name,p.age)
}


func main() {
	var speaker Speak
	kings := People{6,"kings"}
	speaker = kings
	speaker.speak()
}
