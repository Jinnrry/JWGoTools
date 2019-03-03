# JWGoTools
对Go语言的一些常用方法封装


POST示例

```go
package main

import (
"github.com/jiangwei1995910/JWGoTools/tools"
"fmt"
"net/http"
"net/url"
)

func main() {
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

	option.RandomUa = true
	resp, ee := tools.POST(option)

	fmt.Println(resp.Body)
	fmt.Println(ee)


}



```