package server

import (
	"github.com/gavrylenkoIvan/hopper/internal/hopper"
	sbound "github.com/gavrylenkoIvan/hopper/public/serverbound"
)

const (
	StatusState int = 0x01
	LoginState  int = 0x2
)

func (h *Hopper) handshake(conn *hopper.Conn) error {
	defer conn.Close()

	// new conn always starts with handshake packet
	p := new(sbound.Handshake)
	_, _, err := conn.ReadPacket(p)
	if err != nil {
		return err
	}

	switch int(p.NextState) {
	case StatusState:
		return h.status(conn)
	case LoginState:
		return h.login(conn)
	}

	return nil
}
