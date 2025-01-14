//go:build all || race
// +build all race

package peer

import (
	"testing"

	"github.com/jirs5/tracing-proxy/config"
)

func TestFilePeers(t *testing.T) {
	peers := []string{"peer"}

	c := &config.MockConfig{
		GetPeersVal: peers,
	}
	p := newFilePeers(c)

	if d, _ := p.GetPeers(); !(len(d) == 1 && d[0] == "peer") {
		t.Error("received", d, "expected", "[peer]")
	}
}
