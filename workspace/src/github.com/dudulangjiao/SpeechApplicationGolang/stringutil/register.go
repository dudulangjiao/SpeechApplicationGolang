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
	returnCircuit [8]byte
}

func NewEightBitRegister() *EightBitRegister {
	var result = new(EightBitRegister)
	for i := 0; i < 8; i++ {
		result.returnCircuit[i] = '0'
	}
	return result
}

func (e *EightBitRegister) ReadWrite(WriteEnable byte,
	DataInput string) string {
	DataInputByte := []byte(DataInput)
	result := "00000000"
	resultByte := []byte(result)
	for i := 0; i < 8; i++ {
		resultByte[i] = GatedLatch(DataInputByte[i], WriteEnable,
			e.returnCircuit[i])
		e.returnCircuit[i] = resultByte[i]
	}
	result = string(resultByte)
	return result
}
