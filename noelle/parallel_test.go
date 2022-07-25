package noelle

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	easy()
	pipe()
	expect()
}

func easy() {
	p := Active()
	p.Register(func() {
		fmt.Println(111)
	})
	p.Register(func() {
		fmt.Println(222)
	})
	p.Run()
}

/*
	jobA      jobC
	  |        /
	  |       /
	jobB     /
	  |     /
	  |    /
	  result
*/
func pipe() {
	var res string
	var x, y int
	p := Active()
	pipe := p.Pipeline()
	pipe.Register(testJobA).SetReceivers(&res)
	pipe.Register(testJobB, 1, 2).SetReceivers(&x)
	p.Register(testJobC, 3).SetReceivers(&y)
	// block here
	p.Run()
	fmt.Println(res, x, y) //job 3 -3
}

func testJobA() string {
	return "job"
}

func testJobB(x, y int) int {
	return x + y
}

func testJobC(x int) int {
	return -x
}

func expect() {
	p := Active()
	p.Register(exceptionJob)
	p.Except(exceptionHandler, "topic1")
	p.Run()
}

func exceptionHandler(topic string, e interface{}) {
	fmt.Println(topic, e)
}

func exceptionJob() {
	var a map[string]int
	//assignment to entry in nil map
	a["123"] = 1
}
