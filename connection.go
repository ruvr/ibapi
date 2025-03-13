/* connection handle the tcp socket to the TWS or IB Gateway*/

package ibapi

import (
	"log/slog"
	"net"
	"strconv"
)

// IbConnection wrap the tcp connection with TWS or Gateway
type IbConnection struct {
	*net.TCPConn
	host         string
	port         int
	state        int
	numBytesSent int
	numMsgSent   int
	numBytesRecv int
	numMsgRecv   int
}

func (ibconn *IbConnection) Write(bs []byte) (int, error) {
	n, err := ibconn.TCPConn.Write(bs)

	ibconn.numBytesSent += n
	ibconn.numMsgSent++

	log.Debug("conn write", slog.Int("nBytes", n))

	return n, err
}

func (ibconn *IbConnection) Read(bs []byte) (int, error) {
	n, err := ibconn.TCPConn.Read(bs)

	ibconn.numBytesRecv += n
	ibconn.numMsgRecv++

	log.Debug("conn read", slog.Int("nBytes", n))

	return n, err
}

func (ibconn *IbConnection) reset() {
	ibconn.numBytesSent = 0
	ibconn.numBytesRecv = 0
	ibconn.numMsgSent = 0
	ibconn.numMsgRecv = 0
}

func (ibconn *IbConnection) disconnect() error {
	log.Debug("conn disconnect",
		slog.Int("nMsgSent", ibconn.numMsgSent),
		slog.Int("nBytesSent", ibconn.numBytesSent),
		slog.Int("nMsgRecv", ibconn.numMsgRecv),
		slog.Int("nBytesRecv", ibconn.numBytesRecv),
	)
	return ibconn.Close()
}

func (ibconn *IbConnection) connect(host string, port int) error {
	var err error
	var addr *net.TCPAddr
	ibconn.host = host
	ibconn.port = port
	ibconn.reset()

	server := ibconn.host + ":" + strconv.Itoa(port)
	if addr, err = net.ResolveTCPAddr("tcp4", server); err != nil {
		log.Error("failed to resove tcp address", slog.Any("error", err), slog.String("host", server))
		return err
	}

	if ibconn.TCPConn, err = net.DialTCP("tcp4", nil, addr); err != nil {
		log.Error("failed to dial tcp", slog.Any("error", err), slog.Any("address", addr))
		return err
	}

	log.Debug("tcp socket connected", slog.Any("address", ibconn.TCPConn.RemoteAddr()))

	return err
}
