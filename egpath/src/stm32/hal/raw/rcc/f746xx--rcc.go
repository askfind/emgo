// +build f746xx

// Peripheral: RCC_Periph  Reset and Clock Control.
// Instances:
//  RCC  mmap.RCC_BASE
// Registers:
//  0x00 32  CR         Clock control register.
//  0x04 32  PLLCFGR    PLL configuration register.
//  0x08 32  CFGR       Clock configuration register.
//  0x0C 32  CIR        Clock interrupt register.
//  0x10 32  AHB1RSTR   AHB1 peripheral reset register.
//  0x14 32  AHB2RSTR   AHB2 peripheral reset register.
//  0x18 32  AHB3RSTR   AHB3 peripheral reset register.
//  0x20 32  APB1RSTR   APB1 peripheral reset register.
//  0x24 32  APB2RSTR   APB2 peripheral reset register.
//  0x30 32  AHB1ENR    AHB1 peripheral clock register.
//  0x34 32  AHB2ENR    AHB2 peripheral clock register.
//  0x38 32  AHB3ENR    AHB3 peripheral clock register.
//  0x40 32  APB1ENR    APB1 peripheral clock enable register.
//  0x44 32  APB2ENR    APB2 peripheral clock enable register.
//  0x50 32  AHB1LPENR  AHB1 peripheral clock enable in low power mode register.
//  0x54 32  AHB2LPENR  AHB2 peripheral clock enable in low power mode register.
//  0x58 32  AHB3LPENR  AHB3 peripheral clock enable in low power mode register.
//  0x60 32  APB1LPENR  APB1 peripheral clock enable in low power mode register.
//  0x64 32  APB2LPENR  APB2 peripheral clock enable in low power mode register.
//  0x70 32  BDCR       Backup domain control register.
//  0x74 32  CSR        Clock control & status register.
//  0x80 32  SSCGR      Spread spectrum clock generation register.
//  0x84 32  PLLI2SCFGR PLLI2S configuration register.
//  0x88 32  PLLSAICFGR PLLSAI configuration register.
//  0x8C 32  DCKCFGR1   Dedicated Clocks configuration register1.
//  0x90 32  DCKCFGR2   Dedicated Clocks configuration register 2.
// Import:
//  stm32/o/f746xx/mmap
package rcc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	HSION     CR_Bits = 0x01 << 0  //+
	HSIRDY    CR_Bits = 0x01 << 1  //+
	HSITRIM   CR_Bits = 0x1F << 3  //+
	HSITRIM_0 CR_Bits = 0x01 << 3  //  Bit 0.
	HSITRIM_1 CR_Bits = 0x02 << 3  //  Bit 1.
	HSITRIM_2 CR_Bits = 0x04 << 3  //  Bit 2.
	HSITRIM_3 CR_Bits = 0x08 << 3  //  Bit 3.
	HSITRIM_4 CR_Bits = 0x10 << 3  //  Bit 4.
	HSICAL    CR_Bits = 0xFF << 8  //+
	HSICAL_0  CR_Bits = 0x01 << 8  //  Bit 0.
	HSICAL_1  CR_Bits = 0x02 << 8  //  Bit 1.
	HSICAL_2  CR_Bits = 0x04 << 8  //  Bit 2.
	HSICAL_3  CR_Bits = 0x08 << 8  //  Bit 3.
	HSICAL_4  CR_Bits = 0x10 << 8  //  Bit 4.
	HSICAL_5  CR_Bits = 0x20 << 8  //  Bit 5.
	HSICAL_6  CR_Bits = 0x40 << 8  //  Bit 6.
	HSICAL_7  CR_Bits = 0x80 << 8  //  Bit 7.
	HSEON     CR_Bits = 0x01 << 16 //+
	HSERDY    CR_Bits = 0x01 << 17 //+
	HSEBYP    CR_Bits = 0x01 << 18 //+
	CSSON     CR_Bits = 0x01 << 19 //+
	PLLON     CR_Bits = 0x01 << 24 //+
	PLLRDY    CR_Bits = 0x01 << 25 //+
	PLLI2SON  CR_Bits = 0x01 << 26 //+
	PLLI2SRDY CR_Bits = 0x01 << 27 //+
	PLLSAION  CR_Bits = 0x01 << 28 //+
	PLLSAIRDY CR_Bits = 0x01 << 29 //+
)

const (
	HSIONn     = 0
	HSIRDYn    = 1
	HSITRIMn   = 3
	HSICALn    = 8
	HSEONn     = 16
	HSERDYn    = 17
	HSEBYPn    = 18
	CSSONn     = 19
	PLLONn     = 24
	PLLRDYn    = 25
	PLLI2SONn  = 26
	PLLI2SRDYn = 27
	PLLSAIONn  = 28
	PLLSAIRDYn = 29
)

const (
	PLLM       PLLCFGR_Bits = 0x3F << 0 //+
	PLLM_0     PLLCFGR_Bits = 0x01 << 0
	PLLM_1     PLLCFGR_Bits = 0x02 << 0
	PLLM_2     PLLCFGR_Bits = 0x04 << 0
	PLLM_3     PLLCFGR_Bits = 0x08 << 0
	PLLM_4     PLLCFGR_Bits = 0x10 << 0
	PLLM_5     PLLCFGR_Bits = 0x20 << 0
	PLLN       PLLCFGR_Bits = 0x1FF << 6 //+
	PLLN_0     PLLCFGR_Bits = 0x01 << 6
	PLLN_1     PLLCFGR_Bits = 0x02 << 6
	PLLN_2     PLLCFGR_Bits = 0x04 << 6
	PLLN_3     PLLCFGR_Bits = 0x08 << 6
	PLLN_4     PLLCFGR_Bits = 0x10 << 6
	PLLN_5     PLLCFGR_Bits = 0x20 << 6
	PLLN_6     PLLCFGR_Bits = 0x40 << 6
	PLLN_7     PLLCFGR_Bits = 0x80 << 6
	PLLN_8     PLLCFGR_Bits = 0x100 << 6
	PLLP       PLLCFGR_Bits = 0x03 << 16 //+
	PLLP_0     PLLCFGR_Bits = 0x01 << 16
	PLLP_1     PLLCFGR_Bits = 0x02 << 16
	PLLSRC     PLLCFGR_Bits = 0x01 << 22 //+
	PLLSRC_HSE PLLCFGR_Bits = 0x01 << 22
	PLLSRC_HSI PLLCFGR_Bits = 0x00 << 22
	PLLQ       PLLCFGR_Bits = 0x0F << 24 //+
	PLLQ_0     PLLCFGR_Bits = 0x01 << 24
	PLLQ_1     PLLCFGR_Bits = 0x02 << 24
	PLLQ_2     PLLCFGR_Bits = 0x04 << 24
	PLLQ_3     PLLCFGR_Bits = 0x08 << 24
)

