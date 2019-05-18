// 从零开始模拟计算机的逻辑单元
package main

import "fmt"

// NOT逻辑门
func Not(input byte) byte {
	switch input {
	case '1':
		return '0'
	case '0':
		return '1'
	default:
		return 'e'
	}
}

// AND逻辑门
func And(inputFirst, inputSecond byte) byte {
	z := string(inputFirst) + string(inputSecond)
	switch z {
	case "11":
		return '1'
	case "01", "10", "00":
		return '0'
	default:
		return 'e'
	}
}

// OR逻辑门
func Or(inputFirst, inputSecond byte) byte {
	z := string(inputFirst) + string(inputSecond)
	switch z {
	case "11", "01", "10":
		return '1'
	case "00":
		return '0'
	default:
		return 'e'
	}
}

// XOR逻辑门 或 一位加法器
func XOr(inputFirst, inputSecond byte) byte {
	and := And(inputFirst, inputSecond)
	notAnd := Not(and)
	or := Or(inputFirst, inputSecond)
	output := And(notAnd, or)
	return output
}

// 半加器
func HalfAdder(inputFirst, inputSecond byte) (byte, byte) {
	sum := XOr(inputFirst, inputSecond)
	carry := And(inputFirst, inputSecond)
	return carry, sum
}

// 全加器
func FullAdder(inputFirst, inputSecond, inputThird byte) (byte, byte) {
	c1, s := HalfAdder(inputFirst, inputSecond)
	c2, sum := HalfAdder(s, inputThird)
	carry := Or(c1, c2)
	return carry, sum
}

// 8位加法器
func EightBitAdder(inputFirst, inputSecond string) (byte, string) {
	output := "00000000"
	outputByte := []byte(output)
	var carry byte
	inputFirstByte := []byte(inputFirst)
	inputSecondByte := []byte(inputSecond)
	carry, outputByte[7] = HalfAdder(inputFirstByte[7],
		inputSecondByte[7])
	for i := 6; i >= 0; i-- {
		carry, outputByte[i] = FullAdder(carry, inputFirst[i],
			inputSecond[i])
	}
	output = string(outputByte)
	return carry, output
}

// 门锁 Gated Latch
type GatedLatch struct {
	DataInput, WriteEnable, returnCircuit byte
}

func NewGatedLatch() *GatedLatch {
	var result = new(GatedLatch)
	result.returnCircuit = '0'
	return result
}

func (g *GatedLatch) ReadWrite() byte {
	a := Or(And(g.DataInput, g.WriteEnable), g.returnCircuit)
	b := Not(And(Not(g.DataInput), g.WriteEnable))
	c := And(a, b)
	g.returnCircuit = c
	return c
}

// 8位寄存器 8-BIT Register
type EightBitRegister struct {
	WriteEnable   byte
	DataInput     string
	returnCircuit [8]byte
}

func NewEightBitRegister() *EightBitRegister {
	var result = new(EightBitRegister)
	for i := 0; i < 8; i++ {
		result.returnCircuit[i] = '0'
	}
	return result
}

func (e *EightBitRegister) ReadWrite() string {
	DataInputByte := []byte(e.DataInput)
	result := "00000000"
	resultByte := []byte(result)
	for i := 0; i < 8; i++ {
		a := Or(And(DataInputByte[i], e.WriteEnable), e.returnCircuit[i])
		b := Not(And(Not(DataInputByte[i]), e.WriteEnable))
		resultByte[i] = And(a, b)
		e.returnCircuit[i] = resultByte[i]
	}
	result = string(resultByte)
	return result
}

func main() {

	dd1 := "01011110"
	dd2 := "11110001"
	fmt.Println(EightBitAdder(dd1, dd2))

	fmt.Println("\n门锁")
	w := NewGatedLatch()
	// 1 "00" "0"
	w.DataInput = '0'
	w.WriteEnable = '0'
	fmt.Println("Input:", w.DataInput, "Write:", w.WriteEnable,
		"---Result:", w.ReadWrite())

	fmt.Println("\n8位寄存器")
	ww := NewEightBitRegister()

	// 1 Input:"00011001" WriteEnable:'0' Result:"00000000"
	ww.DataInput = "00011001"
	ww.WriteEnable = '0'
	fmt.Println("Input:", ww.DataInput, "Write:",
		string(ww.WriteEnable), "---Result:", ww.ReadWrite())

	// 2 Input:"00011001" WriteEnable:'1' Result:"00011001"
	ww.DataInput = "00011001"
	ww.WriteEnable = '1'
	fmt.Println("Input:", ww.DataInput, "Write:",
		string(ww.WriteEnable), "---Result:", ww.ReadWrite())

	// 3 Input:"00011001" WriteEnable:'0' Result:"00011001"
	ww.DataInput = "00011001"
	ww.WriteEnable = '0'
	fmt.Println("Input:", ww.DataInput, "Write:",
		string(ww.WriteEnable), "---Result:", ww.ReadWrite())

	// 4 Input:"11111001" WriteEnable:'0' Result:"00011001"
	ww.DataInput = "11111001"
	ww.WriteEnable = '1'
	fmt.Println("Input:", ww.DataInput, "Write:",
		string(ww.WriteEnable), "---Result:", ww.ReadWrite())
}
