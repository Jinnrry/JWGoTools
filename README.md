# JWGoTools
对Go语言的一些常用工具方法封装


## 简单POST、GET示例

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
	
	//请求地址
	option.RequestUrl = "http://34.80.94.8/test.php"
	//设置参数
	option.Parameter = url.Values{
		"aa": []string{"nn"},
		"bb":[]string{"cc"},
	}
	//请求头
	option.Headers = map[string]string{
		"Ckpacknum": "2",
	}

    // Socket代理
	//option.SsProxy="127.0.0.1:1080"

    //Http或Https代理
	option.HttpProxy="http://127.0.0.1:1081"

    //自定义Cookie
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



## 保存Session状态

```go
package main

import (
	"JWGoTools/tools"
	"fmt"
	"net/http"
)

func main() {
	option := tools.RequestBuild{}
	option.RequestUrl = "http://myhttpheader.com/"
	
	option.Cookies = []http.Cookie{
		{Name: "cookie", Value: "v"},
	}


	res, _ := tools.GET(option)
	//获取本次请求的全部Cookie，包括自定义和服务端返回
	fmt.Println(res.Cookies)
	
	option2:=tools.RequestBuild{}
	option2.RequestUrl="http://myhttpheader.com/"
	
	//将上一个请求的Cookie设置进去可实现保存Session等状态
	option2.Cookies=res.Cookies

	res,_=tools.POST(option2)

	fmt.Println(res.Cookies)

}


```



## 发送一封邮件


```go
	tools.SendToMail("xxx@qq.com","password","smtp.qq.com:465","xxx@163.com","title","<html>body</html>","html")
```