package stringutil

// ControlUnit
type ControlUnit struct {
	InstRegister     EightBitRegister
	InstAddrRegister FourBitRegister
	Negative         byte
}

func NewControlUnit() *ControlUnit {
	var result = new(ControlUnit)
	result.InstRegister = *NewEightBitRegister()
	result.InstAddrRegister = *NewFourBitRegister()
	result.Negative = '0'
	return result
}

func (c *ControlUnit) ReadInstAddr() string {
	InstAddr := c.InstAddrRegister.ReadWrite('0',
		"0000")
	return InstAddr
}

func (c *ControlUnit) WriteInst(Inst string) string {
	inst := c.InstRegister.ReadWrite('1', Inst)
	return inst
}
