package tlsg10x_test

import (
	"testing"

	"github.com/0ghny/tlsg10x"
	testsupport "github.com/0ghny/tlsg10x/internal/test_support"
	"github.com/stretchr/testify/assert"
)

func Test_SystemInformation_Retrieval(t *testing.T) {
	mockedWebsite := testsupport.NewMockedWebsite()
	httpClient := mockedWebsite.GetHttpClient()
	defer mockedWebsite.Disable()
	fakeSwitch := testsupport.NewFakeSwitch()

	client := tlsg10x.New(fakeSwitch.IPAddress, fakeSwitch.AdminUsr, fakeSwitch.AdminPwd, httpClient)
	sinfo, err := client.SystemInfo()

	assert.Nil(t, err)
	assert.NotNil(t, sinfo)
	assert.Equal(t, sinfo.Name, fakeSwitch.Description)
	assert.Equal(t, sinfo.MacAddress, fakeSwitch.MacAddress)
	assert.Equal(t, sinfo.IPAddress, fakeSwitch.IPAddress)
	assert.Equal(t, sinfo.NetMask, fakeSwitch.Netmask)
	assert.Equal(t, sinfo.Gateway, fakeSwitch.Gateway)
	assert.Equal(t, sinfo.Firmware, fakeSwitch.Firmware)
	assert.Equal(t, sinfo.Hardware, fakeSwitch.Hardware)
}
