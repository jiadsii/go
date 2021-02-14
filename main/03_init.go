package main

import "math"

var Pi float64

//开始执行程序之前对程序进行校验和修复

func init() {
	Pi = 4 * math.Atan(1)
}

