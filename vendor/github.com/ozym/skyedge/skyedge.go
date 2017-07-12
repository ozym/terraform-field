package skyedge

import (
	"encoding/xml"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	DefaultUsername = ""
	DefaultPassword = ""
	DefaultTimeout  = 30 * time.Second
	DefaultPort     = 80
)

type SkyEdge struct {
	client *http.Client

	hostname string
	port     int
	username string
	password string

	mu  sync.Mutex
	err error
}

func Port(port int) func(*SkyEdge) error {
	return func(s *SkyEdge) error {
		s.port = port
		return nil
	}
}

func Username(username string) func(*SkyEdge) error {
	return func(s *SkyEdge) error {
		s.username = username
		return nil
	}
}

func Password(password string) func(*SkyEdge) error {
	return func(s *SkyEdge) error {
		s.password = password
		return nil
	}
}

func Timeout(timeout time.Duration) func(*SkyEdge) error {
	return func(s *SkyEdge) error {
		s.client.Timeout = timeout
		return nil
	}
}

func NewSkyEdge(hostname string, options ...func(*SkyEdge) error) (*SkyEdge, error) {

	s := &SkyEdge{
		client: &http.Client{
			Timeout: DefaultTimeout,
		},
		hostname: hostname,
		port:     DefaultPort,
		username: DefaultUsername,
		password: DefaultPassword,
	}

	for _, option := range options {
		if err := option(s); err != nil {
			s.err = err

			return nil, err
		}
	}

	return s, nil
}

func (s *SkyEdge) Id() string {
	return net.JoinHostPort(s.hostname, strconv.Itoa(s.port))
}

func (s *SkyEdge) Error() error {
	return s.err
}

type Status struct {
	StActiveSw  string
	WebDivMode  string
	StGwCon     string
	StSatLink   string
	StSync      string
	StAuth      string
	StAuthor    string
	StNetLink   string
	StTcpAcc    string
	StObLock    string
	StRpa       string
	StLanPort   string
	StLanPort2  string
	StPwrMode   string
	StOnTime    string
	StActiveSat string
	StMode      string
	StSdCard    string

	WebGwtStatus1 string
	WebGwtId1     string
	WebGwtMac1    string
	WebGwrs       string
	WebOthers     string

	VsatPn string
	VsatId string
	VsatSn string

	HwMBoard      string
	SwBtVer       string
	SwActiveBtVer string
	SwOpVer       string

	NetMac   string
	NetIp    string
	NetMask  string
	NetAIp   string
	NetAMask string

	StCputil string

	AccessMode      string
	RxEbN0          string
	OutboundFreq    string
	ModCod          string
	DrrpLanSent     string
	DrrpLanReceived string
}

