package stringutil

import (
	"fmt"
	"strconv"
)

// 门锁 Gated Latch
func GatedLatch(DataInput, WriteEnable, returnCircuit byte) byte {
	a := Or(And(DataInput, WriteEnable), returnCircuit)
	b := Not(And(Not(DataInput), WriteEnable))
	c := And(a, b)
	return c
}

// 8位寄存器 8-BIT Register
type EightBitRegister struct {
	returnCircuit string
}

func NewEightBitRegister() *EightBitRegister {
	var result = new(EightBitRegister)
	result.returnCircuit = "00000000"
	return result
}

func (e *EightBitRegister) ReadWrite(WriteEnable byte,
	DataInput string) string {
	DataInputByte := []byte(DataInput)
	result := "00000000"
	resultByte := []byte(result)
	returnCircuitByte := []byte(e.returnCircuit)
	for i := 0; i < 8; i++ {
		resultByte[i] = GatedLatch(DataInputByte[i], WriteEnable,
			e.returnCircuit[i])
		returnCircuitByte[i] = resultByte[i]
	}
	e.returnCircuit = string(returnCircuitByte)
	result = string(resultByte)
	return result
}

// 4位寄存器 4-BIT Register
type FourBitRegister struct {
	returnCircuit string
}

func NewFourBitRegister() *FourBitRegister {
	var result = new(FourBitRegister)
	result.returnCircuit = "0000"
	return result
}

// 方法：读写
func (f *FourBitRegister) ReadWrite(WriteEnable byte,
	DataInput string) string {
	DataInputByte := []byte(DataInput)
	result := "0000"
	resultByte := []byte(result)
	returnCircuitByte := []byte(f.returnCircuit)
	for i := 0; i < 4; i++ {
		resultByte[i] = GatedLatch(DataInputByte[i], WriteEnable,
			f.returnCircuit[i])
		returnCircuitByte[i] = resultByte[i]
	}
	f.returnCircuit = string(returnCircuitByte)
	result = string(resultByte)
	return result
}

// 方法：自增1
func (f *FourBitRegister) IncreaseOne() {
	// 计数器加1
	_, uu := FourBitAdder(f.returnCircuit, "0001")
	f.ReadWrite('1', uu)
}

// 4个8位寄存器集合
type Four8BitRegister struct {
	EightBitRegisterSet [4]EightBitRegister
}

func NewFour8BitRegister() *Four8BitRegister {
	var result = new(Four8BitRegister)
	for i := 0; i < 4; i++ {
		result.EightBitRegisterSet[i] = *NewEightBitRegister()
	}
	return result
}

func (f *Four8BitRegister) ReadWrite(WriteEnable byte,
	DataInput, TwoBitAddress string) string {
	address, _ := strconv.ParseInt(TwoBitAddress,
		2, 10)
	fmt.Println("地址：", address)
	result := f.EightBitRegisterSet[address].ReadWrite(WriteEnable,
		DataInput)
	fmt.Println("result:", result)
	return result
}
