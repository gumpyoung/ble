package dev

import (
	"github.com/gumpyoung/ble"
	"github.com/gumpyoung/ble/linux"
)

// DefaultDevice ...
func DefaultDevice() (d ble.Device, err error) {
	return linux.NewDevice()
}
