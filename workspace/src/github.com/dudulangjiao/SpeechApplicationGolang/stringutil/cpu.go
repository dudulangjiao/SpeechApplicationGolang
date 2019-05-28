package stringutil

import "fmt"

// CPU
type CPU struct {
	Control   ControlUnit
	RegisterA EightBitRegister
	RegisterB EightBitRegister
	RegisterC EightBitRegister
	RegisterD EightBitRegister
	Alu       ALU
	Ram       FourBitAddressableMemory
}

func NewCPU() *CPU {
	var result = new(CPU)
	result.RegisterA = *NewEightBitRegister()
	result.RegisterB = *NewEightBitRegister()
	result.RegisterC = *NewEightBitRegister()
	result.RegisterD = *NewEightBitRegister()
	result.Ram = *NewFourBitAddressableMemory()
	result.Control = *NewControlUnith()
	return result
}

func (c *CPU) Execute() {
	for i := 0; i < 10; i++ {

		// fetch phase 取指令阶段
		InstAddr := c.Control.InstAddrRegister.returnCircuit
		fmt.Println("*\n", InstAddr)
		_, instruction := c.Ram.ReadWrite('0',
			'1', "00000000", InstAddr)
		fmt.Println(instruction)
		a := c.Control.InstRegister.
			ReadWrite('1', instruction)

		// execute phase 执行阶段
		switch a[0:4] {
		case "0010":
			_, b := c.Ram.ReadWrite('0', '1',
				"00000000", a[4:8])
			c.RegisterA.ReadWrite('1', b)
			fmt.Println(b)
		case "0001":
			_, d := c.Ram.ReadWrite('0', '1',
				"00000000", a[4:8])
			c.RegisterB.ReadWrite('1', d)
		case "1000":
			bb := c.RegisterB.returnCircuit
			aa := c.RegisterA.returnCircuit
			cc := c.Alu.run(a[0:4], bb, aa)
			c.RegisterA.ReadWrite('1', cc)
		case "0100":
			hj := c.RegisterA.ReadWrite('0',
				"00000000")
			c.Ram.ReadWrite('1', '0', hj, a[4:8])
		default:
			fmt.Println("error")
		}

		// 计数器加1
		ff := c.Control.InstAddrRegister.returnCircuit
		_, uu := FourBitAdder(ff, "0001")
		c.Control.InstAddrRegister.ReadWrite('1', uu)
	}
	fmt.Println("结束")
}
