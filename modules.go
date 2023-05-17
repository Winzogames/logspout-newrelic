package main

import (
	_ "github.com/gliderlabs/logspout/adapters/multiline"
	_ "github.com/gliderlabs/logspout/adapters/raw"
	_ "github.com/Winzogames/logspout-newrelic/syslog"
	_ "github.com/gliderlabs/logspout/healthcheck"
	_ "github.com/gliderlabs/logspout/httpstream"
	_ "github.com/gliderlabs/logspout/routesapi"
	_ "github.com/gliderlabs/logspout/transports/tcp"
	_ "github.com/gliderlabs/logspout/transports/tls"
	_ "github.com/gliderlabs/logspout/transports/udp"
)
