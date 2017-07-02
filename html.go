package html

import (
	"io"
	"fmt"
	"os"
	"golang.org/x/net/html"
)

type Node struct{
	Type NodeType
	Data string
	Attr []Attribute
	FirstChild,NextSibling *Node
}

type NodeType int32

const(
	ErrorNode NodeType = iota
	TextNode 
	DocumentNode
	CommentNode
	DoctypeNOde
)
type Attribute struct{
	Key,Val string
}

func Parse(r io.Reader) (*Node, error)

func visit(links []string, n *html.Node) []string{
	if n.Type == html.ElementNode && n.Data == "a"{
		for _,a := range n.Attr{
			if a.Key == "href"{
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling{
		links = visit(links, c)
	}
	return links
}

func main(){
	for _,url := range os.Args[1:]{
		links,err := findLinks(url)
		if err != nil{
			fmt.Fprintf(os.Stderr,"findlinks2:%v\n",err)
			continue
		}
		for _,link := range links{
			fmt.Println(link)
		}
	}
}

func findLinks(url string) ([]string, error){
	resp,err := htpp.Get(url)
	if err != nil{
		return nil,err
	}
	if resp.StatusCode != http.StatusOk{
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s:%s", url,resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil{
		return nil, fmt.Errorf("parsing %s as HTML:%v", url,err)
	}
	return visit(nil,doc), nil 
}
/*
	如果一个函数所有的返回值都显示的变量名，那么该函数的return语句可以省略
	操作数

	对于大部分函数而言，永远无法确保能否成功运行
	这是因为错误的原因超出了程序员的控制
	任何运行i/o操作的函数都会棉铃出现错误的可能，
*/
/*
	在go中函数被看作第一类值：函数像其他值一样，拥有类型，
	可以被赋值给其他变量，传递给函数，从函数中返回

	匿名函数
	拥有匿名函数的函数只能在包级语法块中被声明，通过函数字面量

	在函数中定义的匿名函数可以访问该函数的完整的词法环境
	意味着在函数中定义的内部函数可以引用该函数的变量
*/

func squares() func() int{
	var x int 
	return func() int {
		x ++
		return x*x
	}
}