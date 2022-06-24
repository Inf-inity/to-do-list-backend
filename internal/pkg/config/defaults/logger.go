package defaults

// Logger contains the logger default values.
var Logger = map[string]interface{}{
	"level":             "info",
	"timeFormat":        "02.01.2006 15:04:05",
	"logDir":            "logs",
	"logDirMode":        0775,
	"logFileNameFormat": "2006-01-02",
	"logFileMode":       0664,
	"refreshSpec":       "@daily",
}
