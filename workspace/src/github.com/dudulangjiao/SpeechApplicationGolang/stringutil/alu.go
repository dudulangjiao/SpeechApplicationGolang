package stringutil

// 半加器
func HalfAdder(inputFirst, inputSecond byte) (byte, byte) {
	sum := XOr(inputFirst, inputSecond)
	carry := And(inputFirst, inputSecond)
	return carry, sum
}

// 全加器
func FullAdder(inputFirst, inputSecond,
	inputThird byte) (byte, byte) {
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

// 4位加法器
func FourBitAdder(inputFirst, inputSecond string) (byte, string) {
	output := "0000"
	outputByte := []byte(output)
	var carry byte
	inputFirstByte := []byte(inputFirst)
	inputSecondByte := []byte(inputSecond)
	carry, outputByte[3] = HalfAdder(inputFirstByte[3],
		inputSecondByte[3])
	for i := 2; i >= 0; i-- {
		carry, outputByte[i] = FullAdder(carry, inputFirst[i],
			inputSecond[i])
	}
	output = string(outputByte)
	return carry, output
}

// 8位减法器 -128~127
func EightBitSub(inputSecond, inputFirst string) (byte, string) {
	output := "00000000"
	outputByte := []byte(output)
	inputSecondByte := []byte(inputSecond)
	//fmt.Println("inputSecond", inputSecond)
	//fmt.Println("inputSecondByte", inputSecondByte)
	for i := 0; i < 8; i++ {
		//fmt.Println("减法循环")
		//fmt.Println(i)
		//fmt.Println(inputSecondByte[i])
		outputByte[i] = Not(inputSecondByte[i])
		//fmt.Println("outputByte:", outputByte[i])
	}
	//fmt.Println("减法循环结束")
	output = string(outputByte)
	//fmt.Println(output)
	_, a := EightBitAdder(inputFirst, output)
	_, result := EightBitAdder(a, "00000001")
	m := []byte(result)
	return m[1], result
}

type ALU struct {
}

func (a *ALU) run(OpCode, Input1, Input2 string) (Negative byte,
	Output string) {
	switch OpCode {
	case "1000":
		_, c := EightBitAdder(Input1, Input2)
		return '0', c
	case "1001":
		//fmt.Println("减法")
		//fmt.Println("减数：", Input1)
		//fmt.Println("被减数：", Input2)
		m, c := EightBitSub(Input1, Input2)
		//fmt.Println("减后结果：", c)
		return m, c
	default:
		return '0', "error"
	}
}
