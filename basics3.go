// A Tour of Go : More types: structs, slices, and maps
// https://go-tour-jp.appspot.com/moretypes/1
package main

import (
	"fmt"
	"math"
)

func main() {

	{
		// ポインタ
		// C言語とは異なり、ポインタ演算は無い
		var i = 10
		var p *int = &i
		fmt.Println("1-1. ", i, *p, p)

		*p = 20
		fmt.Println("1-2. ", i, *p, p)
	}

	{
		// 構造体
		type Vertex struct {
			X int
			Y int
		}
		v := Vertex{1, 2}
		fmt.Println("2-1. ", v.X, v.Y)

		// 構造体のポインタの場合でも、要素にアクセスする際は、. を使う
		vp := &v
		vp.X = 10
		fmt.Println("2-2. ", vp, v)
	}

	{
		// Array (固定長配列)
		var ia [10]int
		var sa [2]string = [2]string{"Hello", "World"}
		fa := [3]float64{}
		fmt.Println("3. ", ia, sa, fa)
	}

	{
		// Slice
		// コロンで区切られた２つのインデックスの範囲
		i := [6]int{2, 3, 5, 7, 11, 13}
		var s []int = i[2:4] // インデックスが２から３までを取得。４は含まない
		fmt.Println("4-1. ", s)

		// スライスの要素を変更すると、元の配列の要素も変更される
		s[0] = 100
		fmt.Println("4-2. ", i, s)

		// 配列を作成して、それを参照するスライスを作成する
		r := []bool{true, false, true, false, false}
		fmt.Println("4-3. ", r)

		// 組み込みの make 関数を使用したスライス作成
		var is = make([]int, 5)
		fmt.Println("4-4. ", is)

		// スライスは、他のスライスを含む任意の型を含むことができる
		board := [][]string{
			[]string{"-", "-", "-"},
			[]string{"-", "-", "-"},
			[]string{"-", "-", "-"},
		}
		board[0][0] = "X"
		board[2][2] = "O"
		board[1][2] = "X"
		board[1][0] = "O"
		board[0][2] = "X"

		for i := 0; i < len(board); i++ {
			fmt.Println("4-5. ", board[i])
		}

		// スライスへ要素を追加
		var ss []int
		fmt.Println("4-6-1. ", ss)

		ss = append(ss, 0)
		fmt.Println("4-6-2. ", ss)

		ss = append(ss, 2, 3, 4)
		fmt.Println("4-6-3. ", ss)
	}

	{
		// Range
		// スライスや配列をひとつずつ反復処理するために使う
		pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
		for i, v := range pow {
			fmt.Println("5-1. ", i, v)
		}

		// インデックスだけ取得
		for i := range pow {
			fmt.Println("5-2. ", i)
		}

		// 値だけ取得
		for _, v := range pow {
			fmt.Println("5-3. ", v)
		}
	}

	{
		// map
		var m = make(map[string]int)
		m["hoge"] = 1
		fmt.Println("6-1. ", m)

		// 要素の削除
		delete(m, "hoge")
		fmt.Println("6-2. ", m)

		// キーに対する要素があるかどうか
		if _, ok := m["hoge"]; ok == true {
			fmt.Println("6-3. ", "key:hoge is found")
		} else {
			fmt.Println("6-3. ", "key:hoge is not found")
		}
	}

	{
		// function values
		// 関数も、関数の引数に取ったり、戻り値として利用できる
		var fn func(x, y float64) float64
		fn = hoge1
		fmt.Println("7-1. ", compute(fn))

		fn = hoge2
		fmt.Println("7-2. ", compute(fn))
	}

}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func hoge1(x, y float64) float64 {
	return x * y
}

func hoge2(x, y float64) float64 {
	return math.Pow(x, y)
}
