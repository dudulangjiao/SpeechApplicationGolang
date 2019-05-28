// 从零开始模拟计算机
package main

import (
	"fmt"
	"github.com/dudulangjiao/SpeechApplicationGolang/stringutil"
)

func main() {
	/*
		eightBitAddressableMem := stringutil.NewEightBitAddressableMemory()

		fmt.Println("\n8位可寻址内存")
		eightBitAddressableMem.ReadWrite('1', '0',
			"11110001", "11000110")

		eightBitAddressableMem.ReadWrite('0', '1',
			"00000001", "11000110")
		eightBitAddressableMem.ReadWrite('1', '1',
			"00000001", "11000110")

		eightBitAddressableMem.ReadWrite('1', '0',
			"00000011", "11000110")

		eightBitAddressableMem.ReadWrite('0', '0',
			"00000111", "11000110")

		fourBitAddressableMem := stringutil.NewFourBitAddressableMemory()

		fmt.Println("\n4位可寻址内存")
		fourBitAddressableMem.ReadWrite('1', '0',
			"11110001", "1100")

		fourBitAddressableMem.ReadWrite('0', '1',
			"00000001", "1100")
		fourBitAddressableMem.ReadWrite('1', '1',
			"00000001", "1100")

		fourBitAddressableMem.ReadWrite('1', '0',
			"00000011", "1100")

		fourBitAddressableMem.ReadWrite('0', '0',
			"00000111", "1100")

		fmt.Println("\nCPU")
		ram := stringutil.NewFourBitAddressableMemory()
		ram.ReadWrite('1', '0',
			"00101110", "0000") // ADDRESS:0
		ram.ReadWrite('1', '0',
			"00011111", "0001") // ADDRESS:1
		ram.ReadWrite('1', '0',
			"10000100", "0010") // ADDRESS:2
		ram.ReadWrite('1', '0',
			"01001101", "0011") // ADDRESS:3
		ram.ReadWrite('1', '0',
			"00000011", "1110") // ADDRESS:14
		ram.ReadWrite('1', '0',
			"00001110", "1111") // ADDRESS:15

		aRegister := stringutil.NewEightBitRegister()
		//bRegister := stringutil.NewEightBitRegister()
		//cRegister := stringutil.NewEightBitRegister()
		//dRegister := stringutil.NewEightBitRegister()

		instructionAddressRegister := stringutil.NewFourBitRegister()
		instructionRegister := stringutil.NewEightBitRegister()

		_, ramOutput0 := ram.ReadWrite('0', '1',
			"00000000",
			instructionAddressRegister.ReadWrite('0',
				"0000"))
		iRV := instructionRegister.ReadWrite('1', ramOutput0)
	    bc := stringutil.LoadA(iRV[0:4])
		_, bf := ram.ReadWrite('0', bc,"00000000",
			iRV[4:8])
	    result := aRegister.ReadWrite(bc, bf)

	    fmt.Println("\nLOAD_A")
		fmt.Println("\n", result)

	*/
	fmt.Println("\nCPU")

	cpu1 := stringutil.NewCPU()
	cpu1.Ram.ReadWrite('1', '0',
		"00101110", "0000") // ADDRESS:0
	cpu1.Ram.ReadWrite('1', '0',
		"00011111", "0001") // ADDRESS:1
	cpu1.Ram.ReadWrite('1', '0',
		"10000100", "0010") // ADDRESS:2
	cpu1.Ram.ReadWrite('1', '0',
		"01001101", "0011") // ADDRESS:3
	cpu1.Ram.ReadWrite('1', '0',
		"00000011", "1110") // ADDRESS:14
	cpu1.Ram.ReadWrite('1', '0',
		"00001110", "1111") // ADDRESS:15
	fmt.Println("\n\ncpu1.Execute()开始")

	cpu1.Execute()

	fmt.Println("*******************")
	fmt.Println(cpu1.Ram.ReadWrite('0', '1',
		"00000000", "1101"))
}
