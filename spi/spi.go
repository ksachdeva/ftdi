package spi

// #include "native/libMPSSE_spi.h"
// #include <stdlib.h>
import "C"
import "fmt"

// DeviceHandle is the handle to the channel
type DeviceHandle struct {
	handlePtr *C.FT_HANDLE
}

// DeviceChannelInfo is the information related to the channel
type DeviceChannelInfo struct {
	SerialNumber string
	ptr          *C.FT_DEVICE_LIST_INFO_NODE
}

// GetNumChannels returns the number of channels available
func GetNumChannels() (int, error) {
	var numOfChannels C.uint32
	status := C.SPI_GetNumChannels(&numOfChannels)
	if status != 0 {
		return -1, fmt.Errorf("an error occurred %g", status)
	}

	return int(numOfChannels), nil
}

// OpenChannel takes an index of the channel to be opened
func OpenChannel(channelIndex int) (handle *DeviceHandle, err error) {
	var handlePtr C.FT_HANDLE
	status := C.SPI_OpenChannel(C.uint32(channelIndex), &handlePtr)
	if status != 0 {
		return nil, fmt.Errorf("an error occurred %g", status)
	}
	return &DeviceHandle{handlePtr: &handlePtr}, nil
}

// CloseChannel closes the channel
func CloseChannel(handle *DeviceHandle) (err error) {
	status := C.SPI_CloseChannel(*handle.handlePtr)
	if status != 0 {
		return fmt.Errorf("an error occurred %g", status)
	}
	return nil
}

// GetChannelInfo returns the channel info
func GetChannelInfo(channelIndex int) (channelInfo *DeviceChannelInfo, err error) {
	var ptr C.FT_DEVICE_LIST_INFO_NODE
	status := C.SPI_GetChannelInfo(C.uint32(channelIndex), &ptr)
	if status != 0 {
		return nil, fmt.Errorf("an error occurred %g", status)
	}
	deviceInfo := &DeviceChannelInfo{ptr: &ptr}
	deviceInfo.SerialNumber = C.GoString(&ptr.SerialNumber[0])

	return deviceInfo, nil
}

// InitChannel initializes the channel
func InitChannel(deviceHandle *DeviceHandle) (err error) {
	var channelConfig C.ChannelConfig
	channelConfig.ClockRate = 500 * 1000
	channelConfig.LatencyTimer = 10
	channelConfig.configOptions = 0x00000020
	channelConfig.Pin = 0

	status := C.SPI_InitChannel(*deviceHandle.handlePtr, &channelConfig)
	if status != 0 {
		return fmt.Errorf("an error occurred %g", status)
	}
	return nil
}

// Write send the data to the slave SPI device
func Write(deviceHandle *DeviceHandle, data []uint8) (dataTransferred int, err error) {
	var sizeTransferred C.uint32
	var options C.uint32 = 0x00000002
	status := C.SPI_Write(*deviceHandle.handlePtr, (*C.uint8)(&data[0]), C.uint32(len(data)), &sizeTransferred, options)
	if status != 0 {
		return -1, fmt.Errorf("an error occurred %g", status)
	}
	return int(sizeTransferred), nil
}
