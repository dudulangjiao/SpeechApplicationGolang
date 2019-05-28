package stringutil

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

func (e *FourBitRegister) ReadWrite(WriteEnable byte,
	DataInput string) string {
	DataInputByte := []byte(DataInput)
	result := "0000"
	resultByte := []byte(result)
	returnCircuitByte := []byte(e.returnCircuit)
	for i := 0; i < 4; i++ {
		resultByte[i] = GatedLatch(DataInputByte[i], WriteEnable,
			e.returnCircuit[i])
		returnCircuitByte[i] = resultByte[i]
	}
	e.returnCircuit = string(returnCircuitByte)
	result = string(resultByte)
	return result
}
