package main

import (
	"fmt"

	spi "github.com/ksachdeva/ftdi/spi"
)

func main() {

	channels, err := spi.GetNumChannels()

	if err == nil {
		fmt.Printf("Number of channels %d\n", channels)
	} else {
		fmt.Printf("Failed to get the number of channels")
		return
	}

	var ptr *spi.ChannelInfo

	ptr, err = spi.GetChannelInfo(0)

	if err == nil {
		fmt.Println(ptr.SerialNumber)
	} else {
		fmt.Printf("Failed to get the channel information\n")
		return
	}

	var handle *spi.ChannelHandle

	handle, err = spi.OpenChannel(0)

	if err == nil {
		// we may have the handle
	} else {
		fmt.Printf("Failed to open the channel\n")
		return
	}

	var channelConfig spi.ChannelConfiguration
	channelConfig.ClockRate = 500 * 1000
	channelConfig.LatencyTimer = 10
	channelConfig.ConfigOptions = spi.ChipSelectIsActiveLow | spi.ChipSelectIsDBUS3

	err = spi.InitChannel(handle, channelConfig)

	if err == nil {

	} else {
		fmt.Printf("Failed to initialize the channel - %s\n", err)
		return
	}

	var sizeTransferred int
	var dataToTransfer = []byte{0x48, 0x45, 0x4C, 0x4C, 0x4F, 0x20, 0x57, 0x4F, 0x52, 0x4C, 0x44, 0x0A}
	sizeTransferred, err = spi.Write(handle, dataToTransfer, spi.InputSizeIsInBytes|spi.EnableCSAtStartOfTransfer)

	if err == nil {
		fmt.Printf("Number of bytes that were transferred %d\n", sizeTransferred)
	} else {
		fmt.Printf("Failed to send the data - %s\n", err)
		return
	}

	err = spi.CloseChannel(handle)

	if err == nil {

	} else {
		fmt.Printf("Failed to close the channel - %s\n", err)
	}
}
