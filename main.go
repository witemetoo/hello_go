package main

// import (
// 	"os"
// 	"fmt"
// 	"log"
// 	"io"
// 	"sync"
// )

// import "fmt"

// func main(){
// 	fmt.Println("hello,世界")
// }

// import (
// 	// "fmt"
// 	// "os"
// 	// "bufio"
// 	// "strings"
// 	// "io/ioutil"
// 	// "image"
// 	// "image/color"
// 	// "image/gif"
// 	// "math"
// 	// "math/rand"
// 	// "io"
// 	// "net/http" 

// 	// 并发获取多个url
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// 	"time"
// 	"bufio"
// 	"unicode/utf8"
// 	"unicode"
// 	"encoding/json"
// 	"log"
// )
// // := 短变量声明
// func main(){
// 	// var s,sep string
// 	// for i := 1; i<len(os.Args);i++{
// 	// 	s += sep + os.Args[i]
// 	// 	sep = " "
// 	// }
// 	// fmt.Println(s)
// 	// other_echo()
// 	// dup1()
// 	// get_url()
// 	get_mut_url()
// }


// // other_echo 函数循环的每次迭代，range产生一对值：索引以及该处的索引值
// // 由于例中不需要索引 所以把它赋予临时变量temp 但go 中不使用无用的临时变量 所以用空标识符_ 
// func other_echo(){
// 	s,sep := "", ""
// 	for _,arg := range os.Args[1:]{
// 		s += sep + arg
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

// // func dup1(){
// // 	counts := make(map[string]int)
// // 	input := bufio.NewScanner(os.Stdin)
// // 	for input.Scan(){
// // 		counts[input.Text()] ++
// // 	}
// // 	for line,n :=range counts{
// // 		if n>1 {
// // 			fmt.Printf("%d\t %s\n",n,line)
// // 		}
// // 	}
// // }

// // // os.open 函数返回两个值 ，一个是被打开的文件 第二个是内置的error类型的值
// // // 函数和包级别变量可以任意顺序声明，并不影响其使用
// // // map是由make函数创建的数据结构引用，map作为参数传递给函数时 函数接受的是这个引用的拷贝
// // func dup2(){
// // 	counts := make(map[string]int)
// // 	files := os.Args[1:]
// // 	if len(files) == 0{
// // 		countLines(os.Stdin,counts)
// // 	}else{
// // 		for _,args := range files{
// // 			f,err := os.Open(args)
// // 			if err != nil{
// // 				fmt.Fprintf(os.Stderr, "dup2 :%v\n", err)
// // 				continue
// // 			}
// // 			countLines(f,counts)
// // 			f.Close()
// // 		}
// // 	}
// // 	for line,n := range counts{
// // 		if n>1{
// // 			fmt.Printf("%d \t %s\n", n,line)
// // 		}
// // 	}
// // }

// // func countLines(f*os.File,counts map[string]int){
// // 	input := bufio.NewScanner(f)
// // 	for input.Scan(){
// // 		counts[input.Text()]++
// // 	}
// // }

// // func dup3(){
// // 	counts := make(map[string]int)
// // 	for _,filename := range os.Args[1:]{
// // 		data,err := ioutil.ReadFile(filename)
// // 		if err != nil{
// // 			fmt.Fprintf(os.Stderr,"dup3:%v\n",err)
// // 			continue
// // 		}
// // 		for _,line := range strings.Split(string(data), "\n"){
// // 			counts[line] ++
// // 		}
// // 	}
// // 	for line,n := range counts {
// // 		if n>1{
// // 			fmt.Printf("%d\t%s\n", n, line)
// // 		}
// // 	}
// // }

// // // 复合声明 color.Color()
// // var palette = []color.Color{color.White,color.Black}

// // const (
// // 	whiteIndes = 0
// // 	blackIndex = 1
// // )

// // // func main(){
// // // 	lissajous(os.Stdout)
// // // }

// // func lissajous(out io.Writer){
// // 	// 程序中的常量声明给出了一系列的常量值 常量是指在程序编译后始终都不变的值
// // 	// 常量声明和变量声明一般都会出现在包 级别 所以常量在整个包中都是可以共享的
// // 	const(
// // 		cycles = 5
// // 		res = 0.001
// // 		size = 100
// // 		nframes = 64
// // 		delay = 8
// // 	)

// // 	freq := rand.Float64() * 3.0
// // 	// git.GIF{}为复合声明,会把loopcount设置为nframes anim是是符合声明的结构体
// // 	anim := gif.GIF{LoopCount:nframes}
// // 	phase := 0.0
// // 	for i := 0; i < nframes; i++{
// // 		rect := imag.Rect(0,0,2*size+1,2*size+1)
// // 		img := image.NewPaletted(rect,palette)
// // 		for t:= 0.0; t < cycles*2*math.Pi; t+=res{
// // 			x := math.Sin(t)
// // 			y := math.Sin(t * freq + phase)
// // 			img.SetColorIndex(size + int(x*size + 0.5), size + int(y * size + 0.5),
// // 			blackIndex
// // 			)
// // 		}
// // 		phase += 0.1
// // 		anim.Delay = append(anim.Delay,delay)
// // 		anim.Image = append(anim.Image,img)
// // 	}
// // 	gif.EncodeAll(out,&anim)
// // }