func (s *SkyEdge) Status() (*Status, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.err != nil {
		return nil, s.err
	}

	r := &Status{}

	for _, p := range []string{"status", "receivers", "info", "cpu", "ebn0", "telemetry"} {
		u := url.URL{
			Scheme:   "http",
			Host:     net.JoinHostPort(s.hostname, strconv.Itoa(s.port)),
			Path:     "/cgi-bin/vsat.cgi",
			RawQuery: "action=" + p,
		}

		request, err := http.NewRequest("GET", u.String(), nil)
		if err != nil {
			s.err = err

			return nil, err
		}
		request.SetBasicAuth(s.username, s.password)

		if err := clientHandleResponseBody(s.client, request, func(body []byte) error {
			var res values
			if err := xml.Unmarshal(body, &res); err != nil {
				return err
			}
			for _, v := range res.Values {
				switch strings.TrimSpace(v.Id) {
				case "ST_ACTIVE_SW":
					r.StActiveSw = strings.TrimSpace(v.Val)
				case "WEB_DIV_MODE":
					r.WebDivMode = strings.TrimSpace(v.Val)
				case "ST_GW_CON":
					r.StGwCon = strings.TrimSpace(v.Val)
				case "ST_SAT_LINK":
					r.StSatLink = strings.TrimSpace(v.Val)
				case "ST_SYNC":
					r.StSync = strings.TrimSpace(v.Val)
				case "ST_AUTH":
					r.StAuth = strings.TrimSpace(v.Val)
				case "ST_AUTHOR":
					r.StAuthor = strings.TrimSpace(v.Val)
				case "ST_NET_LINK":
					r.StNetLink = strings.TrimSpace(v.Val)
				case "ST_TCP_ACC":
					r.StTcpAcc = strings.TrimSpace(v.Val)
				case "ST_RPA":
					r.StRpa = strings.TrimSpace(v.Val)
				case "ST_OB_LOCK":
					r.StObLock = strings.TrimSpace(v.Val)
				case "ST_LAN_PORT":
					r.StLanPort = strings.TrimSpace(v.Val)
				case "ST_LAN_PORT2":
					r.StLanPort2 = strings.TrimSpace(v.Val)
				case "ST_PWR_MODE":
					r.StPwrMode = strings.TrimSpace(v.Val)
				case "ST_ON_TIME":
					r.StOnTime = strings.TrimSpace(v.Val)
				case "ST_ACTIVE_SAT":
					r.StActiveSat = strings.TrimSpace(v.Val)
				case "ST_MODE":
					r.StMode = strings.TrimSpace(v.Val)
				case "ST_SD_CARD":
					r.StSdCard = strings.TrimSpace(v.Val)
				case "WEB_GWT_STATUS1":
					r.WebGwtStatus1 = strings.TrimSpace(v.Val)
				case "WEB_GWT_ID1":
					r.WebGwtId1 = strings.TrimSpace(v.Val)
				case "WEB_GWT_MAC1":
					r.WebGwtMac1 = strings.TrimSpace(v.Val)
				case "WEB_GWRS":
					r.WebGwrs = strings.TrimSpace(v.Val)
				case "WEB_OTHERS":
					r.WebOthers = strings.TrimSpace(v.Val)
				case "VSAT_PN":
					r.VsatPn = strings.TrimSpace(v.Val)
				case "VSAT_ID":
					r.VsatId = strings.TrimSpace(v.Val)
				case "VSAT_SN":
					r.VsatSn = strings.TrimSpace(v.Val)
				case "HW_M_BOARD":
					r.HwMBoard = strings.TrimSpace(v.Val)
				case "SW_BT_VER":
					r.SwBtVer = strings.TrimSpace(v.Val)
				case "SW_ACTIVE_BT_VER":
					r.SwActiveBtVer = strings.TrimSpace(v.Val)
				case "SW_OP_VER":
					r.SwOpVer = strings.TrimSpace(v.Val)
				case "NET_MAC":
					r.NetMac = strings.TrimSpace(v.Val)
				case "NET_IP":
					r.NetIp = strings.TrimSpace(v.Val)
				case "NET_MASK":
					r.NetMask = strings.TrimSpace(v.Val)
				case "NET_A_IP":
					r.NetAIp = strings.TrimSpace(v.Val)
				case "NET_A_MASK":
					r.NetAMask = strings.TrimSpace(v.Val)
				case "ST_CPU_UTIL":
					r.StCputil = strings.TrimSpace(v.Val)
				case "AccessMode":
					r.AccessMode = strings.TrimSpace(v.Val)
				case "RxEbN0":
					r.AccessMode = strings.TrimSpace(v.Val)
				case "OUTBOUND_FREQ":
					r.OutboundFreq = strings.TrimSpace(v.Val)
				case "MODCOD":
					r.ModCod = strings.TrimSpace(v.Val)
				case "DRPP_LAN_SENT":
					r.DrrpLanSent = strings.TrimSpace(v.Val)
				case "DRPP_LAN_RECEIVED":
					r.DrrpLanReceived = strings.TrimSpace(v.Val)
				}
			}
			return nil
		}); err != nil {
			s.err = err

			return nil, err
		}
	}

	return r, nil
}
