package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/*
	传入url 发送请求
	返回请求后的页面信息
	返回错误信息 error
*/

var rateLimiter = time.Tick(100 * time.Microsecond)

func Fetch(url string) ([]byte, error) {

	<-rateLimiter

	// 发送http请求
	resp, err := http.Get(url)

	// 如果出错 return err
	if err != nil {
		return nil, err
	}

	// 延时机制
	defer resp.Body.Close()

	// 如果http 错误，返回code码提示信息
	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

// 判断页面code, 如果不是utf-8, 转换成utf-8

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
