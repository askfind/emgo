// +build f40_41xxx

package spi

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f40_41xxx/mmap"
)

type SPI_Periph struct {
	CR1     CR1
	_       uint16
	CR2     CR2
	_       uint16
	SR      SR
	_       uint16
	DR      DR
	_       uint16
	CRCPR   CRCPR
	_       uint16
	RXCRCR  RXCRCR
	_       uint16
	TXCRCR  TXCRCR
	_       uint16
	I2SCFGR I2SCFGR
	_       uint16
	I2SPR   I2SPR
}

func (p *SPI_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var I2S2ext = (*SPI_Periph)(unsafe.Pointer(uintptr(mmap.I2S2ext_BASE)))

var SPI2 = (*SPI_Periph)(unsafe.Pointer(uintptr(mmap.SPI2_BASE)))

var SPI3 = (*SPI_Periph)(unsafe.Pointer(uintptr(mmap.SPI3_BASE)))

var I2S3ext = (*SPI_Periph)(unsafe.Pointer(uintptr(mmap.I2S3ext_BASE)))

var SPI1 = (*SPI_Periph)(unsafe.Pointer(uintptr(mmap.SPI1_BASE)))

var SPI4 = (*SPI_Periph)(unsafe.Pointer(uintptr(mmap.SPI4_BASE)))

var SPI5 = (*SPI_Periph)(unsafe.Pointer(uintptr(mmap.SPI5_BASE)))

var SPI6 = (*SPI_Periph)(unsafe.Pointer(uintptr(mmap.SPI6_BASE)))

type CR1_Bits uint16

func (b CR1_Bits) Field(mask CR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR1_Bits) J(v int) CR1_Bits {
	return CR1_Bits(bits.Make32(v, uint32(mask)))
}

type CR1 struct{ mmio.U16 }

func (r *CR1) Bits(mask CR1_Bits) CR1_Bits { return CR1_Bits(r.U16.Bits(uint16(mask))) }
func (r *CR1) StoreBits(mask, b CR1_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CR1) SetBits(mask CR1_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *CR1) ClearBits(mask CR1_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *CR1) Load() CR1_Bits              { return CR1_Bits(r.U16.Load()) }
func (r *CR1) Store(b CR1_Bits)            { r.U16.Store(uint16(b)) }

type CR1_Mask struct{ mmio.UM16 }

func (rm CR1_Mask) Load() CR1_Bits   { return CR1_Bits(rm.UM16.Load()) }
func (rm CR1_Mask) Store(b CR1_Bits) { rm.UM16.Store(uint16(b)) }

func (p *SPI_Periph) CPHA() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(CPHA)}}
}

func (p *SPI_Periph) CPOL() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(CPOL)}}
}

func (p *SPI_Periph) MSTR() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(MSTR)}}
}

func (p *SPI_Periph) BR() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(BR)}}
}

func (p *SPI_Periph) SPE() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(SPE)}}
}

func (p *SPI_Periph) LSBFIRST() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(LSBFIRST)}}
}

func (p *SPI_Periph) SSI() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(SSI)}}
}

func (p *SPI_Periph) SSM() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(SSM)}}
}

func (p *SPI_Periph) RXONLY() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(RXONLY)}}
}

func (p *SPI_Periph) DFF() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(DFF)}}
}

func (p *SPI_Periph) CRCNEXT() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(CRCNEXT)}}
}

func (p *SPI_Periph) CRCEN() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(CRCEN)}}
}

func (p *SPI_Periph) BIDIOE() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(BIDIOE)}}
}

func (p *SPI_Periph) BIDIMODE() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(BIDIMODE)}}
}

type CR2_Bits uint16

func (b CR2_Bits) Field(mask CR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR2_Bits) J(v int) CR2_Bits {
	return CR2_Bits(bits.Make32(v, uint32(mask)))
}

type CR2 struct{ mmio.U16 }

func (r *CR2) Bits(mask CR2_Bits) CR2_Bits { return CR2_Bits(r.U16.Bits(uint16(mask))) }
func (r *CR2) StoreBits(mask, b CR2_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CR2) SetBits(mask CR2_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *CR2) ClearBits(mask CR2_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *CR2) Load() CR2_Bits              { return CR2_Bits(r.U16.Load()) }
func (r *CR2) Store(b CR2_Bits)            { r.U16.Store(uint16(b)) }

