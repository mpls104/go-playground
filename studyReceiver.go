package main

import "fmt"

type Score int
// レシーバ変数　
func (s Score) Show() { fmt.Printf("点数は%d点です\n", s) }

func main() {
	var myScore Score = 100
	myScore.Show()
}