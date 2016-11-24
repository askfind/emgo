// Peripheral: DMA2D_Periph  DMA2D Controller.
// Instances:
//  DMA2D  mmap.DMA2D_BASE
// Registers:
//  0x00  32  CR          Control Register.
//  0x04  32  ISR         Interrupt Status Register.
//  0x08  32  IFCR        Interrupt Flag Clear Register.
//  0x0C  32  FGMAR       Foreground Memory Address Register.
//  0x10  32  FGOR        Foreground Offset Register.
//  0x14  32  BGMAR       Background Memory Address Register.
//  0x18  32  BGOR        Background Offset Register.
//  0x1C  32  FGPFCCR     Foreground PFC Control Register.
//  0x20  32  FGCOLR      Foreground Color Register.
//  0x24  32  BGPFCCR     Background PFC Control Register.
//  0x28  32  BGCOLR      Background Color Register.
//  0x2C  32  FGCMAR      Foreground CLUT Memory Address Register.
//  0x30  32  BGCMAR      Background CLUT Memory Address Register.
//  0x34  32  OPFCCR      Output PFC Control Register.
//  0x38  32  OCOLR       Output Color Register.
//  0x3C  32  OMAR        Output Memory Address Register.
//  0x40  32  OOR         Output Offset Register.
//  0x44  32  NLR         Number of Line Register.
//  0x48  32  LWR         Line Watermark Register.
//  0x4C  32  AMTCR       AHB Master Timer Configuration Register.
//  0x400 32  FGCLUT[256] Foreground CLUT.
//  0x800 32  BGCLUT[256] Background CLUT.
// Import:
//  stm32/o/f40_41xxx/mmap
package dma2d

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	START CR_Bits = 0x01 << 0  //+ Start transfer.
	SUSP  CR_Bits = 0x01 << 1  //+ Suspend transfer.
	ABORT CR_Bits = 0x01 << 2  //+ Abort transfer.
	TEIE  CR_Bits = 0x01 << 8  //+ Transfer Error Interrupt Enable.
	TCIE  CR_Bits = 0x01 << 9  //+ Transfer Complete Interrupt Enable.
	TWIE  CR_Bits = 0x01 << 10 //+ Transfer Watermark Interrupt Enable.
	CAEIE CR_Bits = 0x01 << 11 //+ CLUT Access Error Interrupt Enable.
	CTCIE CR_Bits = 0x01 << 12 //+ CLUT Transfer Complete Interrupt Enable.
	CEIE  CR_Bits = 0x01 << 13 //+ Configuration Error Interrupt Enable.
	MODE  CR_Bits = 0x03 << 16 //+ DMA2D Mode.
)

const (
	STARTn = 0
	SUSPn  = 1
	ABORTn = 2
	TEIEn  = 8
	TCIEn  = 9
	TWIEn  = 10
	CAEIEn = 11
	CTCIEn = 12
	CEIEn  = 13
	MODEn  = 16
)

const (
	TEIF  ISR_Bits = 0x01 << 0 //+ Transfer Error Interrupt Flag.
	TCIF  ISR_Bits = 0x01 << 1 //+ Transfer Complete Interrupt Flag.
	TWIF  ISR_Bits = 0x01 << 2 //+ Transfer Watermark Interrupt Flag.
	CAEIF ISR_Bits = 0x01 << 3 //+ CLUT Access Error Interrupt Flag.
	CTCIF ISR_Bits = 0x01 << 4 //+ CLUT Transfer Complete Interrupt Flag.
	CEIF  ISR_Bits = 0x01 << 5 //+ Configuration Error Interrupt Flag.
)

const (
	TEIFn  = 0
	TCIFn  = 1
	TWIFn  = 2
	CAEIFn = 3
	CTCIFn = 4
	CEIFn  = 5
)

