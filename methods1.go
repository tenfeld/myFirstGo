package main

import (
	"fmt"
	"math"
)

// go にはクラスの仕組みはないが、型にメソッドを定義できる
// このメソッドは、func キーワードとメソッド名の間に、特別なレシーバー引数をとる
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ポインタレシーバーでもメソッドを宣言できる
// レシーバ自身を更新する場合、ポインタにしないと更新できない
// ポインタレシーバーの場合、メソッド呼び出し毎に、変数のコピーをしないので、基本ポインタを使う方がいいらしい
func (v *Vertex) Scale(f float64) {

	if v == nil {
		fmt.Println("<nil>")
		return
	}
	v.X = v.X * f
	v.Y = v.Y * f
}

// struct だけでなく、任意の型にメソッドを追加できる。
// ただし、レシーバーを伴うメソッドの宣言は、同じパッケージにある型に対してのみできる
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// interface 型は、メソッドのシグニチャの集まりで定義する
// interface を実装することを明示的に宣言する必要はない(implements キーワードは必要ない)。
type Abser interface {
	Abs() float64
	//hoge() int
}

// MyFloat型に、error型の組み込みインターフェイスを実装してみる
// type error interface {
// 	Error() string
// }
func (f *MyFloat) Error() string {
	return fmt.Sprintf("MyFloat Error")
}

func run() error {
	f := MyFloat(0)
	var err error
	err = &f
	return err
}

func main() {

	{
		v := Vertex{3, 4}
		fmt.Println("1. ", v.Abs())
	}

	{
		f := MyFloat(-math.Sqrt2)
		fmt.Println("2. ", f.Abs())
	}

	{
		v := Vertex{3, 4}
		v.Scale(10)
		fmt.Println("3-1 ", v)

		// ポインタからでも追加したメソッドを呼ぶことができる
		// (*p).Scale()と自動的に解釈してくれる
		p := &v
		p.Scale(10)
		fmt.Println("3-2. ", p)
	}

	{
		var a Abser
		f := MyFloat(-math.Sqrt2)
		v := Vertex{3, 4}

		a = f
		fmt.Println("4-1. ", a.Abs())

		a = v
		fmt.Println("4-2. ", a.Abs())

		a = &v
		fmt.Println("4-3. ", a.Abs())
	}

	{
		// 0個のメソッドを指定されたインターフェース型は、
		// 空のインターフェースと呼ばれ、任意の型を保持できる
		var i interface{}
		fmt.Println("5-1. ", describe(i))

		i = 10
		fmt.Println("5-2. ", describe(i))

		i = "Hello"
		fmt.Println("5-3. ", describe(i))

		i = func() {}
		fmt.Println("5-4. ", describe(i))
	}

	{
		// インターフェースの値が、特定の型を保持しているか
		var i interface{} = "Hello"
		s, ok := i.(string)
		fmt.Println("6-1. ", s, ok)

		f, ok := i.(float64)
		fmt.Println("6-2. ", f, ok)

		i = func(v int) float64 { return 0.0 }
		v, ok := i.(func(v int) float64)
		fmt.Println("6-3. ", v, ok)
	}

	{
		// switch で型アサーションを直列に使用できる
		var i interface{} = 123
		switch i.(type) {
		case int:
			fmt.Println("7. ", "type: int")
		case string:
			fmt.Println("7. ", "type: string")
		default:
			fmt.Println("7. ", "type: unknown")
		}
	}

	{
		// error インターフェイスの実装
		fmt.Println("8. ", run())
	}
}

func describe(i interface{}) string {
	return fmt.Sprintf("(%v, %T)", i, i)
}
