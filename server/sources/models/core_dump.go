package models

type CoreDumpStatus int

const (
	WAITING CoreDumpStatus = 0
	ACTIVE                 = 1
	FIXED                  = 2
	REJECTED               = 3
)

// CoreDump TODO: need to add JSON extensions
type CoreDump struct {
	osInfo OSInfo
	appInfo AppInfo
	status CoreDumpStatus
	coreDumpData string
	timestamp uint64
}
