package tlsg10x_test

import (
	"testing"

	"github.com/0ghny/tlsg10x"
	testsupport "github.com/0ghny/tlsg10x/internal/test_support"
	"github.com/stretchr/testify/assert"
)

func Test_Login_Success(t *testing.T) {
	mockedWebsite := testsupport.NewMockedWebsite()
	httpClient := mockedWebsite.GetHttpClient()
	defer mockedWebsite.Disable()
	fakeSwitch := testsupport.NewFakeSwitch()

	client := tlsg10x.New(fakeSwitch.IPAddress,
		fakeSwitch.AdminUsr, fakeSwitch.AdminPwd, httpClient)
	err := client.Login()

	assert.Nil(t, err)
}

func Test_Login_Failed(t *testing.T) {
	mockedWebsite := testsupport.NewMockedWebsite()
	httpClient := mockedWebsite.GetHttpClient()
	defer mockedWebsite.Disable()
	fakeSwitch := testsupport.NewFakeSwitch()

	client := tlsg10x.New(fakeSwitch.IPAddress, "asdfasdf", "asdfasdf", httpClient)
	err := client.Login()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "invalid username or password")
}
