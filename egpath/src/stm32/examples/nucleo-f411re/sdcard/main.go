package main

import (
	"delay"
	"fmt"
	"rtos"

	"sdcard"

	"stm32/hal/dma"
	"stm32/hal/gpio"
	"stm32/hal/irq"
	"stm32/hal/sdmmc"
	"stm32/hal/system"
	"stm32/hal/system/timer/systick"
)

var sd *sdmmc.Driver

func init() {
	system.Setup96(8) // Setups USB/SDIO/RNG clock to 48 MHz
	systick.Setup(2e6)

	// GPIO

	gpio.A.EnableClock(true)
	// irq := gpio.A.Pin(0)
	cmd := gpio.A.Pin(6)
	//d1 := gpio.A.Pin(8)
	//d2 := gpio.A.Pin(9)

	gpio.B.EnableClock(true)
	//d3 := gpio.B.Pin(5)
	d0 := gpio.B.Pin(7)
	clk := gpio.B.Pin(15)

	cfg := gpio.Config{Mode: gpio.Alt, Speed: gpio.High, Pull: gpio.PullUp}
	for _, pin := range []gpio.Pin{clk, cmd, d0 /*d1, d2, d3*/} {
		pin.Setup(&cfg)
		pin.SetAltFunc(gpio.SDIO)
	}

	d := dma.DMA2
	d.EnableClock(true)
	sd = sdmmc.NewDriver(sdmmc.SDIO, d.Channel(6, 4))
	sd.Periph().EnableClock(true)
	sd.Periph().Enable()

	rtos.IRQ(irq.SDIO).Enable()
}

func main() {
	delay.Millisec(200) // For SWO output

	ocr := sdcard.V31 | sdcard.V32 | sdcard.V33 | sdcard.HCXC
	v2 := true

	// Set SDIO_CK to no more than 400 kHz (max. open-drain freq). Clock must be
	// continuously enabled (pwrsave = false) to allow correct initialisation.
	sd.SetFreq(400e3, false)

	// SD card power-up takes maximum of 1 ms or 74 SDIO_CK cycles.
	delay.Millisec(1)

	// Reset.
	sd.Cmd(sdcard.CMD0())
	checkErr("CMD0", sd.Err(true), 0)

	// CMD0 may require up to 8 SDIO_CK cycles to reset the card.
	delay.Millisec(1)

	// Verify card interface operating condition.
	vhs, pattern := sd.Cmd(sdcard.CMD8(sdcard.V27_36, 0xAC)).R7()
	if err := sd.Err(true); err != nil {
		if err == sdcard.ErrCmdTimeout {
			ocr &^= sdcard.HCXC
			v2 = false
		} else {
			checkErr("CMD8", err, 0)
		}
	} else if vhs != sdcard.V27_36 || pattern != 0xAC {
		fmt.Printf("CMD8 bad response: %x, %x\n", vhs, pattern)
		for {
		}
	}
	fmt.Printf("\nPhysical layer version 2.00+: %t\n", v2)

	fmt.Printf("Initializing SD card ")
	var oca sdcard.OCR
	for i := 0; oca&sdcard.PWUP == 0 && i < 20; i++ {
		sd.Cmd(sdcard.CMD55(0))
		oca = sd.Cmd(sdcard.ACMD41(ocr)).R3()
		checkErr("ACMD41", sd.Err(true), 0)
		fmt.Printf(".")
		delay.Millisec(50)
	}
	if oca&sdcard.PWUP == 0 {
		fmt.Printf(" timeout\n")
		for {
		}
	}
	fmt.Printf(" OK\n\n")
	fmt.Printf("Operation Conditions Register: 0x%08X\n\n", oca)

	// Read Card Identification Register.
	cid := sd.Cmd(sdcard.CMD2()).R2CID()
	checkErr("CMD2", sd.Err(true), 0)

	printCID(cid)

	// Generate new Relative Card Address.
	rca, _ := sd.Cmd(sdcard.CMD3()).R6()
	checkErr("CMD3", sd.Err(true), 0)

	fmt.Printf("Relative Card Address: 0x%04X\n\n", rca)

	// After CMD3 card is in Data Transfer Mode (Standby State) and SDIO_CK can
	// be set to no more than 25 MHz (max. push-pull freq). Clock power save
	// mode can be enabled.
	sd.SetFreq(25e6, true)

	// Read Card Specific Data register.
	csd := sd.Cmd(sdcard.CMD9(rca)).R2CSD()
	checkErr("CMD9", sd.Err(true), 0)

	printCSD(csd)

	// Select card (put into Transfer State).
	cs := sd.Cmd(sdcard.CMD7(rca)).R1()
	checkErr("CMD7", sd.Err(true), cs)

	// Disable 50k pull-up resistor on D3/CD.
	sd.Cmd(sdcard.CMD55(rca))
	cs = sd.Cmd(sdcard.ACMD42(false)).R1()
	checkErr("ACMD42", sd.Err(true), cs)

	if ocr&sdcard.HCXC == 0 {
		// Set block size to 512 B for version < 2 or SDSC card.
		cs = sd.Cmd(sdcard.CMD16(512)).R1()
		checkErr("CMD16", sd.Err(true), cs)
	}

	buf := make(sdcard.Data, 512/8)

	for n := uint(0); n < 16; n++ {
		addr := n
		if oca&sdcard.HCXC == 0 {
			addr *= 512
		}
		sd.Data(sdcard.Recv|sdcard.Block512, buf)
		cs = sd.Cmd(sdcard.CMD17(addr)).R1()
		checkErr("CMD17", sd.Err(true), cs)
		fmt.Printf("%d:\n", n)
		s := buf.Bytes()
		for i := 0; i < len(s); i += 16 {
			fmt.Printf("%02x\n", s[i:i+16])
			delay.Millisec(50) // ST-LINK V2-1 SWO is too slow.
		}
	}
}

func sdioISR() {
	sd.ISR()
}

//c:__attribute__((section(".ISRs")))
var ISRs = [...]func(){
	irq.SDIO: sdioISR,
}
