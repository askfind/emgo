// Peripheral: PWR_Periph  Power Control.
// Instances:
//  PWR  mmap.PWR_BASE
// Registers:
//  0x00 32  CR1  Power control register 1.
//  0x04 32  CSR1 Power control/status register 2.
//  0x08 32  CR2  Power control register 2.
//  0x0C 32  CSR2 Power control/status register 2.
// Import:
//  stm32/o/f746xx/mmap
package pwr

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	LPDS     CR1_Bits = 0x01 << 0  //+ Low-Power Deepsleep.
	PDDS     CR1_Bits = 0x01 << 1  //+ Power Down Deepsleep.
	CSBF     CR1_Bits = 0x01 << 3  //+ Clear Standby Flag.
	PVDE     CR1_Bits = 0x01 << 4  //+ Power Voltage Detector Enable.
	PLS      CR1_Bits = 0x07 << 5  //+ PLS[2:0] bits (PVD Level Selection).
	PLS_0    CR1_Bits = 0x01 << 5  //  Bit 0.
	PLS_1    CR1_Bits = 0x02 << 5  //  Bit 1.
	PLS_2    CR1_Bits = 0x04 << 5  //  Bit 2.
	PLS_LEV0 CR1_Bits = 0x00 << 5  //  PVD level 0.
	PLS_LEV1 CR1_Bits = 0x01 << 5  //  PVD level 1.
	PLS_LEV2 CR1_Bits = 0x02 << 5  //  PVD level 2.
	PLS_LEV3 CR1_Bits = 0x03 << 5  //  PVD level 3.
	PLS_LEV4 CR1_Bits = 0x04 << 5  //  PVD level 4.
	PLS_LEV5 CR1_Bits = 0x05 << 5  //  PVD level 5.
	PLS_LEV6 CR1_Bits = 0x06 << 5  //  PVD level 6.
	PLS_LEV7 CR1_Bits = 0x07 << 5  //  PVD level 7.
	DBP      CR1_Bits = 0x01 << 8  //+ Disable Backup Domain write protection.
	FPDS     CR1_Bits = 0x01 << 9  //+ Flash power down in Stop mode.
	LPUDS    CR1_Bits = 0x01 << 10 //+ Low-power regulator in deepsleep under-drive mode.
	MRUDS    CR1_Bits = 0x01 << 11 //+ Main regulator in deepsleep under-drive mode.
	ADCDC1   CR1_Bits = 0x01 << 13 //+ Refer to AN4073 on how to use this bit.
	VOS      CR1_Bits = 0x03 << 14 //+ VOS[1:0] bits (Regulator voltage scaling output selection).
	VOS_0    CR1_Bits = 0x01 << 14 //  Bit 0.
	VOS_1    CR1_Bits = 0x02 << 14 //  Bit 1.
	ODEN     CR1_Bits = 0x01 << 16 //+ Over Drive enable.
	ODSWEN   CR1_Bits = 0x01 << 17 //+ Over Drive switch enabled.
	UDEN     CR1_Bits = 0x03 << 18 //+ Under Drive enable in stop mode.
	UDEN_0   CR1_Bits = 0x01 << 18 //  Bit 0.
	UDEN_1   CR1_Bits = 0x02 << 18 //  Bit 1.
)

const (
	LPDSn   = 0
	PDDSn   = 1
	CSBFn   = 3
	PVDEn   = 4
	PLSn    = 5
	DBPn    = 8
	FPDSn   = 9
	LPUDSn  = 10
	MRUDSn  = 11
	ADCDC1n = 13
	VOSn    = 14
	ODENn   = 16
	ODSWENn = 17
	UDENn   = 18
)

const (
	WUIF    CSR1_Bits = 0x01 << 0  //+ Wake up internal Flag.
	SBF     CSR1_Bits = 0x01 << 1  //+ Standby Flag.
	PVDO    CSR1_Bits = 0x01 << 2  //+ PVD Output.
	BRR     CSR1_Bits = 0x01 << 3  //+ Backup regulator ready.
	EIWUP   CSR1_Bits = 0x01 << 8  //+ Enable internal wakeup.
	BRE     CSR1_Bits = 0x01 << 9  //+ Backup regulator enable.
	VOSRDY  CSR1_Bits = 0x01 << 14 //+ Regulator voltage scaling output selection ready.
	ODRDY   CSR1_Bits = 0x01 << 16 //+ Over Drive generator ready.
	ODSWRDY CSR1_Bits = 0x01 << 17 //+ Over Drive Switch ready.
	UDRDY   CSR1_Bits = 0x03 << 18 //+ Under Drive ready.
)

