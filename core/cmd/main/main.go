package main

import (
	"github.com/ogundejoseph/docs-library/core/internal/controller"
	"github.com/ogundejoseph/docs-library/core/internal/model"
)

func main() {
	model.Init()
	controller.Start()
}