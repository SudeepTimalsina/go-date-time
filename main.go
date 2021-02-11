package main

import (
	"fmt"
	"time"
)

func main() {
	input := "2017-08-31 10:14:59"
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, input)
	fmt.Println(t)
	fmt.Println(t.Format("02-Jan-2006"))
	fmt.Println(t.Format("02-January-2006"))

	loc, _ := time.LoadLocation("Japan")
	fmt.Println(time.Now())
	fmt.Println(time.Now().In(loc).Format("2006-01-02 15:04:05"))

	ktm, _ := time.LoadLocation("Asia/Kathmandu")
	fmt.Println(time.Now().In(ktm).Format("2006-01-02 15:04:05"))

	sng, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(time.Now().In(sng).Format("2006-01-02 15:04:05"))

	utc, _ := time.LoadLocation("UTC")
	fmt.Println(time.Now().In(utc).Format("2006-01-02 15:04:05"))
}

