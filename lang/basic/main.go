package main

import "fmt"


func manipulateByPointer(){
	f := func(v *uint32){
		*v += 5
	}

	var x uint32 = 5
	fmt.Printf("before x: %d\n", x)

	f(&x)

	fmt.Printf("after x: %d\n", x)
}

// Range で生成される変数は for内で使い回される
// このため、rangeで生成される変数をキャプチャしてはいけない
func valuesOfRangeSyntax(){
	s := []string{"a", "b", "c", "d"}
	for i, v := range s {
		fmt.Printf("loop %d\n", i)
		fmt.Printf("i: %d, v: %s\n", i, v)
		fmt.Printf("address i: %+v, v: %+v\n", &i, &v)
	}
}

func main(){
	//manipulateByPointer()
	valuesOfRangeSyntax()
}
