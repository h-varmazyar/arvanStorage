package service

import (
	"fmt"
	"net"
)

var (
	serves = make(map[uint16]ServeFunc)
)

type ServeFunc func(listener net.Listener) error

func Serve(port uint16, f ServeFunc) {
	serves[port] = f
}

func (serve ServeFunc) Listen(port uint16) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return err
	}
	defer func() {
		_ = listener.Close()
	}()
	return serve(listener)
}