const (
	PLLMn   = 0
	PLLNn   = 6
	PLLPn   = 16
	PLLSRCn = 22
	PLLQn   = 24
)

const (
	SW          CFGR_Bits = 0x03 << 0  //+ SW[1:0] bits (System clock Switch).
	SW_0        CFGR_Bits = 0x01 << 0  //  Bit 0.
	SW_1        CFGR_Bits = 0x02 << 0  //  Bit 1.
	SW_HSI      CFGR_Bits = 0x00 << 0  //  HSI selected as system clock.
	SW_HSE      CFGR_Bits = 0x01 << 0  //  HSE selected as system clock.
	SW_PLL      CFGR_Bits = 0x02 << 0  //  PLL selected as system clock.
	SWS         CFGR_Bits = 0x03 << 2  //+ SWS[1:0] bits (System Clock Switch Status).
	SWS_0       CFGR_Bits = 0x01 << 2  //  Bit 0.
	SWS_1       CFGR_Bits = 0x02 << 2  //  Bit 1.
	SWS_HSI     CFGR_Bits = 0x00 << 2  //  HSI oscillator used as system clock.
	SWS_HSE     CFGR_Bits = 0x01 << 2  //  HSE oscillator used as system clock.
	SWS_PLL     CFGR_Bits = 0x02 << 2  //  PLL used as system clock.
	HPRE        CFGR_Bits = 0x0F << 4  //+ HPRE[3:0] bits (AHB prescaler).
	HPRE_0      CFGR_Bits = 0x01 << 4  //  Bit 0.
	HPRE_1      CFGR_Bits = 0x02 << 4  //  Bit 1.
	HPRE_2      CFGR_Bits = 0x04 << 4  //  Bit 2.
	HPRE_3      CFGR_Bits = 0x08 << 4  //  Bit 3.
	HPRE_DIV1   CFGR_Bits = 0x00 << 4  //  SYSCLK not divided.
	HPRE_DIV2   CFGR_Bits = 0x08 << 4  //  SYSCLK divided by 2.
	HPRE_DIV4   CFGR_Bits = 0x09 << 4  //  SYSCLK divided by 4.
	HPRE_DIV8   CFGR_Bits = 0x0A << 4  //  SYSCLK divided by 8.
	HPRE_DIV16  CFGR_Bits = 0x0B << 4  //  SYSCLK divided by 16.
	HPRE_DIV64  CFGR_Bits = 0x0C << 4  //  SYSCLK divided by 64.
	HPRE_DIV128 CFGR_Bits = 0x0D << 4  //  SYSCLK divided by 128.
	HPRE_DIV256 CFGR_Bits = 0x0E << 4  //  SYSCLK divided by 256.
	HPRE_DIV512 CFGR_Bits = 0x0F << 4  //  SYSCLK divided by 512.
	PPRE1       CFGR_Bits = 0x07 << 10 //+ PRE1[2:0] bits (APB1 prescaler).
	PPRE1_0     CFGR_Bits = 0x01 << 10 //  Bit 0.
	PPRE1_1     CFGR_Bits = 0x02 << 10 //  Bit 1.
	PPRE1_2     CFGR_Bits = 0x04 << 10 //  Bit 2.
	PPRE1_DIV1  CFGR_Bits = 0x00 << 10 //  HCLK not divided.
	PPRE1_DIV2  CFGR_Bits = 0x04 << 10 //  HCLK divided by 2.
	PPRE1_DIV4  CFGR_Bits = 0x05 << 10 //  HCLK divided by 4.
	PPRE1_DIV8  CFGR_Bits = 0x06 << 10 //  HCLK divided by 8.
	PPRE1_DIV16 CFGR_Bits = 0x07 << 10 //  HCLK divided by 16.
	PPRE2       CFGR_Bits = 0x07 << 13 //+ PRE2[2:0] bits (APB2 prescaler).
	PPRE2_0     CFGR_Bits = 0x01 << 13 //  Bit 0.
	PPRE2_1     CFGR_Bits = 0x02 << 13 //  Bit 1.
	PPRE2_2     CFGR_Bits = 0x04 << 13 //  Bit 2.
	PPRE2_DIV1  CFGR_Bits = 0x00 << 13 //  HCLK not divided.
	PPRE2_DIV2  CFGR_Bits = 0x04 << 13 //  HCLK divided by 2.
	PPRE2_DIV4  CFGR_Bits = 0x05 << 13 //  HCLK divided by 4.
	PPRE2_DIV8  CFGR_Bits = 0x06 << 13 //  HCLK divided by 8.
	PPRE2_DIV16 CFGR_Bits = 0x07 << 13 //  HCLK divided by 16.
	RTCPRE      CFGR_Bits = 0x1F << 16 //+
	RTCPRE_0    CFGR_Bits = 0x01 << 16
	RTCPRE_1    CFGR_Bits = 0x02 << 16
	RTCPRE_2    CFGR_Bits = 0x04 << 16
	RTCPRE_3    CFGR_Bits = 0x08 << 16
	RTCPRE_4    CFGR_Bits = 0x10 << 16
	MCO1        CFGR_Bits = 0x03 << 21 //+
	MCO1_0      CFGR_Bits = 0x01 << 21
	MCO1_1      CFGR_Bits = 0x02 << 21
	I2SSRC      CFGR_Bits = 0x01 << 23 //+
	MCO1PRE     CFGR_Bits = 0x07 << 24 //+
	MCO1PRE_0   CFGR_Bits = 0x01 << 24
	MCO1PRE_1   CFGR_Bits = 0x02 << 24
	MCO1PRE_2   CFGR_Bits = 0x04 << 24
	MCO2PRE     CFGR_Bits = 0x07 << 27 //+
	MCO2PRE_0   CFGR_Bits = 0x01 << 27
	MCO2PRE_1   CFGR_Bits = 0x02 << 27
	MCO2PRE_2   CFGR_Bits = 0x04 << 27
	MCO2        CFGR_Bits = 0x03 << 30 //+
	MCO2_0      CFGR_Bits = 0x01 << 30
	MCO2_1      CFGR_Bits = 0x02 << 30
)

const (
	SWn      = 0
	SWSn     = 2
	HPREn    = 4
	PPRE1n   = 10
	PPRE2n   = 13
	RTCPREn  = 16
	MCO1n    = 21
	I2SSRCn  = 23
	MCO1PREn = 24
	MCO2PREn = 27
	MCO2n    = 30
)

