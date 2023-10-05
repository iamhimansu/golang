package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	println(presentTime.Format("01-02-2006 Monday"))
	// in sequential manner it is
	// Jan 2 15:04:05 2006 MST
	// An easy way to remember this value is that it holds, when presented
	// in this order, the values (lined up with the elements above):
	// 1 2  3  4  5  6  -7
	// There are some wrinkles illustrated below.

	//creating timestamp from date
	date := time.Date(2020, time.February, 7, 23, 0, 0, 0, time.UTC)

	fmt.Println(date.UnixMilli())
}
