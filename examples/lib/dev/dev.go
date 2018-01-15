package dev

import "github.com/gumpyoung/ble"

// NewDevice ...
func NewDevice(impl string) (d ble.Device, err error) {
	return DefaultDevice()
}
