// +build f746xx

package spdifrx

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"mmio"
	"unsafe"

	"stm32/o/f746xx/mmap"
)

type SPDIFRX_Periph struct {
	CR   CR
	IMR  IMR
	SR   SR
	IFCR IFCR
	DR   DR
	CSR  CSR
	DIR  DIR
}

func (p *SPDIFRX_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var SPDIFRX = (*SPDIFRX_Periph)(unsafe.Pointer(uintptr(mmap.SPDIFRX_BASE)))

type CR_Bits uint32

type CR struct{ mmio.U32 }

func (r *CR) Bits(mask CR_Bits) CR_Bits { return CR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR) StoreBits(mask, b CR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR) SetBits(mask CR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CR) ClearBits(mask CR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CR) Load() CR_Bits             { return CR_Bits(r.U32.Load()) }
func (r *CR) Store(b CR_Bits)           { r.U32.Store(uint32(b)) }

type CR_Mask struct{ mmio.UM32 }

func (rm CR_Mask) Load() CR_Bits   { return CR_Bits(rm.UM32.Load()) }
func (rm CR_Mask) Store(b CR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SPDIFRX_Periph) SPDIFEN() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(SPDIFEN)}}
}

func (p *SPDIFRX_Periph) RXDMAEN() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(RXDMAEN)}}
}

func (p *SPDIFRX_Periph) RXSTEO() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(RXSTEO)}}
}

func (p *SPDIFRX_Periph) DRFMT() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(DRFMT)}}
}

func (p *SPDIFRX_Periph) PMSK() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(PMSK)}}
}

func (p *SPDIFRX_Periph) VMSK() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(VMSK)}}
}

func (p *SPDIFRX_Periph) CUMSK() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(CUMSK)}}
}

func (p *SPDIFRX_Periph) PTMSK() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(PTMSK)}}
}

func (p *SPDIFRX_Periph) CBDMAEN() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(CBDMAEN)}}
}

func (p *SPDIFRX_Periph) CHSEL() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(CHSEL)}}
}

func (p *SPDIFRX_Periph) NBTR() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(NBTR)}}
}

func (p *SPDIFRX_Periph) WFA() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(WFA)}}
}

func (p *SPDIFRX_Periph) INSEL() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(INSEL)}}
}

type IMR_Bits uint32

type IMR struct{ mmio.U32 }

func (r *IMR) Bits(mask IMR_Bits) IMR_Bits { return IMR_Bits(r.U32.Bits(uint32(mask))) }
func (r *IMR) StoreBits(mask, b IMR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IMR) SetBits(mask IMR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *IMR) ClearBits(mask IMR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *IMR) Load() IMR_Bits              { return IMR_Bits(r.U32.Load()) }
func (r *IMR) Store(b IMR_Bits)            { r.U32.Store(uint32(b)) }

type IMR_Mask struct{ mmio.UM32 }

func (rm IMR_Mask) Load() IMR_Bits   { return IMR_Bits(rm.UM32.Load()) }
func (rm IMR_Mask) Store(b IMR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SPDIFRX_Periph) RXNEIE() IMR_Mask {
	return IMR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(RXNEIE)}}
}

func (p *SPDIFRX_Periph) CSRNEIE() IMR_Mask {
	return IMR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(CSRNEIE)}}
}

func (p *SPDIFRX_Periph) PERRIE() IMR_Mask {
	return IMR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(PERRIE)}}
}

func (p *SPDIFRX_Periph) OVRIE() IMR_Mask {
	return IMR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(OVRIE)}}
}

func (p *SPDIFRX_Periph) SBLKIE() IMR_Mask {
	return IMR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(SBLKIE)}}
}

func (p *SPDIFRX_Periph) SYNCDIE() IMR_Mask {
	return IMR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(SYNCDIE)}}
}

func (p *SPDIFRX_Periph) IFEIE() IMR_Mask {
	return IMR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(IFEIE)}}
}

type SR_Bits uint32

type SR struct{ mmio.U32 }

func (r *SR) Bits(mask SR_Bits) SR_Bits { return SR_Bits(r.U32.Bits(uint32(mask))) }
func (r *SR) StoreBits(mask, b SR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SR) SetBits(mask SR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *SR) ClearBits(mask SR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *SR) Load() SR_Bits             { return SR_Bits(r.U32.Load()) }
func (r *SR) Store(b SR_Bits)           { r.U32.Store(uint32(b)) }

type SR_Mask struct{ mmio.UM32 }

