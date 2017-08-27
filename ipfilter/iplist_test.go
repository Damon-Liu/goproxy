package ipfilter

import (
	"bytes"
	"net"
	"testing"

	"github.com/shell909090/goproxy/tunnel"
)

const iplist = "10.0.0.0 255.0.0.0\n172.16.0.0 255.240.0.0\n192.168.0.0 255.255.0.0"

func TestIPList(t *testing.T) {
	tunnel.SetLogging()

	buf := bytes.NewBufferString(iplist)
	filter, err := ReadIPList(buf)
	if err != nil {
		t.Fatalf("ReadIPList failed: %s", err)
	}

	if !filter.Contain(net.ParseIP("192.168.1.1")) {
		t.Fatalf("Contain wrong1.")
	}

	if !filter.Contain(net.ParseIP("10.8.0.1")) {
		t.Fatalf("Contain wrong2.")
	}

	if filter.Contain(net.ParseIP("211.80.90.25")) {
		t.Fatalf("Contain wrong3.")
	}
}