// // 获取url
// func get_url(){
// 	for _,url := range os.Args[1:]{
// 		resp, err := http.Get(url)
// 		if err != nil{
// 			fmt.Fprintf(os.Stderr, "fetch:%v\n",err)
// 			os.Exit(1)
// 		}
// 		b, err := ioutil.ReadAll(resp.Body)
// 		resp.Body.Close()
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr,"fetch: reading %s: %v\n", url,err)
// 			os.Exit(1)
// 		}
// 		fmt.Printf("%s",b)
// 	}

// }


// // 并发获取多个url 并发获取多个url的时间不会超过 最长的单个时间

// func get_mut_url(){
// 	start := time.Now()
// 	ch := make(chan string)
// 	for _,url := range os.Args[1:]{
// 		go fetch(url,ch)
// 	}
// 	for range  os.Args[1:]{
// 		fmt.Println(<-ch)
// 	}
// 	fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds())
// }
// func fetch(url string,ch chan<- string){
// 	start := time.Now()
// 	resp,err := http.Get(url)
// 	if err != nil{
// 		ch <- fmt.Sprint(err)
// 		return
// 	}
// 	nbytes,err := io.Copy(ioutil.Discard,resp.Body)
// 	resp.Body.Close()
// 	if err != nil{
// 		ch <- fmt.Sprintf("while reading %s:%v" ,url,err)
// 		return
// 	}
// 	secs := time.Since(start).Seconds()
// 	ch <- fmt.Sprintf("%.2fs %7d %s", secs,nbytes,url)
// }


// /*
// 	go 语言中函数名、变量名、常量名、类型名、语句标号和包名 

// 	如果一个名字是大写字母开头（必须是函数外部定义的包级名字；包级函数名本身
// 	也是包级名字）

// 	声明：
// 	声明语句定义了程序的各种实体对象一级部分或全部属性 
// 	go 语言有四种类型声明语句：var const type func分别对应：变量、常量、类型、函数
// 	每个源文件以包名的声明语句开始，说明该源文件属于哪个包

// 	包一级声明语句声明的名字可在整个包对应的每个源文件中访问
// 	而不仅仅是在其声明语句所在的源文件中访问


// 	var 声明语句可以创造一个特定类型的变量，然后变量附加一个名字，并且设置变量的初始化值
// 	var 变量名字 类型＝表达式
// 	如果初始化表达式被省略，数值类型变量对应的零值是0 布尔类型对应的零值是false
// 	字符串类型对应的零值是空字符串，接口或者引用类型(slice.map.chan)
// 	变量对应的零值是nil 
// 	数组或者结构体等聚合类型对应的零值是每个元素对应该类型的零值

// 	var i,j,k int
// 	var b,f,s = true,2.3,"four"

// 	一组变量也可以通过调用一个函数，由函数返回的多个返回值初始化
// 	var f,err = os.Open(name)

// 	简短变量声明
// 	简短变量声明语句的形式可用于声明和初始化局部变量 变量的类型根据表达式来自动推导

// 	简短变量声明语句也可以用函数返回值来声明和初始化变量
// 	f,err := os.Open(name)
// 	在某种情况下 简短变量声明左边变量可能并不是全部都是刚刚声明的 那么简短声明语句只有赋值行为来
// 	in,err := os.Open(infile)
// 	...
// 	out,err := os.Create(outfile)
// 	简短声明语句至少包含一个新的声明变量 下面的代码不能编译通过
// 	f,err := os.Open(infile)
// 	...
// 	f,err := os.Create(outfile)


// 	指针
// 	一个变量对应一个保存了 变量对应类型值 的内存空间
// 	变量有时候被称为可寻址的值 任何类型的指针的零值都是 nil 如果 p! =nil 
// 	测试为真 那么p是指向 某个有效变量 

// 	new函数
// 	new()函数将创建一个T类型的匿名变量，初始化为T类型的零值
// 	然后返回变量地址，返回指针类型为*T


// 	变量的生命周期
// 	变量的生命周期指的是程序运行期间变量有效存在的时间间隔
// 	包级变量的生命周期和整个程序的运行周期是一致的
// 	而局部变量的声明周期则是动态的：从每次创建一个新变量的声明语句开始
// 	直到该变量不在被引用为止
// 	因为一个变量的有效周期只取决于是否可达，因此一个循环迭代内部的局部
// 	变量的生命周期可能超出其局部作用域。同时，局部变量可能在函数返回之后依然存在
// 	变量生命周期是从时间上描述变量 作用域是从空间上描述变量


// 	类型
// 	变量或表达式的类型定义了对应存储值的属性特征
// 	一个类型声明语句创建了一个新的类型名称，和现有类型具有相同的底层结构
// 	新命名的类型提供了一个方法，用来分隔不同概念的类型
// 	type 类型名 底层类型

