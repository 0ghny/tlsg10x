package tlsg10x

type PortStatus int64

const (
	PortDisabled PortStatus = 0
	PortEnabled  PortStatus = 1
)

func (s PortStatus) String() string {
	switch s {
	case 1:
		return "Enabled"
	case 2:
		return "Disabled"
	}
	return "unknown"
}
