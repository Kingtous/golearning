package main

// Go语言实现了反射，所谓反射就是能检查程序在运行时的状态。
//我们一般用到的包是reflect包。
//如何运用reflect包，官方的这篇文章详细的讲解了reflect包的实现原理，
//laws of reflection
//使用reflect一般分成三步，下面简要的讲解一下：
//要去反射是一个类型的值(这些值都实现了空interface)，
//首先需要把它转化成reflect对象(reflect.Type或者reflect.Value，根据不同的情况调用不同的函数)。

import "reflect"

func main() {
	var x float64 = 3.4
	// 反射的字段必须是可修改的，如果x是const，则无法反射
	p := reflect.ValueOf(&x)
	v := p.Elem()
	// 必须调用.Elem才能setFloat
	v.SetFloat(7.1)
}