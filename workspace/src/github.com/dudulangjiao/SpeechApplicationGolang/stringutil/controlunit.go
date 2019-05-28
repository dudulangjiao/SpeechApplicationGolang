package stringutil

// ControlUnit
type ControlUnit struct {
	InstRegister     EightBitRegister
	InstAddrRegister FourBitRegister
}

func NewControlUnith() *ControlUnit {
	var result = new(ControlUnit)
	result.InstRegister = *NewEightBitRegister()
	result.InstAddrRegister = *NewFourBitRegister()
	return result
}
