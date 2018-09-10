package main

import "fmt"

func main(){
	var test1 = []int{1,2,3,4,5}
	test2 := test1
	fmt.Println(test1)
	fmt.Println(test2)

	test1[0] = 0
	fmt.Println(test1)
	fmt.Println(test2)
}

/*
[1 2 3 4 5]
[1 2 3 4 5]
[0 2 3 4 5]
[0 2 3 4 5]
*/

/*
スライスは一度初期化されると、その要素を実際に所有する参照先配列との関連を常に保ちます。
そのためスライスは参照先配列および、同一の配列から作られた別のスライスとメモリを共有します。
これとは対照的に異なる配列は常に異なるメモリ領域を有します。


スライスは参照型です。すなわちスライスを別のスライスに代入したときは、
双方が同じ配列を参照します。たとえば関数がスライスを引数として取るとき、
関数内でスライスの要素に変更を加えると、ポインタ渡しのように関数の呼び元からも変更内容が参照できます
*/