type CR2_Mask struct{ mmio.UM16 }

func (rm CR2_Mask) Load() CR2_Bits   { return CR2_Bits(rm.UM16.Load()) }
func (rm CR2_Mask) Store(b CR2_Bits) { rm.UM16.Store(uint16(b)) }

func (p *SPI_Periph) RXDMAEN() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(RXDMAEN)}}
}

func (p *SPI_Periph) TXDMAEN() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(TXDMAEN)}}
}

func (p *SPI_Periph) SSOE() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(SSOE)}}
}

func (p *SPI_Periph) ERRIE() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(ERRIE)}}
}

func (p *SPI_Periph) RXNEIE() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(RXNEIE)}}
}

func (p *SPI_Periph) TXEIE() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(TXEIE)}}
}

type SR_Bits uint16

func (b SR_Bits) Field(mask SR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR_Bits) J(v int) SR_Bits {
	return SR_Bits(bits.Make32(v, uint32(mask)))
}

type SR struct{ mmio.U16 }

func (r *SR) Bits(mask SR_Bits) SR_Bits { return SR_Bits(r.U16.Bits(uint16(mask))) }
func (r *SR) StoreBits(mask, b SR_Bits) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *SR) SetBits(mask SR_Bits)      { r.U16.SetBits(uint16(mask)) }
func (r *SR) ClearBits(mask SR_Bits)    { r.U16.ClearBits(uint16(mask)) }
func (r *SR) Load() SR_Bits             { return SR_Bits(r.U16.Load()) }
func (r *SR) Store(b SR_Bits)           { r.U16.Store(uint16(b)) }

type SR_Mask struct{ mmio.UM16 }

func (rm SR_Mask) Load() SR_Bits   { return SR_Bits(rm.UM16.Load()) }
func (rm SR_Mask) Store(b SR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *SPI_Periph) RXNE() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(RXNE)}}
}

func (p *SPI_Periph) TXE() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(TXE)}}
}

func (p *SPI_Periph) CHSIDE() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(CHSIDE)}}
}

func (p *SPI_Periph) UDR() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(UDR)}}
}

func (p *SPI_Periph) CRCERR() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(CRCERR)}}
}

func (p *SPI_Periph) MODF() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(MODF)}}
}

func (p *SPI_Periph) OVR() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(OVR)}}
}

func (p *SPI_Periph) BSY() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(BSY)}}
}

type DR_Bits uint16

func (b DR_Bits) Field(mask DR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DR_Bits) J(v int) DR_Bits {
	return DR_Bits(bits.Make32(v, uint32(mask)))
}

type DR struct{ mmio.U16 }

func (r *DR) Bits(mask DR_Bits) DR_Bits { return DR_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR) StoreBits(mask, b DR_Bits) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR) SetBits(mask DR_Bits)      { r.U16.SetBits(uint16(mask)) }
func (r *DR) ClearBits(mask DR_Bits)    { r.U16.ClearBits(uint16(mask)) }
func (r *DR) Load() DR_Bits             { return DR_Bits(r.U16.Load()) }
func (r *DR) Store(b DR_Bits)           { r.U16.Store(uint16(b)) }

type DR_Mask struct{ mmio.UM16 }

func (rm DR_Mask) Load() DR_Bits   { return DR_Bits(rm.UM16.Load()) }
func (rm DR_Mask) Store(b DR_Bits) { rm.UM16.Store(uint16(b)) }

type CRCPR_Bits uint16

func (b CRCPR_Bits) Field(mask CRCPR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CRCPR_Bits) J(v int) CRCPR_Bits {
	return CRCPR_Bits(bits.Make32(v, uint32(mask)))
}

type CRCPR struct{ mmio.U16 }

func (r *CRCPR) Bits(mask CRCPR_Bits) CRCPR_Bits { return CRCPR_Bits(r.U16.Bits(uint16(mask))) }
func (r *CRCPR) StoreBits(mask, b CRCPR_Bits)    { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CRCPR) SetBits(mask CRCPR_Bits)         { r.U16.SetBits(uint16(mask)) }
func (r *CRCPR) ClearBits(mask CRCPR_Bits)       { r.U16.ClearBits(uint16(mask)) }
func (r *CRCPR) Load() CRCPR_Bits                { return CRCPR_Bits(r.U16.Load()) }
func (r *CRCPR) Store(b CRCPR_Bits)              { r.U16.Store(uint16(b)) }

