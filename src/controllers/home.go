package controllers

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, I am home controller!")
	w.WriteHeader(http.StatusOK)
 }