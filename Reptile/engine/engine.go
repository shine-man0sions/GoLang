package engine

import (
	"GoLang/Reptile/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		// 打印城市的url
		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error"+"fetching url %s: %v", r.Url, err)
			continue
		}
		ParseResult := r.ParserFunc(body)
		requests = append(requests, ParseResult.Request...)

		for _, item := range ParseResult.Items {
			log.Printf("got item %v", item)
		}
	}
}
