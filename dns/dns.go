package dns

import (
	"fmt"

	"github.com/miekg/dns"
	"go.uber.org/zap"

	"github.com/sp-yduck/dnsity/utils"
)

var logger *zap.SugaredLogger
var records map[string]string

func init() {
	logger = utils.NewLogger("dns")
	records = make(map[string]string)
	if err := configLoader("config.yaml"); err != nil {
		logger.Errorf("cannot load config file")
	}
}

func Run() error {
	dns.HandleFunc(".", handleDNSRequest)
	logger.Infof("register new dns handle func")

	server := &dns.Server{Addr: ":53", Net: "udp"}
	ch := make(chan error, 1)
	go func() error {
		err := server.ListenAndServe()
		if err != nil {
			logger.Errorf("Failed to start server: %s", err.Error())
			return err
		}
		return nil
	}()

	select {
	case err := <-ch:
		return err
	default:
		return nil
	}
}

func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	logger.Debugf("received query: %v", r)

	msg := new(dns.Msg)
	msg.SetReply(r)

	switch r.Opcode {
	case dns.OpcodeQuery:
		parseQuery(msg)
		logger.Debugf("Answer : ", msg.Answer)
	}

	err := w.WriteMsg(msg)
	if err != nil {
		logger.Errorf("Failed to send response: %v", err)
	}
}

func parseQuery(m *dns.Msg) {
	for _, q := range m.Question {
		switch q.Qtype {
		case dns.TypeA:
			logger.Debugf("Query for %s\n", q.Name)
			ip, ok := records[q.Name]
			if !ok {
				m.Rcode = dns.RcodeNameError
				continue
			}
			if ip != "" {
				rr, err := dns.NewRR(fmt.Sprintf("%s A %s", q.Name, ip))
				if err == nil {
					m.Answer = append(m.Answer, rr)
				}
			}
		}
	}
}

func RegisterRecord(name string, ip string) {
	name += "."
	records[name] = ip
}

func Records() *map[string]string {
	return &records
}
