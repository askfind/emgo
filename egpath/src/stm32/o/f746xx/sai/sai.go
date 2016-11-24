// Peripheral: SAI_Periph  Serial Audio Interface.
// Instances:
//  SAI1  mmap.SAI1_BASE
//  SAI2  mmap.SAI2_BASE
// Registers:
//  0x00 32  GCR Global configuration register.
// Import:
//  stm32/o/f746xx/mmap
package sai

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	SYNCIN    GCR_Bits = 0x03 << 0 //+ SYNCIN[1:0] bits (Synchronization Inputs).
	SYNCIN_0  GCR_Bits = 0x01 << 0 //  Bit 0.
	SYNCIN_1  GCR_Bits = 0x02 << 0 //  Bit 1.
	SYNCOUT   GCR_Bits = 0x03 << 4 //+ SYNCOUT[1:0] bits (Synchronization Outputs).
	SYNCOUT_0 GCR_Bits = 0x01 << 4 //  Bit 0.
	SYNCOUT_1 GCR_Bits = 0x02 << 4 //  Bit 1.
)

const (
	SYNCINn  = 0
	SYNCOUTn = 4
)
