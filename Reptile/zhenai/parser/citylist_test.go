package parser

import (
	"io/ioutil"
	"testing"
)

// 测试citylist

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)

	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	expectedCities := []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}

	if len(result.Request) != resultSize {
		t.Errorf("result should have %d"+"request; but had %d", resultSize, len(result.Request))
	}

	for i, url := range expectedUrls {
		if result.Request[i].Url != url {
			t.Errorf("expected url #%d: %s; but "+"was %s", i, url, result.Request[i].Url)
		}
	}

	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected city #%d: %s; but "+"was %s", i, city, result.Items[i].(string))
		}
	}

}
