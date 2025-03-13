package ibapi

import (
	"fmt"
	"log/slog"
	"testing"
)

func TestConnection(t *testing.T) {
	fmt.Println("connection testing!")
	conn := &IbConnection{}
	conn.connect("127.0.0.1", 4002)
	buf := make([]byte, 4096)
	_, err := conn.Read(buf)
	if err != nil {
		log.Error("read error", slog.Any("error", err))
	} else {
		fmt.Println(string(buf))
	}
	conn.disconnect()
}
