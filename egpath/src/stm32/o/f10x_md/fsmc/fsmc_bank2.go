// Peripheral: FSMC_Bank2_Periph  Flexible Static Memory Controller Bank2.
// Instances:
//  FSMC_Bank2  mmap.FSMC_Bank2_R_BASE
// Registers:
//  0x00 32  PCR2
//  0x04 32  SR2
//  0x08 32  PMEM2
//  0x0C 32  PATT2
//  0x14 32  ECCR2
// Import:
//  stm32/o/f10x_md/mmap
package fsmc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	PWAITEN PCR2_Bits = 0x01 << 1  //+ Wait feature enable bit.
	PBKEN   PCR2_Bits = 0x01 << 2  //+ PC Card/NAND Flash memory bank enable bit.
	PTYP    PCR2_Bits = 0x01 << 3  //+ Memory type.
	PWID    PCR2_Bits = 0x03 << 4  //+ PWID[1:0] bits (NAND Flash databus width).
	PWID_0  PCR2_Bits = 0x01 << 4  //  Bit 0.
	PWID_1  PCR2_Bits = 0x02 << 4  //  Bit 1.
	ECCEN   PCR2_Bits = 0x01 << 6  //+ ECC computation logic enable bit.
	TCLR    PCR2_Bits = 0x0F << 9  //+ TCLR[3:0] bits (CLE to RE delay).
	TCLR_0  PCR2_Bits = 0x01 << 9  //  Bit 0.
	TCLR_1  PCR2_Bits = 0x02 << 9  //  Bit 1.
	TCLR_2  PCR2_Bits = 0x04 << 9  //  Bit 2.
	TCLR_3  PCR2_Bits = 0x08 << 9  //  Bit 3.
	TAR     PCR2_Bits = 0x0F << 13 //+ TAR[3:0] bits (ALE to RE delay).
	TAR_0   PCR2_Bits = 0x01 << 13 //  Bit 0.
	TAR_1   PCR2_Bits = 0x02 << 13 //  Bit 1.
	TAR_2   PCR2_Bits = 0x04 << 13 //  Bit 2.
	TAR_3   PCR2_Bits = 0x08 << 13 //  Bit 3.
	ECCPS   PCR2_Bits = 0x07 << 17 //+ ECCPS[1:0] bits (ECC page size).
	ECCPS_0 PCR2_Bits = 0x01 << 17 //  Bit 0.
	ECCPS_1 PCR2_Bits = 0x02 << 17 //  Bit 1.
	ECCPS_2 PCR2_Bits = 0x04 << 17 //  Bit 2.
)

const (
	PWAITENn = 1
	PBKENn   = 2
	PTYPn    = 3
	PWIDn    = 4
	ECCENn   = 6
	TCLRn    = 9
	TARn     = 13
	ECCPSn   = 17
)

const (
	IRS   SR2_Bits = 0x01 << 0 //+ Interrupt Rising Edge status.
	ILS   SR2_Bits = 0x01 << 1 //+ Interrupt Level status.
	IFS   SR2_Bits = 0x01 << 2 //+ Interrupt Falling Edge status.
	IREN  SR2_Bits = 0x01 << 3 //+ Interrupt Rising Edge detection Enable bit.
	ILEN  SR2_Bits = 0x01 << 4 //+ Interrupt Level detection Enable bit.
	IFEN  SR2_Bits = 0x01 << 5 //+ Interrupt Falling Edge detection Enable bit.
	FEMPT SR2_Bits = 0x01 << 6 //+ FIFO empty.
)

const (
	IRSn   = 0
	ILSn   = 1
	IFSn   = 2
	IRENn  = 3
	ILENn  = 4
	IFENn  = 5
	FEMPTn = 6
)