const (
	CTEIF  IFCR_Bits = 0x01 << 0 //+ Clears Transfer Error Interrupt Flag.
	CTCIF  IFCR_Bits = 0x01 << 1 //+ Clears Transfer Complete Interrupt Flag.
	CTWIF  IFCR_Bits = 0x01 << 2 //+ Clears Transfer Watermark Interrupt Flag.
	CAECIF IFCR_Bits = 0x01 << 3 //+ Clears CLUT Access Error Interrupt Flag.
	CCTCIF IFCR_Bits = 0x01 << 4 //+ Clears CLUT Transfer Complete Interrupt Flag.
	CCEIF  IFCR_Bits = 0x01 << 5 //+ Clears Configuration Error Interrupt Flag.
)

const (
	CTEIFn  = 0
	CTCIFn  = 1
	CTWIFn  = 2
	CAECIFn = 3
	CCTCIFn = 4
	CCEIFn  = 5
)

const (
	MA FGMAR_Bits = 0xFFFFFFFF << 0 //+ Memory Address.
)

const (
	MAn = 0
)

const (
	LO FGOR_Bits = 0x3FFF << 0 //+ Line Offset.
)

const (
	LOn = 0
)

const (
	MA BGMAR_Bits = 0xFFFFFFFF << 0 //+ Memory Address.
)

const (
	MAn = 0
)

const (
	LO BGOR_Bits = 0x3FFF << 0 //+ Line Offset.
)

const (
	LOn = 0
)

const (
	CM    FGPFCCR_Bits = 0x0F << 0  //+ Input color mode CM[3:0].
	CM_0  FGPFCCR_Bits = 0x01 << 0  //  Input color mode CM bit 0.
	CM_1  FGPFCCR_Bits = 0x02 << 0  //  Input color mode CM bit 1.
	CM_2  FGPFCCR_Bits = 0x04 << 0  //  Input color mode CM bit 2.
	CM_3  FGPFCCR_Bits = 0x08 << 0  //  Input color mode CM bit 3.
	CCM   FGPFCCR_Bits = 0x01 << 4  //+ CLUT Color mode.
	START FGPFCCR_Bits = 0x01 << 5  //+ Start.
	CS    FGPFCCR_Bits = 0xFF << 8  //+ CLUT size.
	AM    FGPFCCR_Bits = 0x03 << 16 //+ Alpha mode AM[1:0].
	AM_0  FGPFCCR_Bits = 0x01 << 16 //  Alpha mode AM bit 0.
	AM_1  FGPFCCR_Bits = 0x02 << 16 //  Alpha mode AM bit 1.
	ALPHA FGPFCCR_Bits = 0xFF << 24 //+ Alpha value.
	CM_3  FGPFCCR_Bits = 0x08 << 0  //  Input color mode CM bit 3.
)

const (
	CMn    = 0
	CCMn   = 4
	STARTn = 5
	CSn    = 8
	AMn    = 16
	ALPHAn = 24
)

const (
	BLUE  FGCOLR_Bits = 0xFF << 0  //+ Blue Value.
	GREEN FGCOLR_Bits = 0xFF << 8  //+ Green Value.
	RED   FGCOLR_Bits = 0xFF << 16 //+ Red Value.
)

const (
	BLUEn  = 0
	GREENn = 8
	REDn   = 16
)

const (
	CM    BGPFCCR_Bits = 0x0F << 0  //+ Input color mode CM[3:0].
	CM_0  BGPFCCR_Bits = 0x01 << 0  //  Input color mode CM bit 0.
	CM_1  BGPFCCR_Bits = 0x02 << 0  //  Input color mode CM bit 1.
	CM_2  BGPFCCR_Bits = 0x04 << 0  //  Input color mode CM bit 2.
	CCM   BGPFCCR_Bits = 0x01 << 4  //+ CLUT Color mode.
	START BGPFCCR_Bits = 0x01 << 5  //+ Start.
	CS    BGPFCCR_Bits = 0xFF << 8  //+ CLUT size.
	AM    BGPFCCR_Bits = 0x03 << 16 //+ Alpha mode AM[1:0].
	AM_0  BGPFCCR_Bits = 0x01 << 16 //  Alpha mode AM bit 0.
	AM_1  BGPFCCR_Bits = 0x02 << 16 //  Alpha mode AM bit 1.
	ALPHA BGPFCCR_Bits = 0xFF << 24 //+ Alpha value.
)

