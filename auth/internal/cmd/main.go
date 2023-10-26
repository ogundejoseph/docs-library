package main

import (
	"github.com/ogundejoseph/docs-library/auth/internal/model"
	"github.com/ogundejoseph/docs-library/auth/internal/routes"
)


func main() {
	model.Setup()
	routes.Setup()
}