// 	包和文件
// 	go语言的包和其他语言的库或者模块的概念类似
// 	都是为了支持模块化、封装、单独编译和代码重用
// 	每个源文件都是以包的声明语句开始的，用来指名包的名字。当包
// 	被导入时

// 	包的初始化
// 	包的初始化首先就是解决包级变量的依赖顺序，然后按照包级变量出现
// 	的顺序依次初始化。如果包中含多个.go文件 go 语言的构建工具首先将文件根据
// 	文件名排序。然后依次调用编译器进行编译

// 	作用域
// 	一个声明语句将程序中的尸体和一个名字关联
// 	声明语句的作用域是指源代码中可以有效使用这个名字的范围
// 	语法块是由华括号包含的一系列语句


// 	第三章 基础数据类型
// 	go 语言将数据类型分为四类：基础类型、复合类型、引用类型、接口类型
// 	int8 int16 int32 int64四种截然不同大小的有符号整数类型
// 	uint8 uint16 uint32 uint64四种无符号整数类型

// 	尽管go 提供了无符号数和运算，即使数值本身不可能出现负数还是倾向于使用有符号的int
// 	类型，就像数组长度那样 

// 	浮点数
// 	go 语言提供了两种精度的浮点数，float32和float64
// 	math.MaxFloat32 表示float32能表示最大的数
// 	math.MaxFloat64 表示float64能表示最大的数
// 	复数
// 	go语言提供了两种精度的复数类型：complex64和coplex128

// 	布尔类型
// 	and or 有短路行为：如果运算符左边值已经可以确定整个布尔类型的
// 	值，那么运算符右边的值将不在被求值
// 	下面表达式是安全的
// 	x !="" && s[0] == 'x'
// 	布尔类型不会隐式的转化为 0、1
// 	func btoi(b bool) int{
// 		if b{
// 			return 1
// 		}
// 		return 0
// 	}

// 	字符串
// 	字符串是一个不可以改变的字节序列
// 	字符串的值是不可变的：一个字符串包含的字节序列永远不会被改变
// 	对于普通中文字符串 每个汉字占用3个字节
// 	对于rune 汉字占用一个字节 其他字符也是占用一个字节


// 	字符串面值
// 	字符串值也可以用字符串面值方式编写，
	
// 	字符串和byte切片
// 	标准库中有四个包对字符串处理：
// 	bytes:
// 	string:包含许多如字符串查询、替换、比较、拆分、截断、合并
// 	strconv 提供布尔类型、整数类型、浮点数类型、对应的字符串转化

// 	常量
// 	常量表达式的值是在编译期计算的，而不是在运行期。每种的潜在都是基础类型：
// 	boolen string 或者数字
// 	常量生成器：iota 类似于c中的枚举类型
// 	在一个const声明语句中，在第一个声明常量所在的行，iota将会被置为0
// 	然后每个有常量声明的行加1



// */
// func comma(s string) string{
// 	n := len(s)
// 	if n <= 3{
// 		return s 
// 	}
// 	return comma(s[:n-3] + "," + s[n-3:])
// }

// type Weekday int
// const (
// 	Sunday Weekday = iota
// 	Monday
// 	Tuesday
// 	Wednesday
// 	Thursday
// 	Friday
// )

// /*
// 	复合类型数据
// 	数组：
// 	数组是由一个由固定长度的特定类型元素组成的序列
// 	因为数组长度是固定的，所以在go 语言中常使用的是slice

// 	var a [3]int array of 3 integers
	
// 	当调用一个函数的时候，函数的每个调用参数将会被赋值给函数内部的
// 	参数变量，所以函数接收的是复制的副本，并不是调用的变量，
// 	所以函数参数传递的机制导致传递大的数组类型是低效的

// 	slice（切片）
// 	切片代表变长的序列，
// 	slice 底层引用了数组对象
// 	一个slice由三部分组成，指针、长度、容量 slice之间不能进行比较 对于
// 	字节类型 则可以通过标准库 bytes.Equal来进行比较

// 	内置的make函数创建一个指定元素类型、长度和容量的slice
// 	在底层make 创建了一个匿名的数组变量，然后返回一个slice；
// 	之友通过返回的slice才能一用底层匿名的数组变量
// 	make([]T,len)
// 	make([]T, len, cap)
// 	容量部分可以省略
// 	append 函数用于向slice追加元素

// 	在循环中使用使用append函数 构建一个由九个rune 字符构成
// 	的slice 我们可以通过go语言内置的[]rune("hello,世界")
// 	转化操作完成。append 函数理解slice
	



// */
// var runes []rune
//  for _,r := range "hello world"{
// 	 runes = append(runes,r)
//  }
//  fmt.Printf("%q\n", runes)

// var str string
// str = "hello world"
// ch := str[0]
// len := len(str)

// func appendInt(x []int, y int) []int{
// 	var z []int
// 	zlen := len(x) + 1
// 	if zlen <= cap(x){
// 		z = x[:zlen]
// 	}else {
// 		zcap := zlen
// 		if zcap < 2 * len(x){
// 			zcap = 2 * len(x)
// 		}
// 		z = make([]int,zlen,zcap)
// 		copy(z,x)
// 	}
// 	z[len(x)] = y
// 	return z
// }

