package application

// Manifest contains a description of the app.
type Manifest struct {
	Name           string
	DisplayName    string
	ConfigDefaults []map[string]interface{}
}
