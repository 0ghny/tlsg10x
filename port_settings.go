package tlsg10x

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	portSettingsUrl = "http://%s/PortSettingRpm.htm"
)

// Ports Settings
type PortSettings struct {
	Name        string
	State       PortStatus
	TrunkInfo   int
	SpeedCfg    LinkStatus
	SpeedAct    LinkStatus
	FlowCtrlCfg OnOff
	FlowCtrlAct OnOff
}

// get port settings and info
func (t *Client) PortsSettings() ([]PortSettings, error) {
	content, err := t.getPage(fmt.Sprintf(portSettingsUrl, t.Host))
	if err != nil {
		return nil, err
	}

	// Total Ports
	total, err := t.getFirstValue(content, `(?m)var max_port_num = (\d)+`)
	if err != nil {
		return nil, err
	}
	totalPorts, err := strconv.Atoi(total)
	if err != nil {
		return nil, err
	}

	// state
	stateData, _ := t.getFirstValue(content, `state:[\s]?\[([0-9,]+)\],`)
	state := strings.Split(stateData, ",")
	// trunkinfo
	trunkInfoData, _ := t.getFirstValue(content, `trunk_info:[\s]?\[([0-9,]+)\],`)
	trunkInfo := strings.Split(trunkInfoData, ",")
	// spd_cfg
	spdCfgData, _ := t.getFirstValue(content, `spd_cfg:[\s]?\[([0-9,]+)\],`)
	spdCfg := strings.Split(spdCfgData, ",")
	// spd_act
	spdActData, _ := t.getFirstValue(content, `spd_act:[\s]?\[([0-9,]+)\],`)
	spdAct := strings.Split(spdActData, ",")
	// fc_cfg
	fcCfgData, _ := t.getFirstValue(content, `fc_cfg:[\s]?\[([0-9,]+)\],`)
	fcCfg := strings.Split(fcCfgData, ",")
	// fc_act
	fcActData, _ := t.getFirstValue(content, `fc_act:[\s]?\[([0-9,]+)\]`)
	fcAct := strings.Split(fcActData, ",")

	var portSettings []PortSettings
	for i := 0; i < totalPorts; i++ {
		var settings PortSettings

		// name
		settings.Name = fmt.Sprintf("Port%d", i+1)

		// status
		v, err := strconv.Atoi(state[i])
		if err != nil {
			return nil, fmt.Errorf("error converting state value %s to int", state[i])
		}
		settings.State = PortStatus(v)

		// trunk info
		v, err = strconv.Atoi(trunkInfo[i])
		if err != nil {
			return nil, fmt.Errorf("error converting trunk info value %s to int", trunkInfo[i])
		}
		settings.TrunkInfo = v

		// Speed Cfg
		v, err = strconv.Atoi(spdCfg[i])
		if err != nil {
			return nil, fmt.Errorf("error converting speed cfg value %s to int", state[i])
		}
		settings.SpeedCfg = LinkStatus(v)

		// Speed Act
		v, err = strconv.Atoi(spdAct[i])
		if err != nil {
			return nil, fmt.Errorf("error converting speed act value %s to int", state[i])
		}
		settings.SpeedAct = LinkStatus(v)

		// FlowControl Cfg
		v, err = strconv.Atoi(fcCfg[i])
		if err != nil {
			return nil, fmt.Errorf("error converting flow control cfg value %s to int", state[i])
		}
		settings.FlowCtrlCfg = OnOff(v)

		// FlowControl Act
		v, err = strconv.Atoi(fcAct[i])
		if err != nil {
			return nil, fmt.Errorf("error converting flow control act value %s to int", state[i])
		}
		settings.FlowCtrlAct = OnOff(v)

		portSettings = append(portSettings, settings)
	}

	return portSettings, nil
}
