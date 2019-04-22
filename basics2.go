// A Tour of Go : Flow control statements: for, if, else, switch and defer
// https://to-tour-jp.appspot.com.list/flowcontrol/1
package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {

	{
		// for ループ
		sum := 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		fmt.Println("1-1. ", sum)

		// 初期化と後処理ステートメントの記述は省略可
		// セミコロンも省略可
		for sum < 1000 {
			sum += sum
		}
		fmt.Println("1-2. ", sum)
	}

	{
		// if ステートメント
		fmt.Println("2-1. ", sqrt(2), sqrt(-4))
		fmt.Println("2-2. ", pow(3, 2, 10), pow(3, 3, 20))
	}

	{
		// switch ステートメント
		// Goでは選択されたcaseだけを実行して、それに続くcaseは実行されない
		// caseの最後に必要なbreakステートメントが、goでは自動的に提供される
		switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("3-1. ", "OS X")
		case "linux":
			fmt.Println("3-1. ", "Linux")
		default:
			fmt.Printf("3-1. %s\n", os)
		}

		// switch case は上から下へとcase を評価する
		// case 条件が一致すれば、そこで停止する
		today := time.Now().Weekday()
		switch time.Saturday {
		case today + 0:
			fmt.Println("3-2. ", "Today")
		case today + 2:
			fmt.Println("3-2. ", "Tomorrow")
		case today + 2:
			fmt.Println("3-2. ", "In two days")
		default:
			fmt.Println("3-2. ", "Too far away")
		}
	}

	{
		// defer
		// defer ステートメントは、defer へ渡した関数の実行を、呼び出し元の関数の終わりまで遅延させる。
		// 渡した関数が複数ある場合、その呼び出しはスタックされ、last-in-first-out(最後に追加したものから)の順番で実行される
		defer fmt.Println("4-1. ", "Hello")
		defer fmt.Println("4-2. ", "World")
		fmt.Println("4-3. ", "!!")
	}

}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// if ステートメントは、for ループのように、評価するための簡単なステートメントを書くことができる
	// ここで宣言された変数は、if のスコープ内でのみ有効
	if v := math.Pow(x, n); v < lim {
		return v
	} else {

		// }
		// else {
		// だとエラー扱い

		//fmt.Printf("\t %g >= %g\n", v, lim)
	}
	return lim
}
