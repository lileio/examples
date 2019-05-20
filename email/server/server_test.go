package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/lileio/examples/email"
	"github.com/lileio/lile"
)

var s = EmailServer{}
var cli email.EmailClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		email.RegisterEmailServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = email.NewEmailClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