const (
	LSIRDYF     CIR_Bits = 0x01 << 0  //+
	LSERDYF     CIR_Bits = 0x01 << 1  //+
	HSIRDYF     CIR_Bits = 0x01 << 2  //+
	HSERDYF     CIR_Bits = 0x01 << 3  //+
	PLLRDYF     CIR_Bits = 0x01 << 4  //+
	PLLI2SRDYF  CIR_Bits = 0x01 << 5  //+
	PLLSAIRDYF  CIR_Bits = 0x01 << 6  //+
	CSSF        CIR_Bits = 0x01 << 7  //+
	LSIRDYIE    CIR_Bits = 0x01 << 8  //+
	LSERDYIE    CIR_Bits = 0x01 << 9  //+
	HSIRDYIE    CIR_Bits = 0x01 << 10 //+
	HSERDYIE    CIR_Bits = 0x01 << 11 //+
	PLLRDYIE    CIR_Bits = 0x01 << 12 //+
	PLLI2SRDYIE CIR_Bits = 0x01 << 13 //+
	PLLSAIRDYIE CIR_Bits = 0x01 << 14 //+
	LSIRDYC     CIR_Bits = 0x01 << 16 //+
	LSERDYC     CIR_Bits = 0x01 << 17 //+
	HSIRDYC     CIR_Bits = 0x01 << 18 //+
	HSERDYC     CIR_Bits = 0x01 << 19 //+
	PLLRDYC     CIR_Bits = 0x01 << 20 //+
	PLLI2SRDYC  CIR_Bits = 0x01 << 21 //+
	PLLSAIRDYC  CIR_Bits = 0x01 << 22 //+
	CSSC        CIR_Bits = 0x01 << 23 //+
)

const (
	LSIRDYFn     = 0
	LSERDYFn     = 1
	HSIRDYFn     = 2
	HSERDYFn     = 3
	PLLRDYFn     = 4
	PLLI2SRDYFn  = 5
	PLLSAIRDYFn  = 6
	CSSFn        = 7
	LSIRDYIEn    = 8
	LSERDYIEn    = 9
	HSIRDYIEn    = 10
	HSERDYIEn    = 11
	PLLRDYIEn    = 12
	PLLI2SRDYIEn = 13
	PLLSAIRDYIEn = 14
	LSIRDYCn     = 16
	LSERDYCn     = 17
	HSIRDYCn     = 18
	HSERDYCn     = 19
	PLLRDYCn     = 20
	PLLI2SRDYCn  = 21
	PLLSAIRDYCn  = 22
	CSSCn        = 23
)

const (
	GPIOARST  AHB1RSTR_Bits = 0x01 << 0  //+
	GPIOBRST  AHB1RSTR_Bits = 0x01 << 1  //+
	GPIOCRST  AHB1RSTR_Bits = 0x01 << 2  //+
	GPIODRST  AHB1RSTR_Bits = 0x01 << 3  //+
	GPIOERST  AHB1RSTR_Bits = 0x01 << 4  //+
	GPIOFRST  AHB1RSTR_Bits = 0x01 << 5  //+
	GPIOGRST  AHB1RSTR_Bits = 0x01 << 6  //+
	GPIOHRST  AHB1RSTR_Bits = 0x01 << 7  //+
	GPIOIRST  AHB1RSTR_Bits = 0x01 << 8  //+
	GPIOJRST  AHB1RSTR_Bits = 0x01 << 9  //+
	GPIOKRST  AHB1RSTR_Bits = 0x01 << 10 //+
	CRCRST    AHB1RSTR_Bits = 0x01 << 12 //+
	DMA1RST   AHB1RSTR_Bits = 0x01 << 21 //+
	DMA2RST   AHB1RSTR_Bits = 0x01 << 22 //+
	DMA2DRST  AHB1RSTR_Bits = 0x01 << 23 //+
	ETHMACRST AHB1RSTR_Bits = 0x01 << 25 //+
	OTGHRST   AHB1RSTR_Bits = 0x01 << 29 //+
)

const (
	GPIOARSTn  = 0
	GPIOBRSTn  = 1
	GPIOCRSTn  = 2
	GPIODRSTn  = 3
	GPIOERSTn  = 4
	GPIOFRSTn  = 5
	GPIOGRSTn  = 6
	GPIOHRSTn  = 7
	GPIOIRSTn  = 8
	GPIOJRSTn  = 9
	GPIOKRSTn  = 10
	CRCRSTn    = 12
	DMA1RSTn   = 21
	DMA2RSTn   = 22
	DMA2DRSTn  = 23
	ETHMACRSTn = 25
	OTGHRSTn   = 29
)

const (
	DCMIRST  AHB2RSTR_Bits = 0x01 << 0 //+
	RNGRST   AHB2RSTR_Bits = 0x01 << 6 //+
	OTGFSRST AHB2RSTR_Bits = 0x01 << 7 //+
)

const (
	DCMIRSTn  = 0
	RNGRSTn   = 6
	OTGFSRSTn = 7
)

const (
	FMCRST  AHB3RSTR_Bits = 0x01 << 0 //+
	QSPIRST AHB3RSTR_Bits = 0x01 << 1 //+
)

const (
	FMCRSTn  = 0
	QSPIRSTn = 1
)

const (
	TIM2RST    APB1RSTR_Bits = 0x01 << 0  //+
	TIM3RST    APB1RSTR_Bits = 0x01 << 1  //+
	TIM4RST    APB1RSTR_Bits = 0x01 << 2  //+
	TIM5RST    APB1RSTR_Bits = 0x01 << 3  //+
	TIM6RST    APB1RSTR_Bits = 0x01 << 4  //+
	TIM7RST    APB1RSTR_Bits = 0x01 << 5  //+
	TIM12RST   APB1RSTR_Bits = 0x01 << 6  //+
	TIM13RST   APB1RSTR_Bits = 0x01 << 7  //+
	TIM14RST   APB1RSTR_Bits = 0x01 << 8  //+
	LPTIM1RST  APB1RSTR_Bits = 0x01 << 9  //+
	WWDGRST    APB1RSTR_Bits = 0x01 << 11 //+
	SPI2RST    APB1RSTR_Bits = 0x01 << 14 //+
	SPI3RST    APB1RSTR_Bits = 0x01 << 15 //+
	SPDIFRXRST APB1RSTR_Bits = 0x01 << 16 //+
	USART2RST  APB1RSTR_Bits = 0x01 << 17 //+
	USART3RST  APB1RSTR_Bits = 0x01 << 18 //+
	UART4RST   APB1RSTR_Bits = 0x01 << 19 //+
	UART5RST   APB1RSTR_Bits = 0x01 << 20 //+
	I2C1RST    APB1RSTR_Bits = 0x01 << 21 //+
	I2C2RST    APB1RSTR_Bits = 0x01 << 22 //+
	I2C3RST    APB1RSTR_Bits = 0x01 << 23 //+
	I2C4RST    APB1RSTR_Bits = 0x01 << 24 //+
	CAN1RST    APB1RSTR_Bits = 0x01 << 25 //+
	CAN2RST    APB1RSTR_Bits = 0x01 << 26 //+
	CECRST     APB1RSTR_Bits = 0x01 << 27 //+
	PWRRST     APB1RSTR_Bits = 0x01 << 28 //+
	DACRST     APB1RSTR_Bits = 0x01 << 29 //+
	UART7RST   APB1RSTR_Bits = 0x01 << 30 //+
	UART8RST   APB1RSTR_Bits = 0x01 << 31 //+
)