func (rm SR_Mask) Load() SR_Bits   { return SR_Bits(rm.UM32.Load()) }
func (rm SR_Mask) Store(b SR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SPDIFRX_Periph) RXNE() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(RXNE)}}
}

func (p *SPDIFRX_Periph) CSRNE() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(CSRNE)}}
}

func (p *SPDIFRX_Periph) PERR() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(PERR)}}
}

func (p *SPDIFRX_Periph) OVR() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(OVR)}}
}

func (p *SPDIFRX_Periph) SBD() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(SBD)}}
}

func (p *SPDIFRX_Periph) SYNCD() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(SYNCD)}}
}

func (p *SPDIFRX_Periph) FERR() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(FERR)}}
}

func (p *SPDIFRX_Periph) SERR() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(SERR)}}
}

func (p *SPDIFRX_Periph) TERR() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(TERR)}}
}

func (p *SPDIFRX_Periph) WIDTH5() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(WIDTH5)}}
}

type IFCR_Bits uint32

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

func (p *SPDIFRX_Periph) PERRCF() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(PERRCF)}}
}

func (p *SPDIFRX_Periph) OVRCF() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(OVRCF)}}
}

func (p *SPDIFRX_Periph) SBDCF() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(SBDCF)}}
}

func (p *SPDIFRX_Periph) SYNCDCF() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(SYNCDCF)}}
}

type DR_Bits uint32

type DR struct{ mmio.U32 }

func (r *DR) Bits(mask DR_Bits) DR_Bits { return DR_Bits(r.U32.Bits(uint32(mask))) }
func (r *DR) StoreBits(mask, b DR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DR) SetBits(mask DR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DR) ClearBits(mask DR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DR) Load() DR_Bits             { return DR_Bits(r.U32.Load()) }
func (r *DR) Store(b DR_Bits)           { r.U32.Store(uint32(b)) }

type DR_Mask struct{ mmio.UM32 }

func (rm DR_Mask) Load() DR_Bits   { return DR_Bits(rm.UM32.Load()) }
func (rm DR_Mask) Store(b DR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SPDIFRX_Periph) DR() DR_Mask {
	return DR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(DR)}}
}

func (p *SPDIFRX_Periph) PT() DR_Mask {
	return DR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(PT)}}
}

func (p *SPDIFRX_Periph) C() DR_Mask {
	return DR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(C)}}
}

func (p *SPDIFRX_Periph) U() DR_Mask {
	return DR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(U)}}
}

func (p *SPDIFRX_Periph) V() DR_Mask {
	return DR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(V)}}
}

func (p *SPDIFRX_Periph) PE() DR_Mask {
	return DR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(PE)}}
}

type CSR_Bits uint32

type CSR struct{ mmio.U32 }

func (r *CSR) Bits(mask CSR_Bits) CSR_Bits { return CSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CSR) StoreBits(mask, b CSR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSR) SetBits(mask CSR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CSR) ClearBits(mask CSR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CSR) Load() CSR_Bits              { return CSR_Bits(r.U32.Load()) }
func (r *CSR) Store(b CSR_Bits)            { r.U32.Store(uint32(b)) }

type CSR_Mask struct{ mmio.UM32 }

func (rm CSR_Mask) Load() CSR_Bits   { return CSR_Bits(rm.UM32.Load()) }
func (rm CSR_Mask) Store(b CSR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SPDIFRX_Periph) USR() CSR_Mask {
	return CSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(USR)}}
}

func (p *SPDIFRX_Periph) CS() CSR_Mask {
	return CSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(CS)}}
}

func (p *SPDIFRX_Periph) SOB() CSR_Mask {
	return CSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(SOB)}}
}

type DIR_Bits uint32

type DIR struct{ mmio.U32 }

func (r *DIR) Bits(mask DIR_Bits) DIR_Bits { return DIR_Bits(r.U32.Bits(uint32(mask))) }
func (r *DIR) StoreBits(mask, b DIR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DIR) SetBits(mask DIR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *DIR) ClearBits(mask DIR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *DIR) Load() DIR_Bits              { return DIR_Bits(r.U32.Load()) }
func (r *DIR) Store(b DIR_Bits)            { r.U32.Store(uint32(b)) }

type DIR_Mask struct{ mmio.UM32 }

func (rm DIR_Mask) Load() DIR_Bits   { return DIR_Bits(rm.UM32.Load()) }
func (rm DIR_Mask) Store(b DIR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SPDIFRX_Periph) THI() DIR_Mask {
	return DIR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(THI)}}
}

func (p *SPDIFRX_Periph) TLO() DIR_Mask {
	return DIR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(TLO)}}
}
