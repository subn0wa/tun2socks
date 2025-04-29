package proxy

import (
	"context"
	"errors"
	"net"

	M "github.com/subn0wa/tun2socks/metadata"
	"github.com/subn0wa/tun2socks/proxy/proto"
)

var _ Proxy = (*Base)(nil)

type Base struct {
	addr  string
	proto proto.Proto
}

func (b *Base) Addr() string {
	return b.addr
}

func (b *Base) Proto() proto.Proto {
	return b.proto
}

func (b *Base) DialContext(context.Context, *M.Metadata) (net.Conn, error) {
	return nil, errors.ErrUnsupported
}

func (b *Base) DialUDP(*M.Metadata) (net.PacketConn, error) {
	return nil, errors.ErrUnsupported
}
