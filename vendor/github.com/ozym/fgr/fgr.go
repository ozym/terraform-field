package fgr

import (
	//"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/soniah/gosnmp"
)

const (
	DefaultCommunity = "public"
	DefaultPort      = 161
	DefaultTimeout   = 30 * time.Second
	DefaultRetries   = 3
)

type Fgr struct {
	snmp *gosnmp.GoSNMP
	mu   sync.Mutex
	err  error
}

func Port(port int) func(*Fgr) error {
	return func(a *Fgr) error {
		a.snmp.Port = uint16(port)
		return nil
	}
}

func Community(community string) func(*Fgr) error {
	return func(a *Fgr) error {
		a.snmp.Community = community
		return nil
	}
}

func Timeout(timeout time.Duration) func(*Fgr) error {
	return func(a *Fgr) error {
		a.snmp.Timeout = timeout
		return nil
	}
}

func Retries(retries int) func(*Fgr) error {
	return func(a *Fgr) error {
		a.snmp.Retries = retries
		return nil
	}
}

func NewFgr(target string, options ...func(*Fgr) error) (*Fgr, error) {

	a := &Fgr{
		snmp: &gosnmp.GoSNMP{
			Target:    target,
			Port:      DefaultPort,
			Community: DefaultCommunity,
			Version:   gosnmp.Version1,
			Timeout:   DefaultTimeout,
			Retries:   DefaultRetries,
		},
	}

	for _, option := range options {
		if err := option(a); err != nil {
			a.err = err

			return nil, err
		}
	}

	return a, nil
}

func (a *Fgr) Id() string {
	return net.JoinHostPort(a.snmp.Target, strconv.Itoa(int(a.snmp.Port)))
}

func (a *Fgr) Error() error {
	return a.err
}

func (a *Fgr) Connect() error {
	return a.snmp.Connect()
}

type Status struct {
	Model    string
	Version  string
	Firmware string
	Contact  string
	Name     string
	Location string

	Signal          int32
	Noise           int32
	SupplyVoltage   int32
	RxRate          int32
	ReflectedPower  int32
	Temperature     int32
	Range           int32
	TxRate          int32
	SNDelta         int32
	VendorString    string
	ConnectedTo     string
	UpstreamSignal  int32
	UpstreamNoise   int32
	DisconnectCount int32
	PacketRxCount   int32
	PacketTxCount   int32
	DroppedCount    int32
	BadCount        int32
}

func (a *Fgr) Status() (*Status, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.err != nil {
		return nil, a.err
	}

	if err := a.snmp.Connect(); err != nil {
		a.err = err

		return nil, err
	}

	s := &Status{}
	if err := a.snmp.Walk(".1.3.6.1.2.1.1", func(pdu gosnmp.SnmpPDU) error {
		if oid := strings.Split(pdu.Name, "."); len(oid) > 8 {
			switch pdu.Type {
			case gosnmp.OctetString:
				if pdu.Value.([]byte) != nil {
					switch {
					case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.2.1.1.1.0"):
						s.Version = string(pdu.Value.([]byte))
						l := strings.Replace((string)(pdu.Value.([]byte)), "Freewave Technologies", "Freewave", -1)
						l = strings.Replace(strings.Replace(strings.Replace(l, " (", ";", -1), " ;", ";", -1), ")", "", -1)
						if f := strings.Split(l, ";"); len(f) > 0 {
							s.Model = f[0]
							for _, j := range f[1:] {
								if k := strings.Fields(j); len(k) > 2 {
									switch k[0] {
									case "wireless":
										s.Firmware = k[2]
									}
								}
							}
						}
					case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.2.1.1.4"):
						s.Contact = string(pdu.Value.([]byte))
					case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.2.1.1.5"):
						s.Name = string(pdu.Value.([]byte))
					case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.2.1.1.6"):
						s.Location = string(pdu.Value.([]byte))
					}
				}
			}
		}
		return nil
	}); err != nil {
		a.err = err

		return nil, err
	}

	if err := a.snmp.Walk(".1.3.6.1.4.1.29956.3.1.1", func(pdu gosnmp.SnmpPDU) error {
		if oid := strings.Split(pdu.Name, "."); len(oid) > 8 {
			switch pdu.Type {
			case gosnmp.OctetString:
				switch {
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.29956.3.1.1.1.1.11.1"):
					s.VendorString = string(pdu.Value.([]byte))
				}
			case gosnmp.Integer:
				switch {
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.29956.3.1.1.1.1.2.1"):
					s.Signal = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.29956.3.1.1.1.1.3.1"):
					if pdu.Value.(int) > 0 {
						s.Noise = int32(-pdu.Value.(int))
					} else {
						s.Noise = int32(pdu.Value.(int))
					}
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.29956.3.1.1.1.1.4.1"):
					switch s.Version {
					case `Freewave Technologies FGRplus (firmware version 2.12 7/23/2008; wireless version 2.65s ; boot version 1)`:
						s.SupplyVoltage = int32(7.885 * float64(pdu.Value.(int)))
					default:
						s.SupplyVoltage = int32(pdu.Value.(int))
					}
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.29956.3.1.1.1.1.5.1"):
					s.RxRate = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.29956.3.1.1.1.1.6.1"):
					s.ReflectedPower = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.29956.3.1.1.1.1.7.1"):
					s.Temperature = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.29956.3.1.1.1.1.8.1"):
					s.Range = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.29956.3.1.1.1.1.9.1"):
					s.TxRate = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.29956.3.1.1.1.1.10.1"):
					s.SNDelta = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.29956.3.1.1.1.1.12.1"):
					s.ConnectedTo = strconv.Itoa(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.29956.3.1.1.1.1.13.1"):
					s.UpstreamSignal = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.29956.3.1.1.1.1.14.1"):
					s.UpstreamNoise = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.29956.3.1.1.1.1.15.1"):
					s.DisconnectCount = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.29956.3.1.1.1.1.16.1"):
					s.PacketRxCount = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.29956.3.1.1.1.1.17.1"):
					s.PacketTxCount = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.29956.3.1.1.1.1.18.1"):
					s.DroppedCount = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.29956.3.1.1.1.1.19.1"):
					s.BadCount = int32(pdu.Value.(int))
				}
			}
		}
		return nil
	}); err != nil {
		a.err = err

		return nil, err
	}

	return s, nil
}
