// A Tour of Go : Packages, variables, and functions
// https://go-tour-jp.appspot.com/basics/1
package main

import (
	"fmt"
	"math/cmplx"
	"math/rand"
)

// varステートメントで変数を宣言する
var c, python, java bool

func main() {

	{
		// パッケージ名はインポートパスの最後の要素と同じ名前になる
		fmt.Println("1. ", "My Favorite number is", rand.Intn(10))
	}

	{
		fmt.Println("2. ", "add(10, 11) = ", add(10, 11))
	}

	{
		// 関数から複数の戻り値を受け取る
		a, b := swap("World!!", "Hello")
		fmt.Println("3. ", a, b)
	}

	{
		var i int
		fmt.Println("4. ", i, c, python, java)
	}

	{
		// var宣言では、変数毎に初期化することができる
		// 初期化している場合、型の宣言を省略できる
		var c, python, java = true, false, "no!"
		fmt.Println("5. ", c, python, java)
	}

	{
		// 関数内ではvar宣言の代わりに、:=を使い、暗黙的な型宣言ができる
		var i, j int = 1, 2
		k := 3
		c, python, java := true, false, "no!"
		fmt.Println("6. ", i, j, k, c, python, java)
	}

	{
		// go言語の基本型(組み込み型)
		// bool
		// string
		// int, int8, int16, int32, int64
		// uint, uint8, uint16, uint32, uint64, uintptr
		// byte (uint8の別名)
		// rune (int32の別名。Unicodeのコードポイントを表す)
		// float32, float64
		// complex64, complex128 (複素数)
		// int, uint, uintptr型は32bitシステムでは32bit, 64bitシステムでは64bitになる
		// 符号なし整数の型は、使うための特別な理由がない限り控える
		var (
			ToBe   bool       = false
			MaxInt uint64     = 1<<64 - 1
			z      complex128 = cmplx.Sqrt(-5 + 12i)
		)
		fmt.Printf("7-1. Type: %T Value: %v\n", ToBe, ToBe)
		fmt.Printf("7-2. Type: %T Value: %v\n", MaxInt, MaxInt)
		fmt.Printf("7-3. Type: %T Value: %v\n", z, z)
	}

	{
		// 型変換
		var i int = 42
		var f float64 = float64(i)
		var u uint = uint(f)

		// 明示的な型を指定せずに変数を宣言する場合、変数の型は右側の変数から型推論される
		j := i

		fmt.Println("8. ", i, f, u, j)
	}

	{
		// 定数はconstキーワードを使う
		// 文字、文字列、boolean、数値のみ定数にできる
		const PI = 3.14
		fmt.Println("9. Happy", PI, "Day")
	}
}

// 頭文字が大文字で始まる場合、外部から参照できる(public)。小文字だとprivate
// goは変数名の後ろに型名を書く
func add(x int, y int) int {
	return x + y
}

// 関数は複数の戻り値を返すことができる
func swap(x, y string) (string, string) {
	return y, x
}

// 戻り値となる変数名に名前をつけることができる
// 名前をつけると、returnステートメントに何も書かずに戻すことができる
// 長い関数で使うと読みやすさに悪影響があるので注意
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
