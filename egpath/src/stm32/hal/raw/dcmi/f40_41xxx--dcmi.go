// +build f40_41xxx

// Peripheral: DCMI_Periph  DCMI.
// Instances:
//  DCMI  mmap.DCMI_BASE
// Registers:
//  0x00 32  CR      Control register 1.
//  0x04 32  SR      Status register.
//  0x08 32  RISR    Raw interrupt status register.
//  0x0C 32  IER     Interrupt enable register.
//  0x10 32  MISR    Masked interrupt status register.
//  0x14 32  ICR     Interrupt clear register.
//  0x18 32  ESCR    Embedded synchronization code register.
//  0x1C 32  ESUR    Embedded synchronization unmask register.
//  0x20 32  CWSTRTR Crop window start.
//  0x24 32  CWSIZER Crop window size.
//  0x28 32  DR      Data register.
// Import:
//  stm32/o/f40_41xxx/mmap
package dcmi

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	CAPTURE CR_Bits = 0x01 << 0  //+
	CM      CR_Bits = 0x01 << 1  //+
	CROP    CR_Bits = 0x01 << 2  //+
	JPEG    CR_Bits = 0x01 << 3  //+
	ESS     CR_Bits = 0x01 << 4  //+
	PCKPOL  CR_Bits = 0x01 << 5  //+
	HSPOL   CR_Bits = 0x01 << 6  //+
	VSPOL   CR_Bits = 0x01 << 7  //+
	FCRC_0  CR_Bits = 0x01 << 8  //+
	FCRC_1  CR_Bits = 0x01 << 9  //+
	EDM_0   CR_Bits = 0x01 << 10 //+
	EDM_1   CR_Bits = 0x01 << 11 //+
	CRE     CR_Bits = 0x01 << 12 //+
	ENABLE  CR_Bits = 0x01 << 14 //+
)

const (
	CAPTUREn = 0
	CMn      = 1
	CROPn    = 2
	JPEGn    = 3
	ESSn     = 4
	PCKPOLn  = 5
	HSPOLn   = 6
	VSPOLn   = 7
	FCRC_0n  = 8
	FCRC_1n  = 9
	EDM_0n   = 10
	EDM_1n   = 11
	CREn     = 12
	ENABLEn  = 14
)

const (
	HSYNC SR_Bits = 0x01 << 0 //+
	VSYNC SR_Bits = 0x01 << 1 //+
	FNE   SR_Bits = 0x01 << 2 //+
)

const (
	HSYNCn = 0
	VSYNCn = 1
	FNEn   = 2
)

const (
	FRAME_IE IER_Bits = 0x01 << 0 //+
	OVR_IE   IER_Bits = 0x01 << 1 //+
	ERR_IE   IER_Bits = 0x01 << 2 //+
	VSYNC_IE IER_Bits = 0x01 << 3 //+
	LINE_IE  IER_Bits = 0x01 << 4 //+
)

const (
	FRAME_IEn = 0
	OVR_IEn   = 1
	ERR_IEn   = 2
	VSYNC_IEn = 3
	LINE_IEn  = 4
)

const (
	FRAME_ISC ICR_Bits = 0x01 << 0 //+
	OVR_ISC   ICR_Bits = 0x01 << 1 //+
	ERR_ISC   ICR_Bits = 0x01 << 2 //+
	VSYNC_ISC ICR_Bits = 0x01 << 3 //+
	LINE_ISC  ICR_Bits = 0x01 << 4 //+
)

const (
	FRAME_ISCn = 0
	OVR_ISCn   = 1
	ERR_ISCn   = 2
	VSYNC_ISCn = 3
	LINE_ISCn  = 4
)

const (
	FSC ESCR_Bits = 0xFF << 0  //+
	LSC ESCR_Bits = 0xFF << 8  //+
	LEC ESCR_Bits = 0xFF << 16 //+
	FEC ESCR_Bits = 0xFF << 24 //+
)

const (
	FSCn = 0
	LSCn = 8
	LECn = 16
	FECn = 24
)

const (
	FSU ESUR_Bits = 0xFF << 0  //+
	LSU ESUR_Bits = 0xFF << 8  //+
	LEU ESUR_Bits = 0xFF << 16 //+
	FEU ESUR_Bits = 0xFF << 24 //+
)

const (
	FSUn = 0
	LSUn = 8
	LEUn = 16
	FEUn = 24
)

const (
	BYTE0 DR_Bits = 0xFF << 0  //+
	BYTE1 DR_Bits = 0xFF << 8  //+
	BYTE2 DR_Bits = 0xFF << 16 //+
	BYTE3 DR_Bits = 0xFF << 24 //+
)

const (
	BYTE0n = 0
	BYTE1n = 8
	BYTE2n = 16
	BYTE3n = 24
)
