package main

import (
	"GoLang/Reptile/engine"
	"GoLang/Reptile/scheduler"
	"GoLang/Reptile/zhenai/parser"
)

// main function
func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
