package tlsg10x

import (
	"fmt"
)

const (
	logoutUrl = "http://%s/Logout.htm"
)

// Logouts from site
func (t *Client) Logout() error {
	_, err := t.client.Get(fmt.Sprintf(logoutUrl, t.Host))

	if err != nil {
		return err
	}
	return nil
}