const (
	TIM2RSTn    = 0
	TIM3RSTn    = 1
	TIM4RSTn    = 2
	TIM5RSTn    = 3
	TIM6RSTn    = 4
	TIM7RSTn    = 5
	TIM12RSTn   = 6
	TIM13RSTn   = 7
	TIM14RSTn   = 8
	LPTIM1RSTn  = 9
	WWDGRSTn    = 11
	SPI2RSTn    = 14
	SPI3RSTn    = 15
	SPDIFRXRSTn = 16
	USART2RSTn  = 17
	USART3RSTn  = 18
	UART4RSTn   = 19
	UART5RSTn   = 20
	I2C1RSTn    = 21
	I2C2RSTn    = 22
	I2C3RSTn    = 23
	I2C4RSTn    = 24
	CAN1RSTn    = 25
	CAN2RSTn    = 26
	CECRSTn     = 27
	PWRRSTn     = 28
	DACRSTn     = 29
	UART7RSTn   = 30
	UART8RSTn   = 31
)

const (
	TIM1RST   APB2RSTR_Bits = 0x01 << 0  //+
	TIM8RST   APB2RSTR_Bits = 0x01 << 1  //+
	USART1RST APB2RSTR_Bits = 0x01 << 4  //+
	USART6RST APB2RSTR_Bits = 0x01 << 5  //+
	ADCRST    APB2RSTR_Bits = 0x01 << 8  //+
	SDMMC1RST APB2RSTR_Bits = 0x01 << 11 //+
	SPI1RST   APB2RSTR_Bits = 0x01 << 12 //+
	SPI4RST   APB2RSTR_Bits = 0x01 << 13 //+
	SYSCFGRST APB2RSTR_Bits = 0x01 << 14 //+
	TIM9RST   APB2RSTR_Bits = 0x01 << 16 //+
	TIM10RST  APB2RSTR_Bits = 0x01 << 17 //+
	TIM11RST  APB2RSTR_Bits = 0x01 << 18 //+
	SPI5RST   APB2RSTR_Bits = 0x01 << 20 //+
	SPI6RST   APB2RSTR_Bits = 0x01 << 21 //+
	SAI1RST   APB2RSTR_Bits = 0x01 << 22 //+
	SAI2RST   APB2RSTR_Bits = 0x01 << 23 //+
	LTDCRST   APB2RSTR_Bits = 0x01 << 26 //+
)

const (
	TIM1RSTn   = 0
	TIM8RSTn   = 1
	USART1RSTn = 4
	USART6RSTn = 5
	ADCRSTn    = 8
	SDMMC1RSTn = 11
	SPI1RSTn   = 12
	SPI4RSTn   = 13
	SYSCFGRSTn = 14
	TIM9RSTn   = 16
	TIM10RSTn  = 17
	TIM11RSTn  = 18
	SPI5RSTn   = 20
	SPI6RSTn   = 21
	SAI1RSTn   = 22
	SAI2RSTn   = 23
	LTDCRSTn   = 26
)

const (
	GPIOAEN     AHB1ENR_Bits = 0x01 << 0  //+
	GPIOBEN     AHB1ENR_Bits = 0x01 << 1  //+
	GPIOCEN     AHB1ENR_Bits = 0x01 << 2  //+
	GPIODEN     AHB1ENR_Bits = 0x01 << 3  //+
	GPIOEEN     AHB1ENR_Bits = 0x01 << 4  //+
	GPIOFEN     AHB1ENR_Bits = 0x01 << 5  //+
	GPIOGEN     AHB1ENR_Bits = 0x01 << 6  //+
	GPIOHEN     AHB1ENR_Bits = 0x01 << 7  //+
	GPIOIEN     AHB1ENR_Bits = 0x01 << 8  //+
	GPIOJEN     AHB1ENR_Bits = 0x01 << 9  //+
	GPIOKEN     AHB1ENR_Bits = 0x01 << 10 //+
	CRCEN       AHB1ENR_Bits = 0x01 << 12 //+
	BKPSRAMEN   AHB1ENR_Bits = 0x01 << 18 //+
	DTCMRAMEN   AHB1ENR_Bits = 0x01 << 20 //+
	DMA1EN      AHB1ENR_Bits = 0x01 << 21 //+
	DMA2EN      AHB1ENR_Bits = 0x01 << 22 //+
	DMA2DEN     AHB1ENR_Bits = 0x01 << 23 //+
	ETHMACEN    AHB1ENR_Bits = 0x01 << 25 //+
	ETHMACTXEN  AHB1ENR_Bits = 0x01 << 26 //+
	ETHMACRXEN  AHB1ENR_Bits = 0x01 << 27 //+
	ETHMACPTPEN AHB1ENR_Bits = 0x01 << 28 //+
	OTGHSEN     AHB1ENR_Bits = 0x01 << 29 //+
	OTGHSULPIEN AHB1ENR_Bits = 0x01 << 30 //+
)

const (
	GPIOAENn     = 0
	GPIOBENn     = 1
	GPIOCENn     = 2
	GPIODENn     = 3
	GPIOEENn     = 4
	GPIOFENn     = 5
	GPIOGENn     = 6
	GPIOHENn     = 7
	GPIOIENn     = 8
	GPIOJENn     = 9
	GPIOKENn     = 10
	CRCENn       = 12
	BKPSRAMENn   = 18
	DTCMRAMENn   = 20
	DMA1ENn      = 21
	DMA2ENn      = 22
	DMA2DENn     = 23
	ETHMACENn    = 25
	ETHMACTXENn  = 26
	ETHMACRXENn  = 27
	ETHMACPTPENn = 28
	OTGHSENn     = 29
	OTGHSULPIENn = 30
)

const (
	DCMIEN  AHB2ENR_Bits = 0x01 << 0 //+
	RNGEN   AHB2ENR_Bits = 0x01 << 6 //+
	OTGFSEN AHB2ENR_Bits = 0x01 << 7 //+
)

const (
	DCMIENn  = 0
	RNGENn   = 6
	OTGFSENn = 7
)

const (
	FMCEN  AHB3ENR_Bits = 0x01 << 0 //+
	QSPIEN AHB3ENR_Bits = 0x01 << 1 //+
)

const (
	FMCENn  = 0
	QSPIENn = 1
)

