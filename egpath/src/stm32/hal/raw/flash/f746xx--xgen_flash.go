// +build f746xx

package flash

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f746xx/mmap"
)

type FLASH_Periph struct {
	ACR     ACR
	KEYR    KEYR
	OPTKEYR OPTKEYR
	SR      SR
	CR      CR
	OPTCR   [2]OPTCR
}

func (p *FLASH_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var FLASH = (*FLASH_Periph)(unsafe.Pointer(uintptr(mmap.FLASH_R_BASE)))

type ACR_Bits uint32

func (b ACR_Bits) Field(mask ACR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ACR_Bits) J(v int) ACR_Bits {
	return ACR_Bits(bits.Make32(v, uint32(mask)))
}

type ACR struct{ mmio.U32 }

func (r *ACR) Bits(mask ACR_Bits) ACR_Bits { return ACR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ACR) StoreBits(mask, b ACR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ACR) SetBits(mask ACR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ACR) ClearBits(mask ACR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ACR) Load() ACR_Bits              { return ACR_Bits(r.U32.Load()) }
func (r *ACR) Store(b ACR_Bits)            { r.U32.Store(uint32(b)) }

type ACR_Mask struct{ mmio.UM32 }

func (rm ACR_Mask) Load() ACR_Bits   { return ACR_Bits(rm.UM32.Load()) }
func (rm ACR_Mask) Store(b ACR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) LATENCY() ACR_Mask {
	return ACR_Mask{mmio.UM32{&p.ACR.U32, uint32(LATENCY)}}
}

func (p *FLASH_Periph) PRFTEN() ACR_Mask {
	return ACR_Mask{mmio.UM32{&p.ACR.U32, uint32(PRFTEN)}}
}

func (p *FLASH_Periph) ARTEN() ACR_Mask {
	return ACR_Mask{mmio.UM32{&p.ACR.U32, uint32(ARTEN)}}
}

func (p *FLASH_Periph) ARTRST() ACR_Mask {
	return ACR_Mask{mmio.UM32{&p.ACR.U32, uint32(ARTRST)}}
}

type KEYR_Bits uint32

func (b KEYR_Bits) Field(mask KEYR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask KEYR_Bits) J(v int) KEYR_Bits {
	return KEYR_Bits(bits.Make32(v, uint32(mask)))
}

type KEYR struct{ mmio.U32 }

func (r *KEYR) Bits(mask KEYR_Bits) KEYR_Bits { return KEYR_Bits(r.U32.Bits(uint32(mask))) }
func (r *KEYR) StoreBits(mask, b KEYR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *KEYR) SetBits(mask KEYR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *KEYR) ClearBits(mask KEYR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *KEYR) Load() KEYR_Bits               { return KEYR_Bits(r.U32.Load()) }
func (r *KEYR) Store(b KEYR_Bits)             { r.U32.Store(uint32(b)) }

type KEYR_Mask struct{ mmio.UM32 }

func (rm KEYR_Mask) Load() KEYR_Bits   { return KEYR_Bits(rm.UM32.Load()) }
func (rm KEYR_Mask) Store(b KEYR_Bits) { rm.UM32.Store(uint32(b)) }

type OPTKEYR_Bits uint32

func (b OPTKEYR_Bits) Field(mask OPTKEYR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OPTKEYR_Bits) J(v int) OPTKEYR_Bits {
	return OPTKEYR_Bits(bits.Make32(v, uint32(mask)))
}

type OPTKEYR struct{ mmio.U32 }

func (r *OPTKEYR) Bits(mask OPTKEYR_Bits) OPTKEYR_Bits { return OPTKEYR_Bits(r.U32.Bits(uint32(mask))) }
func (r *OPTKEYR) StoreBits(mask, b OPTKEYR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OPTKEYR) SetBits(mask OPTKEYR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *OPTKEYR) ClearBits(mask OPTKEYR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *OPTKEYR) Load() OPTKEYR_Bits                  { return OPTKEYR_Bits(r.U32.Load()) }
func (r *OPTKEYR) Store(b OPTKEYR_Bits)                { r.U32.Store(uint32(b)) }

type OPTKEYR_Mask struct{ mmio.UM32 }

func (rm OPTKEYR_Mask) Load() OPTKEYR_Bits   { return OPTKEYR_Bits(rm.UM32.Load()) }
func (rm OPTKEYR_Mask) Store(b OPTKEYR_Bits) { rm.UM32.Store(uint32(b)) }

type SR_Bits uint32

func (b SR_Bits) Field(mask SR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR_Bits) J(v int) SR_Bits {
	return SR_Bits(bits.Make32(v, uint32(mask)))
}

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

func (p *FLASH_Periph) EOP() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(EOP)}}
}

