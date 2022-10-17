package parser

import (
	"GoLang/Reptile/model"
	"io/ioutil"
	"testing"
)

// 测试

func TestParserProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "安静的雪")

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Name:   "安静的雪",
		Gender: "女",
		Age:    34,
		Height: 162,
		Weight: 57,
		Income: "3000-5000元",
	}

	if profile != expected {
		t.Errorf("result should contain 1"+"element; but was %v", result.Items[0])
	}

}