const (
	TIM2EN    APB1ENR_Bits = 0x01 << 0  //+
	TIM3EN    APB1ENR_Bits = 0x01 << 1  //+
	TIM4EN    APB1ENR_Bits = 0x01 << 2  //+
	TIM5EN    APB1ENR_Bits = 0x01 << 3  //+
	TIM6EN    APB1ENR_Bits = 0x01 << 4  //+
	TIM7EN    APB1ENR_Bits = 0x01 << 5  //+
	TIM12EN   APB1ENR_Bits = 0x01 << 6  //+
	TIM13EN   APB1ENR_Bits = 0x01 << 7  //+
	TIM14EN   APB1ENR_Bits = 0x01 << 8  //+
	LPTIM1EN  APB1ENR_Bits = 0x01 << 9  //+
	WWDGEN    APB1ENR_Bits = 0x01 << 11 //+
	SPI2EN    APB1ENR_Bits = 0x01 << 14 //+
	SPI3EN    APB1ENR_Bits = 0x01 << 15 //+
	SPDIFRXEN APB1ENR_Bits = 0x01 << 16 //+
	USART2EN  APB1ENR_Bits = 0x01 << 17 //+
	USART3EN  APB1ENR_Bits = 0x01 << 18 //+
	UART4EN   APB1ENR_Bits = 0x01 << 19 //+
	UART5EN   APB1ENR_Bits = 0x01 << 20 //+
	I2C1EN    APB1ENR_Bits = 0x01 << 21 //+
	I2C2EN    APB1ENR_Bits = 0x01 << 22 //+
	I2C3EN    APB1ENR_Bits = 0x01 << 23 //+
	I2C4EN    APB1ENR_Bits = 0x01 << 24 //+
	CAN1EN    APB1ENR_Bits = 0x01 << 25 //+
	CAN2EN    APB1ENR_Bits = 0x01 << 26 //+
	CECEN     APB1ENR_Bits = 0x01 << 27 //+
	PWREN     APB1ENR_Bits = 0x01 << 28 //+
	DACEN     APB1ENR_Bits = 0x01 << 29 //+
	UART7EN   APB1ENR_Bits = 0x01 << 30 //+
	UART8EN   APB1ENR_Bits = 0x01 << 31 //+
)

const (
	TIM2ENn    = 0
	TIM3ENn    = 1
	TIM4ENn    = 2
	TIM5ENn    = 3
	TIM6ENn    = 4
	TIM7ENn    = 5
	TIM12ENn   = 6
	TIM13ENn   = 7
	TIM14ENn   = 8
	LPTIM1ENn  = 9
	WWDGENn    = 11
	SPI2ENn    = 14
	SPI3ENn    = 15
	SPDIFRXENn = 16
	USART2ENn  = 17
	USART3ENn  = 18
	UART4ENn   = 19
	UART5ENn   = 20
	I2C1ENn    = 21
	I2C2ENn    = 22
	I2C3ENn    = 23
	I2C4ENn    = 24
	CAN1ENn    = 25
	CAN2ENn    = 26
	CECENn     = 27
	PWRENn     = 28
	DACENn     = 29
	UART7ENn   = 30
	UART8ENn   = 31
)

const (
	TIM1EN   APB2ENR_Bits = 0x01 << 0  //+
	TIM8EN   APB2ENR_Bits = 0x01 << 1  //+
	USART1EN APB2ENR_Bits = 0x01 << 4  //+
	USART6EN APB2ENR_Bits = 0x01 << 5  //+
	ADC1EN   APB2ENR_Bits = 0x01 << 8  //+
	ADC2EN   APB2ENR_Bits = 0x01 << 9  //+
	ADC3EN   APB2ENR_Bits = 0x01 << 10 //+
	SDMMC1EN APB2ENR_Bits = 0x01 << 11 //+
	SPI1EN   APB2ENR_Bits = 0x01 << 12 //+
	SPI4EN   APB2ENR_Bits = 0x01 << 13 //+
	SYSCFGEN APB2ENR_Bits = 0x01 << 14 //+
	TIM9EN   APB2ENR_Bits = 0x01 << 16 //+
	TIM10EN  APB2ENR_Bits = 0x01 << 17 //+
	TIM11EN  APB2ENR_Bits = 0x01 << 18 //+
	SPI5EN   APB2ENR_Bits = 0x01 << 20 //+
	SPI6EN   APB2ENR_Bits = 0x01 << 21 //+
	SAI1EN   APB2ENR_Bits = 0x01 << 22 //+
	SAI2EN   APB2ENR_Bits = 0x01 << 23 //+
	LTDCEN   APB2ENR_Bits = 0x01 << 26 //+
)

const (
	TIM1ENn   = 0
	TIM8ENn   = 1
	USART1ENn = 4
	USART6ENn = 5
	ADC1ENn   = 8
	ADC2ENn   = 9
	ADC3ENn   = 10
	SDMMC1ENn = 11
	SPI1ENn   = 12
	SPI4ENn   = 13
	SYSCFGENn = 14
	TIM9ENn   = 16
	TIM10ENn  = 17
	TIM11ENn  = 18
	SPI5ENn   = 20
	SPI6ENn   = 21
	SAI1ENn   = 22
	SAI2ENn   = 23
	LTDCENn   = 26
)

const (
	GPIOALPEN     AHB1LPENR_Bits = 0x01 << 0  //+
	GPIOBLPEN     AHB1LPENR_Bits = 0x01 << 1  //+
	GPIOCLPEN     AHB1LPENR_Bits = 0x01 << 2  //+
	GPIODLPEN     AHB1LPENR_Bits = 0x01 << 3  //+
	GPIOELPEN     AHB1LPENR_Bits = 0x01 << 4  //+
	GPIOFLPEN     AHB1LPENR_Bits = 0x01 << 5  //+
	GPIOGLPEN     AHB1LPENR_Bits = 0x01 << 6  //+
	GPIOHLPEN     AHB1LPENR_Bits = 0x01 << 7  //+
	GPIOILPEN     AHB1LPENR_Bits = 0x01 << 8  //+
	GPIOJLPEN     AHB1LPENR_Bits = 0x01 << 9  //+
	GPIOKLPEN     AHB1LPENR_Bits = 0x01 << 10 //+
	CRCLPEN       AHB1LPENR_Bits = 0x01 << 12 //+
	AXILPEN       AHB1LPENR_Bits = 0x01 << 13 //+
	FLITFLPEN     AHB1LPENR_Bits = 0x01 << 15 //+
	SRAM1LPEN     AHB1LPENR_Bits = 0x01 << 16 //+
	SRAM2LPEN     AHB1LPENR_Bits = 0x01 << 17 //+
	BKPSRAMLPEN   AHB1LPENR_Bits = 0x01 << 18 //+
	DTCMLPEN      AHB1LPENR_Bits = 0x01 << 20 //+
	DMA1LPEN      AHB1LPENR_Bits = 0x01 << 21 //+
	DMA2LPEN      AHB1LPENR_Bits = 0x01 << 22 //+
	DMA2DLPEN     AHB1LPENR_Bits = 0x01 << 23 //+
	ETHMACLPEN    AHB1LPENR_Bits = 0x01 << 25 //+
	ETHMACTXLPEN  AHB1LPENR_Bits = 0x01 << 26 //+
	ETHMACRXLPEN  AHB1LPENR_Bits = 0x01 << 27 //+
	ETHMACPTPLPEN AHB1LPENR_Bits = 0x01 << 28 //+
	OTGHSLPEN     AHB1LPENR_Bits = 0x01 << 29 //+
	OTGHSULPILPEN AHB1LPENR_Bits = 0x01 << 30 //+
)