func (p *FLASH_Periph) OPERR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(OPERR)}}
}

func (p *FLASH_Periph) WRPERR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(WRPERR)}}
}

func (p *FLASH_Periph) PGAERR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(PGAERR)}}
}

func (p *FLASH_Periph) PGPERR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(PGPERR)}}
}

func (p *FLASH_Periph) ERSERR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(ERSERR)}}
}

func (p *FLASH_Periph) BSY() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(BSY)}}
}

type CR_Bits uint32

func (b CR_Bits) Field(mask CR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR_Bits) J(v int) CR_Bits {
	return CR_Bits(bits.Make32(v, uint32(mask)))
}

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

func (p *FLASH_Periph) PG() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PG)}}
}

func (p *FLASH_Periph) SER() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(SER)}}
}

func (p *FLASH_Periph) MER() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(MER)}}
}

func (p *FLASH_Periph) SNB() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(SNB)}}
}

func (p *FLASH_Periph) PSIZE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PSIZE)}}
}

func (p *FLASH_Periph) STRT() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(STRT)}}
}

func (p *FLASH_Periph) EOPIE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(EOPIE)}}
}

func (p *FLASH_Periph) ERRIE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ERRIE)}}
}

func (p *FLASH_Periph) LOCK() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(LOCK)}}
}

type OPTCR_Bits uint32

func (b OPTCR_Bits) Field(mask OPTCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OPTCR_Bits) J(v int) OPTCR_Bits {
	return OPTCR_Bits(bits.Make32(v, uint32(mask)))
}

type OPTCR struct{ mmio.U32 }

func (r *OPTCR) Bits(mask OPTCR_Bits) OPTCR_Bits { return OPTCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *OPTCR) StoreBits(mask, b OPTCR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OPTCR) SetBits(mask OPTCR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *OPTCR) ClearBits(mask OPTCR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *OPTCR) Load() OPTCR_Bits                { return OPTCR_Bits(r.U32.Load()) }
func (r *OPTCR) Store(b OPTCR_Bits)              { r.U32.Store(uint32(b)) }

type OPTCR_Mask struct{ mmio.UM32 }

func (rm OPTCR_Mask) Load() OPTCR_Bits   { return OPTCR_Bits(rm.UM32.Load()) }
func (rm OPTCR_Mask) Store(b OPTCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) OPTLOCK(n int) OPTCR_Mask {
	return OPTCR_Mask{mmio.UM32{&p.OPTCR[n].U32, uint32(OPTLOCK)}}
}

func (p *FLASH_Periph) OPTSTRT(n int) OPTCR_Mask {
	return OPTCR_Mask{mmio.UM32{&p.OPTCR[n].U32, uint32(OPTSTRT)}}
}

func (p *FLASH_Periph) BOR_LEV(n int) OPTCR_Mask {
	return OPTCR_Mask{mmio.UM32{&p.OPTCR[n].U32, uint32(BOR_LEV)}}
}

func (p *FLASH_Periph) WWDG_SW(n int) OPTCR_Mask {
	return OPTCR_Mask{mmio.UM32{&p.OPTCR[n].U32, uint32(WWDG_SW)}}
}

func (p *FLASH_Periph) IWDG_SW(n int) OPTCR_Mask {
	return OPTCR_Mask{mmio.UM32{&p.OPTCR[n].U32, uint32(IWDG_SW)}}
}

func (p *FLASH_Periph) nRST_STOP(n int) OPTCR_Mask {
	return OPTCR_Mask{mmio.UM32{&p.OPTCR[n].U32, uint32(nRST_STOP)}}
}

func (p *FLASH_Periph) nRST_STDBY(n int) OPTCR_Mask {
	return OPTCR_Mask{mmio.UM32{&p.OPTCR[n].U32, uint32(nRST_STDBY)}}
}

func (p *FLASH_Periph) RDP(n int) OPTCR_Mask {
	return OPTCR_Mask{mmio.UM32{&p.OPTCR[n].U32, uint32(RDP)}}
}

func (p *FLASH_Periph) nWRP(n int) OPTCR_Mask {
	return OPTCR_Mask{mmio.UM32{&p.OPTCR[n].U32, uint32(nWRP)}}
}

func (p *FLASH_Periph) IWDG_STDBY(n int) OPTCR_Mask {
	return OPTCR_Mask{mmio.UM32{&p.OPTCR[n].U32, uint32(IWDG_STDBY)}}
}

func (p *FLASH_Periph) IWDG_STOP(n int) OPTCR_Mask {
	return OPTCR_Mask{mmio.UM32{&p.OPTCR[n].U32, uint32(IWDG_STOP)}}
}
