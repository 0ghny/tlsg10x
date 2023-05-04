package tlsg10x

import "fmt"

const (
	createBackupUrl = "http://%s/config_back.cgi?btnBackup=Backup+Config"
)

// creates a backup of the connected device
func (t *Client) CreateBackup() ([]byte, error) {
	return t.getBytes(fmt.Sprintf(createBackupUrl, t.Host))
}
