package parser

import (
	"log"
	"regexp"
)

// 城市列表解析器 城市名称+url
func ParseCityList(contents []byte) engine.ParseResult {

	// 使用正则匹配匹配出名称和url
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`)

	matches := re.FindAllSubmatch(contents, -1)

	// 循环打印城市列表
	for _, m := range matches {
		log.Printf("City: %s, URL: %s", m[1], m[2])
	}
}
