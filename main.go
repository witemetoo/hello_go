package main

// import "fmt"

// func main(){
// 	fmt.Println("hello,世界")
// }

import (
	// "fmt"
	// "os"
	// "bufio"
	// "strings"
	// "io/ioutil"
	// "image"
	// "image/color"
	// "image/gif"
	// "math"
	// "math/rand"
	// "io"
	// "net/http" 

	// 并发获取多个url
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)
// := 短变量声明
func main(){
	// var s,sep string
	// for i := 1; i<len(os.Args);i++{
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// fmt.Println(s)
	// other_echo()
	// dup1()
	// get_url()
	get_mut_url()
}


// other_echo 函数循环的每次迭代，range产生一对值：索引以及该处的索引值
// 由于例中不需要索引 所以把它赋予临时变量temp 但go 中不使用无用的临时变量 所以用空标识符_ 
func other_echo(){
	s,sep := "", ""
	for _,arg := range os.Args[1:]{
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// func dup1(){
// 	counts := make(map[string]int)
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan(){
// 		counts[input.Text()] ++
// 	}
// 	for line,n :=range counts{
// 		if n>1 {
// 			fmt.Printf("%d\t %s\n",n,line)
// 		}
// 	}
// }

// // os.open 函数返回两个值 ，一个是被打开的文件 第二个是内置的error类型的值
// // 函数和包级别变量可以任意顺序声明，并不影响其使用
// // map是由make函数创建的数据结构引用，map作为参数传递给函数时 函数接受的是这个引用的拷贝
// func dup2(){
// 	counts := make(map[string]int)
// 	files := os.Args[1:]
// 	if len(files) == 0{
// 		countLines(os.Stdin,counts)
// 	}else{
// 		for _,args := range files{
// 			f,err := os.Open(args)
// 			if err != nil{
// 				fmt.Fprintf(os.Stderr, "dup2 :%v\n", err)
// 				continue
// 			}
// 			countLines(f,counts)
// 			f.Close()
// 		}
// 	}
// 	for line,n := range counts{
// 		if n>1{
// 			fmt.Printf("%d \t %s\n", n,line)
// 		}
// 	}
// }

// func countLines(f*os.File,counts map[string]int){
// 	input := bufio.NewScanner(f)
// 	for input.Scan(){
// 		counts[input.Text()]++
// 	}
// }

// func dup3(){
// 	counts := make(map[string]int)
// 	for _,filename := range os.Args[1:]{
// 		data,err := ioutil.ReadFile(filename)
// 		if err != nil{
// 			fmt.Fprintf(os.Stderr,"dup3:%v\n",err)
// 			continue
// 		}
// 		for _,line := range strings.Split(string(data), "\n"){
// 			counts[line] ++
// 		}
// 	}
// 	for line,n := range counts {
// 		if n>1{
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

// // 复合声明 color.Color()
// var palette = []color.Color{color.White,color.Black}

// const (
// 	whiteIndes = 0
// 	blackIndex = 1
// )

// // func main(){
// // 	lissajous(os.Stdout)
// // }

// func lissajous(out io.Writer){
// 	// 程序中的常量声明给出了一系列的常量值 常量是指在程序编译后始终都不变的值
// 	// 常量声明和变量声明一般都会出现在包 级别 所以常量在整个包中都是可以共享的
// 	const(
// 		cycles = 5
// 		res = 0.001
// 		size = 100
// 		nframes = 64
// 		delay = 8
// 	)

// 	freq := rand.Float64() * 3.0
// 	// git.GIF{}为复合声明,会把loopcount设置为nframes anim是是符合声明的结构体
// 	anim := gif.GIF{LoopCount:nframes}
// 	phase := 0.0
// 	for i := 0; i < nframes; i++{
// 		rect := imag.Rect(0,0,2*size+1,2*size+1)
// 		img := image.NewPaletted(rect,palette)
// 		for t:= 0.0; t < cycles*2*math.Pi; t+=res{
// 			x := math.Sin(t)
// 			y := math.Sin(t * freq + phase)
// 			img.SetColorIndex(size + int(x*size + 0.5), size + int(y * size + 0.5),
// 			blackIndex
// 			)
// 		}
// 		phase += 0.1
// 		anim.Delay = append(anim.Delay,delay)
// 		anim.Image = append(anim.Image,img)
// 	}
// 	gif.EncodeAll(out,&anim)
// }


// 获取url
func get_url(){
	for _,url := range os.Args[1:]{
		resp, err := http.Get(url)
		if err != nil{
			fmt.Fprintf(os.Stderr, "fetch:%v\n",err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch: reading %s: %v\n", url,err)
			os.Exit(1)
		}
		fmt.Printf("%s",b)
	}

}


// 并发获取多个url 并发获取多个url的时间不会超过 最长的单个时间

func get_mut_url(){
	start := time.Now()
	ch := make(chan string)
	for _,url := range os.Args[1:]{
		go fetch(url,ch)
	}
	for range  os.Args[1:]{
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds())
}
func fetch(url string,ch chan<- string){
	start := time.Now()
	resp,err := http.Get(url)
	if err != nil{
		ch <- fmt.Sprint(err)
		return
	}
	nbytes,err := io.Copy(ioutil.Discard,resp.Body)
	resp.Body.Close()
	if err != nil{
		ch <- fmt.Sprintf("while reading %s:%v" ,url,err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs,nbytes,url)
}


/*
	go 语言中函数名、变量名、常量名、类型名、语句标号和包名 

	如果一个名字是大写字母开头（必须是函数外部定义的包级名字；包级函数名本身
	也是包级名字）

	声明：
	声明语句定义了程序的各种实体对象一级部分或全部属性 
	go 语言有四种类型声明语句：var const type func分别对应：变量、常量、类型、函数
	每个源文件以包名的声明语句开始，说明该源文件属于哪个包

	包一级声明语句声明的名字可在整个包对应的每个源文件中访问
	而不仅仅是在其声明语句所在的源文件中访问


	var 声明语句可以创造一个特定类型的变量，然后变量附加一个名字，并且设置变量的初始化值
	var 变量名字 类型＝表达式
	如果初始化表达式被省略，数值类型变量对应的零值是0 布尔类型对应的零值是false
	字符串类型对应的零值是空字符串，接口或者引用类型(slice.map.chan)
	变量对应的零值是nil 
	数组或者结构体等聚合类型对应的零值是每个元素对应该类型的零值

	var i,j,k int
	var b,f,s = true,2.3,"four"

	一组变量也可以通过调用一个函数，由函数返回的多个返回值初始化
	var f,err = os.Open(name)

	简短变量声明
	简短变量声明语句的形式可用于声明和初始化局部变量 变量的类型根据表达式来自动推导

	简短变量声明语句也可以用函数返回值来声明和初始化变量
	f,err := os.Open(name)
	在某种情况下 简短变量声明左边变量可能并不是全部都是刚刚声明的 那么简短声明语句只有赋值行为来
	in,err := os.Open(infile)
	...
	out,err := os.Create(outfile)
	简短声明语句至少包含一个新的声明变量 下面的代码不能编译通过
	f,err := os.Open(infile)
	...
	f,err := os.Create(outfile)


	指针
	一个变量对应一个保存了 变量对应类型值 的内存空间
	变量有时候被称为可寻址的值 任何类型的指针的零值都是 nil 如果 p! =nil 
	测试为真 那么p是指向 某个有效变量 

	new函数
	new()函数将创建一个T类型的匿名变量，初始化为T类型的零值
	然后返回变量地址，返回指针类型为*T


	变量的生命周期
	变量的生命周期指的是程序运行期间变量有效存在的时间间隔
	包级变量的生命周期和整个程序的运行周期是一致的
	而局部变量的声明周期则是动态的：从每次创建一个新变量的声明语句开始
	直到该变量不在被引用为止
	因为一个变量的有效周期只取决于是否可达，因此一个循环迭代内部的局部
	变量的生命周期可能超出其局部作用域。同时，局部变量可能在函数返回之后依然存在
	变量生命周期是从时间上描述变量 作用域是从空间上描述变量


	类型
	变量或表达式的类型定义了对应存储值的属性特征
	一个类型声明语句创建了一个新的类型名称，和现有类型具有相同的底层结构
	新命名的类型提供了一个方法，用来分隔不同概念的类型
	type 类型名 底层类型

	包和文件
	go语言的包和其他语言的库或者模块的概念类似
	都是为了支持模块化、封装、单独编译和代码重用
	每个源文件都是以包的声明语句开始的，用来指名包的名字。当包
	被导入时

	

*/