const (
	GPIOALPENn     = 0
	GPIOBLPENn     = 1
	GPIOCLPENn     = 2
	GPIODLPENn     = 3
	GPIOELPENn     = 4
	GPIOFLPENn     = 5
	GPIOGLPENn     = 6
	GPIOHLPENn     = 7
	GPIOILPENn     = 8
	GPIOJLPENn     = 9
	GPIOKLPENn     = 10
	CRCLPENn       = 12
	AXILPENn       = 13
	FLITFLPENn     = 15
	SRAM1LPENn     = 16
	SRAM2LPENn     = 17
	BKPSRAMLPENn   = 18
	DTCMLPENn      = 20
	DMA1LPENn      = 21
	DMA2LPENn      = 22
	DMA2DLPENn     = 23
	ETHMACLPENn    = 25
	ETHMACTXLPENn  = 26
	ETHMACRXLPENn  = 27
	ETHMACPTPLPENn = 28
	OTGHSLPENn     = 29
	OTGHSULPILPENn = 30
)

const (
	DCMILPEN  AHB2LPENR_Bits = 0x01 << 0 //+
	RNGLPEN   AHB2LPENR_Bits = 0x01 << 6 //+
	OTGFSLPEN AHB2LPENR_Bits = 0x01 << 7 //+
)

const (
	DCMILPENn  = 0
	RNGLPENn   = 6
	OTGFSLPENn = 7
)

const (
	FMCLPEN  AHB3LPENR_Bits = 0x01 << 0 //+
	QSPILPEN AHB3LPENR_Bits = 0x01 << 1 //+
)

const (
	FMCLPENn  = 0
	QSPILPENn = 1
)

const (
	TIM2LPEN    APB1LPENR_Bits = 0x01 << 0  //+
	TIM3LPEN    APB1LPENR_Bits = 0x01 << 1  //+
	TIM4LPEN    APB1LPENR_Bits = 0x01 << 2  //+
	TIM5LPEN    APB1LPENR_Bits = 0x01 << 3  //+
	TIM6LPEN    APB1LPENR_Bits = 0x01 << 4  //+
	TIM7LPEN    APB1LPENR_Bits = 0x01 << 5  //+
	TIM12LPEN   APB1LPENR_Bits = 0x01 << 6  //+
	TIM13LPEN   APB1LPENR_Bits = 0x01 << 7  //+
	TIM14LPEN   APB1LPENR_Bits = 0x01 << 8  //+
	LPTIM1LPEN  APB1LPENR_Bits = 0x01 << 9  //+
	WWDGLPEN    APB1LPENR_Bits = 0x01 << 11 //+
	SPI2LPEN    APB1LPENR_Bits = 0x01 << 14 //+
	SPI3LPEN    APB1LPENR_Bits = 0x01 << 15 //+
	SPDIFRXLPEN APB1LPENR_Bits = 0x01 << 16 //+
	USART2LPEN  APB1LPENR_Bits = 0x01 << 17 //+
	USART3LPEN  APB1LPENR_Bits = 0x01 << 18 //+
	UART4LPEN   APB1LPENR_Bits = 0x01 << 19 //+
	UART5LPEN   APB1LPENR_Bits = 0x01 << 20 //+
	I2C1LPEN    APB1LPENR_Bits = 0x01 << 21 //+
	I2C2LPEN    APB1LPENR_Bits = 0x01 << 22 //+
	I2C3LPEN    APB1LPENR_Bits = 0x01 << 23 //+
	I2C4LPEN    APB1LPENR_Bits = 0x01 << 24 //+
	CAN1LPEN    APB1LPENR_Bits = 0x01 << 25 //+
	CAN2LPEN    APB1LPENR_Bits = 0x01 << 26 //+
	CECLPEN     APB1LPENR_Bits = 0x01 << 27 //+
	PWRLPEN     APB1LPENR_Bits = 0x01 << 28 //+
	DACLPEN     APB1LPENR_Bits = 0x01 << 29 //+
	UART7LPEN   APB1LPENR_Bits = 0x01 << 30 //+
	UART8LPEN   APB1LPENR_Bits = 0x01 << 31 //+
)

const (
	TIM2LPENn    = 0
	TIM3LPENn    = 1
	TIM4LPENn    = 2
	TIM5LPENn    = 3
	TIM6LPENn    = 4
	TIM7LPENn    = 5
	TIM12LPENn   = 6
	TIM13LPENn   = 7
	TIM14LPENn   = 8
	LPTIM1LPENn  = 9
	WWDGLPENn    = 11
	SPI2LPENn    = 14
	SPI3LPENn    = 15
	SPDIFRXLPENn = 16
	USART2LPENn  = 17
	USART3LPENn  = 18
	UART4LPENn   = 19
	UART5LPENn   = 20
	I2C1LPENn    = 21
	I2C2LPENn    = 22
	I2C3LPENn    = 23
	I2C4LPENn    = 24
	CAN1LPENn    = 25
	CAN2LPENn    = 26
	CECLPENn     = 27
	PWRLPENn     = 28
	DACLPENn     = 29
	UART7LPENn   = 30
	UART8LPENn   = 31
)

const (
	TIM1LPEN   APB2LPENR_Bits = 0x01 << 0  //+
	TIM8LPEN   APB2LPENR_Bits = 0x01 << 1  //+
	USART1LPEN APB2LPENR_Bits = 0x01 << 4  //+
	USART6LPEN APB2LPENR_Bits = 0x01 << 5  //+
	ADC1LPEN   APB2LPENR_Bits = 0x01 << 8  //+
	ADC2LPEN   APB2LPENR_Bits = 0x01 << 9  //+
	ADC3LPEN   APB2LPENR_Bits = 0x01 << 10 //+
	SDMMC1LPEN APB2LPENR_Bits = 0x01 << 11 //+
	SPI1LPEN   APB2LPENR_Bits = 0x01 << 12 //+
	SPI4LPEN   APB2LPENR_Bits = 0x01 << 13 //+
	SYSCFGLPEN APB2LPENR_Bits = 0x01 << 14 //+
	TIM9LPEN   APB2LPENR_Bits = 0x01 << 16 //+
	TIM10LPEN  APB2LPENR_Bits = 0x01 << 17 //+
	TIM11LPEN  APB2LPENR_Bits = 0x01 << 18 //+
	SPI5LPEN   APB2LPENR_Bits = 0x01 << 20 //+
	SPI6LPEN   APB2LPENR_Bits = 0x01 << 21 //+
	SAI1LPEN   APB2LPENR_Bits = 0x01 << 22 //+
	SAI2LPEN   APB2LPENR_Bits = 0x01 << 23 //+
	LTDCLPEN   APB2LPENR_Bits = 0x01 << 26 //+
)

