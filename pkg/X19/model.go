package X19

import (
	"net/url"
)

const (
	// CgiPath default url path to make requests.
	CgiPath string = "/cgi-bin/"
	// SummaryUrl endpoint to get summary information
	SummaryUrl string = "summary.cgi"
	// StatsUrl endpoint to get stats information
	StatsUrl string = "stats.cgi"
	// HistoryLogUrl endpoint to get historical log of the device
	HistoryLogUrl string = "hlog.cgi"
	// CurrentLogUrl endpoint to get current log of the device
	CurrentLogUrl string = "log.cgi"
	// BlinkStatusUrl endpoint to get status of led blinking
	BlinkStatusUrl string = "get_blink_status.cgi"
	// BlinkUrl  endpoint to enable or disabled led blinking
	BlinkUrl string = "blink.cgi"
	// RebootUrl endpoint to make device reboot
	RebootUrl string = "reboot.cgi"
	// MinerConfUrl endpoint to get current device config
	MinerConfUrl string = "get_miner_conf.cgi"
	// SystemInfoUrl endpoint to get current system information
	SystemInfoUrl string = "get_system_info.cgi"
	// NetworkInfoUrl endpoint to get current network information
	NetworkInfoUrl string = "get_network_info.cgi"
	ChartUrl       string = "chart.cgi"
	// PoolsUrl endpoint to get information about pools
	PoolsUrl string = "pools.cgi"
	// SetMinerConfUrl endpoint to set specific configuration of the device
	SetMinerConfUrl string = "set_miner_conf.cgi"
)

type Device struct {
	Target  url.URL `json:"target"`
	User    string  `json:"user,omitempty"`
	Pass    string  `json:"pass,omitempty"`
	Payload string  `json:"payload,omitempty"`
}

type Stats struct {
	STATUS struct {
		STATUS     string `json:"STATUS"`
		When       int    `json:"when"`
		Msg        string `json:"Msg"`
		ApiVersion string `json:"api_version"`
	} `json:"STATUS"`
	INFO struct {
		MinerVersion string `json:"miner_version"`
		CompileTime  string `json:"CompileTime"`
		Type         string `json:"type"`
	} `json:"INFO"`
	STATS []struct {
		Elapsed   int     `json:"elapsed"`
		Rate5S    float64 `json:"rate_5s"`
		Rate30M   float64 `json:"rate_30m"`
		RateAvg   float64 `json:"rate_avg"`
		RateIdeal float64 `json:"rate_ideal"`
		RateUnit  string  `json:"rate_unit"`
		ChainNum  int     `json:"chain_num"`
		FanNum    int     `json:"fan_num"`
		Fan       []int   `json:"fan"`
		HwpTotal  float64 `json:"hwp_total"`
		MinerMode int     `json:"miner-mode"`
		FreqLevel int     `json:"freq-level"`
		Chain     []struct {
			Index        int     `json:"index"`
			FreqAvg      int     `json:"freq_avg"`
			RateIdeal    float64 `json:"rate_ideal"`
			RateReal     float64 `json:"rate_real"`
			AsicNum      int     `json:"asic_num"`
			Asic         string  `json:"asic"`
			TempPic      []int   `json:"temp_pic"`
			TempPcb      []int   `json:"temp_pcb"`
			TempChip     []int   `json:"temp_chip"`
			Hw           int     `json:"hw"`
			EepromLoaded bool    `json:"eeprom_loaded"`
			Sn           string  `json:"sn"`
			Hwp          float64 `json:"hwp"`
		} `json:"chain"`
	} `json:"STATS"`
}

type Pools struct {
	STATUS struct {
		STATUS     string `json:"STATUS,omitempty"`
		When       int    `json:"when,omitempty"`
		Msg        string `json:"Msg,omitempty"`
		ApiVersion string `json:"api_version,omitempty"`
	} `json:"STATUS"`
	INFO struct {
		MinerVersion string `json:"miner_version"`
		CompileTime  string `json:"CompileTime"`
		Type         string `json:"type"`
	} `json:"INFO"`
	POOLS []struct {
		Index     int    `json:"index"`
		Url       string `json:"url"`
		User      string `json:"user"`
		Status    string `json:"status"`
		Priority  int    `json:"priority"`
		Getworks  int    `json:"getworks"`
		Accepted  int    `json:"accepted"`
		Rejected  int    `json:"rejected"`
		Discarded int    `json:"discarded"`
		Stale     int    `json:"stale"`
		Diff      string `json:"diff"`
		Diff1     int    `json:"diff1"`
		Diffa     int64  `json:"diffa"`
		Diffr     int    `json:"diffr"`
		Diffs     int    `json:"diffs"`
		Lsdiff    int    `json:"lsdiff"`
		Lstime    string `json:"lstime"`
	} `json:"POOLS"`
}

