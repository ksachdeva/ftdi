package spi

// #include "native/libMPSSE_spi.h"
// #include <stdlib.h>
import "C"
import "fmt"

// TransferOptions specify the properties of data being sent and behavior of CS line
type TransferOptions uint32

// ChannelConfigOptions specifies the various properties of the channel
type ChannelConfigOptions uint32

// FTErrCode is the error code returned by libftdi
type FTErrCode int

const (
	// ErrInvalidHandle for invalid handle
	ErrInvalidHandle FTErrCode = 1
	// ErrDeviceNotFound when device is not connected/found
	ErrDeviceNotFound FTErrCode = 2
	// ErrDeviceNotOpened when channel is not opened
	ErrDeviceNotOpened FTErrCode = 3
	// ErrIOError for an I/O error
	ErrIOError FTErrCode = 4
	// ErrInsufficientResources for not having sufficient resources
	ErrInsufficientResources FTErrCode = 5
	// ErrInvalidParameter for passing an invalid parameter
	ErrInvalidParameter FTErrCode = 6
	// ErrInvalidBaudRate for invalid baud rate
	ErrInvalidBaudRate FTErrCode = 7

	// ErrDeviceNotOpenedForErase for when device is not opened for erased
	ErrDeviceNotOpenedForErase FTErrCode = 8
	// ErrDeviceNotOpenedForWrite for when device is not opened for write
	ErrDeviceNotOpenedForWrite FTErrCode = 9
	// ErrFailedToWriteDevice for when failed to write to device
	ErrFailedToWriteDevice FTErrCode = 10
	// ErrEEPROMReadFailed for when failed to do eeprom read
	ErrEEPROMReadFailed FTErrCode = 11
	// ErrEEPROMWriteFailed for when failed to do eeprom write
	ErrEEPROMWriteFailed FTErrCode = 12
	// ErrEEPROMEraseFailed for when eeprom erase failed
	ErrEEPROMEraseFailed FTErrCode = 13
	// ErrEEPROMNotPresent for when eeprom is not present
	ErrEEPROMNotPresent FTErrCode = 14
	// ErrEEPROMNotProgrammed for when eeprom is not programmed
	ErrEEPROMNotProgrammed FTErrCode = 15
	// ErrInvalidArgs for invalid arguements
	ErrInvalidArgs FTErrCode = 16
	// ErrNotSupported for when operation is not supported
	ErrNotSupported FTErrCode = 17
	// ErrOtherError for when the error is unknown
	ErrOtherError FTErrCode = 18
)

const (
	// InputSizeIsInBytes let your specify the size provided is in bytes
	InputSizeIsInBytes TransferOptions = 0x00000000
	// InputSizeIsInBits let your specify the size provided is in bits
	InputSizeIsInBits TransferOptions = 0x00000001
	// EnableCSAtStartOfTransfer will enable the CS at the start of the transfer
	EnableCSAtStartOfTransfer TransferOptions = 0x00000002
	// DisableCSAtEndOfTransfer will disable the CS at the end of the transfer
	DisableCSAtEndOfTransfer TransferOptions = 0x00000004
)

const (
	// Mode0 is for SPI Mode 0
	Mode0 ChannelConfigOptions = 0x00000000
	// Mode1 is for SPI Mode 1
	Mode1 ChannelConfigOptions = 0x00000001
	// Mode2 is for SPI Mode 2
	Mode2 ChannelConfigOptions = 0x00000002
	// Mode3 is for SPI Mode 3
	Mode3 ChannelConfigOptions = 0x00000003

	// ChipSelectIsDBUS3 will configure DBUS3 to be used as chip select line
	ChipSelectIsDBUS3 ChannelConfigOptions = 0x00000000
	// ChipSelectIsDBUS4 will configure DBUS4 to be used as chip select line
	ChipSelectIsDBUS4 ChannelConfigOptions = 0x00000004
	// ChipSelectIsDBUS5 will configure DBUS5 to be used as chip select line
	ChipSelectIsDBUS5 ChannelConfigOptions = 0x00000008
	// ChipSelectIsDBUS6 will configure DBUS6 to be used as chip select line
	ChipSelectIsDBUS6 ChannelConfigOptions = 0x0000000C
	// ChipSelectIsDBUS7 will configure DBUS7 to be used as chip select line
	ChipSelectIsDBUS7 ChannelConfigOptions = 0x00000010

	// ChipSelectIsActiveLow will ..
	ChipSelectIsActiveLow ChannelConfigOptions = 0x00000020
)