const (
	TIM1LPENn   = 0
	TIM8LPENn   = 1
	USART1LPENn = 4
	USART6LPENn = 5
	ADC1LPENn   = 8
	ADC2LPENn   = 9
	ADC3LPENn   = 10
	SDMMC1LPENn = 11
	SPI1LPENn   = 12
	SPI4LPENn   = 13
	SYSCFGLPENn = 14
	TIM9LPENn   = 16
	TIM10LPENn  = 17
	TIM11LPENn  = 18
	SPI5LPENn   = 20
	SPI6LPENn   = 21
	SAI1LPENn   = 22
	SAI2LPENn   = 23
	LTDCLPENn   = 26
)

const (
	LSEON    BDCR_Bits = 0x01 << 0 //+
	LSERDY   BDCR_Bits = 0x01 << 1 //+
	LSEBYP   BDCR_Bits = 0x01 << 2 //+
	LSEDRV   BDCR_Bits = 0x03 << 3 //+
	LSEDRV_0 BDCR_Bits = 0x01 << 3
	LSEDRV_1 BDCR_Bits = 0x02 << 3
	RTCSEL   BDCR_Bits = 0x03 << 8 //+
	RTCSEL_0 BDCR_Bits = 0x01 << 8
	RTCSEL_1 BDCR_Bits = 0x02 << 8
	RTCEN    BDCR_Bits = 0x01 << 15 //+
	BDRST    BDCR_Bits = 0x01 << 16 //+
)

const (
	LSEONn  = 0
	LSERDYn = 1
	LSEBYPn = 2
	LSEDRVn = 3
	RTCSELn = 8
	RTCENn  = 15
	BDRSTn  = 16
)

const (
	LSION    CSR_Bits = 0x01 << 0  //+
	LSIRDY   CSR_Bits = 0x01 << 1  //+
	RMVF     CSR_Bits = 0x01 << 24 //+
	BORRSTF  CSR_Bits = 0x01 << 25 //+
	PINRSTF  CSR_Bits = 0x01 << 26 //+
	PORRSTF  CSR_Bits = 0x01 << 27 //+
	SFTRSTF  CSR_Bits = 0x01 << 28 //+
	IWDGRSTF CSR_Bits = 0x01 << 29 //+
	WWDGRSTF CSR_Bits = 0x01 << 30 //+
	LPWRRSTF CSR_Bits = 0x01 << 31 //+
)

const (
	LSIONn    = 0
	LSIRDYn   = 1
	RMVFn     = 24
	BORRSTFn  = 25
	PINRSTFn  = 26
	PORRSTFn  = 27
	SFTRSTFn  = 28
	IWDGRSTFn = 29
	WWDGRSTFn = 30
	LPWRRSTFn = 31
)

const (
	MODPER    SSCGR_Bits = 0x1FFF << 0  //+
	INCSTEP   SSCGR_Bits = 0x7FFF << 13 //+
	SPREADSEL SSCGR_Bits = 0x01 << 30   //+
	SSCGEN    SSCGR_Bits = 0x01 << 31   //+
)

const (
	MODPERn    = 0
	INCSTEPn   = 13
	SPREADSELn = 30
	SSCGENn    = 31
)

const (
	PLLI2SN   PLLI2SCFGR_Bits = 0x1FF << 6 //+
	PLLI2SN_0 PLLI2SCFGR_Bits = 0x01 << 6
	PLLI2SN_1 PLLI2SCFGR_Bits = 0x02 << 6
	PLLI2SN_2 PLLI2SCFGR_Bits = 0x04 << 6
	PLLI2SN_3 PLLI2SCFGR_Bits = 0x08 << 6
	PLLI2SN_4 PLLI2SCFGR_Bits = 0x10 << 6
	PLLI2SN_5 PLLI2SCFGR_Bits = 0x20 << 6
	PLLI2SN_6 PLLI2SCFGR_Bits = 0x40 << 6
	PLLI2SN_7 PLLI2SCFGR_Bits = 0x80 << 6
	PLLI2SN_8 PLLI2SCFGR_Bits = 0x100 << 6
	PLLI2SP   PLLI2SCFGR_Bits = 0x03 << 16 //+
	PLLI2SP_0 PLLI2SCFGR_Bits = 0x01 << 16
	PLLI2SP_1 PLLI2SCFGR_Bits = 0x02 << 16
	PLLI2SQ   PLLI2SCFGR_Bits = 0x0F << 24 //+
	PLLI2SQ_0 PLLI2SCFGR_Bits = 0x01 << 24
	PLLI2SQ_1 PLLI2SCFGR_Bits = 0x02 << 24
	PLLI2SQ_2 PLLI2SCFGR_Bits = 0x04 << 24
	PLLI2SQ_3 PLLI2SCFGR_Bits = 0x08 << 24
	PLLI2SR   PLLI2SCFGR_Bits = 0x07 << 28 //+
	PLLI2SR_0 PLLI2SCFGR_Bits = 0x01 << 28
	PLLI2SR_1 PLLI2SCFGR_Bits = 0x02 << 28
	PLLI2SR_2 PLLI2SCFGR_Bits = 0x04 << 28
)

const (
	PLLI2SNn = 6
	PLLI2SPn = 16
	PLLI2SQn = 24
	PLLI2SRn = 28
)

const (
	PLLSAIN   PLLSAICFGR_Bits = 0x1FF << 6 //+
	PLLSAIN_0 PLLSAICFGR_Bits = 0x01 << 6
	PLLSAIN_1 PLLSAICFGR_Bits = 0x02 << 6
	PLLSAIN_2 PLLSAICFGR_Bits = 0x04 << 6
	PLLSAIN_3 PLLSAICFGR_Bits = 0x08 << 6
	PLLSAIN_4 PLLSAICFGR_Bits = 0x10 << 6
	PLLSAIN_5 PLLSAICFGR_Bits = 0x20 << 6
	PLLSAIN_6 PLLSAICFGR_Bits = 0x40 << 6
	PLLSAIN_7 PLLSAICFGR_Bits = 0x80 << 6
	PLLSAIN_8 PLLSAICFGR_Bits = 0x100 << 6
	PLLSAIP   PLLSAICFGR_Bits = 0x03 << 16 //+
	PLLSAIP_0 PLLSAICFGR_Bits = 0x01 << 16
	PLLSAIP_1 PLLSAICFGR_Bits = 0x02 << 16
	PLLSAIQ   PLLSAICFGR_Bits = 0x0F << 24 //+
	PLLSAIQ_0 PLLSAICFGR_Bits = 0x01 << 24
	PLLSAIQ_1 PLLSAICFGR_Bits = 0x02 << 24
	PLLSAIQ_2 PLLSAICFGR_Bits = 0x04 << 24
	PLLSAIQ_3 PLLSAICFGR_Bits = 0x08 << 24
	PLLSAIR   PLLSAICFGR_Bits = 0x07 << 28 //+
	PLLSAIR_0 PLLSAICFGR_Bits = 0x01 << 28
	PLLSAIR_1 PLLSAICFGR_Bits = 0x02 << 28
	PLLSAIR_2 PLLSAICFGR_Bits = 0x04 << 28
)

