package tlsg10x

// Defines the api operations
type Api interface {
	PortsStats() ([]PortStats, error)
	PortsSettings() ([]PortSettings, error)
	CreateBackup() ([]byte, error)
	Login()
}
