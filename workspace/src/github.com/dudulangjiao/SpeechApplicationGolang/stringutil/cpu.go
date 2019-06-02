package stringutil

import (
	"fmt"
)

// CPU
type CPU struct {
	Control   ControlUnit
	Registers Four8BitRegister
	Alu       ALU
	Ram       FourBitAddressableMemory
}

func NewCPU() *CPU {
	var result = new(CPU)
	result.Registers = *NewFour8BitRegister()
	result.Ram = *NewFourBitAddressableMemory()
	result.Control = *NewControlUnit()
	return result
}

func (c *CPU) Execute() {
	fmt.Println("循环开始")
	for i := 0; i < 10; i++ {
		// fetch phase 取指令阶段

		InstAddr := c.Control.ReadInstAddr()
		fmt.Println("*\n", InstAddr)
		_, instruction := c.Ram.ReadWrite('0',
			'1', "00000000", InstAddr)
		fmt.Println("命令：", instruction)
		a := c.Control.WriteInst(instruction)

		// execute phase 执行阶段
		switch a[0:4] {
		case "0010": // LOAD_A
			_, b := c.Ram.ReadWrite('0', '1',
				"00000000", a[4:8])
			c.Registers.ReadWrite('1', b, "00")
			c.Control.InstAddrRegister.IncreaseOne()
			fmt.Println(i)
		case "0001": // LOAD_B
			_, d := c.Ram.ReadWrite('0', '1',
				"00000000", a[4:8])
			c.Registers.ReadWrite('1', d, "01")
			c.Control.InstAddrRegister.IncreaseOne()
		case "1000": // ADD
			fmt.Println("ADD")
			aa := c.Registers.ReadWrite('0',
				"00000000", a[4:6])
			bb := c.Registers.ReadWrite('0',
				"00000000", a[6:8])
			_, m := c.Alu.run(a[0:4], aa, bb)
			c.Registers.ReadWrite('1', m, a[6:8])
			c.Control.InstAddrRegister.IncreaseOne()
		case "1001": // SUB
			aa := c.Registers.ReadWrite('0',
				"00000000", a[4:6])
			fmt.Println("aa", aa)
			bb := c.Registers.ReadWrite('0',
				"00000000", a[6:8])
			fmt.Println("bb", bb)
			m, d := c.Alu.run(a[0:4], aa, bb)
			fmt.Println("sub", d)
			c.Registers.ReadWrite('1', d, a[6:8])
			fmt.Println("Negative:", string(m))
			c.Control.Negative = m
			c.Control.InstAddrRegister.IncreaseOne()
		case "0100": // STORE_A
			hj := c.Registers.ReadWrite('0',
				"00000000", "00")
			c.Ram.ReadWrite('1', '0', hj, a[4:8])
			c.Control.InstAddrRegister.IncreaseOne()
		case "0011": // JUMP
			fmt.Println("JUMP")
			fmt.Println("Negative", c.Control.Negative)
			fmt.Println("a[4:8]", a[4:8])
			c.Control.InstAddrRegister.ReadWrite('1',
				a[4:8])
		case "0110": // JUMP_NEG
			fmt.Println("JUMP_NEG")
			fmt.Println("Negative", c.Control.Negative)
			fmt.Println("a[6:8]", a[4:8])
			if c.Control.Negative == '1' {
				c.Control.InstAddrRegister.ReadWrite('1',
					a[4:8])

			} else {
				c.Control.InstAddrRegister.IncreaseOne()
				//fmt.Println("自增：", c.Control.InstAddrRegister.returnCircuit)
			}
			fmt.Println("qu：", c.Control.InstAddrRegister.returnCircuit)
		case "0111": // HALT
			break
		default:
			fmt.Println("error")
		}
	}
	fmt.Println("结束")
}
