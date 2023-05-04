package tlsg10x

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

const (
	logonUrl = "http://%s/logon.cgi"
)

// do a login into switch webpage
// it calls to http://<host>/logon.cgi
func (t *Client) Login() error {
	resp, err := t.client.
		PostForm(fmt.Sprintf(logonUrl, t.Host),
			url.Values{"username": {t.Username}, "password": {t.Password}, "logon": {"Login"}})
	// debug request? comment above, uncomment below
	// req, err := http.NewRequest("POST", fmt.Sprintf(logonUrl, t.Host), bytes.NewBuffer([]byte(url.Values{"username": {t.Username}, "password": {t.Password}, "logon": {"Login"}}.Encode())))
	// dRequest, err := httputil.DumpRequest(req, true)
	// fmt.Println(string(dRequest))
	if err != nil {
		return fmt.Errorf("error processing login request")
	}

	// Invalid login contains...in body
	// var logonInfo = new Array(
	// 	1,
	// 	0,0);
	defer resp.Body.Close()
	responseBody, _ := io.ReadAll(resp.Body)
	//responseBodyStr := string(responseBody)
	invalidLogin, err := regexp.Match(`logonInfo = new Array\(\n1,\n0,0\);`, responseBody)
	//invalidLogin := strings.Contains(responseBodyStr, "logonInfo = new Array(\n1,\n0,0)")

	if err != nil || resp.StatusCode != http.StatusOK || invalidLogin {
		return fmt.Errorf("invalid username or password")
	}

	return nil
}
