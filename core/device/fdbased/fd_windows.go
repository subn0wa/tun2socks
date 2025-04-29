package fdbased

import (
	"errors"

	"github.com/subn0wa/tun2socks/core/device"
)

func Open(name string, mtu uint32, offset int) (device.Device, error) {
	return nil, errors.ErrUnsupported
}
