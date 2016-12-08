// +build f746xx

package pwr

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f746xx/mmap"
)

type PWR_Periph struct {
	CR1  CR1
	CSR1 CSR1
	CR2  CR2
	CSR2 CSR2
}

func (p *PWR_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var PWR = (*PWR_Periph)(unsafe.Pointer(uintptr(mmap.PWR_BASE)))

type CR1_Bits uint32

func (b CR1_Bits) Field(mask CR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR1_Bits) J(v int) CR1_Bits {
	return CR1_Bits(bits.Make32(v, uint32(mask)))
}

type CR1 struct{ mmio.U32 }

func (r *CR1) Bits(mask CR1_Bits) CR1_Bits { return CR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR1) StoreBits(mask, b CR1_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR1) SetBits(mask CR1_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CR1) ClearBits(mask CR1_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CR1) Load() CR1_Bits              { return CR1_Bits(r.U32.Load()) }
func (r *CR1) Store(b CR1_Bits)            { r.U32.Store(uint32(b)) }

type CR1_Mask struct{ mmio.UM32 }

func (rm CR1_Mask) Load() CR1_Bits   { return CR1_Bits(rm.UM32.Load()) }
func (rm CR1_Mask) Store(b CR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *PWR_Periph) LPDS() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(LPDS)}}
}

func (p *PWR_Periph) PDDS() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(PDDS)}}
}

func (p *PWR_Periph) CSBF() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(CSBF)}}
}

func (p *PWR_Periph) PVDE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(PVDE)}}
}

func (p *PWR_Periph) PLS() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(PLS)}}
}

func (p *PWR_Periph) DBP() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(DBP)}}
}

func (p *PWR_Periph) FPDS() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(FPDS)}}
}

func (p *PWR_Periph) LPUDS() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(LPUDS)}}
}

func (p *PWR_Periph) MRUDS() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(MRUDS)}}
}

func (p *PWR_Periph) ADCDC1() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(ADCDC1)}}
}

func (p *PWR_Periph) VOS() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(VOS)}}
}

func (p *PWR_Periph) ODEN() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(ODEN)}}
}

func (p *PWR_Periph) ODSWEN() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(ODSWEN)}}
}

func (p *PWR_Periph) UDEN() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(UDEN)}}
}

type CSR1_Bits uint32

func (b CSR1_Bits) Field(mask CSR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSR1_Bits) J(v int) CSR1_Bits {
	return CSR1_Bits(bits.Make32(v, uint32(mask)))
}

type CSR1 struct{ mmio.U32 }

func (r *CSR1) Bits(mask CSR1_Bits) CSR1_Bits { return CSR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *CSR1) StoreBits(mask, b CSR1_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSR1) SetBits(mask CSR1_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *CSR1) ClearBits(mask CSR1_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *CSR1) Load() CSR1_Bits               { return CSR1_Bits(r.U32.Load()) }
func (r *CSR1) Store(b CSR1_Bits)             { r.U32.Store(uint32(b)) }

type CSR1_Mask struct{ mmio.UM32 }

func (rm CSR1_Mask) Load() CSR1_Bits   { return CSR1_Bits(rm.UM32.Load()) }
func (rm CSR1_Mask) Store(b CSR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *PWR_Periph) WUIF() CSR1_Mask {
	return CSR1_Mask{mmio.UM32{&p.CSR1.U32, uint32(WUIF)}}
}

func (p *PWR_Periph) SBF() CSR1_Mask {
	return CSR1_Mask{mmio.UM32{&p.CSR1.U32, uint32(SBF)}}
}

func (p *PWR_Periph) PVDO() CSR1_Mask {
	return CSR1_Mask{mmio.UM32{&p.CSR1.U32, uint32(PVDO)}}
}

func (p *PWR_Periph) BRR() CSR1_Mask {
	return CSR1_Mask{mmio.UM32{&p.CSR1.U32, uint32(BRR)}}
}

func (p *PWR_Periph) EIWUP() CSR1_Mask {
	return CSR1_Mask{mmio.UM32{&p.CSR1.U32, uint32(EIWUP)}}
}

func (p *PWR_Periph) BRE() CSR1_Mask {
	return CSR1_Mask{mmio.UM32{&p.CSR1.U32, uint32(BRE)}}
}

func (p *PWR_Periph) VOSRDY() CSR1_Mask {
	return CSR1_Mask{mmio.UM32{&p.CSR1.U32, uint32(VOSRDY)}}
}

func (p *PWR_Periph) ODRDY() CSR1_Mask {
	return CSR1_Mask{mmio.UM32{&p.CSR1.U32, uint32(ODRDY)}}
}

func (p *PWR_Periph) ODSWRDY() CSR1_Mask {
	return CSR1_Mask{mmio.UM32{&p.CSR1.U32, uint32(ODSWRDY)}}
}

func (p *PWR_Periph) UDRDY() CSR1_Mask {
	return CSR1_Mask{mmio.UM32{&p.CSR1.U32, uint32(UDRDY)}}
}