type Summary struct {
	STATUS struct {
		STATUS     string `json:"STATUS"`
		When       int    `json:"when"`
		Msg        string `json:"Msg"`
		ApiVersion string `json:"api_version"`
	} `json:"STATUS"`
	INFO struct {
		MinerVersion string `json:"miner_version"`
		CompileTime  string `json:"CompileTime"`
		Type         string `json:"type"`
	} `json:"INFO"`
	SUMMARY []struct {
		Elapsed   int     `json:"elapsed"`
		Rate5S    float64 `json:"rate_5s"`
		Rate30M   float64 `json:"rate_30m"`
		RateAvg   float64 `json:"rate_avg"`
		RateIdeal float64 `json:"rate_ideal"`
		RateUnit  string  `json:"rate_unit"`
		HwAll     int     `json:"hw_all"`
		Bestshare int64   `json:"bestshare"`
		Status    []struct {
			Type   string `json:"type"`
			Status string `json:"status"`
			Code   int    `json:"code"`
			Msg    string `json:"msg"`
		} `json:"status"`
	} `json:"SUMMARY"`
}

type Chart struct {
	STATUS struct {
		STATUS     string `json:"STATUS"`
		When       int    `json:"when"`
		Msg        string `json:"Msg"`
		ApiVersion string `json:"api_version"`
	} `json:"STATUS"`
	INFO struct {
		MinerVersion string `json:"miner_version"`
		CompileTime  string `json:"CompileTime"`
		Type         string `json:"type"`
	} `json:"INFO"`
	RATE []struct {
		Unit   string   `json:"unit"`
		XAxis  []string `json:"xAxis"`
		Series []struct {
			Name string `json:"name"`
			Data []int  `json:"data"`
		} `json:"series"`
	} `json:"RATE"`
}

type SystemInfo struct {
	Minertype               string `json:"minertype"`
	Nettype                 string `json:"nettype"`
	Netdevice               string `json:"netdevice"`
	Macaddr                 string `json:"macaddr"`
	Hostname                string `json:"hostname"`
	Ipaddress               string `json:"ipaddress"`
	Netmask                 string `json:"netmask"`
	Gateway                 string `json:"gateway"`
	Dnsservers              string `json:"dnsservers"`
	SystemMode              string `json:"system_mode"`
	SystemKernelVersion     string `json:"system_kernel_version"`
	SystemFilesystemVersion string `json:"system_filesystem_version"`
	FirmwareType            string `json:"firmware_type"`
}

type NetworkInfo struct {
	Nettype        string `json:"nettype"`
	Netdevice      string `json:"netdevice"`
	Macaddr        string `json:"macaddr"`
	Ipaddress      string `json:"ipaddress"`
	Netmask        string `json:"netmask"`
	ConfNettype    string `json:"conf_nettype"`
	ConfHostname   string `json:"conf_hostname"`
	ConfIpaddress  string `json:"conf_ipaddress"`
	ConfNetmask    string `json:"conf_netmask"`
	ConfGateway    string `json:"conf_gateway"`
	ConfDnsservers string `json:"conf_dnsservers"`
}

type MinerConf struct {
	Pools []struct {
		Url  string `json:"url"`
		User string `json:"user"`
		Pass string `json:"pass"`
	} `json:"pools"`
	ApiListen        bool   `json:"api-listen"`
	ApiNetwork       bool   `json:"api-network"`
	ApiGroups        string `json:"api-groups"`
	ApiAllow         string `json:"api-allow"`
	BitmainFanCtrl   bool   `json:"bitmain-fan-ctrl"`
	BitmainFanPwm    string `json:"bitmain-fan-pwm"`
	BitmainUseVil    bool   `json:"bitmain-use-vil"`
	BitmainFreq      string `json:"bitmain-freq"`
	BitmainVoltage   string `json:"bitmain-voltage"`
	BitmainCcdelay   string `json:"bitmain-ccdelay"`
	BitmainPwth      string `json:"bitmain-pwth"`
	BitmainWorkMode  string `json:"bitmain-work-mode"`
	BitmainFreqLevel string `json:"bitmain-freq-level"`
}

type SetMinerConf struct {
	BitmainFanCtrl string `json:"bitmain-fan-ctrl,omitempty"`
	BitmainFanPwm  string `json:"bitmain-fan-pwm,omitempty"`
	MinerMode      string `json:"miner-mode,omitempty"`
	FreqLevel      string `json:"freq-level,omitempty"`
	Pools          []struct {
		Url  string `json:"url,omitempty"`
		User string `json:"user,omitempty"`
		Pass string `json:"pass,omitempty"`
	} `json:"pools,omitempty"`
}

type SetMinerConfResponse struct {
	Stats string `json:"stats"`
	Code  string `json:"code"`
	Msg   string `json:"msg"`
}

type Blink struct {
	Blink    string `json:"blink"`
	Response struct {
		Code    string `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
		Error   error  `json:"error,omitempty"`
	} `json:"response,omitempty"`
	BlinkStatus struct {
		Blink bool `json:"blink,omitempty"`
	}
}

type Logs struct {
	HistoryLog string `json:"hlog"`
	CurrentLog string `json:"log"`
}
