package main

import (
	"log"

	"github.com/quiemill/gogogo/internal/app/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
