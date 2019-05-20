package main

import (
	_ "net/http/pprof"

	"github.com/lileio/examples/email"
	"github.com/lileio/examples/email/email/cmd"
	"github.com/lileio/examples/email/server"
	"github.com/lileio/fromenv"
	"github.com/lileio/lile"
	"github.com/lileio/logr"
	"github.com/lileio/pubsub"
	"github.com/lileio/pubsub/middleware/defaults"
	"google.golang.org/grpc"
)

func main() {
	logr.SetLevelFromEnv()
	s := &server.EmailServer{}

	lile.Name("email")
	lile.Server(func(g *grpc.Server) {
		email.RegisterEmailServer(g, s)
	})

	pubsub.SetClient(&pubsub.Client{
		ServiceName: lile.GlobalService().Name,
		Provider:    fromenv.PubSubProvider(),
		Middleware:  defaults.Middleware,
	})

	cmd.Execute()
}
