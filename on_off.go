package tlsg10x

type OnOff int64

const (
	Off OnOff = 0
	On  OnOff = 1
)

func (s OnOff) String() string {
	switch s {
	case 0:
		return "Off"
	case 1:
		return "On"
	}
	return "unknown"
}
