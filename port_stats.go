package tlsg10x

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	portStatsUrl = "http://%s/PortStatisticsRpm.htm"
)

type PortStats struct {
	Name       string
	State      PortStatus
	LinkStatus LinkStatus
	TxGoodPkt  int
	TxBadPkt   int
	RxGoodPkt  int
	RxBadPkt   int
}

// get ports statistics
func (t *Client) PortsStats() ([]PortStats, error) {
	portsSettings, err := t.PortsSettings()
	if err != nil {
		return nil, err
	}

	totalPorts := len(portsSettings)
	content, err := t.getPage(fmt.Sprintf(portStatsUrl, t.Host))
	if err != nil {
		return nil, err
	}
	// state
	stateData, _ := t.getFirstValue(content, `state:[\s]?\[([0-9,]+)\],`)
	state := strings.Split(stateData, ",")
	// link_status
	linkStatusData, _ := t.getFirstValue(content, `link_status:[\s]?\[([0-9,]+)\],`)
	linkStatus := strings.Split(linkStatusData, ",")
	// pkts
	// pkts are in blocks of i, i+1, i+2, i+3
	pktsData, _ := t.getFirstValue(content, `pkts:[\s]?\[([0-9,]+)\]`)
	pkts := strings.Split(pktsData, ",")
	// last two entries are not measurements
	pkts = pkts[:len(pkts)-2]

	// Guard
	if len(state) <= 0 || len(linkStatus) <= 0 || len(pkts) <= 0 {
		return nil, fmt.Errorf("error getting port stats information, state, linkstatus or pkts are 0")
	}

	var allPortStats []PortStats
	for i := 0; i < totalPorts; i++ {
		var stats PortStats

		// Name

		stats.Name = fmt.Sprintf("Port%d", i+1)

		// status
		v, err := strconv.Atoi(state[i])
		if err != nil {
			return nil, fmt.Errorf("error converting state value %s to int", state[i])
		}
		stats.State = PortStatus(v)

		// link_status
		v, err = strconv.Atoi(linkStatus[i])
		if err != nil {
			return nil, fmt.Errorf("error converting link status value %s to int", linkStatus[i])
		}
		stats.LinkStatus = LinkStatus(v)
		index := i * 4
		// take care with edges...
		if index+3 > len(pkts) {
			return nil, fmt.Errorf("pkts not having enaugh elements (%d) to process current pkts in position %d", len(pkts), index)
		}
		// tx good
		v, err = strconv.Atoi(pkts[index])
		if err != nil {
			return nil, fmt.Errorf("error converting tx good value %s to int", pkts[index])
		}
		stats.TxGoodPkt = v
		// tx bad
		v, err = strconv.Atoi(pkts[index+1])
		if err != nil {
			return nil, fmt.Errorf("error converting tx bad value %s to int", pkts[index+1])
		}
		stats.TxBadPkt = v
		// rx good
		v, err = strconv.Atoi(pkts[index+2])
		if err != nil {
			return nil, fmt.Errorf("error converting rx good value %s to int", pkts[index+2])
		}
		stats.RxGoodPkt = v
		// rx bad
		v, err = strconv.Atoi(pkts[index+3])
		if err != nil {
			return nil, fmt.Errorf("error converting rx bad value %s to int", pkts[index+3])
		}
		stats.RxBadPkt = v

		allPortStats = append(allPortStats, stats)
	}

	return allPortStats, nil
}
