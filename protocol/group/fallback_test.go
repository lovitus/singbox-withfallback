package group

import (
	"context"
	"errors"
	"net"
	"testing"

	"github.com/sagernet/sing-box/adapter"
	"github.com/sagernet/sing-box/common/interrupt"
	M "github.com/sagernet/sing/common/metadata"
	N "github.com/sagernet/sing/common/network"
)

type testOutbound struct {
	tag      string
	networks []string
}

func (t *testOutbound) Type() string {
	return "test"
}

func (t *testOutbound) Tag() string {
	return t.tag
}

func (t *testOutbound) Network() []string {
	return t.networks
}

func (t *testOutbound) Dependencies() []string {
	return nil
}

func (t *testOutbound) DialContext(context.Context, string, M.Socksaddr) (net.Conn, error) {
	return nil, errors.New("not implemented")
}

func (t *testOutbound) ListenPacket(context.Context, M.Socksaddr) (net.PacketConn, error) {
	return nil, errors.New("not implemented")
}

func TestFallbackThresholdSwitching(t *testing.T) {
	a := &testOutbound{tag: "a", networks: []string{N.NetworkTCP, N.NetworkUDP}}
	b := &testOutbound{tag: "b", networks: []string{N.NetworkTCP, N.NetworkUDP}}
	g := &FallbackGroup{
		outbounds:           []adapter.Outbound{a, b},
		failureThreshold:    3,
		successThreshold:    3,
		outboundStatus:      map[string]fallbackOutboundState{"a": {available: true}, "b": {available: true}},
		interruptGroup:      interrupt.NewGroup(),
		selectedOutboundTCP: nil,
		selectedOutboundUDP: nil,
	}

	g.performUpdateCheck()
	if g.selectedOutboundTCP != a {
		t.Fatalf("expected initial TCP outbound a, got %v", g.selectedOutboundTCP)
	}

	g.markOutboundUnavailable("a")
	g.performUpdateCheck()
	if g.selectedOutboundTCP != a {
		t.Fatalf("expected TCP outbound a before reaching failure threshold")
	}

	g.markOutboundUnavailable("a")
	g.performUpdateCheck()
	if g.selectedOutboundTCP != a {
		t.Fatalf("expected TCP outbound a before reaching failure threshold")
	}

	g.markOutboundUnavailable("a")
	g.performUpdateCheck()
	if g.selectedOutboundTCP != b {
		t.Fatalf("expected TCP outbound switched to b after failure threshold")
	}

	g.markOutboundAvailable("a")
	g.performUpdateCheck()
	if g.selectedOutboundTCP != b {
		t.Fatalf("expected TCP outbound remain b before reaching success threshold")
	}

	g.markOutboundAvailable("a")
	g.performUpdateCheck()
	if g.selectedOutboundTCP != b {
		t.Fatalf("expected TCP outbound remain b before reaching success threshold")
	}

	g.markOutboundAvailable("a")
	g.performUpdateCheck()
	if g.selectedOutboundTCP != a {
		t.Fatalf("expected TCP outbound switched back to a after success threshold")
	}
}

func TestFallbackSelectByNetworkAndAvailability(t *testing.T) {
	a := &testOutbound{tag: "a", networks: []string{N.NetworkTCP}}
	b := &testOutbound{tag: "b", networks: []string{N.NetworkUDP}}
	c := &testOutbound{tag: "c", networks: []string{N.NetworkTCP, N.NetworkUDP}}
	g := &FallbackGroup{
		outbounds: []adapter.Outbound{a, b, c},
		outboundStatus: map[string]fallbackOutboundState{
			"a": {available: false},
			"b": {available: true},
			"c": {available: true},
		},
		interruptGroup: interrupt.NewGroup(),
	}

	selected, ok := g.Select(N.NetworkTCP)
	if !ok || selected != c {
		t.Fatalf("expected available TCP outbound c, got %v (ok=%v)", selected, ok)
	}

	g.outboundStatus["c"] = fallbackOutboundState{available: false}
	selected, ok = g.Select(N.NetworkTCP)
	if ok || selected != a {
		t.Fatalf("expected fallback TCP outbound a when all unavailable, got %v (ok=%v)", selected, ok)
	}
}
