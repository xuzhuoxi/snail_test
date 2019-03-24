//
//Created by xuzhuoxi
//on 2019-03-24.
//@author xuzhuoxi
//
package internel

import "fmt"

type testA struct {
	A string
	B int
	C bool
}

type I interface {
	Func1()
	Func2()
	Func3()
}

type B struct{}

func (b *B) Func1() {
	b.Func3()
}

func (*B) Func3() {
	fmt.Println("B3")
}

type S struct {
	B
}

//func (s *S) Func1() {
//	fmt.Println("S1")
//}

func (s *S) Func2() {
	s.Func1()
}

func (s *S) Func3() {
	fmt.Println("S3")
}