const (
	MEMSET2    PMEM2_Bits = 0xFF << 0  //+ MEMSET2[7:0] bits (Common memory 2 setup time).
	MEMSET2_0  PMEM2_Bits = 0x01 << 0  //  Bit 0.
	MEMSET2_1  PMEM2_Bits = 0x02 << 0  //  Bit 1.
	MEMSET2_2  PMEM2_Bits = 0x04 << 0  //  Bit 2.
	MEMSET2_3  PMEM2_Bits = 0x08 << 0  //  Bit 3.
	MEMSET2_4  PMEM2_Bits = 0x10 << 0  //  Bit 4.
	MEMSET2_5  PMEM2_Bits = 0x20 << 0  //  Bit 5.
	MEMSET2_6  PMEM2_Bits = 0x40 << 0  //  Bit 6.
	MEMSET2_7  PMEM2_Bits = 0x80 << 0  //  Bit 7.
	MEMWAIT2   PMEM2_Bits = 0xFF << 8  //+ MEMWAIT2[7:0] bits (Common memory 2 wait time).
	MEMWAIT2_0 PMEM2_Bits = 0x01 << 8  //  Bit 0.
	MEMWAIT2_1 PMEM2_Bits = 0x02 << 8  //  Bit 1.
	MEMWAIT2_2 PMEM2_Bits = 0x04 << 8  //  Bit 2.
	MEMWAIT2_3 PMEM2_Bits = 0x08 << 8  //  Bit 3.
	MEMWAIT2_4 PMEM2_Bits = 0x10 << 8  //  Bit 4.
	MEMWAIT2_5 PMEM2_Bits = 0x20 << 8  //  Bit 5.
	MEMWAIT2_6 PMEM2_Bits = 0x40 << 8  //  Bit 6.
	MEMWAIT2_7 PMEM2_Bits = 0x80 << 8  //  Bit 7.
	MEMHOLD2   PMEM2_Bits = 0xFF << 16 //+ MEMHOLD2[7:0] bits (Common memory 2 hold time).
	MEMHOLD2_0 PMEM2_Bits = 0x01 << 16 //  Bit 0.
	MEMHOLD2_1 PMEM2_Bits = 0x02 << 16 //  Bit 1.
	MEMHOLD2_2 PMEM2_Bits = 0x04 << 16 //  Bit 2.
	MEMHOLD2_3 PMEM2_Bits = 0x08 << 16 //  Bit 3.
	MEMHOLD2_4 PMEM2_Bits = 0x10 << 16 //  Bit 4.
	MEMHOLD2_5 PMEM2_Bits = 0x20 << 16 //  Bit 5.
	MEMHOLD2_6 PMEM2_Bits = 0x40 << 16 //  Bit 6.
	MEMHOLD2_7 PMEM2_Bits = 0x80 << 16 //  Bit 7.
	MEMHIZ2    PMEM2_Bits = 0xFF << 24 //+ MEMHIZ2[7:0] bits (Common memory 2 databus HiZ time).
	MEMHIZ2_0  PMEM2_Bits = 0x01 << 24 //  Bit 0.
	MEMHIZ2_1  PMEM2_Bits = 0x02 << 24 //  Bit 1.
	MEMHIZ2_2  PMEM2_Bits = 0x04 << 24 //  Bit 2.
	MEMHIZ2_3  PMEM2_Bits = 0x08 << 24 //  Bit 3.
	MEMHIZ2_4  PMEM2_Bits = 0x10 << 24 //  Bit 4.
	MEMHIZ2_5  PMEM2_Bits = 0x20 << 24 //  Bit 5.
	MEMHIZ2_6  PMEM2_Bits = 0x40 << 24 //  Bit 6.
	MEMHIZ2_7  PMEM2_Bits = 0x80 << 24 //  Bit 7.
)

const (
	MEMSET2n  = 0
	MEMWAIT2n = 8
	MEMHOLD2n = 16
	MEMHIZ2n  = 24
)

