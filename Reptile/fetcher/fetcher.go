package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
	返回[] byte 具体内容
*/

func Fetch(url string) ([]byte, error) {

	// 获取get请求结果
	resp, err := http.Get(url)

	// 如果请求错误，返回错误信息
	if err != nil {
		return nil, err
	}

	// 延迟机制，在当前函数执行之后才执行
	defer resp.Body.Close()

	// 判断http返回状态码的值，如果状态吗错误，返回状态码
	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("Wrong status code: %d", resp.StatusCode)
	}

	// 如果返回数据有错，返回错误
	if err != nil {
		log.Printf("Fetcher error: %d\n", err)
		return nil, err
	}

	// 返回抓包结果
	return ioutil.ReadAll(resp.Body)
}
