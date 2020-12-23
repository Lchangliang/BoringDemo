package main

import (
	"serialzation/serial"
)

type Person struct {
	Id int64
	Name string
	Age int
}

type DemoClass_1 struct {
	serialzer serial.Serialzable
}

type DemoClass_2 struct {
	deserialzer serial.Deserializable
}

func main() {
	var demo1 DemoClass_1
	person := Person{1,"wuweiwei",20}
	demo1.serialzer = &serial.Serialzation{}
	tmp := demo1.serialzer.Serialize(person)
	println(tmp)
	var demo2 DemoClass_2
	demo2.deserialzer = &serial.Serialzation{}
	person_tmp := Person{}
	demo2.deserialzer.Deserializable(tmp, &person_tmp)
	println(person_tmp == person)
}