type CRCPR_Mask struct{ mmio.UM16 }

func (rm CRCPR_Mask) Load() CRCPR_Bits   { return CRCPR_Bits(rm.UM16.Load()) }
func (rm CRCPR_Mask) Store(b CRCPR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *SPI_Periph) CRCPOLY() CRCPR_Mask {
	return CRCPR_Mask{mmio.UM16{&p.CRCPR.U16, uint16(CRCPOLY)}}
}

type RXCRCR_Bits uint16

func (b RXCRCR_Bits) Field(mask RXCRCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RXCRCR_Bits) J(v int) RXCRCR_Bits {
	return RXCRCR_Bits(bits.Make32(v, uint32(mask)))
}

type RXCRCR struct{ mmio.U16 }

func (r *RXCRCR) Bits(mask RXCRCR_Bits) RXCRCR_Bits { return RXCRCR_Bits(r.U16.Bits(uint16(mask))) }
func (r *RXCRCR) StoreBits(mask, b RXCRCR_Bits)     { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RXCRCR) SetBits(mask RXCRCR_Bits)          { r.U16.SetBits(uint16(mask)) }
func (r *RXCRCR) ClearBits(mask RXCRCR_Bits)        { r.U16.ClearBits(uint16(mask)) }
func (r *RXCRCR) Load() RXCRCR_Bits                 { return RXCRCR_Bits(r.U16.Load()) }
func (r *RXCRCR) Store(b RXCRCR_Bits)               { r.U16.Store(uint16(b)) }

type RXCRCR_Mask struct{ mmio.UM16 }

func (rm RXCRCR_Mask) Load() RXCRCR_Bits   { return RXCRCR_Bits(rm.UM16.Load()) }
func (rm RXCRCR_Mask) Store(b RXCRCR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *SPI_Periph) RXCRC() RXCRCR_Mask {
	return RXCRCR_Mask{mmio.UM16{&p.RXCRCR.U16, uint16(RXCRC)}}
}

type TXCRCR_Bits uint16

func (b TXCRCR_Bits) Field(mask TXCRCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TXCRCR_Bits) J(v int) TXCRCR_Bits {
	return TXCRCR_Bits(bits.Make32(v, uint32(mask)))
}

type TXCRCR struct{ mmio.U16 }

func (r *TXCRCR) Bits(mask TXCRCR_Bits) TXCRCR_Bits { return TXCRCR_Bits(r.U16.Bits(uint16(mask))) }
func (r *TXCRCR) StoreBits(mask, b TXCRCR_Bits)     { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *TXCRCR) SetBits(mask TXCRCR_Bits)          { r.U16.SetBits(uint16(mask)) }
func (r *TXCRCR) ClearBits(mask TXCRCR_Bits)        { r.U16.ClearBits(uint16(mask)) }
func (r *TXCRCR) Load() TXCRCR_Bits                 { return TXCRCR_Bits(r.U16.Load()) }
func (r *TXCRCR) Store(b TXCRCR_Bits)               { r.U16.Store(uint16(b)) }

type TXCRCR_Mask struct{ mmio.UM16 }

func (rm TXCRCR_Mask) Load() TXCRCR_Bits   { return TXCRCR_Bits(rm.UM16.Load()) }
func (rm TXCRCR_Mask) Store(b TXCRCR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *SPI_Periph) TXCRC() TXCRCR_Mask {
	return TXCRCR_Mask{mmio.UM16{&p.TXCRCR.U16, uint16(TXCRC)}}
}

type I2SCFGR_Bits uint16

func (b I2SCFGR_Bits) Field(mask I2SCFGR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask I2SCFGR_Bits) J(v int) I2SCFGR_Bits {
	return I2SCFGR_Bits(bits.Make32(v, uint32(mask)))
}

type I2SCFGR struct{ mmio.U16 }

