package caddynet

import (
	// plug in the server
	_ "github.com/nielsavonds/caddy-net/caddynet/netserver"
	// // plug in the standard directives
	_ "github.com/nielsavonds/caddy-net/caddynet/host"
)
