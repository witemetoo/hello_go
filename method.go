package method
import (
	"time"
	"fmt"
	"math"
	"image/color"
	"bytes"
)

/*
	一个对象其实就是一个简单的值或者一个变量，在这个对象
	中会包含一些方法，而一个方法则是一个一个和特殊类关联de函数
	一个面向对象的程序会用方法来表达其属性和对应的操作，这样使用这个对象的用户就不用
	直接去操作对象，而是借助方法来做这些事情

	oop编程的关键两个关键点
	封装和组合
	方法的声明，在其名字之前放上一个变量，即是一个方法
*/

const day = 24 * time.Hour
fmt.Println(day.Seconds())

type Point struct{x,y float64}

func (p Point) Distance(q Point) float64{
	return math.Hypot(q.x - p.x, q.y - p.y)
}

/*
	在 Distance函数中附加的参数 P，叫做方法的接收器
	早期的面向对象语言留下的遗产调用一个方法称为 “向一个对象发送消息”
	
*/

p := Point{1,2}
q := Point{4,6}

fmt.Println(Distance(p,q))
fmt.Println(p.Distance(q))
/*
	上面的两个 Println 函数 第一个是调用包级别的函数geometry.Distance
	第二个 是调用 Point.Distance的方法
	p.Distance 的表达式叫做选择器，因为他会选择合适的对象
	p这个对象的 Distance方法来执行
	在go中可以给任何 类型定义方法 可以为简单的数值、字符串、
	slice、map、来定义一些附加的行为很方便。方法可以被声明到任意类型 只要不是
	一个指针或者一个interface

	基于指针对象的方法
	当调用一个函数时、会对其每个参数值进行拷贝，如果一个函数需要更新一个变量
	或者其中的一个参数太大我们希望能够避免进行这种默认的拷贝
	为此我们需要用到指针
*/

func (p *Point) ScaleBy(factor float64){
	p.x *= factor
	p.y *= factor
}

/*
	只有类型point和指向他们的指针(*Point),才可能会出现在
	接收器声明里的两种接收器
	如果一个类型名本身是一个指针的话，是不允许出现在接收器中的
*/
 type P *int
func (P) f(){} //complie error:invalid receiver type
/*
	通过嵌入结构体来扩展类型
*/
type Point struct{x,y float64}
type ColorPoint struct{
	Point
	Color color.RGBA
}
/*
	方法值和方法表达式
	p.Distance叫做选择器，选择器会返回一个方法“值” 一个将方法
	绑定到特定接收器变量的函数，这个函数可以不通过特定指定其
	接收器即可被调用
*/

p := Point{1,2}
q := Point{4,6}

distanceFromP := p.Distance
fmt.Println(distanceFromP(q))
var origin Point
fmt.Println(distanceFromP(origin))

scaleP := p.ScaleBy
scaleP(2)

/*
	一个bit数组通常会用一个无符号或者称之为 ‘字’的
	slice 来表示，每个元素的每一位都表示集合里的一个值
	当集合的第i位被设置时，表示这个集合包含元素i
*/

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x % 64)
	return word < len(s.words) && s.words[word] & (1 <<bit) != 0
}

func (s * IntSet) Add(x int){
	word, bit := x/64, uint(x%64)
	for word >= len(s.words){
		s.words = append(s.words, 0)
	}
	s.words[word] != 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet){
	for i,tword := range t.words{
		if i < len(s.words){
			s.words[i] |= tword
	}else{
		s.words = append(s.words,tword)
	}
	}
}
/*
	封装 
	一个对象的变量或者方法如果对调用方是不可见的话
	一般就被定义为封装
	go语言中只有一种控制可见性的手段：大写字母的标识符会
	从定义它们的包中被导出 小写字母则不会，这种限制包括成员de
	方式同样适用于struct或者一个类型的方法 如果我们想封装一个对象 必须dingyi一个struct

*/

/*
	接口
	接口类型是对其他类型行为的抽象概括；因为接口类型不会
	和特定的实现细节绑定 在一起，通过这种抽象的方式我们可以让
	函数更加灵活和更具有适应能力

	接口类型不会暴露出它所代表的对象的内部的结构和这个对象支持的基础
	操作的集合；他们只会展示他们自己的方法
*/
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error){
	*c += ByteCounter(len(p))
	return len(p), nil
}

/*
	接口类型具体描述了一系列方法的集合 一个实现这些方法的具体类型是
	这个接口类型的实例
	接口指定规则：表达一个类型属于某个接口只要这个类型实现
	这个接口
	interface{}空接口类型，因为空接口类型对实现它的类型
	没有要求，所以我们可以将任意一个值赋给空接口类型
*/
var any interface{}
any = true
any = 12.34
any = "hello"
any = map[string]int{"one":1}
any = new(bytes.Buffer)


/*
	goroutines
	go中每个并发执行单元叫做一个goroutline
	语法支持：go语句是一个普通的函数或者方法调用前加上
	关键字go。go语句会使其语句在一个新创建的goroutine中运行

*/
