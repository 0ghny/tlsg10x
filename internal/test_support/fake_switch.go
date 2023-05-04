package testsupport

type FakeSwitch struct {
	AdminUsr    string
	AdminPwd    string
	Description string
	MacAddress  string
	IPAddress   string
	Netmask     string
	Gateway     string
	Firmware    string
	Hardware    string
}

// creates a new fake switch with default data
func NewFakeSwitch() *FakeSwitch {
	return &FakeSwitch{
		Description: "fake switch",
		MacAddress:  "11:11:11:11:11:11",
		IPAddress:   "192.168.1.2",
		Netmask:     "255.255.255.0",
		Gateway:     "192.168.1.1",
		Firmware:    "1.0.0 Build 20230218 Rel.50633",
		Hardware:    "TL-SG108E 6.0",
		AdminUsr:    "admin",
		AdminPwd:    "admin123",
	}
}
