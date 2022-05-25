// Package multistream implements a peerstream transport using
// go-multistream to select the underlying stream muxer
//
// Deprecated: This package has moved into go-libp2p as a sub-package: github.com/libp2p/go-libp2p/p2p/muxer/muxer-multistream.
package multistream

import (
	muxer_multistream "github.com/libp2p/go-libp2p/p2p/muxer/muxer-multistream"
)

// Deprecated: use github.com/libp2p/go-libp2p/p2p/muxer/muxer-multistream.DefaultNegoiateTimeout instead.
var DefaultNegotiateTimeout = muxer_multistream.DefaultNegotiateTimeout

// Deprecated: use github.com/libp2p/go-libp2p/p2p/muxer/muxer-multistream.Transport instead.
type Transport = muxer_multistream.Transport

// Deprecated: use github.com/libp2p/go-libp2p/p2p/muxer/muxer-multistream.NewBlankTransport instead.
func NewBlankTransport() *Transport {
	return muxer_multistream.NewBlankTransport()
}
