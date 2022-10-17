package main

import (
	"GoLang/Reptile/engine"
	"GoLang/Reptile/zhenai/parser"
)

// main function
func main() {
	engine.Run(engine.Request{
		Url:        "http://www.supei.com/www/searchcity_1.htm",
		ParserFunc: parser.ParseCityList,
	})
}
