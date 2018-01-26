package main

import (
	"fmt"

	spi "github.com/ksachdeva/ftdi/spi"
)

func main() {
	println("Hello World !")

	// var NumChannels []uint32

	//	status := m.SPI_GetNumChannels(NumChannels)

	channels, err := spi.GetNumChannels()

	if err == nil {
		fmt.Printf("Number of channels %d\n", channels)
	} else {
		fmt.Printf("Failed to get the number of channels")
		return
	}

	var ptr *spi.DeviceChannelInfo

	ptr, err = spi.GetChannelInfo(0)

	if err == nil {
		fmt.Println(ptr.SerialNumber)
	} else {
		fmt.Printf("Failed to get the channel information")
		return
	}

	var handle *spi.DeviceHandle

	handle, err = spi.OpenChannel(0)

	if err == nil {
		// we may have the handle
	} else {
		fmt.Printf("Failed to open the channel")
		return
	}

	err = spi.InitChannel(handle)

	if err == nil {

	} else {
		fmt.Printf("Failed to initialize the channel - %s", err)
	}

	err = spi.CloseChannel(handle)

	if err == nil {

	} else {
		fmt.Printf("Failed to close the channel - %s", err)
	}
}
