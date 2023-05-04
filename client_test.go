package tlsg10x_test

import (
	"testing"

	"github.com/0ghny/tlsg10x"
	testsupport "github.com/0ghny/tlsg10x/internal/test_support"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	fakeSwitch := testsupport.NewFakeSwitch()
	client := tlsg10x.New(fakeSwitch.IPAddress, fakeSwitch.AdminUsr, fakeSwitch.AdminPwd, nil)

	assert.NotNil(t, client)
	assert.Equal(t, client.Host, fakeSwitch.IPAddress)
	assert.Equal(t, client.Username, fakeSwitch.AdminUsr)
	assert.Equal(t, client.Password, fakeSwitch.AdminPwd)
}
