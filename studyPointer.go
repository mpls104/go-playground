package main

import (
	"fmt"
)

func main() {

	// int型のポインタ変数
	var pointer *int
	var n int = 100

	// nのアドレスを代入
	pointer = &n
	fmt.Println("nのアドレス：", &n)
	fmt.Println("pointerの値：", pointer)

	fmt.Println("nの値：", n)
	// *(間接参照演算子）を利用して、ポインタの中身を取得
	fmt.Println("pointerの中身：", *pointer) //100

	a, b := 10, 10

	// aはそのまま、bはアドレス演算子をつけて呼び出す
	called(a, &b)

	fmt.Println("値渡し:", a)
	fmt.Println("ポインタ渡し:", b)
} 

func called(a int, b *int) {
	a = a + 1
	*b = *b + 1
}