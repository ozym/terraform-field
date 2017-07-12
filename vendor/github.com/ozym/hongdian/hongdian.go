package hongdian

import (
	"bufio"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	DefaultPort    = 23
	DefaultTimeout = 30 * time.Second
)

type Hongdian struct {
	hostname string
	port     int

	password string

	timeout time.Duration

	mu  sync.Mutex
	err error
}

func Port(port int) func(*Hongdian) error {
	return func(h *Hongdian) error {
		h.port = port
		return nil
	}
}

func Password(password string) func(*Hongdian) error {
	return func(h *Hongdian) error {
		h.password = password
		return nil
	}
}

func Timeout(timeout time.Duration) func(*Hongdian) error {
	return func(h *Hongdian) error {
		h.timeout = timeout
		return nil
	}
}

/*
func Retries(retries int) func(*Hongdian) error {
	return func(a *Hongdian) error {
		a.snmp.Retries = retries
		return nil
	}
}
*/

func NewHongdian(hostname string, options ...func(*Hongdian) error) (*Hongdian, error) {

	h := &Hongdian{
		hostname: hostname,
		port:     DefaultPort,
		timeout:  DefaultTimeout,
	}

	for _, option := range options {
		if err := option(h); err != nil {
			h.err = err

			return nil, err
		}
	}

	return h, nil
}

func (h *Hongdian) Id() string {
	return net.JoinHostPort(h.hostname, strconv.Itoa(int(h.port)))
}

func (h *Hongdian) Error() error {
	return h.err
}

type Status struct {
	RouterModel     string
	HardwareVersion string
	SoftwareVersion string
	SerialNumber    string

	ModemStatus    string
	ModemSignal    int
	IpAddress      string
	DeviceName     string
	VendorId       string
	ProductId      string
	ProductSn      string
	OnlineTime     string
	SystemTime     string
	ModemInterface int
	SimStatus      string
	CardStatus     string
	NetworkType    string
}

func (h *Hongdian) Status() (*Status, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.err != nil {
		return nil, h.err
	}

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(h.hostname, strconv.Itoa(h.port)), h.timeout)
	if err != nil {
		h.err = err

		return nil, err
	}
	defer conn.Close()

	conn.Write([]byte(strings.Join([]string{h.password, "show version", "enable", "show modem-information", "exit"}, "\n") + "\n"))

	var lines []string

	details := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		fields := strings.FieldsFunc(scanner.Text(), func(c rune) bool { return c == ':' })
		if len(fields) > 1 {
			details[strings.TrimSpace(fields[0])] = strings.TrimSpace(strings.Join(fields[1:], ":"))
		}
	}
	if err := scanner.Err(); err != nil {
		h.err = err

		return nil, err
	}

	s := &Status{}

	for k, v := range details {
		switch k {
		case "Router Model":
			s.RouterModel = v
		case "Hardware Version":
			s.HardwareVersion = v
		case "Software Version":
			s.SoftwareVersion = v
		case "Serial Number":
			s.SerialNumber = v
		case "Modem Status":
			s.ModemStatus = v
		case "Modem Signal":
			if i, err := strconv.Atoi(v); err == nil {
				s.ModemSignal = i
			}
		case "IP Address":
			s.IpAddress = v
		case "Device Name":
			s.DeviceName = v
		case "Vendor ID":
			s.VendorId = v
		case "Product ID":
			s.ProductId = v
		case "Product SN":
			s.ProductSn = v
		case "Online Time":
			s.OnlineTime = v
		case "System Time":
			s.SystemTime = v
		case "Modem Interface":
			if i, err := strconv.Atoi(v); err == nil {
				s.ModemInterface = i
			}
		case "SIM Status":
			s.SimStatus = v
		case "Card Status":
			s.CardStatus = v
		case "Network Type":
			s.NetworkType = v
		}
	}

	return s, nil
}
