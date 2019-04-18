package main

import "fmt"

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}

func main() {
	c := &Customer{"Barak Obama", &Log{"yes we can!"}}
	c.Log().Add("After me the world will be a better place!")
	fmt.Println(c.Log())
}
