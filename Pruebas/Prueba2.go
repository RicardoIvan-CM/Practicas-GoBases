package main

import (
	"errors"
	"fmt"
)

var err1 = errors.New("Error número 1")

func x() error {
	return fmt.Errorf("info extra del error: %w", err1)
}

func main() {
	e := x()
	coincidence := errors.Is(e, err1)
	fmt.Println(coincidence)
}