const (
	WUIFn    = 0
	SBFn     = 1
	PVDOn    = 2
	BRRn     = 3
	EIWUPn   = 8
	BREn     = 9
	VOSRDYn  = 14
	ODRDYn   = 16
	ODSWRDYn = 17
	UDRDYn   = 18
)

const (
	CWUPF1 CR2_Bits = 0x01 << 0  //+ Clear Wakeup Pin Flag for PA0.
	CWUPF2 CR2_Bits = 0x01 << 1  //+ Clear Wakeup Pin Flag for PA2.
	CWUPF3 CR2_Bits = 0x01 << 2  //+ Clear Wakeup Pin Flag for PC1.
	CWUPF4 CR2_Bits = 0x01 << 3  //+ Clear Wakeup Pin Flag for PC13.
	CWUPF5 CR2_Bits = 0x01 << 4  //+ Clear Wakeup Pin Flag for PI8.
	CWUPF6 CR2_Bits = 0x01 << 5  //+ Clear Wakeup Pin Flag for PI11.
	WUPP1  CR2_Bits = 0x01 << 8  //+ Wakeup Pin Polarity bit for PA0.
	WUPP2  CR2_Bits = 0x01 << 9  //+ Wakeup Pin Polarity bit for PA2.
	WUPP3  CR2_Bits = 0x01 << 10 //+ Wakeup Pin Polarity bit for PC1.
	WUPP4  CR2_Bits = 0x01 << 11 //+ Wakeup Pin Polarity bit for PC13.
	WUPP5  CR2_Bits = 0x01 << 12 //+ Wakeup Pin Polarity bit for PI8.
	WUPP6  CR2_Bits = 0x01 << 13 //+ Wakeup Pin Polarity bit for PI11.
)

const (
	CWUPF1n = 0
	CWUPF2n = 1
	CWUPF3n = 2
	CWUPF4n = 3
	CWUPF5n = 4
	CWUPF6n = 5
	WUPP1n  = 8
	WUPP2n  = 9
	WUPP3n  = 10
	WUPP4n  = 11
	WUPP5n  = 12
	WUPP6n  = 13
)

const (
	WUPF1 CSR2_Bits = 0x01 << 0  //+ Wakeup Pin Flag for PA0.
	WUPF2 CSR2_Bits = 0x01 << 1  //+ Wakeup Pin Flag for PA2.
	WUPF3 CSR2_Bits = 0x01 << 2  //+ Wakeup Pin Flag for PC1.
	WUPF4 CSR2_Bits = 0x01 << 3  //+ Wakeup Pin Flag for PC13.
	WUPF5 CSR2_Bits = 0x01 << 4  //+ Wakeup Pin Flag for PI8.
	WUPF6 CSR2_Bits = 0x01 << 5  //+ Wakeup Pin Flag for PI11.
	EWUP1 CSR2_Bits = 0x01 << 8  //+ Enable Wakeup Pin PA0.
	EWUP2 CSR2_Bits = 0x01 << 9  //+ Enable Wakeup Pin PA2.
	EWUP3 CSR2_Bits = 0x01 << 10 //+ Enable Wakeup Pin PC1.
	EWUP4 CSR2_Bits = 0x01 << 11 //+ Enable Wakeup Pin PC13.
	EWUP5 CSR2_Bits = 0x01 << 12 //+ Enable Wakeup Pin PI8.
	EWUP6 CSR2_Bits = 0x01 << 13 //+ Enable Wakeup Pin PI11.
)

const (
	WUPF1n = 0
	WUPF2n = 1
	WUPF3n = 2
	WUPF4n = 3
	WUPF5n = 4
	WUPF6n = 5
	EWUP1n = 8
	EWUP2n = 9
	EWUP3n = 10
	EWUP4n = 11
	EWUP5n = 12
	EWUP6n = 13
)
