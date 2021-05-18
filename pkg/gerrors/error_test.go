package gerrors

import (
	"log"
	"testing"
)

func TestStack(t *testing.T) {
	s := stack()
	log.Println(s)
}