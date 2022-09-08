package entities

import "time"

type CoreDumpStatus int

const (
	ToDoCoreDumpStatus CoreDumpStatus = iota
	InProgressCoreDumpStatus
	SolvedCoreDumpStatus
	RejectedCoreDumpStatus
)

type CoreDump struct {
	osInfo    *OSInfo
	appInfo   *AppInfo
	status    CoreDumpStatus
	data      string
	timestamp time.Time
}

func NewCoreDump() *CoreDump {
	return &CoreDump{}
}

func (c *CoreDump) OSInfo() *OSInfo {
	return c.osInfo
}

func (c *CoreDump) AppInfo() *AppInfo {
	return c.appInfo
}

func (c *CoreDump) Status() CoreDumpStatus {
	return c.status
}

func (c *CoreDump) Data() string {
	return c.data
}

func (c *CoreDump) Timestamp() time.Time {
	return c.timestamp
}

func (c *CoreDump) SetOSInfo(info *OSInfo) {
	c.osInfo = info
}

func (c *CoreDump) SetAppInfo(info *AppInfo) {
	c.appInfo = info
}

func (c *CoreDump) SetStatus(status CoreDumpStatus) {
	c.status = status
}

func (c *CoreDump) SetData(data string) {
	c.data = data
}

func (c *CoreDump) SetTimestamp(timestamp time.Time) {
	c.timestamp = timestamp
}