func (r *I2SCFGR) Bits(mask I2SCFGR_Bits) I2SCFGR_Bits { return I2SCFGR_Bits(r.U16.Bits(uint16(mask))) }
func (r *I2SCFGR) StoreBits(mask, b I2SCFGR_Bits)      { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *I2SCFGR) SetBits(mask I2SCFGR_Bits)           { r.U16.SetBits(uint16(mask)) }
func (r *I2SCFGR) ClearBits(mask I2SCFGR_Bits)         { r.U16.ClearBits(uint16(mask)) }
func (r *I2SCFGR) Load() I2SCFGR_Bits                  { return I2SCFGR_Bits(r.U16.Load()) }
func (r *I2SCFGR) Store(b I2SCFGR_Bits)                { r.U16.Store(uint16(b)) }

type I2SCFGR_Mask struct{ mmio.UM16 }

func (rm I2SCFGR_Mask) Load() I2SCFGR_Bits   { return I2SCFGR_Bits(rm.UM16.Load()) }
func (rm I2SCFGR_Mask) Store(b I2SCFGR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *SPI_Periph) CHLEN() I2SCFGR_Mask {
	return I2SCFGR_Mask{mmio.UM16{&p.I2SCFGR.U16, uint16(CHLEN)}}
}

func (p *SPI_Periph) DATLEN() I2SCFGR_Mask {
	return I2SCFGR_Mask{mmio.UM16{&p.I2SCFGR.U16, uint16(DATLEN)}}
}

func (p *SPI_Periph) CKPOL() I2SCFGR_Mask {
	return I2SCFGR_Mask{mmio.UM16{&p.I2SCFGR.U16, uint16(CKPOL)}}
}

func (p *SPI_Periph) I2SSTD() I2SCFGR_Mask {
	return I2SCFGR_Mask{mmio.UM16{&p.I2SCFGR.U16, uint16(I2SSTD)}}
}

func (p *SPI_Periph) PCMSYNC() I2SCFGR_Mask {
	return I2SCFGR_Mask{mmio.UM16{&p.I2SCFGR.U16, uint16(PCMSYNC)}}
}

func (p *SPI_Periph) I2SCFG() I2SCFGR_Mask {
	return I2SCFGR_Mask{mmio.UM16{&p.I2SCFGR.U16, uint16(I2SCFG)}}
}

func (p *SPI_Periph) I2SE() I2SCFGR_Mask {
	return I2SCFGR_Mask{mmio.UM16{&p.I2SCFGR.U16, uint16(I2SE)}}
}

func (p *SPI_Periph) I2SMOD() I2SCFGR_Mask {
	return I2SCFGR_Mask{mmio.UM16{&p.I2SCFGR.U16, uint16(I2SMOD)}}
}

type I2SPR_Bits uint16

func (b I2SPR_Bits) Field(mask I2SPR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask I2SPR_Bits) J(v int) I2SPR_Bits {
	return I2SPR_Bits(bits.Make32(v, uint32(mask)))
}

type I2SPR struct{ mmio.U16 }

func (r *I2SPR) Bits(mask I2SPR_Bits) I2SPR_Bits { return I2SPR_Bits(r.U16.Bits(uint16(mask))) }
func (r *I2SPR) StoreBits(mask, b I2SPR_Bits)    { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *I2SPR) SetBits(mask I2SPR_Bits)         { r.U16.SetBits(uint16(mask)) }
func (r *I2SPR) ClearBits(mask I2SPR_Bits)       { r.U16.ClearBits(uint16(mask)) }
func (r *I2SPR) Load() I2SPR_Bits                { return I2SPR_Bits(r.U16.Load()) }
func (r *I2SPR) Store(b I2SPR_Bits)              { r.U16.Store(uint16(b)) }

type I2SPR_Mask struct{ mmio.UM16 }

func (rm I2SPR_Mask) Load() I2SPR_Bits   { return I2SPR_Bits(rm.UM16.Load()) }
func (rm I2SPR_Mask) Store(b I2SPR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *SPI_Periph) I2SDIV() I2SPR_Mask {
	return I2SPR_Mask{mmio.UM16{&p.I2SPR.U16, uint16(I2SDIV)}}
}

func (p *SPI_Periph) ODD() I2SPR_Mask {
	return I2SPR_Mask{mmio.UM16{&p.I2SPR.U16, uint16(ODD)}}
}

func (p *SPI_Periph) MCKOE() I2SPR_Mask {
	return I2SPR_Mask{mmio.UM16{&p.I2SPR.U16, uint16(MCKOE)}}
}