// FTError is an error class containing the reason for the error
type FTError struct {
	Code FTErrCode
}

func (e FTError) Error() string {
	return fmt.Sprintf("Error Code: %v", e.Code)
}

// ChannelConfiguration specifies how channel is to be intialized
type ChannelConfiguration struct {
	ClockRate     uint32
	LatencyTimer  byte
	ConfigOptions ChannelConfigOptions
}

// ChannelHandle is the handle to the channel
type ChannelHandle struct {
	handlePtr *C.FT_HANDLE
}

// ChannelInfo is the information related to the channel
type ChannelInfo struct {
	SerialNumber string
	ptr          *C.FT_DEVICE_LIST_INFO_NODE
}

// GetNumChannels returns the number of channels available
func GetNumChannels() (int, error) {
	var numOfChannels C.uint32
	status := C.SPI_GetNumChannels(&numOfChannels)
	if status != 0 {
		return -1, FTError{Code: FTErrCode(status)}
	}

	return int(numOfChannels), nil
}

// OpenChannel takes an index of the channel to be opened
func OpenChannel(channelIndex int) (handle *ChannelHandle, err error) {
	var handlePtr C.FT_HANDLE
	status := C.SPI_OpenChannel(C.uint32(channelIndex), &handlePtr)
	if status != 0 {
		return nil, FTError{Code: FTErrCode(status)}
	}
	return &ChannelHandle{handlePtr: &handlePtr}, nil
}

// CloseChannel closes the channel
func CloseChannel(handle *ChannelHandle) (err error) {
	status := C.SPI_CloseChannel(*handle.handlePtr)
	if status != 0 {
		return FTError{Code: FTErrCode(status)}
	}
	return nil
}

// GetChannelInfo returns the channel info
func GetChannelInfo(channelIndex int) (channelInfo *ChannelInfo, err error) {
	var ptr C.FT_DEVICE_LIST_INFO_NODE
	status := C.SPI_GetChannelInfo(C.uint32(channelIndex), &ptr)
	if status != 0 {
		return nil, FTError{Code: FTErrCode(status)}
	}
	deviceInfo := &ChannelInfo{ptr: &ptr}
	deviceInfo.SerialNumber = C.GoString(&ptr.SerialNumber[0])

	return deviceInfo, nil
}

// InitChannel initializes the channel
func InitChannel(deviceHandle *ChannelHandle, chanConfig ChannelConfiguration) (err error) {
	var channelConfig C.ChannelConfig
	channelConfig.ClockRate = C.uint32(chanConfig.ClockRate)
	channelConfig.LatencyTimer = C.uint8(chanConfig.LatencyTimer)
	channelConfig.configOptions = C.uint32(chanConfig.ConfigOptions)
	channelConfig.Pin = 0

	status := C.SPI_InitChannel(*deviceHandle.handlePtr, &channelConfig)
	if status != 0 {
		return FTError{Code: FTErrCode(status)}
	}
	return nil
}

// Write send the data to the slave SPI device
func Write(deviceHandle *ChannelHandle, data []uint8, transferOptions TransferOptions) (dataTransferred int, err error) {
	var sizeTransferred C.uint32
	var options C.uint32 = C.uint32(transferOptions)
	status := C.SPI_Write(*deviceHandle.handlePtr, (*C.uint8)(&data[0]), C.uint32(len(data)), &sizeTransferred, options)
	if status != 0 {
		return -1, FTError{Code: FTErrCode(status)}
	}
	return int(sizeTransferred), nil
}
