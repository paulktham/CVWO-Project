package main

import (
	"net/http"

	"github.com/paulktham/CVWO-Project/tree/main/BACK1/internal/router"
)

func main() {
	r := router.SetUp()
	http.ListenAndServe(":8080", r)
}