const (
	PLLSAINn = 6
	PLLSAIPn = 16
	PLLSAIQn = 24
	PLLSAIRn = 28
)

const (
	PLLI2SDIVQ   DCKCFGR1_Bits = 0x1F << 0 //+
	PLLI2SDIVQ_0 DCKCFGR1_Bits = 0x01 << 0
	PLLI2SDIVQ_1 DCKCFGR1_Bits = 0x02 << 0
	PLLI2SDIVQ_2 DCKCFGR1_Bits = 0x04 << 0
	PLLI2SDIVQ_3 DCKCFGR1_Bits = 0x08 << 0
	PLLI2SDIVQ_4 DCKCFGR1_Bits = 0x10 << 0
	PLLSAIDIVQ   DCKCFGR1_Bits = 0x1F << 8 //+
	PLLSAIDIVQ_0 DCKCFGR1_Bits = 0x01 << 8
	PLLSAIDIVQ_1 DCKCFGR1_Bits = 0x02 << 8
	PLLSAIDIVQ_2 DCKCFGR1_Bits = 0x04 << 8
	PLLSAIDIVQ_3 DCKCFGR1_Bits = 0x08 << 8
	PLLSAIDIVQ_4 DCKCFGR1_Bits = 0x10 << 8
	PLLSAIDIVR   DCKCFGR1_Bits = 0x03 << 16 //+
	PLLSAIDIVR_0 DCKCFGR1_Bits = 0x01 << 16
	PLLSAIDIVR_1 DCKCFGR1_Bits = 0x02 << 16
	SAI1SEL      DCKCFGR1_Bits = 0x03 << 20 //+
	SAI1SEL_0    DCKCFGR1_Bits = 0x01 << 20
	SAI1SEL_1    DCKCFGR1_Bits = 0x02 << 20
	SAI2SEL      DCKCFGR1_Bits = 0x03 << 22 //+
	SAI2SEL_0    DCKCFGR1_Bits = 0x01 << 22
	SAI2SEL_1    DCKCFGR1_Bits = 0x02 << 22
	TIMPRE       DCKCFGR1_Bits = 0x01 << 24 //+
)

const (
	PLLI2SDIVQn = 0
	PLLSAIDIVQn = 8
	PLLSAIDIVRn = 16
	SAI1SELn    = 20
	SAI2SELn    = 22
	TIMPREn     = 24
)

const (
	USART1SEL   DCKCFGR2_Bits = 0x03 << 0 //+
	USART1SEL_0 DCKCFGR2_Bits = 0x01 << 0
	USART1SEL_1 DCKCFGR2_Bits = 0x02 << 0
	USART2SEL   DCKCFGR2_Bits = 0x03 << 2 //+
	USART2SEL_0 DCKCFGR2_Bits = 0x01 << 2
	USART2SEL_1 DCKCFGR2_Bits = 0x02 << 2
	USART3SEL   DCKCFGR2_Bits = 0x03 << 4 //+
	USART3SEL_0 DCKCFGR2_Bits = 0x01 << 4
	USART3SEL_1 DCKCFGR2_Bits = 0x02 << 4
	UART4SEL    DCKCFGR2_Bits = 0x03 << 6 //+
	UART4SEL_0  DCKCFGR2_Bits = 0x01 << 6
	UART4SEL_1  DCKCFGR2_Bits = 0x02 << 6
	UART5SEL    DCKCFGR2_Bits = 0x03 << 8 //+
	UART5SEL_0  DCKCFGR2_Bits = 0x01 << 8
	UART5SEL_1  DCKCFGR2_Bits = 0x02 << 8
	USART6SEL   DCKCFGR2_Bits = 0x03 << 10 //+
	USART6SEL_0 DCKCFGR2_Bits = 0x01 << 10
	USART6SEL_1 DCKCFGR2_Bits = 0x02 << 10
	UART7SEL    DCKCFGR2_Bits = 0x03 << 12 //+
	UART7SEL_0  DCKCFGR2_Bits = 0x01 << 12
	UART7SEL_1  DCKCFGR2_Bits = 0x02 << 12
	UART8SEL    DCKCFGR2_Bits = 0x03 << 14 //+
	UART8SEL_0  DCKCFGR2_Bits = 0x01 << 14
	UART8SEL_1  DCKCFGR2_Bits = 0x02 << 14
	I2C1SEL     DCKCFGR2_Bits = 0x03 << 16 //+
	I2C1SEL_0   DCKCFGR2_Bits = 0x01 << 16
	I2C1SEL_1   DCKCFGR2_Bits = 0x02 << 16
	I2C2SEL     DCKCFGR2_Bits = 0x03 << 18 //+
	I2C2SEL_0   DCKCFGR2_Bits = 0x01 << 18
	I2C2SEL_1   DCKCFGR2_Bits = 0x02 << 18
	I2C3SEL     DCKCFGR2_Bits = 0x03 << 20 //+
	I2C3SEL_0   DCKCFGR2_Bits = 0x01 << 20
	I2C3SEL_1   DCKCFGR2_Bits = 0x02 << 20
	I2C4SEL     DCKCFGR2_Bits = 0x03 << 22 //+
	I2C4SEL_0   DCKCFGR2_Bits = 0x01 << 22
	I2C4SEL_1   DCKCFGR2_Bits = 0x02 << 22
	LPTIM1SEL   DCKCFGR2_Bits = 0x03 << 24 //+
	LPTIM1SEL_0 DCKCFGR2_Bits = 0x01 << 24
	LPTIM1SEL_1 DCKCFGR2_Bits = 0x02 << 24
	CECSEL      DCKCFGR2_Bits = 0x01 << 26 //+
	CK48MSEL    DCKCFGR2_Bits = 0x01 << 27 //+
	SDMMC1SEL   DCKCFGR2_Bits = 0x01 << 28 //+
)

const (
	USART1SELn = 0
	USART2SELn = 2
	USART3SELn = 4
	UART4SELn  = 6
	UART5SELn  = 8
	USART6SELn = 10
	UART7SELn  = 12
	UART8SELn  = 14
	I2C1SELn   = 16
	I2C2SELn   = 18
	I2C3SELn   = 20
	I2C4SELn   = 22
	LPTIM1SELn = 24
	CECSELn    = 26
	CK48MSELn  = 27
	SDMMC1SELn = 28
)
