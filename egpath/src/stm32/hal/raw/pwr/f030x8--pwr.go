// +build f030x8

// Peripheral: PWR_Periph  Power Control.
// Instances:
//  PWR  mmap.PWR_BASE
// Registers:
//  0x00 32  CR  Power control register.
//  0x04 32  CSR Power control/status register.
// Import:
//  stm32/o/f030x8/mmap
package pwr

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	LPDS CR = 0x01 << 0 //+ Low-power Deepsleep.
	PDDS CR = 0x01 << 1 //+ Power Down Deepsleep.
	CWUF CR = 0x01 << 2 //+ Clear Wakeup Flag.
	CSBF CR = 0x01 << 3 //+ Clear Standby Flag.
	DBP  CR = 0x01 << 8 //+ Disable Backup Domain write protection.
)

const (
	LPDSn = 0
	PDDSn = 1
	CWUFn = 2
	CSBFn = 3
	DBPn  = 8
)

const (
	WUF   CSR = 0x01 << 0 //+ Wakeup Flag.
	SBF   CSR = 0x01 << 1 //+ Standby Flag.
	EWUP1 CSR = 0x01 << 8 //+ Enable WKUP pin 1.
	EWUP2 CSR = 0x01 << 9 //+ Enable WKUP pin 2.
)

const (
	WUFn   = 0
	SBFn   = 1
	EWUP1n = 8
	EWUP2n = 9
)