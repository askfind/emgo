// +build f746xx

package usb

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"mmio"
	"unsafe"

	"stm32/o/f746xx/mmap"
)

type USB_OTG_Device_Periph struct {
	DCFG       DCFG
	DCTL       DCTL
	DSTS       DSTS
	_          uint32
	DIEPMSK    DIEPMSK
	DOEPMSK    DOEPMSK
	DAINT      DAINT
	DAINTMSK   DAINTMSK
	_          [2]uint32
	DVBUSDIS   DVBUSDIS
	DVBUSPULSE DVBUSPULSE
	DTHRCTL    DTHRCTL
	DIEPEMPMSK DIEPEMPMSK
	DEACHINT   DEACHINT
	DEACHMSK   DEACHMSK
	_          uint32
	DINEP1MSK  DINEP1MSK
	_          [15]uint32
	DOUTEP1MSK DOUTEP1MSK
}

func (p *USB_OTG_Device_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type DCFG_Bits uint32

type DCFG struct{ mmio.U32 }

func (r *DCFG) Bits(mask DCFG_Bits) DCFG_Bits { return DCFG_Bits(r.U32.Bits(uint32(mask))) }
func (r *DCFG) StoreBits(mask, b DCFG_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DCFG) SetBits(mask DCFG_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *DCFG) ClearBits(mask DCFG_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *DCFG) Load() DCFG_Bits               { return DCFG_Bits(r.U32.Load()) }
func (r *DCFG) Store(b DCFG_Bits)             { r.U32.Store(uint32(b)) }

type DCFG_Mask struct{ mmio.UM32 }

func (rm DCFG_Mask) Load() DCFG_Bits   { return DCFG_Bits(rm.UM32.Load()) }
func (rm DCFG_Mask) Store(b DCFG_Bits) { rm.UM32.Store(uint32(b)) }

type DCTL_Bits uint32

type DCTL struct{ mmio.U32 }

func (r *DCTL) Bits(mask DCTL_Bits) DCTL_Bits { return DCTL_Bits(r.U32.Bits(uint32(mask))) }
func (r *DCTL) StoreBits(mask, b DCTL_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DCTL) SetBits(mask DCTL_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *DCTL) ClearBits(mask DCTL_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *DCTL) Load() DCTL_Bits               { return DCTL_Bits(r.U32.Load()) }
func (r *DCTL) Store(b DCTL_Bits)             { r.U32.Store(uint32(b)) }

type DCTL_Mask struct{ mmio.UM32 }

func (rm DCTL_Mask) Load() DCTL_Bits   { return DCTL_Bits(rm.UM32.Load()) }
func (rm DCTL_Mask) Store(b DCTL_Bits) { rm.UM32.Store(uint32(b)) }

type DSTS_Bits uint32

type DSTS struct{ mmio.U32 }

func (r *DSTS) Bits(mask DSTS_Bits) DSTS_Bits { return DSTS_Bits(r.U32.Bits(uint32(mask))) }
func (r *DSTS) StoreBits(mask, b DSTS_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DSTS) SetBits(mask DSTS_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *DSTS) ClearBits(mask DSTS_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *DSTS) Load() DSTS_Bits               { return DSTS_Bits(r.U32.Load()) }
func (r *DSTS) Store(b DSTS_Bits)             { r.U32.Store(uint32(b)) }

type DSTS_Mask struct{ mmio.UM32 }

func (rm DSTS_Mask) Load() DSTS_Bits   { return DSTS_Bits(rm.UM32.Load()) }
func (rm DSTS_Mask) Store(b DSTS_Bits) { rm.UM32.Store(uint32(b)) }

type DIEPMSK_Bits uint32

type DIEPMSK struct{ mmio.U32 }

func (r *DIEPMSK) Bits(mask DIEPMSK_Bits) DIEPMSK_Bits { return DIEPMSK_Bits(r.U32.Bits(uint32(mask))) }
func (r *DIEPMSK) StoreBits(mask, b DIEPMSK_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DIEPMSK) SetBits(mask DIEPMSK_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *DIEPMSK) ClearBits(mask DIEPMSK_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *DIEPMSK) Load() DIEPMSK_Bits                  { return DIEPMSK_Bits(r.U32.Load()) }
func (r *DIEPMSK) Store(b DIEPMSK_Bits)                { r.U32.Store(uint32(b)) }

type DIEPMSK_Mask struct{ mmio.UM32 }

func (rm DIEPMSK_Mask) Load() DIEPMSK_Bits   { return DIEPMSK_Bits(rm.UM32.Load()) }
func (rm DIEPMSK_Mask) Store(b DIEPMSK_Bits) { rm.UM32.Store(uint32(b)) }

type DOEPMSK_Bits uint32

type DOEPMSK struct{ mmio.U32 }

func (r *DOEPMSK) Bits(mask DOEPMSK_Bits) DOEPMSK_Bits { return DOEPMSK_Bits(r.U32.Bits(uint32(mask))) }
func (r *DOEPMSK) StoreBits(mask, b DOEPMSK_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DOEPMSK) SetBits(mask DOEPMSK_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *DOEPMSK) ClearBits(mask DOEPMSK_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *DOEPMSK) Load() DOEPMSK_Bits                  { return DOEPMSK_Bits(r.U32.Load()) }
func (r *DOEPMSK) Store(b DOEPMSK_Bits)                { r.U32.Store(uint32(b)) }

type DOEPMSK_Mask struct{ mmio.UM32 }

func (rm DOEPMSK_Mask) Load() DOEPMSK_Bits   { return DOEPMSK_Bits(rm.UM32.Load()) }
func (rm DOEPMSK_Mask) Store(b DOEPMSK_Bits) { rm.UM32.Store(uint32(b)) }

type DAINT_Bits uint32

type DAINT struct{ mmio.U32 }

func (r *DAINT) Bits(mask DAINT_Bits) DAINT_Bits { return DAINT_Bits(r.U32.Bits(uint32(mask))) }
func (r *DAINT) StoreBits(mask, b DAINT_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DAINT) SetBits(mask DAINT_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *DAINT) ClearBits(mask DAINT_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *DAINT) Load() DAINT_Bits                { return DAINT_Bits(r.U32.Load()) }
func (r *DAINT) Store(b DAINT_Bits)              { r.U32.Store(uint32(b)) }

type DAINT_Mask struct{ mmio.UM32 }

func (rm DAINT_Mask) Load() DAINT_Bits   { return DAINT_Bits(rm.UM32.Load()) }
func (rm DAINT_Mask) Store(b DAINT_Bits) { rm.UM32.Store(uint32(b)) }

type DAINTMSK_Bits uint32

type DAINTMSK struct{ mmio.U32 }

func (r *DAINTMSK) Bits(mask DAINTMSK_Bits) DAINTMSK_Bits {
	return DAINTMSK_Bits(r.U32.Bits(uint32(mask)))
}
func (r *DAINTMSK) StoreBits(mask, b DAINTMSK_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DAINTMSK) SetBits(mask DAINTMSK_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DAINTMSK) ClearBits(mask DAINTMSK_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DAINTMSK) Load() DAINTMSK_Bits             { return DAINTMSK_Bits(r.U32.Load()) }
func (r *DAINTMSK) Store(b DAINTMSK_Bits)           { r.U32.Store(uint32(b)) }

type DAINTMSK_Mask struct{ mmio.UM32 }

func (rm DAINTMSK_Mask) Load() DAINTMSK_Bits   { return DAINTMSK_Bits(rm.UM32.Load()) }
func (rm DAINTMSK_Mask) Store(b DAINTMSK_Bits) { rm.UM32.Store(uint32(b)) }

type DVBUSDIS_Bits uint32

type DVBUSDIS struct{ mmio.U32 }

func (r *DVBUSDIS) Bits(mask DVBUSDIS_Bits) DVBUSDIS_Bits {
	return DVBUSDIS_Bits(r.U32.Bits(uint32(mask)))
}
func (r *DVBUSDIS) StoreBits(mask, b DVBUSDIS_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DVBUSDIS) SetBits(mask DVBUSDIS_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DVBUSDIS) ClearBits(mask DVBUSDIS_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DVBUSDIS) Load() DVBUSDIS_Bits             { return DVBUSDIS_Bits(r.U32.Load()) }
func (r *DVBUSDIS) Store(b DVBUSDIS_Bits)           { r.U32.Store(uint32(b)) }

type DVBUSDIS_Mask struct{ mmio.UM32 }

func (rm DVBUSDIS_Mask) Load() DVBUSDIS_Bits   { return DVBUSDIS_Bits(rm.UM32.Load()) }
func (rm DVBUSDIS_Mask) Store(b DVBUSDIS_Bits) { rm.UM32.Store(uint32(b)) }

type DVBUSPULSE_Bits uint32

type DVBUSPULSE struct{ mmio.U32 }

func (r *DVBUSPULSE) Bits(mask DVBUSPULSE_Bits) DVBUSPULSE_Bits {
	return DVBUSPULSE_Bits(r.U32.Bits(uint32(mask)))
}
func (r *DVBUSPULSE) StoreBits(mask, b DVBUSPULSE_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DVBUSPULSE) SetBits(mask DVBUSPULSE_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DVBUSPULSE) ClearBits(mask DVBUSPULSE_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DVBUSPULSE) Load() DVBUSPULSE_Bits             { return DVBUSPULSE_Bits(r.U32.Load()) }
func (r *DVBUSPULSE) Store(b DVBUSPULSE_Bits)           { r.U32.Store(uint32(b)) }

type DVBUSPULSE_Mask struct{ mmio.UM32 }

func (rm DVBUSPULSE_Mask) Load() DVBUSPULSE_Bits   { return DVBUSPULSE_Bits(rm.UM32.Load()) }
func (rm DVBUSPULSE_Mask) Store(b DVBUSPULSE_Bits) { rm.UM32.Store(uint32(b)) }

type DTHRCTL_Bits uint32

type DTHRCTL struct{ mmio.U32 }

func (r *DTHRCTL) Bits(mask DTHRCTL_Bits) DTHRCTL_Bits { return DTHRCTL_Bits(r.U32.Bits(uint32(mask))) }
func (r *DTHRCTL) StoreBits(mask, b DTHRCTL_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DTHRCTL) SetBits(mask DTHRCTL_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *DTHRCTL) ClearBits(mask DTHRCTL_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *DTHRCTL) Load() DTHRCTL_Bits                  { return DTHRCTL_Bits(r.U32.Load()) }
func (r *DTHRCTL) Store(b DTHRCTL_Bits)                { r.U32.Store(uint32(b)) }

type DTHRCTL_Mask struct{ mmio.UM32 }

func (rm DTHRCTL_Mask) Load() DTHRCTL_Bits   { return DTHRCTL_Bits(rm.UM32.Load()) }
func (rm DTHRCTL_Mask) Store(b DTHRCTL_Bits) { rm.UM32.Store(uint32(b)) }

type DIEPEMPMSK_Bits uint32

type DIEPEMPMSK struct{ mmio.U32 }

func (r *DIEPEMPMSK) Bits(mask DIEPEMPMSK_Bits) DIEPEMPMSK_Bits {
	return DIEPEMPMSK_Bits(r.U32.Bits(uint32(mask)))
}
func (r *DIEPEMPMSK) StoreBits(mask, b DIEPEMPMSK_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DIEPEMPMSK) SetBits(mask DIEPEMPMSK_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DIEPEMPMSK) ClearBits(mask DIEPEMPMSK_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DIEPEMPMSK) Load() DIEPEMPMSK_Bits             { return DIEPEMPMSK_Bits(r.U32.Load()) }
func (r *DIEPEMPMSK) Store(b DIEPEMPMSK_Bits)           { r.U32.Store(uint32(b)) }

type DIEPEMPMSK_Mask struct{ mmio.UM32 }

func (rm DIEPEMPMSK_Mask) Load() DIEPEMPMSK_Bits   { return DIEPEMPMSK_Bits(rm.UM32.Load()) }
func (rm DIEPEMPMSK_Mask) Store(b DIEPEMPMSK_Bits) { rm.UM32.Store(uint32(b)) }

type DEACHINT_Bits uint32

type DEACHINT struct{ mmio.U32 }

func (r *DEACHINT) Bits(mask DEACHINT_Bits) DEACHINT_Bits {
	return DEACHINT_Bits(r.U32.Bits(uint32(mask)))
}
func (r *DEACHINT) StoreBits(mask, b DEACHINT_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DEACHINT) SetBits(mask DEACHINT_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DEACHINT) ClearBits(mask DEACHINT_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DEACHINT) Load() DEACHINT_Bits             { return DEACHINT_Bits(r.U32.Load()) }
func (r *DEACHINT) Store(b DEACHINT_Bits)           { r.U32.Store(uint32(b)) }

type DEACHINT_Mask struct{ mmio.UM32 }

func (rm DEACHINT_Mask) Load() DEACHINT_Bits   { return DEACHINT_Bits(rm.UM32.Load()) }
func (rm DEACHINT_Mask) Store(b DEACHINT_Bits) { rm.UM32.Store(uint32(b)) }

type DEACHMSK_Bits uint32

type DEACHMSK struct{ mmio.U32 }

func (r *DEACHMSK) Bits(mask DEACHMSK_Bits) DEACHMSK_Bits {
	return DEACHMSK_Bits(r.U32.Bits(uint32(mask)))
}
func (r *DEACHMSK) StoreBits(mask, b DEACHMSK_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DEACHMSK) SetBits(mask DEACHMSK_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DEACHMSK) ClearBits(mask DEACHMSK_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DEACHMSK) Load() DEACHMSK_Bits             { return DEACHMSK_Bits(r.U32.Load()) }
func (r *DEACHMSK) Store(b DEACHMSK_Bits)           { r.U32.Store(uint32(b)) }

type DEACHMSK_Mask struct{ mmio.UM32 }

func (rm DEACHMSK_Mask) Load() DEACHMSK_Bits   { return DEACHMSK_Bits(rm.UM32.Load()) }
func (rm DEACHMSK_Mask) Store(b DEACHMSK_Bits) { rm.UM32.Store(uint32(b)) }

type DINEP1MSK_Bits uint32

type DINEP1MSK struct{ mmio.U32 }

func (r *DINEP1MSK) Bits(mask DINEP1MSK_Bits) DINEP1MSK_Bits {
	return DINEP1MSK_Bits(r.U32.Bits(uint32(mask)))
}
func (r *DINEP1MSK) StoreBits(mask, b DINEP1MSK_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DINEP1MSK) SetBits(mask DINEP1MSK_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DINEP1MSK) ClearBits(mask DINEP1MSK_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DINEP1MSK) Load() DINEP1MSK_Bits             { return DINEP1MSK_Bits(r.U32.Load()) }
func (r *DINEP1MSK) Store(b DINEP1MSK_Bits)           { r.U32.Store(uint32(b)) }

type DINEP1MSK_Mask struct{ mmio.UM32 }

func (rm DINEP1MSK_Mask) Load() DINEP1MSK_Bits   { return DINEP1MSK_Bits(rm.UM32.Load()) }
func (rm DINEP1MSK_Mask) Store(b DINEP1MSK_Bits) { rm.UM32.Store(uint32(b)) }

type DOUTEP1MSK_Bits uint32

type DOUTEP1MSK struct{ mmio.U32 }

func (r *DOUTEP1MSK) Bits(mask DOUTEP1MSK_Bits) DOUTEP1MSK_Bits {
	return DOUTEP1MSK_Bits(r.U32.Bits(uint32(mask)))
}
func (r *DOUTEP1MSK) StoreBits(mask, b DOUTEP1MSK_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DOUTEP1MSK) SetBits(mask DOUTEP1MSK_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DOUTEP1MSK) ClearBits(mask DOUTEP1MSK_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DOUTEP1MSK) Load() DOUTEP1MSK_Bits             { return DOUTEP1MSK_Bits(r.U32.Load()) }
func (r *DOUTEP1MSK) Store(b DOUTEP1MSK_Bits)           { r.U32.Store(uint32(b)) }

type DOUTEP1MSK_Mask struct{ mmio.UM32 }

func (rm DOUTEP1MSK_Mask) Load() DOUTEP1MSK_Bits   { return DOUTEP1MSK_Bits(rm.UM32.Load()) }
func (rm DOUTEP1MSK_Mask) Store(b DOUTEP1MSK_Bits) { rm.UM32.Store(uint32(b)) }