const (
	CMn    = 0
	CCMn   = 4
	STARTn = 5
	CSn    = 8
	AMn    = 16
	ALPHAn = 24
)

const (
	BLUE  BGCOLR_Bits = 0xFF << 0  //+ Blue Value.
	GREEN BGCOLR_Bits = 0xFF << 8  //+ Green Value.
	RED   BGCOLR_Bits = 0xFF << 16 //+ Red Value.
)

const (
	BLUEn  = 0
	GREENn = 8
	REDn   = 16
)

const (
	MA FGCMAR_Bits = 0xFFFFFFFF << 0 //+ Memory Address.
)

const (
	MAn = 0
)

const (
	MA BGCMAR_Bits = 0xFFFFFFFF << 0 //+ Memory Address.
)

const (
	MAn = 0
)

const (
	CM   OPFCCR_Bits = 0x07 << 0 //+ Color mode CM[2:0].
	CM_0 OPFCCR_Bits = 0x01 << 0 //  Color mode CM bit 0.
	CM_1 OPFCCR_Bits = 0x02 << 0 //  Color mode CM bit 1.
	CM_2 OPFCCR_Bits = 0x04 << 0 //  Color mode CM bit 2.
)

const (
	CMn = 0
)

const (
	BLUE_1  OCOLR_Bits = 0xFF << 0  //+ BLUE Value.
	GREEN_1 OCOLR_Bits = 0xFF << 8  //+ GREEN Value.
	RED_1   OCOLR_Bits = 0xFF << 16 //+ Red Value.
	ALPHA_1 OCOLR_Bits = 0xFF << 24 //+ Alpha Channel Value.
	BLUE_2  OCOLR_Bits = 0x1F << 0  //  BLUE Value.
	GREEN_2 OCOLR_Bits = 0x7E0 << 0 //  GREEN Value.
	RED_2   OCOLR_Bits = 0xF8 << 8  //  Red Value.
	BLUE_3  OCOLR_Bits = 0x1F << 0  //  BLUE Value.
	GREEN_3 OCOLR_Bits = 0x3E0 << 0 //  GREEN Value.
	RED_3   OCOLR_Bits = 0x7C << 8  //  Red Value.
	ALPHA_3 OCOLR_Bits = 0x80 << 8  //  Alpha Channel Value.
	BLUE_4  OCOLR_Bits = 0x0F << 0  //  BLUE Value.
	GREEN_4 OCOLR_Bits = 0xF0 << 0  //  GREEN Value.
	RED_4   OCOLR_Bits = 0x0F << 8  //  Red Value.
	ALPHA_4 OCOLR_Bits = 0xF0 << 8  //  Alpha Channel Value.
)

const (
	BLUE_1n  = 0
	GREEN_1n = 8
	RED_1n   = 16
	ALPHA_1n = 24
)

const (
	MA OMAR_Bits = 0xFFFFFFFF << 0 //+ Memory Address.
)

const (
	MAn = 0
)

const (
	LO OOR_Bits = 0x3FFF << 0 //+ Line Offset.
)

const (
	LOn = 0
)

const (
	NL NLR_Bits = 0xFFFF << 0  //+ Number of Lines.
	PL NLR_Bits = 0x3FFF << 16 //+ Pixel per Lines.
)

const (
	NLn = 0
	PLn = 16
)

const (
	LW LWR_Bits = 0xFFFF << 0 //+ Line Watermark.
)

const (
	LWn = 0
)

const (
	EN AMTCR_Bits = 0x01 << 0 //+ Enable.
	DT AMTCR_Bits = 0xFF << 8 //+ Dead Time.
)

const (
	ENn = 0
	DTn = 8
)