// /*
// nonempty 函数
// */

// func nonempty(strings []string) []string{
// 	i := 0
// 	for _,s := range strings{
// 		if s!= ""{
// 			strings[i] = s
// 			i ++
// 		}
// 	}
// 	return strings[:i]
// }

// /*
// 	map 
// 	内置的make 函数可以创建一个map
// 	ages := make(map[string]int)
// 	下面两种初始化相同的map
// 	ages := make(map[string]int)
// 	ages["alice"] = 31
// 	ages["charlie"] = 34


// 	ages := map[string]int{
// 		"alice":31,
// 		"charlie":34,
// 	}
// 	map 中的元素并不是一个变量，我们不能对map的元素进行取值操作
// 	_= &ages["bob"]

// 	if age,ok := ages["bob"]; !ok{}

// */

// import "sort"

// var names []string
// for name := range ages{
// 	fmt.Printf("%s\t%d\n", name,ages[name])
// }

// func equal(x,y map[string]int) bool{
// 	if len(x) != len(y){
// 		return false
// 	}
// 	for k,xv = range x{
// 		if yv,ok := y[k]; !ok || yv != xv{ // 判断在y中是否存在 k if !ok 或者 存在但value值不为xv
// 			return false
// 		}
// 	}
// 	return true
// }

// /*
// 	创建一个容量为len(ages) 长度的slice
// 	names := make([]string,0,len(ages))


// */