type CR2_Bits uint32

func (b CR2_Bits) Field(mask CR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR2_Bits) J(v int) CR2_Bits {
	return CR2_Bits(bits.Make32(v, uint32(mask)))
}

type CR2 struct{ mmio.U32 }

func (r *CR2) Bits(mask CR2_Bits) CR2_Bits { return CR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR2) StoreBits(mask, b CR2_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR2) SetBits(mask CR2_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CR2) ClearBits(mask CR2_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CR2) Load() CR2_Bits              { return CR2_Bits(r.U32.Load()) }
func (r *CR2) Store(b CR2_Bits)            { r.U32.Store(uint32(b)) }

type CR2_Mask struct{ mmio.UM32 }

func (rm CR2_Mask) Load() CR2_Bits   { return CR2_Bits(rm.UM32.Load()) }
func (rm CR2_Mask) Store(b CR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *PWR_Periph) CWUPF1() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(CWUPF1)}}
}

func (p *PWR_Periph) CWUPF2() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(CWUPF2)}}
}

func (p *PWR_Periph) CWUPF3() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(CWUPF3)}}
}

func (p *PWR_Periph) CWUPF4() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(CWUPF4)}}
}

func (p *PWR_Periph) CWUPF5() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(CWUPF5)}}
}

func (p *PWR_Periph) CWUPF6() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(CWUPF6)}}
}

func (p *PWR_Periph) WUPP1() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(WUPP1)}}
}

func (p *PWR_Periph) WUPP2() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(WUPP2)}}
}

func (p *PWR_Periph) WUPP3() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(WUPP3)}}
}

func (p *PWR_Periph) WUPP4() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(WUPP4)}}
}

func (p *PWR_Periph) WUPP5() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(WUPP5)}}
}

func (p *PWR_Periph) WUPP6() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(WUPP6)}}
}

type CSR2_Bits uint32

func (b CSR2_Bits) Field(mask CSR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSR2_Bits) J(v int) CSR2_Bits {
	return CSR2_Bits(bits.Make32(v, uint32(mask)))
}

type CSR2 struct{ mmio.U32 }

func (r *CSR2) Bits(mask CSR2_Bits) CSR2_Bits { return CSR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *CSR2) StoreBits(mask, b CSR2_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSR2) SetBits(mask CSR2_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *CSR2) ClearBits(mask CSR2_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *CSR2) Load() CSR2_Bits               { return CSR2_Bits(r.U32.Load()) }
func (r *CSR2) Store(b CSR2_Bits)             { r.U32.Store(uint32(b)) }

type CSR2_Mask struct{ mmio.UM32 }

func (rm CSR2_Mask) Load() CSR2_Bits   { return CSR2_Bits(rm.UM32.Load()) }
func (rm CSR2_Mask) Store(b CSR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *PWR_Periph) WUPF1() CSR2_Mask {
	return CSR2_Mask{mmio.UM32{&p.CSR2.U32, uint32(WUPF1)}}
}

func (p *PWR_Periph) WUPF2() CSR2_Mask {
	return CSR2_Mask{mmio.UM32{&p.CSR2.U32, uint32(WUPF2)}}
}

func (p *PWR_Periph) WUPF3() CSR2_Mask {
	return CSR2_Mask{mmio.UM32{&p.CSR2.U32, uint32(WUPF3)}}
}

func (p *PWR_Periph) WUPF4() CSR2_Mask {
	return CSR2_Mask{mmio.UM32{&p.CSR2.U32, uint32(WUPF4)}}
}

func (p *PWR_Periph) WUPF5() CSR2_Mask {
	return CSR2_Mask{mmio.UM32{&p.CSR2.U32, uint32(WUPF5)}}
}

func (p *PWR_Periph) WUPF6() CSR2_Mask {
	return CSR2_Mask{mmio.UM32{&p.CSR2.U32, uint32(WUPF6)}}
}

func (p *PWR_Periph) EWUP1() CSR2_Mask {
	return CSR2_Mask{mmio.UM32{&p.CSR2.U32, uint32(EWUP1)}}
}

func (p *PWR_Periph) EWUP2() CSR2_Mask {
	return CSR2_Mask{mmio.UM32{&p.CSR2.U32, uint32(EWUP2)}}
}

func (p *PWR_Periph) EWUP3() CSR2_Mask {
	return CSR2_Mask{mmio.UM32{&p.CSR2.U32, uint32(EWUP3)}}
}

func (p *PWR_Periph) EWUP4() CSR2_Mask {
	return CSR2_Mask{mmio.UM32{&p.CSR2.U32, uint32(EWUP4)}}
}

func (p *PWR_Periph) EWUP5() CSR2_Mask {
	return CSR2_Mask{mmio.UM32{&p.CSR2.U32, uint32(EWUP5)}}
}

func (p *PWR_Periph) EWUP6() CSR2_Mask {
	return CSR2_Mask{mmio.UM32{&p.CSR2.U32, uint32(EWUP6)}}
}
