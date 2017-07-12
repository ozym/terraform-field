package airos

import (
	"fmt"
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

type AirOs struct {
	snmp *gosnmp.GoSNMP
	mu   sync.Mutex
	err  error
}

func Port(port int) func(*AirOs) error {
	return func(a *AirOs) error {
		a.snmp.Port = uint16(port)
		return nil
	}
}

func Community(community string) func(*AirOs) error {
	return func(a *AirOs) error {
		a.snmp.Community = community
		return nil
	}
}

func Timeout(timeout time.Duration) func(*AirOs) error {
	return func(a *AirOs) error {
		a.snmp.Timeout = timeout
		return nil
	}
}

func Retries(retries int) func(*AirOs) error {
	return func(a *AirOs) error {
		a.snmp.Retries = retries
		return nil
	}
}

func NewAirOs(target string, options ...func(*AirOs) error) (*AirOs, error) {

	a := &AirOs{
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

func (a *AirOs) Id() string {
	return net.JoinHostPort(a.snmp.Target, strconv.Itoa(int(a.snmp.Port)))
}

func (a *AirOs) Error() error {
	return a.err
}

func (a *AirOs) Connect() error {
	return a.snmp.Connect()
}

type Status struct {
	RadioMode   string
	CountryCode int
	Frequency   int32
	DfsEnabled  int
	TxPower     int32
	Distance    int32
	Antenna     string
}

func (a *AirOs) Status() (*Status, error) {
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

	if err := a.snmp.Walk(".1.3.6.1.4.1.41112.1.4.1", func(pdu gosnmp.SnmpPDU) error {
		if oid := strings.Split(pdu.Name, "."); len(oid) > 8 {
			switch pdu.Type {
			case gosnmp.OctetString:
				switch {
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.1.1.9"):
					if pdu.Value.([]byte) != nil {
						s.Antenna = string(pdu.Value.([]byte))
					}
				}
			case gosnmp.Integer:
				switch {
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.1.1.2"):
					if pdu.Value.(int) != 0 {
						s.RadioMode = func() string {
							switch pdu.Value.(int) {
							case 1:
								return "sta"
							case 2:
								return "ap"
							case 3:
								return "aprepeater"
							case 4:
								return "apwds"
							default:
								return "unknown"
							}
						}()
					}
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.1.1.3"):
					s.CountryCode = pdu.Value.(int)
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.1.1.4"):
					s.Frequency = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.1.1.5"):
					s.DfsEnabled = pdu.Value.(int)
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.1.1.7"):
					s.TxPower = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.1.1.8"):
					s.Distance = int32(pdu.Value.(int))
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

type Wireless struct {
	Ssid         string
	HideSsid     int
	ApMacAddr    string
	Signal       int32
	Rssi         int32
	Ccq          int
	NoiseFloor   int32
	TxRate       int32
	RxRate       int32
	Security     string
	WdsEnabled   int
	ApRepeater   int
	ChannelWidth int32
	StationCount int
}

func (a *AirOs) Wireless() (*Wireless, error) {

	a.mu.Lock()
	defer a.mu.Unlock()

	if a.err != nil {
		return nil, a.err
	}

	if err := a.Connect(); err != nil {
		a.err = err

		return nil, err
	}

	w := &Wireless{}

	if err := a.snmp.Walk(".1.3.6.1.4.1.41112.1.4.5", func(pdu gosnmp.SnmpPDU) error {
		if oid := strings.Split(pdu.Name, "."); len(oid) > 8 {
			switch pdu.Type {
			case gosnmp.OctetString:
				switch {
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.5.1.2"):
					if pdu.Value.([]byte) != nil {
						w.Ssid = string(pdu.Value.([]byte))
					}
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.5.1.4"):
					var parts []string
					for i := 0; i < len(pdu.Value.([]byte)); i++ {
						parts = append(parts, fmt.Sprintf("%02x", pdu.Value.([]byte)[i]))
					}
					w.ApMacAddr = strings.Join(parts, ":")
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.5.1.11"):
					if pdu.Value.([]byte) != nil {
						w.Security = string(pdu.Value.([]byte))
					}
				}
			case gosnmp.Integer:
				switch {
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.5.1.3"):
					w.HideSsid = pdu.Value.(int)
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.5.1.5"):
					w.Signal = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.5.1.6"):
					w.Rssi = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.5.1.7"):
					w.Ccq = pdu.Value.(int)
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.5.1.8"):
					w.NoiseFloor = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.5.1.9"):
					w.TxRate = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.5.1.10"):
					w.RxRate = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.5.1.12"):
					w.WdsEnabled = pdu.Value.(int)
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.5.1.13"):
					w.ApRepeater = pdu.Value.(int)
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.5.1.14"):
					w.ChannelWidth = int32(pdu.Value.(int))
				}
			case gosnmp.Gauge32:
				switch {
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.5.1.15"):
					w.StationCount = int(pdu.Value.(uint))
				}
			}
		}
		return nil
	}); err != nil {
		a.err = err

		return nil, err
	}

	return w, nil
}

type Station struct {
	MacAddr    string
	Name       string
	Signal     int32
	NoiseFloor int32
	Distance   int32
	Ccq        int32
	Amp        int32
	Amq        int32
	Amc        int32
	LastIp     string
	TxRate     int64
	RxRate     int64
	TxBytes    int64
	RxBytes    int64
	ConnTime   int64
}

func (a *AirOs) Stations() ([]Station, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.err != nil {
		return nil, a.err
	}

	if err := a.Connect(); err != nil {
		a.err = err

		return nil, err
	}

	stns := make(map[string]*Station)
	if err := a.snmp.Walk(".1.3.6.1.4.1.41112.1.4.7", func(pdu gosnmp.SnmpPDU) error {
		if oid := strings.Split(pdu.Name, "."); len(oid) > 8 {
			switch pdu.Type {
			case gosnmp.OctetString:
				if pdu.Value.([]byte) != nil {
					switch {
					case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.7.1.1"):
						suffix := strings.Replace(pdu.Name, ".1.3.6.1.4.1.41112.1.4.7.1.1", "", 1)
						if _, ok := stns[suffix]; !ok {
							stns[suffix] = &Station{}
						}
						var parts []string
						for i := 0; i < len(pdu.Value.([]byte)); i++ {
							parts = append(parts, fmt.Sprintf("%02x", pdu.Value.([]byte)[i]))
						}
						stns[suffix].MacAddr = strings.Join(parts, ":")
					case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.7.1.2"):
						suffix := strings.Replace(pdu.Name, ".1.3.6.1.4.1.41112.1.4.7.1.2", "", 1)
						if _, ok := stns[suffix]; !ok {
							stns[suffix] = &Station{}
						}
						stns[suffix].Name = string(pdu.Value.([]byte))
					}
				}
			case gosnmp.Integer:
				switch {
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.7.1.3"):
					suffix := strings.Replace(pdu.Name, ".1.3.6.1.4.1.41112.1.4.7.1.3", "", 1)
					if _, ok := stns[suffix]; !ok {
						stns[suffix] = &Station{}
					}
					stns[suffix].Signal = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.7.1.4"):
					suffix := strings.Replace(pdu.Name, ".1.3.6.1.4.1.41112.1.4.7.1.4", "", 1)
					if _, ok := stns[suffix]; !ok {
						stns[suffix] = &Station{}
					}
					stns[suffix].NoiseFloor = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.7.1.5"):
					suffix := strings.Replace(pdu.Name, ".1.3.6.1.4.1.41112.1.4.7.1.5", "", 1)
					if _, ok := stns[suffix]; !ok {
						stns[suffix] = &Station{}
					}
					stns[suffix].Distance = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.7.1.6"):
					suffix := strings.Replace(pdu.Name, ".1.3.6.1.4.1.41112.1.4.7.1.6", "", 1)
					if _, ok := stns[suffix]; !ok {
						stns[suffix] = &Station{}
					}
					stns[suffix].Ccq = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.7.1.7"):
					suffix := strings.Replace(pdu.Name, ".1.3.6.1.4.1.41112.1.4.7.1.7", "", 1)
					if _, ok := stns[suffix]; !ok {
						stns[suffix] = &Station{}
					}
					stns[suffix].Amp = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.7.1.8"):
					suffix := strings.Replace(pdu.Name, ".1.3.6.1.4.1.41112.1.4.7.1.8", "", 1)
					if _, ok := stns[suffix]; !ok {
						stns[suffix] = &Station{}
					}
					stns[suffix].Amq = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.7.1.9"):
					suffix := strings.Replace(pdu.Name, ".1.3.6.1.4.1.41112.1.4.7.1.9", "", 1)
					if _, ok := stns[suffix]; !ok {
						stns[suffix] = &Station{}
					}
					stns[suffix].Amc = int32(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.7.1.11"):
					suffix := strings.Replace(pdu.Name, ".1.3.6.1.4.1.41112.1.4.7.1.11", "", 1)
					if _, ok := stns[suffix]; !ok {
						stns[suffix] = &Station{}
					}
					stns[suffix].TxRate = int64(pdu.Value.(int))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.7.1.12"):
					suffix := strings.Replace(pdu.Name, ".1.3.6.1.4.1.41112.1.4.7.1.12", "", 1)
					if _, ok := stns[suffix]; !ok {
						stns[suffix] = &Station{}
					}
					stns[suffix].RxRate = int64(pdu.Value.(int))
				}
			case gosnmp.IPAddress:
				switch {
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.7.1.10"):
					suffix := strings.Replace(pdu.Name, ".1.3.6.1.4.1.41112.1.4.7.1.10", "", 1)
					if _, ok := stns[suffix]; !ok {
						stns[suffix] = &Station{}
					}
					stns[suffix].LastIp = pdu.Value.(string)
				}
			case gosnmp.Counter64:
				switch {
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.7.1.13"):
					suffix := strings.Replace(pdu.Name, ".1.3.6.1.4.1.41112.1.4.7.1.13", "", 1)
					if _, ok := stns[suffix]; !ok {
						stns[suffix] = &Station{}
					}
					stns[suffix].TxBytes = int64(pdu.Value.(int64))
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.7.1.14"):
					suffix := strings.Replace(pdu.Name, ".1.3.6.1.4.1.41112.1.4.7.1.14", "", 1)
					if _, ok := stns[suffix]; !ok {
						stns[suffix] = &Station{}
					}
					stns[suffix].RxBytes = int64(pdu.Value.(int64))
				}
			case gosnmp.TimeTicks:
				switch {
				case strings.HasPrefix(string(pdu.Name), ".1.3.6.1.4.1.41112.1.4.7.1.15"):
					suffix := strings.Replace(pdu.Name, ".1.3.6.1.4.1.41112.1.4.7.1.15", "", 1)
					if _, ok := stns[suffix]; !ok {
						stns[suffix] = &Station{}
					}
					stns[suffix].ConnTime = int64(pdu.Value.(int))
				}
			}
		}
		return nil
	}); err != nil {
		a.err = err

		return nil, err
	}

	var res []Station
	for _, v := range stns {
		res = append(res, *v)
	}

	return res, nil
}