// func dedup(){
// 	seen := make(map[string]bool)
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan(){
// 		line := input.Text()
// 		if !seen[line]{
// 			seen[line] = true
// 			fmt.Println(line)
// 		}
// 	}
// 	if err := input.Err(); err != nil{
// 		fmt.Fprintf(os.Stderr,"dedup:%v\n",err)
// 		os.Exit(1)
// 	}
// }

// /*
// 有时候我们需要一个map 或 set的key是slice类型 但是map的key必须是
// 可以比较的类型，但 slice并不满足这个条件
// */

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// 	"unicode"
// 	"unicode/utf8"
// )

// func main1(){
// 	counts := make(map[rune]int)
// 	var utflen [utf8.UTFMax + 1]int 
// 	invalid := 0
// 	in := bufio.NewReader(os.Stdin)
// 	for{
// 		r,n,err := in.ReadRune()
// 		if err == io.EOF{
// 			break
// 		}
// 		if err != nil{
// 			fmt.Fprintf(os.Stderr, "char count:%v\n", err)
// 			os.Exit(1)
// 		}
// 		if r == unicode.ReplacementChar && n == 1{
// 			invalid ++
// 			continue
// 		}
// 		fmt.Printf("rune\tcount\n")
// 		for c,n := range counts{

// 		}
// 	}

// }
// /*
// 	结构体的访问
// 	是一种聚合的数据类型，是由零个或者多个任意类型的值
// 	聚合成的实体。每个值称为结构体的成员
// 	dilbert.Salary == 5000
// 	position := &dilbert.Position
// 	*position = "senior" + *position

// 	var employeeOfTheMonth *Employee = &dilbert
// 	employeeOfTheMoneth.Position += " (proactive team player)"


// 	一个命名为s的结构体将不能包含s类型的成员：因为一个聚合的值不能包含它自身
// 	但是s类型的结构体可以包含*s指针类型的成员,因此可以创建递归的数据结构

// */
// type Employee struct {
// 	ID int
// 	Name string
// 	Address string
// 	DoB time.Time
// 	Position string
// 	Salary int
// 	ManagerID int
// }
// var dilbert Employee

// // 使用二叉树来实现一个插入排序
// type tree struct {
// 	value int
// 	left,right *tree
// }

// func Sort(value []int){
// 	var root * tree
// 	for _,v := range values{
// 		root = add(root ,v)
// 	}
// 	appendValues(values[:0], root)
// }

// func appendValues(values []int, t *tree){
// 	if t != nil{
// 		values = appendValues(values, t.left)
// 		values = append(values, t.value)
// 		values = appendValues(values,t.right)
// 	}
// }

// func add(t *tree, value int) *tree{
// 	if t == nil{
// 		t = new(tree)
// 		t.value = value
// 		return t
// 	}
// 	if value < t.value{
// 		t.left = add(t.left,value)
// 	}else{
// 		t.right = add(t.right, value)
// 	}
// 	return t
// }
// /*
// 	go 语言中的new 和 make是两个用于分配内存的原语
// 	new只分配内存 make用于slice、map、和channel的初始化

// 	new 函数并不初始化内存，只是将其置零 new(T) 荟蔚T类型的项目
// 	分配被志零的存储 并返回它的地址

// 	go 语言中没有c++中的构造函数，对象的创建为一个全局的创建函数来完成
// 	func NewRect(x,y,witch,height float64) *Rect{
// 		return &Rect(x,y,width,height)
// 	}

// 	make 函数make(T, args) 只用来创建slice，map和channel,
// 	并返回一个初始化 类型为T的值 

// */

// type line struct{
// 	x int
// }
// type plane struct {
// 	line
// 	y int
// }
// type space struct{
// 	plane
// 	z int
// }

// t := space{plane{line{3},5},7}

// /*
// 	结构体嵌入和匿名成员
// 	go 语言特性只声明一个成员对应的数据类型而不指定成员的名字
// 	这类成员就叫匿名成员 匿名成员的数据类型必须是命名的类型或者
// 	指向一个命名的类型的指针
	

// */
// type Point struct {
// 	x, y int
// }
// type circle struct{
// 	center Point
// 	Radius int
// }

// type Wheel struct{
// 	circle circle
// 	Spokes int
// }
// var w Wheel
// w.circle.center.x = 8
// w.circle.center.y = 8
// /*
// 	匿名成员访问
// */
// type circle struct{
// 	Point
// 	Radius int
// }
// type wheel struct{
// 	circle
// 	Spokes int
// }

// var w wheel
// w.x = 8

// /*
// 	json 
// 	基础类型可以通过json的数组和对象进行递归组合
// 	一个json数组是一个有序的值序列
// */
// type Movie struct {
// 	Title string
// 	Year int 'json:"released"'
// 	Color bool 'json:"color, omitempty"'
// 	Actors []string
// }

// var mobies = []Movie{
// 	{
// 		Title:"casablanca",
// 		Year:"1942",
// 		Color:false,
// 		Actors:[]string{"Humphrey Bogart", "Ingrid Bergman"}
// 	}
// }
// /*
// 	将类似movies的结构体 slice转为json的过程叫编组
// 	编组通过调用json.Marshal函数完成
// 	在编码时、默认使用go语言结构体的成员名字作为json的对象
// 	只有导出的结构体成员才会被编码
// 	编码的逆操作是解码,对应将json数据解码为go语言的数据结构
// 	go 中叫做 unmarshaling

// 	结构体的成员Tag 可以是任意的字符串面值，通常是key:“value”
// 	键值对序列
// */
// data,err := json.Marshal(movies)
// if err!= nil{
// 	log.Fatal("json marshaling failed: %s",err)
// }
// fmt.Printf("%s\n",data)

// import (
// 	"fmt"
// 	"time"
// )
// func main(){
// 	go spinner(100 * time.Millisecond)
// 	const n = 45
// 	fibN := fib(n)
// 	fmt.Printf("\rFibonacci(%d) = %d\n",n,fibN)
// }

// func spinner(delay time.Duration){
// 	for{
// 		for _,r := range `-\|/`{
// 			fmt.Printf("\r%c",r)
// 			time.Sleep(delay)
// 		}
// 	}
// }

// func fib(x int) int{
// 	if x <2 {
// 		return x
// 	}
// 	return fib(x-1) + fib(x-2)
// }

/*
	主函数返回时，所有的goroutline都会直接被打断
	除了从主函数退出或者直接终止程序之外，没有其他的编程方法
	能够让一个goroutine来打断另一个执行
*/

/*
	并发的clock服务

	由于服务器是典型的需要同时处理很多链接的程序，这些程序一般
	来自彼此独立的客户端
	下面是一个顺序执行的时钟服务器，它每隔一秒将当前时间写到客户端
*/

// import (
// 	"io"
// 	"log"
// 	"net"
// 	"time"
// 	"fmt"
// 	"strings"
// 	"bufio"
// 	"os"
// )


// func main(){
// 	listener, err := net.Listen("tcp", "localhost:8000")
// 	if err != nil{
// 		log.Fatal(err)
// 	}
// 	for {
// 		conn,err := listener.Accept()
// 		if err != nil{
// 			log.Print(err)
// 			continue
// 		}
// 		go handleConn(conn)
// 	}
// }

// func handleConn( c net.Conn){
// 	defer c.Close()
// 	for {
// 		_,err := io.WriteString(c, time.Now().Format("15:04:05\n"))
// 		if err != nil{
// 			return
// 		}
// 		time.Sleep(1 * time.Second)
// 	}
// }

// import (
// 	"io"
// 	"log"
// 	"net"
// 	"os"
// )
// func main(){
// 	conn,err := net.Dial("tcp", "localhost:8000")
// 	if err != nil{
// 		log.Fatal(err)
// 	}
// 	defer conn.Close()
// 	mustCopy(os.Stdout, conn)
// }

// func mustCopy(dst io.Writer, src io.Reader){
// 	if _,err := io.Copy(dst, src); err != nil{
// 		log.Fatal(err)
// 	}
// }

/*
	clock服务器每一个连接都会起一个goroutine
*/

// func echo(c net.Conn, shout string, delay time.Duration){
// 	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
// 	time.Sleep(delay)
// 	fmt.Fprintln(c, "\t", shout)
// 	time.Sleep(delay)
// 	fmt.Fprintln(c, "\t", strings.ToLower(shout))
// }

// func handleConn(c net.Conn){
// 	input := bufio.NewScanner(c)
// 	for input.Scan(){
// 		echo(c, input.Text(), 1 * time.Second)
// 	}
// 	c.Close()
// }

/*
	channels 
	goroutine是go语言程序的并发体，channels他们之间的
	通信机制，一个channels是一个通信机制
*/

// ch := make(chan int) // ch has type 'chan int'
/*
	channel 对应make创建的底层数据结构的引用，复制一个channel
	或者用于函数参数传递时， 只是拷贝了一个channel引用
	channel 的零值也是nil
*/
// ch <- x //a send statemnet
// x = <- ch // a receive expreeion in an assignment statement
// <- ch  //a receive statement; result is discarded

/*
	不带缓存的chanels 

	基于无缓存的channels的发送和接收操作将导致两个goroutine
	做一次同步操作，所以无缓存的channels也叫做同步channels。当通过
	一个无缓存channels发送数据时 接受者收到数据发生在唤醒发送者g
	gotoutine之前
*/


// func main(){
// 	conn, err := net.Dial("tcp", "localhost:8000")
// 	if err != nil{
// 		log.Fatal(err)
// 	}
// 	done := make(chan struct{})
// 	go func(){
// 		io.Copy(os.Stdout, conn)
// 		log.Println("done")
// 		done <- struct{}{}
// 	}()
// 	mustCopy(conn,os.Stdin)
// 	conn.Close()
// 	<- done
// }
/*
	如果消息不携带任何额外的信息，作用仅是为了两个
	goroutine之间的同步 我们可以用struct{}空结构体作为
	channels 元素类型

	串联的channels(pipeline)
	channels 可以用于将多个goroutline链接在一起，一个channels
	输出作为下一个channels的输入，这就是所谓的管道
*/

// import (
// 	"fmt"
// )
// func main(){
// 	naturals := make(chan int)
// 	squares := make(chan int)

// 	go func(){
// 		for x := 0; x < 100 ; x++{
// 			naturals <- x
// 		}
// 		close(naturals)
// 	}()

// 	go func(){
// 		for  x := range naturals{
// 			squares <- x * x 
// 		}
// 	}()

// 	for x := range squares {
// 		fmt.Println(x)
// 	}
// }
/*
	类型  chan<- int 表示一个只发送int的 channel
	<- chan int 表示一个只接收int 的 channel 这种限制将在背编译期检测
*/

// import (
// 	"fmt"
// )

// func counter(out chan<- int){
// 	for x := 0; x < 100; x ++{
// 		out <- x
// 	}
// 	close(out)
// }

// func squarer(out chan <- int, in <- chan int){
// 	for v := range in{
// 		out <- v * v
// 	}
// 	close(out)
// }

// func printer(in <- chan int){
// 	for v := range in{
// 		fmt.Println(v)
// 	}
// }

// func main(){
// 	naturals := make(chan int)
// 	squares := make(chan int)
// 	go counter(naturals)
// 	go squarer(squares,naturals)
// 	printer(squares)
// }
/*
	上例中调用counter(naturals)将chan int 类型的naturals隐士
	的转化为chan <- int 类型 任何双向的channel向单向channel变量赋值操作
	都将导致隐士转化
*/


/*
	带缓存的channels 
	ch = make(chan string, 3)
	cap函数获取channel内部的容量
*/

// import (
// 	""
// )
// func makeThumbnails3(filename []string){
// 	ch := make(chan struct{})
// 	for _,f := range filename{
// 		go func(f string){
// 			thumbnail.ImageFile(f)
// 			ch <- struct{}{}
// 		}(f)
// 	}
// 	for range filenames{
// 		<- ch
// 	}
// }

// import (
// 	"fmt"
// 	"runtime"
// )

// func main(){
// 	fmt.Println("cpus", runtime.NumCPU())
// 	fmt.Println("goroot:", runtime.GOROOT())
// 	fmt.Println("archive:", runtime.GOOS)
// }

/*
	go import下划线的作用 当用下划线导入一个包时
	该包文件里的所有init函数都会被执行
*/

/*
	并发的web爬虫
*/

// import (
// 	"fmt"
// 	"os"
// )

// func main(){
// 	worklist := make(chan []string)

// 	//start with the command-line arguments
// 	go func(){worklist <- os.Args[1:]}()

// 	//crawl the web concurrently
// 	seen := make(map[string]bool)
// 	for list := range worklist{
// 		for _,link := range list{
// 			if !seen[link]{
// 				seen[link] = true
// 				go func(link string){
// 					worklist <- crawl(link)
// 				}(link)
// 			}
// 		}
// 	}
// }

// func crawl(url string) []string{
// 	fmt.Println(url)
// 	list, err := links.Extract(url)
// 	if err != nil{
// 		log.Print(err)
// 	}
// 	return list
// }


/*
	竞争条件 
	当我们没法确定一个事件是在另一个事件的前面或者后面发生的话
	就说明x和y这两个事件是并发的
	一个函数在线程程序中可以正常的工作，如果在并发的情况下
	这个函数依然可以正确的工作的话，那么程序就是并发安全的
	并发安全的函数不需要额外的同步工作

	包的初始化是在程序main函数开始执行之前就完成了额，
*/

// func main(){
// 	conn, err := net.Dial("tap", "localhost:8000")
// 	if err != nil{
// 		log.Fatal(err)
// 	}
// 	done := make(chan struct{})
// 	go func(){
// 		io.Copy(os.Stdout, conn)
// 		log.Println("done")
// 		done <- struct{}{}
// 	}()
// 	mustCopy(conn, os.Stdin)
// 	conn.Close()
// 	<- done
// }

/*
	等待goroutine的完成
*/

// func makeThumbnails3(filename []string){
// 	ch := make(chan struct{})
// 	for _, f := range filename{
// 		go func(f string){
// 			thumbnail.ImageFile(f)
// 			ch <- struct{}{}
// 		}(f)
// 	}
// 	for range filename{
// 		<- ch
// 	}
// }



// func makeThumbnails6(filenames <- chan string) int64{
// 	sizes := make(chan int64)

// 	var wg sync.WaitGroup
// 	for f := range filenames{
// 		wg.Add(1)

// 		go func(f string){
// 			defer wg.Done()
// 			thumb, err := thumbnail.ImageFile(f)
// 			if err != nil{
// 				log.Println(err)
// 				return
// 			}
// 			info,_ := os.Stat(thumb)
// 			sizes <- info.size()
// 		}(f)
// 	}

// 	go func(){
// 		wg.Wait()
// 		close(sizes)
// 	}()

// 	var total int64
// 	for size := range sizes{
// 		total += size
// 	}
// 	return total
// }


// /*
// 	无缓冲信道的数据是先到先出， 但是 无缓冲信道并不储存数据
// 	只负责数据流通
// */

// var ch chan int = make(chan int)

// func foo(id int){
// 	ch <- id
// }

// func main(){
// 	for i := 0; i < 5; i++{
// 		go foo(i)
// 	}

// 	for i := 0; i < 5; i++{
// 		fmt.Printf( <- ch)
// 	}
// }

/*
	等待多gorountine
	使用信道堵塞主线，等待开出去所有goroutine完成
*/

// import (
// 	"fmt"
// )
// var quit chan int

// func foo(id int){
// 	fmt.Println(id)
// 	quit <- 0
// }

// func main(){
// 	count := 1000

// 	quit = make(chan int) //无缓冲

// 	for i := 0; i < count; i++{
// 		go foo(i)
// 	}

// 	for i := 0; i < count; i++{
// 		<- quit
// 	}
// }

/*
	声明一个interface{}空接口，可接收任何类型
	万能切片[]interface{}
	任意map map[string]interface{}
*/
// import (
// 	"os"
// 	"fmt"
// 	"log"
// )

// func main(){
// 	file, err := os.Open("github.go")
// 	if err != nil{
// 		log.Fatal(err)
// 	}
// 	data := make([]byte, 100)
// 	count, err := file.Read(data)
// 	if err != nil{
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("read %d bytes : %q\n", count,data[:count])
// }

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"os"
	"path/filepath"
)

