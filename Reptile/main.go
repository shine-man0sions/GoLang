package main

import (
	"GoLang/Reptile/engine"
	"GoLang/Reptile/zhenai/parser"
)

// main function
func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
