package stringutil

import (
	"fmt"
	"strconv"
)

// 矩阵型门锁
func MatrixGatedLatch(rowWire, columnWire, dataInOut, writeEnable,
	readEnable, returnCircuit byte) (byte, byte) {

	a := And(rowWire, columnWire)
	storedValue := GatedLatch(dataInOut, And(a, writeEnable),
		returnCircuit)
	cInput := And(a, readEnable)
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

	storedValue, output := MatrixGatedLatch(RowWire, ColumnWire,
		DataInOut, WriteEnable, ReadEnable, m.returnCircuit)
	m.returnCircuit = storedValue
	return storedValue, output
}

/*
// 256个门锁集合
type TwoHundredFiftySixGatedLatch struct {
	// 256个门锁集合
	MatrixGatedLatchSet [16][16]MatrixGatedLatchObject
}

func New256GatedLatch() *TwoHundredFiftySixGatedLatch {
	var result = new(TwoHundredFiftySixGatedLatch)
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			result.MatrixGatedLatchSet[i][j] = *NewMatrixGatedLatch()
		}
	}
	return result
}

func (t *TwoHundredFiftySixGatedLatch) ReadWrite(Data, WriteEnable,
	ReadEnable byte, EightBitAddress string) (byte, byte) {

	rowAddress, _ := strconv.ParseInt(EightBitAddress[0:4],
		2, 10)
	columnAddress, _ := strconv.ParseInt(EightBitAddress[4:8],
		2, 10)

	storedValue, output := t.
		MatrixGatedLatchSet[rowAddress][columnAddress].
		ReadWrite('1', '1', Data, WriteEnable,
			ReadEnable)
	return storedValue, output
}

// 8位可寻址内存 EightBitAddressableMemory
type EightBitAddressableMemory struct {
	// 8个256门锁集合
	twoHundredFiftySixBitMemorySet [8]TwoHundredFiftySixGatedLatch
}

func NewEightBitAddressableMemory() *EightBitAddressableMemory {
	var result = new(EightBitAddressableMemory)
	for i := 0; i < 8; i++ {
		result.twoHundredFiftySixBitMemorySet[i] = *New256GatedLatch()
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
		storeValueByte[i], resultByte[i] = a.
			twoHundredFiftySixBitMemorySet[i].
			ReadWrite(m[i], WriteEnable, ReadEnable, EightBitAddress)
	}
	result = string(resultByte)
	storeValue = string(storeValueByte)
	fmt.Println("\nWriteEnable:", string(WriteEnable),
		"ReadEnable:", string(ReadEnable),
		"EightBitData:", EightBitData,
		"EightBitAddress:", EightBitAddress)
	fmt.Println("storeValue:", storeValue, "DataInOut:", result)
	return storeValue, result
}
*/

// 16个门锁集合
type SixteenGatedLatch struct {
	// 16个门锁集合
	MatrixGatedLatchSet [4][4]MatrixGatedLatchObject
}

func NewSixteenGatedLatch() *SixteenGatedLatch {
	var result = new(SixteenGatedLatch)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result.MatrixGatedLatchSet[i][j] = *NewMatrixGatedLatch()
		}
	}
	return result
}

func (t *SixteenGatedLatch) ReadWrite(Data, WriteEnable,
	ReadEnable byte, FourBitAddress string) (byte, byte) {

	rowAddress, _ := strconv.ParseInt(FourBitAddress[0:2],
		2, 10)
	columnAddress, _ := strconv.ParseInt(FourBitAddress[2:4],
		2, 10)

	storedValue, output := t.
		MatrixGatedLatchSet[rowAddress][columnAddress].
		ReadWrite('1', '1', Data, WriteEnable,
			ReadEnable)
	return storedValue, output
}

// 4位可寻址内存 FourBitAddressableMemory
type FourBitAddressableMemory struct {
	// 8个16门锁集合
	fourBitAddressableMemorySet [8]SixteenGatedLatch
}

func NewFourBitAddressableMemory() *FourBitAddressableMemory {
	var result = new(FourBitAddressableMemory)
	for i := 0; i < 8; i++ {
		result.fourBitAddressableMemorySet[i] = *NewSixteenGatedLatch()
	}
	return result
}

func (a *FourBitAddressableMemory) ReadWrite(WriteEnable,
	ReadEnable byte, EightBitData,
	FourBitAddress string) (string, string) {

	m := []byte(EightBitData)
	result := "00000000"
	storeValue := "00000000"
	resultByte := []byte(result)
	storeValueByte := []byte(storeValue)
	for i := 0; i < 8; i++ {
		storeValueByte[i], resultByte[i] = a.
			fourBitAddressableMemorySet[i].
			ReadWrite(m[i], WriteEnable, ReadEnable, FourBitAddress)
	}
	result = string(resultByte)
	storeValue = string(storeValueByte)
	fmt.Println("\nWriteEnable:", string(WriteEnable),
		"ReadEnable:", string(ReadEnable),
		"EightBitData:", EightBitData,
		"FourBitAddress:", FourBitAddress)
	fmt.Println("storeValue:", storeValue, "DataInOut:", result)
	return storeValue, result
}
