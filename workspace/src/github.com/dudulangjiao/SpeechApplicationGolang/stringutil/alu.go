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
