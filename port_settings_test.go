package tlsg10x_test

import (
	"testing"

	"github.com/0ghny/tlsg10x"
	testsupport "github.com/0ghny/tlsg10x/internal/test_support"
	"github.com/stretchr/testify/assert"
)

func Test_PortSettings_Retrieval(t *testing.T) {
	mockedWebsite := testsupport.NewMockedWebsite()
	httpClient := mockedWebsite.GetHttpClient()
	defer mockedWebsite.Disable()
	fakeSwitch := testsupport.NewFakeSwitch()

	client := tlsg10x.New(fakeSwitch.IPAddress, fakeSwitch.AdminUsr, fakeSwitch.AdminPwd, httpClient)
	pSettings, err := client.PortsSettings()

	assert.Nil(t, err)
	assert.NotNil(t, pSettings)
	assert.Equal(t, len(pSettings), 8)
	assert.Equal(t, pSettings[0].Name, "Port1")
	assert.Equal(t, pSettings[0].State.String(), "Enabled")
	assert.Equal(t, pSettings[0].SpeedCfg.String(), "Auto")
}
