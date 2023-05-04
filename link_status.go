package tlsg10x

type LinkStatus int64

const (
	LinkDown LinkStatus = 0
	Auto     LinkStatus = 1
	At10MH   LinkStatus = 2
	At10MF   LinkStatus = 3
	At100MH  LinkStatus = 4
	At100MF  LinkStatus = 5
	At1000MF LinkStatus = 6
)

func (s LinkStatus) String() string {
	switch s {
	case 0:
		return "Link Down"
	case 1:
		return "Auto"
	case 2:
		return "10MH"
	case 3:
		return "10MF"
	case 4:
		return "100MH"
	case 5:
		return "100MF"
	case 6:
		return "1000MF"
	}
	return "unknown"
}
