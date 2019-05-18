// 从零开始模拟计算机的逻辑单元
package main

import "fmt"

// NOT逻辑门
func Not(input string) string {
	switch input {
	case '1':
		return '0'
	case '0':
		return '1'
	default:
		return "error"
	}
}

// AND逻辑门
func And(inputFirst, inputSecond string) string {
	z := inputFirst + inputSecond
	switch z {
	case "11":
		return "1"
	case "01", "10", "00":
		return "0"
	default:
		return "error"
	}
}

// OR逻辑门
func Or(inputFirst, inputSecond string) string {
	z := inputFirst + inputSecond
	switch z {
	case "11", "01", "10":
		return "1"
	case "00":
		return "0"
	default:
		return "error"
	}
}

// XOR逻辑门 或 一位加法器
func XOr(inputFirst, inputSecond string) string {
	and := And(inputFirst, inputSecond)
	notAnd := Not(and)
	or := Or(inputFirst, inputSecond)
	output := And(notAnd, or)
	return output
}

// 半加器
func HalfAdder(inputFirst, inputSecond string) (string, string) {
	sum := XOr(inputFirst, inputSecond)
	carry := And(inputFirst, inputSecond)
	return carry, sum
}

// 全加器
func FullAdder(inputFirst, inputSecond, inputThird string) (string, string) {
	c1, s := HalfAdder(inputFirst, inputSecond)
	c2, sum := HalfAdder(s, inputThird)
	carry := Or(c1, c2)
	return carry, sum
}

// 8位加法器
func EightBitAdder(inputFirst, inputSecond [8]string) (string, [8]string) {
	var output [8]string
	var carry string
	carry, output[7] = HalfAdder(inputFirst[7], inputSecond[7])
	for i := 6; i >= 0; i-- {
		carry, output[i] = FullAdder(carry, inputFirst[i], inputSecond[i])
	}
	return carry, output
}

// 门锁 Gated Latch
type GatedLatch struct {
	DataInput, WriteEnable, returnCircuit string
}

func NewGatedLatch() *GatedLatch {
	var result = new(GatedLatch)
	result.returnCircuit = "0"
	return result
}

func (g *GatedLatch) ReadWrite() string {
	a := Or(And(g.DataInput, g.WriteEnable), g.returnCircuit)
	b := Not(And(Not(g.DataInput), g.WriteEnable))
	c := And(a, b)
	g.returnCircuit = c
	return c
}

// 8位寄存器 8-BIT Register
type EightBitRegister struct {
	WriteEnable   string
	DataInput     string
	returnCircuit [8]string
}

func NewEightBitRegister() *EightBitRegister {
	var result = new(EightBitRegister)
	for i := 0; i < 8; i++ {
		result.returnCircuit[i] = "0"
	}
	return result
}

func (e *EightBitRegister) ReadWrite() string {
	DataInputByte := []byte(e.DataInput)
	var result string
	for i := 0; i < 8; i++ {
		a := Or(And(DataInputByte[i], e.WriteEnable), e.returnCircuit[i])
		b := Not(And(Not(DataInputByte[i]), e.WriteEnable))
		result[i] = And(a, b)
		e.returnCircuit[i] = result[i]
	}
	return result
}

func main() {
	q := "1"
	b := "1"
	v := "0"
	fmt.Println(And(q, b))
	fmt.Println(Not(q))
	fmt.Println(XOr(q, b))
	fmt.Println(HalfAdder(q, b))
	fmt.Println(FullAdder(q, b, v))
	dd1 := [8]string{"0", "1", "0", "1", "1", "1", "1", "0"}
	dd2 := [8]string{"1", "1", "1", "0", "1", "0", "1", "1"}
	fmt.Println(EightBitAdder(dd1, dd2))

	fmt.Println("\n门锁")
	w := NewGatedLatch()
	// 1 "00" "0"
	w.DataInput = "0"
	w.WriteEnable = "0"
	fmt.Println("Input:", w.DataInput, "Write:", w.WriteEnable,
		"---Result:", w.ReadWrite())

	fmt.Println("\n8位寄存器")
	w := NewEightBitRegister()
	// 1 "00011001" "00011001"
	f := "00011001"
	for value := range f {
		w.DataInput
	}
	w.DataInput = "0"
	w.WriteEnable = "0"
	fmt.Println("Input:", w.DataInput, "Write:", w.WriteEnable,
		"---Result:", w.ReadWrite())
}