func sayhelloName(w http.ResponseWriter, r * http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Scheme)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form{
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello astaxie!")
}

func main(){
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil{
		log.Fatal("ListenAndServer:", err)
	}
}

/*
	Request 用户请求的信息，用来解析用户的请求信息，包括post、get、cookie、
	url
	Response：服务器反馈给客户端的信息
	Conn：用户每次请求链接
	Handler：处理请求和生成返回信息的处理逻辑
*/

/*
	slice和数组中的元素类型都相同，但可以通过定义 空接口interface{} 绕过这种限制
*/

/*
	有效切片索引遍历:直接修改底层数组
	复制切片遍历：

*/
slice1 = make([]int, 3)
for i := range slice1{
	break
}

func main(){
	if len(os.Args) == 1 || os.Args[1] == "-h" ||
	os.Args[1] == "--help" {
		fmt.Printf("usage:%s file\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	separators := []string{"\t", "*", "|", "."}
	linesRead, lines := readUpToNlines(os.Args[1],5)
	counts := createCounts(lines, separators, linesRead)
	separator := guerssSep(counts, separators, linesRead)
	report(separator)
}

func createCounts(lines, separators []string, linesRead int) [][]int{
	counts := make([][])int, len(separators))
	for sepIndex := range separators{
		counts[sepIndex] = make([]int, linesRead)
		for lineIndex, line := range lines(
			counts[sepIndex][linesIndex] = strings.Count(line, separators[sepIndex])
		)
	}
	return counts
}

/*
	一个字符串可以转化成一个[]byte(底层为utf-8编码)或者一个[]rune
	(底层为 unicode码点) 
	使用并发最简单的方式 就是用一个goroutine来准备工作，然后让另外一个
	goroutine来执行处理，让主goroutine和一些通道来安排一切事情

	对于通道的使用，第一，我们只有在后面要检查通道是否关闭
	第二：应该由发送方的 goroutine关闭通道，
*/

/*
	自定义类型 如果要将一个接受自定义类型底层类型的函数，就必须
	要先将其转换为底层类型（无需成本，在编译期完成的）

	方法：可以为任何类型添加一个或者多个方法。一个方法的接收者总是一个该类型
	的值，或者只是该类型的指针 

	类型方法集是指可以被该类型的值调用的所有方法组成 无论这些方法
	接受的是一个值还是指针
*/

type Part struct{
	Id int
	Name string
}

func (part *Part) LowerCase(){
	part.Name = string.ToLower(part.Name)
}

func (part *Part) UpperCase(){
	part.Name = strings.ToUpper(part.Name)
}

func (part Part) String() string{
	return fmt.Sprintf("<<%d %q>>", part.Id, part.Name)
}

func (part Part) HasPrefix(prefix string) bool{
	return strings.HasPrefix(part.Name, prefix)
}

/*
	方法表达式是一个必须将方法类型座位第一个参数的函数

	在go语言中，接口是一个自定义类型，声明了一个或者多个方法签名，接口完全抽象的
	因此不能将其实例化
	interface()类型是声明了空方法集的接口类型。无论包含或者不
	包含方法，任何一个值都满足interface{}类型， 毕竟一个值有方法，那么其
	方法包含空方法集以及它实际包含的方法
*/

jeky11 := StringPair("Henry", "Jeky11")
hydee := StringPair("Edward", "Hyde")
point := Point(5, -3)

fmt.Println("Before:", jeky11, hyde, point)
jeky11.Exchange()
point.Exchange()
fmt.Println("after #1:", jeky11, hyde, point)
exchangThese(&jeky11, &hyde, &point)
fmt.Println("after #2:", jeky11, hyde, point)

points := [][2]int{{4,6},{},{-7,11},{14,-8}}
for _, point := range points{
	fmt.Printf("%d, %d", point[0], point[1])
}

/*
	结构体的聚合与嵌入
*/
type Person struct{
	Title string
	Forenames []string
	Surname string
}

type Author1 struct{
	Names Person
	Title []string
	YearBorn int
}

/*
	为了嵌入一个匿名字段，使用了嵌入类型(或者接口)的名字而
	未声明的一个变量名， 可以直接访问这些字段，或者为了与
	外围结构体的字段名区分开， 使用类型或者接口的名字访问嵌入字段的字段

*/
type Author2 struct{
	Person // 匿名字段(嵌入)
	Title []string // 具体字段
	YearBorn int // 具体字段
}

/*
	嵌入接口
	结构体除了可以聚合和前乳具体的类型外，也可以聚合和
	嵌入接口(反之在接口中聚合或者嵌入结构体是行不通的
	因为接口是完全抽象的概念)
*/

type Optioner interface{
	Name() string
	IsValid() bool
}

type OptionCommon struct{
	ShortName string "short option name"
	LogNmae string "long option name"
}

type IntOption struct{
	OptionCommon
	Value, Min, Max int
}

// 实现接口的两种方法
func (option IntOption) Name() string {
	return name(option.ShortNmae, optiion.LongNmae)
}

func (option IntOption) IsValid() bool {
	return option.Min <= option.Value && option.Value <= option.Max
}

func name(shortName, longName, string) string {
	if longName == ""{
		return shortName
	}
	return longName
}
// FloatOption 匿名接口的实现 所以在创建Optioner字段时 必须
// 得给一个满足该接口的值 
type FloatOption struct{
	Optioner // 匿名字段(接口嵌入， 需要具体类型)
	Value float64 // 具体字段(集合)
}

// GenericOption 实现了 Name 和 IsValid 两个函数  
type GenericOption struct{
	OptionCommon
}

func (option GenericOption) Name() string{
	return name(option.ShortName, option.LogNmae)
}

func (option GenericOption) IsValid() bool {
	return true
}

/*
	对于FloatOption类型 嵌有一个Optioner类型字段， 因此FloatOption
	值需要一个 具体的类型来满足该字段的Optioner 接口， 可以通过给
	FloatOption值的Optioner字段赋一个GenericOption类型来实现
*/

fileOption := StringOption{OptionCommon{"f", "file"}, "index.html"}
topOption := IntOption{
	OptionCommon:OptionCommon{"t", "top"},
	Max: 100,
}
sizeOption :=Option{GenericOption{OptionCommon{"s", "size"}}, 19.5}

for _, option := range[]Optioner{topOption, fileOption, sizeOption}{
	fmt.Print("name=", option.Name(), ".valid=", optin.IsValid())
	fmt.Print(".valie=")
	switch option := option.(type){ 
		case IntOption:
			fmt.Print(option.Value, ".main=", option.Min, ".max=", option.Max, "\n")
		case StringOption:
			fmt.Println(option.Value)
		case FloatOption:
			fmt.Println(option.Value)	

	}
}


/*
	嵌入方法的匿名值
	如果一个嵌入字段带方法，则可以在外部结构中直接调用它
	并且只有嵌入的字段会作为接收者传递给这些方法
*/

type Tasks atruct{
	slice []string
	Count // count是匿名类型
}

func (tasks *Tasks) Add(task string) {
	task.slice = append(tasks.slice, task)
	task.Increment() // 访问的是 tasks.Count.Increment()
}


type FuzzyBool struct{
	value float32
}

func main(){
	a, _ := fu
}






