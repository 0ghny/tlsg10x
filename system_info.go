package tlsg10x

import "fmt"

const (
	systemInfoUrl = "http://%s/SystemInfoRpm.htm"
)

// System information
type SystemInfo struct {
	Name       string
	MacAddress string
	IPAddress  string
	NetMask    string
	Gateway    string
	Firmware   string
	Hardware   string
}

// retrieve switch system information
func (t *Client) SystemInfo() (*SystemInfo, error) {
	content, err := t.getPage(fmt.Sprintf(systemInfoUrl, t.Host))
	if err != nil {
		return nil, err
	}

	systemInfo := &SystemInfo{}

	// descriStr
	systemInfo.Name, _ = t.getFirstValue(content, `descriStr:\[[\n]?\"(.*)\"[\n]?\],`)
	// macStr
	systemInfo.MacAddress, _ = t.getFirstValue(content, `macStr:\[[\n]?\"(.*)\"[\n]?\],`)
	// ipStr
	systemInfo.IPAddress, _ = t.getFirstValue(content, `ipStr:\[[\n]?\"(.*)\"[\n]?\],`)
	// netmaskStr
	systemInfo.NetMask, _ = t.getFirstValue(content, `netmaskStr:\[[\n]?\"(.*)\"[\n]?\],`)
	// gatewayStr
	systemInfo.Gateway, _ = t.getFirstValue(content, `gatewayStr:\[[\n]?\"(.*)\"[\n]?\],`)
	// firmwareStr
	systemInfo.Firmware, _ = t.getFirstValue(content, `firmwareStr:\[[\n]?\"(.*)\"[\n]?\],`)
	// hardwareStr
	systemInfo.Hardware, _ = t.getFirstValue(content, `hardwareStr:\[[\n]?\"(.*)\"[\n]?\]`)

	return systemInfo, nil
}
