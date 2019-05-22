// 从零开始模拟计算机的逻辑单元
package main

import (
	"fmt"
	"strconv"
)

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
func GatedLatch(DataInput, WriteEnable, returnCircuit byte) byte {
	a := Or(And(DataInput, WriteEnable), returnCircuit)
	b := Not(And(Not(DataInput), WriteEnable))
	c := And(a, b)
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
		resultByte[i] = GatedLatch(DataInputByte[i], e.WriteEnable,
			e.returnCircuit[i])
		e.returnCircuit[i] = resultByte[i]
	}
	result = string(resultByte)
	return result
}

// 矩阵型门锁
func MatrixGatedLatch(RowWire, ColumnWire, DataInOut, WriteEnable,
	ReadEnable, returnCircuit byte) (byte, byte) {
	a := And(RowWire, ColumnWire)
	storedValue := GatedLatch(DataInOut, And(a, WriteEnable), returnCircuit)
	cInput := And(a, ReadEnable)
	output := And(storedValue, cInput)
	return storedValue, output
}

type MatrixGatedLatchObject struct {
	returnCircuit byte
}

func NewMatrixGatedLatch() *MatrixGatedLatchObject {
	var result = new(MatrixGatedLatchObject)
	result.returnCircuit = '0'
	return result
}

func (m *MatrixGatedLatchObject) ReadWrite(RowWire, ColumnWire,
	DataInOut, WriteEnable, ReadEnable byte) (byte, byte) {
	storedValue, output := MatrixGatedLatch(RowWire, ColumnWire, DataInOut,
		WriteEnable, ReadEnable, m.returnCircuit)
	m.returnCircuit = storedValue
	/*
		fmt.Println("RowWire:", string(RowWire),
			"ColumnWire:", string(ColumnWire),
			"WriteEnable:", string(WriteEnable),
			"ReadEnable:", string(ReadEnable),
			"DataInOut:", string(DataInOut),
			"returnCircuit:", string(storedValue))
		fmt.Println("DataInOut:", string(output))
	*/
	return storedValue, output
}

// 256个门锁集合
type TwoHundredFiftySixBitMemory struct {
	// 256个门锁集合
	MatrixGatedLatchSet [16][16]MatrixGatedLatchObject
}

func New256BitMemory() *TwoHundredFiftySixBitMemory {
	var result = new(TwoHundredFiftySixBitMemory)
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			result.MatrixGatedLatchSet[i][j] = *NewMatrixGatedLatch()
		}
	}
	return result
}

func (t *TwoHundredFiftySixBitMemory) ReadWrite(Data, WriteEnable,
	ReadEnable byte, EightBitAddress string) (byte, byte) {
	rowAddress, _ := strconv.ParseInt(EightBitAddress[0:4],
		2, 10)

	columnAddress, _ := strconv.ParseInt(EightBitAddress[4:8],
		2, 10)

	storedValue, output := t.MatrixGatedLatchSet[rowAddress][columnAddress].
		ReadWrite('1', '1', Data, WriteEnable,
			ReadEnable)
	fmt.Println("EightBitAddress:", EightBitAddress,
		"WriteEnable:", string(WriteEnable),
		"ReadEnable:", string(ReadEnable),
		"Data:", string(Data), "storedValue:", string(storedValue))

	return storedValue, output
}

// 8位可寻址内存 EightBitAddressableMemory
type EightBitAddressableMemory struct {
	// 8个256位内存
	TwoHundredFiftySixBitMemorySet [8]TwoHundredFiftySixBitMemory
}

func NewEightBitAddressableMemory() *EightBitAddressableMemory {
	var result = new(EightBitAddressableMemory)
	for i := 0; i < 8; i++ {
		for j := 0; j < 16; j++ {
			for k := 0; k < 16; k++ {
				result.TwoHundredFiftySixBitMemorySet[i].
					MatrixGatedLatchSet[j][k].returnCircuit = '0'
			}

		}
	}
	return result
}

func (a *EightBitAddressableMemory) ReadWrite(WriteEnable,
	ReadEnable byte, EightBitData,
	EightBitAddress string) (string, string) {
	m := []byte(EightBitData)
	result := "00000000"
	storeValue := "00000000"
	resultByte := []byte(result)
	storeValueByte := []byte(storeValue)
	for i := 0; i < 8; i++ {
		storeValueByte[i], resultByte[i] = a.TwoHundredFiftySixBitMemorySet[i].
			ReadWrite(m[i], WriteEnable, ReadEnable, EightBitAddress)
	}
	result = string(resultByte)
	storeValue = string(storeValueByte)
	return storeValue, result
}

func main() {
	/*
		Address := "00011110"
		rowAddress, _ := strconv.ParseInt(Address[0:4],  2,
			10)
		columnAddress, _ := strconv.ParseInt(Address[4:8], 2,
			10)
		fmt.Println(rowAddress)
		fmt.Println(columnAddress)

		dd1 := "01011110"
		dd2 := "11110001"
		fmt.Println("8位加法器")
		fmt.Println(EightBitAdder(dd1, dd2))

		fmt.Println("\n8位寄存器")
		ww := NewEightBitRegister()

		// 1 Input:"00011001" WriteEnable:'0' Result:"00000000"
		ww.DataInput = "00011001"
		ww.WriteEnable = '0'
		fmt.Println("Input:", ww.DataInput, "Write:",
			string(ww.WriteEnable), "---Result:", ww.ReadWrite())

		fmt.Println("\n矩阵型门锁")
		juzhen := NewMatrixGatedLatch()

		// 1 RC:11 WriteRead:00 DAtaInOut:0
		juzhen.ReadWrite('1', '1',
			'0', '0', '0')

		hhhh := New256BitMemory()
		_, hul := hhhh.ReadWrite('1', '1', '0',
			"00111101")
		fmt.Println("\n256位内存")
		fmt.Println(hul)
	*/
	ffff := NewEightBitAddressableMemory()
	sss, yyy := ffff.ReadWrite('1', '0',
		"11110001", "11000110")
	fmt.Println("\n8位可寻址内存")

	fmt.Println("storeValue:", sss, "output:", yyy)

	sss, yyy = ffff.ReadWrite('0', '1',
		"00000001", "11000110")
	fmt.Println("\n8位可寻址内存")

	fmt.Println("storeValue:", sss, "output:", yyy)

	sss, yyy = ffff.ReadWrite('1', '1',
		"00000001", "11000110")
	fmt.Println("\n8位可寻址内存")

	fmt.Println("storeValue:", sss, "output:", yyy)

	sss, yyy = ffff.ReadWrite('0', '0',
		"00000001", "11000110")
	fmt.Println("\n8位可寻址内存")

	fmt.Println("storeValue:", sss, "output:", yyy)
}
