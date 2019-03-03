# JWGoTools
对Go语言的一些常用方法封装


POST、GET示例

```go
package main

import (
	"github.com/jiangwei1995910/JWGoTools/tools"
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	// 设置请求参数
	option := tools.RequestBuild{}
	option.RequestUrl = "http://34.80.94.8/test.php"
	option.Parameter = url.Values{
		"aa": []string{"nn"},
		"bb":[]string{"cc"},
	}
	option.Headers = map[string]string{
		"Ckpacknum": "2",
	}

	//option.SsProxy="127.0.0.1:1080"

	option.HttpProxy="http://127.0.0.1:1081"

	option.Cookies = []http.Cookie{
		http.Cookie{Name: "name", Value: "V"},
		http.Cookie{Name: "name2", Value: "V"},
	}
	//使用随机UA
	option.RandomUa = true
	
	//发送POST请求
	resp, ee := tools.POST(option)

	fmt.Println(resp.Body)
	fmt.Println(ee)
	
	//发送GET请求
	resp2, ee2 := tools.GET(option)

	fmt.Println(resp2.Body)
	fmt.Println(ee2)
}




```