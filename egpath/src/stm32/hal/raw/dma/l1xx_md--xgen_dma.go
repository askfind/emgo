// +build l1xx_md

package dma

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l1xx_md/mmap"
)

type DMA_Periph struct {
	ISR  ISR
	IFCR IFCR
}

func (p *DMA_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var DMA1 = (*DMA_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_BASE)))

var DMA2 = (*DMA_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_BASE)))

type ISR_Bits uint32

func (b ISR_Bits) Field(mask ISR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ISR_Bits) J(v int) ISR_Bits {
	return ISR_Bits(bits.Make32(v, uint32(mask)))
}

type ISR struct{ mmio.U32 }

func (r *ISR) Bits(mask ISR_Bits) ISR_Bits { return ISR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ISR) StoreBits(mask, b ISR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ISR) SetBits(mask ISR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ISR) ClearBits(mask ISR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ISR) Load() ISR_Bits              { return ISR_Bits(r.U32.Load()) }
func (r *ISR) Store(b ISR_Bits)            { r.U32.Store(uint32(b)) }

type ISR_Mask struct{ mmio.UM32 }

func (rm ISR_Mask) Load() ISR_Bits   { return ISR_Bits(rm.UM32.Load()) }
func (rm ISR_Mask) Store(b ISR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Periph) GIF1() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(GIF1)}}
}

func (p *DMA_Periph) TCIF1() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TCIF1)}}
}

func (p *DMA_Periph) HTIF1() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(HTIF1)}}
}

func (p *DMA_Periph) TEIF1() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TEIF1)}}
}

func (p *DMA_Periph) GIF2() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(GIF2)}}
}

func (p *DMA_Periph) TCIF2() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TCIF2)}}
}

func (p *DMA_Periph) HTIF2() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(HTIF2)}}
}

func (p *DMA_Periph) TEIF2() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TEIF2)}}
}

func (p *DMA_Periph) GIF3() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(GIF3)}}
}

func (p *DMA_Periph) TCIF3() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TCIF3)}}
}

func (p *DMA_Periph) HTIF3() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(HTIF3)}}
}

func (p *DMA_Periph) TEIF3() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TEIF3)}}
}

func (p *DMA_Periph) GIF4() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(GIF4)}}
}

func (p *DMA_Periph) TCIF4() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TCIF4)}}
}

func (p *DMA_Periph) HTIF4() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(HTIF4)}}
}

func (p *DMA_Periph) TEIF4() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TEIF4)}}
}

func (p *DMA_Periph) GIF5() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(GIF5)}}
}

func (p *DMA_Periph) TCIF5() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TCIF5)}}
}

func (p *DMA_Periph) HTIF5() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(HTIF5)}}
}

func (p *DMA_Periph) TEIF5() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TEIF5)}}
}

func (p *DMA_Periph) GIF6() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(GIF6)}}
}

func (p *DMA_Periph) TCIF6() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TCIF6)}}
}

func (p *DMA_Periph) HTIF6() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(HTIF6)}}
}

func (p *DMA_Periph) TEIF6() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TEIF6)}}
}

func (p *DMA_Periph) GIF7() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(GIF7)}}
}

func (p *DMA_Periph) TCIF7() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TCIF7)}}
}

func (p *DMA_Periph) HTIF7() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(HTIF7)}}
}

func (p *DMA_Periph) TEIF7() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TEIF7)}}
}

type IFCR_Bits uint32

func (b IFCR_Bits) Field(mask IFCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IFCR_Bits) J(v int) IFCR_Bits {
	return IFCR_Bits(bits.Make32(v, uint32(mask)))
}

type IFCR struct{ mmio.U32 }

func (r *IFCR) Bits(mask IFCR_Bits) IFCR_Bits { return IFCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *IFCR) StoreBits(mask, b IFCR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IFCR) SetBits(mask IFCR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *IFCR) ClearBits(mask IFCR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *IFCR) Load() IFCR_Bits               { return IFCR_Bits(r.U32.Load()) }
func (r *IFCR) Store(b IFCR_Bits)             { r.U32.Store(uint32(b)) }

type IFCR_Mask struct{ mmio.UM32 }

func (rm IFCR_Mask) Load() IFCR_Bits   { return IFCR_Bits(rm.UM32.Load()) }
func (rm IFCR_Mask) Store(b IFCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Periph) CGIF1() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CGIF1)}}
}

func (p *DMA_Periph) CTCIF1() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CTCIF1)}}
}

func (p *DMA_Periph) CHTIF1() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CHTIF1)}}
}

func (p *DMA_Periph) CTEIF1() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CTEIF1)}}
}

func (p *DMA_Periph) CGIF2() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CGIF2)}}
}

func (p *DMA_Periph) CTCIF2() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CTCIF2)}}
}

func (p *DMA_Periph) CHTIF2() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CHTIF2)}}
}

func (p *DMA_Periph) CTEIF2() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CTEIF2)}}
}

func (p *DMA_Periph) CGIF3() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CGIF3)}}
}

func (p *DMA_Periph) CTCIF3() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CTCIF3)}}
}

func (p *DMA_Periph) CHTIF3() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CHTIF3)}}
}

func (p *DMA_Periph) CTEIF3() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CTEIF3)}}
}

func (p *DMA_Periph) CGIF4() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CGIF4)}}
}

func (p *DMA_Periph) CTCIF4() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CTCIF4)}}
}

func (p *DMA_Periph) CHTIF4() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CHTIF4)}}
}

func (p *DMA_Periph) CTEIF4() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CTEIF4)}}
}

func (p *DMA_Periph) CGIF5() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CGIF5)}}
}

func (p *DMA_Periph) CTCIF5() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CTCIF5)}}
}

func (p *DMA_Periph) CHTIF5() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CHTIF5)}}
}

func (p *DMA_Periph) CTEIF5() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CTEIF5)}}
}

func (p *DMA_Periph) CGIF6() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CGIF6)}}
}

func (p *DMA_Periph) CTCIF6() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CTCIF6)}}
}

func (p *DMA_Periph) CHTIF6() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CHTIF6)}}
}

func (p *DMA_Periph) CTEIF6() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CTEIF6)}}
}

func (p *DMA_Periph) CGIF7() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CGIF7)}}
}

func (p *DMA_Periph) CTCIF7() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CTCIF7)}}
}

func (p *DMA_Periph) CHTIF7() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CHTIF7)}}
}

func (p *DMA_Periph) CTEIF7() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{&p.IFCR.U32, uint32(CTEIF7)}}
}
