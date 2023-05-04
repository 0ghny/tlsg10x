package testsupport

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/h2non/gock"
)

// A mocked website
type MockedWebsite struct {
	fSwitch *FakeSwitch
}

// creates a new mocked website
func NewMockedWebsite() *MockedWebsite {
	fakeSwitch := NewFakeSwitch()
	gock.New(fmt.Sprintf("http://%s", fakeSwitch.IPAddress)).
		Post("/logon.cgi").
		BodyString(
			url.Values{"username": {fakeSwitch.AdminUsr}, "password": {fakeSwitch.AdminPwd}, "logon": {"Login"}}.Encode(),
		).
		Reply(200)

	gock.New(fmt.Sprintf("http://%s", fakeSwitch.IPAddress)).
		Get("/SystemInfoRpm.htm").
		Reply(200).
		BodyString(CreateSystemInfoResponse(*fakeSwitch))

	gock.New(fmt.Sprintf("http://%s", fakeSwitch.IPAddress)).
		Get("/PortSettingRpm.htm").
		Reply(200).
		BodyString(CreatePortSettingResponse())

	gock.New(fmt.Sprintf("http://%s", fakeSwitch.IPAddress)).
		Get("/config_back.cgi?btnBackup=Backup+Config").
		Reply(200).
		BodyString(string(CreateBackupResponse()))

	gock.New(fmt.Sprintf("http://%s", fakeSwitch.IPAddress)).
		Get("PortStatisticsRpm.htm").
		Reply(200).
		BodyString(string(CreatePortStatsResponse()))

	return &MockedWebsite{
		fSwitch: fakeSwitch,
	}
}

func (t *MockedWebsite) Disable() {
	gock.Off()
}
func (t *MockedWebsite) InterceptClient(c *http.Client) {
	gock.InterceptClient(c)
}

func (t *MockedWebsite) GetHttpClient() *http.Client {
	client := http.DefaultClient
	t.InterceptClient(client)
	return client
}

func CreateSystemInfoResponse(s FakeSwitch) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<script>
var info_ds = {
descriStr:[
"%s"
],
macStr:[
"%s"
],
ipStr:[
"%s"
],
netmaskStr:[
"%s"
],
gatewayStr:[
"%s"
],
firmwareStr:[
"%s"
],
hardwareStr:[
"%s"
]
};`, s.Description,
		s.MacAddress,
		s.IPAddress,
		s.Netmask,
		s.Gateway,
		s.Firmware,
		s.Hardware)
}

func CreatePortSettingResponse() string {
	return `
<!DOCTYPE html>
<script>
var max_port_num = 8;
var port_middle_num  = 16;
var all_info = {
state:[1,1,1,1,1,1,1,1,0,0],
trunk_info:[0,0,0,0,0,0,0,0,0,0],
spd_cfg:[1,1,1,1,1,1,1,1,0,0],
spd_act:[6,0,0,6,6,6,0,0,0,0],
fc_cfg:[0,0,0,0,0,0,0,0,0,0],
fc_act:[0,0,0,0,0,0,0,0,0,0]
};
var tip = "";
</script>`
}

func CreateBackupResponse() []byte {
	return []byte("content of the backup")
}

func CreatePortStatsResponse() string {
	return `
var max_port_num = 8;
var port_middle_num  = 16;
var all_info = {
state:[1,1,1,1,1,1,1,1,0,0],
link_status:[6,0,0,6,6,6,0,0,0,0],
pkts:[186836605,0,183507103,13009519,0,0,0,0,0,0,0,0,61427943,0,71765069,0,17541146,0,15320967,0,151886175,0,129202625,0,0,0,0,0,0,0,0,0,0,0]
};`
}
