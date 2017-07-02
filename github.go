package github

import (
	"time"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"encoding/json"
)

func SearchIssues(terms []string)(*IssuesSearchResult,error){
	q := url.QueryEscape(strings.Join(terms," "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil{
		return nil,err
	}

	if resp.StatusCode != http.StatusOK{
		resp.Body.Close()
		return nil,fmt.Errorf("search query failed:%s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result);
	err != nil{
		resp.Body.Close()
		return nil,err
	}
	resp.Body.Close()
	return &result,nil 
}
	


const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct{
	TotalCount int 'json:"total_count"'
	Items []*Issue
}

type Issue struct{
	Number int
	HTMLURL string 'json:"html_url"'
	Title string
	Statle string
	User *User
	CreatedAt time.Time 'json:"created_at"'
	Body string
}

type User struct{
	Login string
	HTMLURL string 'json:"html_url"'
}
/*
	上面的结构体 即使对应的json对象名是小写字母，每个结构体的成员也是声明为
	大小字母开头的,因为json成员和go结构体成员名字不相同
	因此需要 go语言结构体成员tag 来指定json名字
*/

/*
	文本和html模版

*/
/*
	函数
	函数可以将一个语句打包为一个单元
	实参 通过值的方式传递，因此函数的形参是实参的拷贝，对
	形参的改变不会影响到实参但 引用类型、如：slice, map, function,
	channel实参就可能被修改
	函数在有返回值的情况下 每个返回值被声明成一个局部变量，并根据返回值
	的类型 将其初始化为0
*/
