package main

import "fmt"

/**
 * @author:libing.niu
 * @date: 2020/2/14 12:10
 */

//声明接口
type Calculation interface {
	calcution(a1, a2 int) int //求和
}

//加法运算结构体
type SumCalculation struct {
	x int
	y int
}

//乘法运算
type MultiplicationCalculation struct {
	SumCalculation
}

//加法运算结构体实现加法运算
func (sum SumCalculation) calcution(a1, a2 int) int {
	return sum.x*a1 + sum.y*a2
}

//乘法运算
func (mu MultiplicationCalculation) calcution(a1, a2 int) int {
	return (mu.x + a1) * (mu.y + a2)
}
func main() {

	var fo Calculation
	fo = SumCalculation{x: 10, y: 10}
	fo = MultiplicationCalculation{SumCalculation{x: 1, y: 2}}
	calcution := fo.calcution(90, 10)
	fmt.Println(calcution)
}
