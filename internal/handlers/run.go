package handlers

import (
	"fmt"
)

func Run(h *Handler) error {
	fmt.Println("http handler running")
	return nil
}
