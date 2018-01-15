package dev

import (
	"github.com/gumpyoung/ble"
	"github.com/gumpyoung/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice() (d ble.Device, err error) {
	return darwin.NewDevice()
}