const (
	ATTSET2    PATT2_Bits = 0xFF << 0  //+ ATTSET2[7:0] bits (Attribute memory 2 setup time).
	ATTSET2_0  PATT2_Bits = 0x01 << 0  //  Bit 0.
	ATTSET2_1  PATT2_Bits = 0x02 << 0  //  Bit 1.
	ATTSET2_2  PATT2_Bits = 0x04 << 0  //  Bit 2.
	ATTSET2_3  PATT2_Bits = 0x08 << 0  //  Bit 3.
	ATTSET2_4  PATT2_Bits = 0x10 << 0  //  Bit 4.
	ATTSET2_5  PATT2_Bits = 0x20 << 0  //  Bit 5.
	ATTSET2_6  PATT2_Bits = 0x40 << 0  //  Bit 6.
	ATTSET2_7  PATT2_Bits = 0x80 << 0  //  Bit 7.
	ATTWAIT2   PATT2_Bits = 0xFF << 8  //+ ATTWAIT2[7:0] bits (Attribute memory 2 wait time).
	ATTWAIT2_0 PATT2_Bits = 0x01 << 8  //  Bit 0.
	ATTWAIT2_1 PATT2_Bits = 0x02 << 8  //  Bit 1.
	ATTWAIT2_2 PATT2_Bits = 0x04 << 8  //  Bit 2.
	ATTWAIT2_3 PATT2_Bits = 0x08 << 8  //  Bit 3.
	ATTWAIT2_4 PATT2_Bits = 0x10 << 8  //  Bit 4.
	ATTWAIT2_5 PATT2_Bits = 0x20 << 8  //  Bit 5.
	ATTWAIT2_6 PATT2_Bits = 0x40 << 8  //  Bit 6.
	ATTWAIT2_7 PATT2_Bits = 0x80 << 8  //  Bit 7.
	ATTHOLD2   PATT2_Bits = 0xFF << 16 //+ ATTHOLD2[7:0] bits (Attribute memory 2 hold time).
	ATTHOLD2_0 PATT2_Bits = 0x01 << 16 //  Bit 0.
	ATTHOLD2_1 PATT2_Bits = 0x02 << 16 //  Bit 1.
	ATTHOLD2_2 PATT2_Bits = 0x04 << 16 //  Bit 2.
	ATTHOLD2_3 PATT2_Bits = 0x08 << 16 //  Bit 3.
	ATTHOLD2_4 PATT2_Bits = 0x10 << 16 //  Bit 4.
	ATTHOLD2_5 PATT2_Bits = 0x20 << 16 //  Bit 5.
	ATTHOLD2_6 PATT2_Bits = 0x40 << 16 //  Bit 6.
	ATTHOLD2_7 PATT2_Bits = 0x80 << 16 //  Bit 7.
	ATTHIZ2    PATT2_Bits = 0xFF << 24 //+ ATTHIZ2[7:0] bits (Attribute memory 2 databus HiZ time).
	ATTHIZ2_0  PATT2_Bits = 0x01 << 24 //  Bit 0.
	ATTHIZ2_1  PATT2_Bits = 0x02 << 24 //  Bit 1.
	ATTHIZ2_2  PATT2_Bits = 0x04 << 24 //  Bit 2.
	ATTHIZ2_3  PATT2_Bits = 0x08 << 24 //  Bit 3.
	ATTHIZ2_4  PATT2_Bits = 0x10 << 24 //  Bit 4.
	ATTHIZ2_5  PATT2_Bits = 0x20 << 24 //  Bit 5.
	ATTHIZ2_6  PATT2_Bits = 0x40 << 24 //  Bit 6.
	ATTHIZ2_7  PATT2_Bits = 0x80 << 24 //  Bit 7.
)

const (
	ATTSET2n  = 0
	ATTWAIT2n = 8
	ATTHOLD2n = 16
	ATTHIZ2n  = 24
)

const (
	ECC2 ECCR2_Bits = 0xFFFFFFFF << 0 //+ ECC result.
)

const (
	ECC2n = 0
)
