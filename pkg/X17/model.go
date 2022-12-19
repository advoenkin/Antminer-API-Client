package X17

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
	Curtime                 string `json:"curtime"`
	Uptime                  string `json:"uptime"`
	Loadaverage             string `json:"loadaverage"`
	MemTotal                string `json:"mem_total"`
	MemUsed                 string `json:"mem_used"`
	MemFree                 string `json:"mem_free"`
	MemBuffers              string `json:"mem_buffers"`
	MemCached               string `json:"mem_cached"`
	SystemMode              string `json:"system_mode"`
	AntHwv                  string `json:"ant_hwv"`
	SystemKernelVersion     string `json:"system_kernel_version"`
	SystemFilesystemVersion string `json:"system_filesystem_version"`
	CgminerVersion          string `json:"cgminer_version"`
}
