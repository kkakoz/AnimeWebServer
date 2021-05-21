package gerrors

import (
	"log"
	"testing"
)

func test1() {
	test2()
}

func test2() {
	test3()
}

func test3() {
	s := stack()
	log.Println(s)
}


func TestStack(t *testing.T) {
	test1()
}