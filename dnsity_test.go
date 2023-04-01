package main

import (
	"net"
	"testing"

	"github.com/miekg/dns"
)

func TestDNSServer(t *testing.T) {

	// case 1
	query1 := createDNSQuery("www.example.com.", dns.TypeA)
	resp1 := sendDNSQuery(query1)
	if resp1.Rcode != dns.RcodeSuccess {
		t.Errorf("unexpected Rcode: %v", resp1.Rcode)
	}
	if len(resp1.Answer) != 1 {
		t.Errorf("unexpected number of answers: %v", len(resp1.Answer))
	}
	if resp1.Answer[0].Header().Name != "www.example.com." {
		t.Errorf("unexpected answer name: %v", resp1.Answer[0].Header().Name)
	}
	if resp1.Answer[0].Header().Rrtype != dns.TypeA {
		t.Errorf("unexpected answer type: %v", resp1.Answer[0].Header().Rrtype)
	}
	if resp1.Answer[0].(*dns.A).A.String() != "192.0.2.1" {
		t.Errorf("unexpected answer value: %v", resp1.Answer[0].(*dns.A).A.String())
	}

	// case 2
	query2 := createDNSQuery("nosuchhost.example.com.", dns.TypeA)
	resp2 := sendDNSQuery(query2)
	if resp2.Rcode != dns.RcodeNameError {
		t.Errorf("unexpected Rcode: %v", resp2.Rcode)
	}
	if len(resp2.Answer) != 0 {
		t.Errorf("unexpected number of answers: %v", len(resp2.Answer))
	}
}

func createDNSQuery(name string, qtype uint16) *dns.Msg {
	query := new(dns.Msg)
	query.SetQuestion(name, qtype)
	query.RecursionDesired = true
	return query
}

func sendDNSQuery(query *dns.Msg) *dns.Msg {
	conn, err := net.Dial("udp", "127.0.0.1:53")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	queryMsg, err := query.Pack()
	if err != nil {
		panic(err)
	}

	if _, err := conn.Write(queryMsg); err != nil {
		panic(err)
	}

	respMsg := make([]byte, 1024)
	n, err := conn.Read(respMsg)
	if err != nil {
		panic(err)
	}

	resp := new(dns.Msg)
	if err := resp.Unpack(respMsg[:n]); err != nil {
		panic(err)
	}

	return resp
}
