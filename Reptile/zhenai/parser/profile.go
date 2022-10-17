package parser

import (
	"GoLang/Reptile/engine"
	"GoLang/Reptile/model"
	"log"
	"regexp"
)

var birthRe = regexp.MustCompile(`<td class=uservalue>([^<]+)&nbsp;[^>]+</td>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}

	profile.Name = name

	profile.Birth = extractString(contents, birthRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	log.Printf("---------%v", result)
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
