package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CoreDumpStatus int

const (
	ToDoCoreDumpStatus CoreDumpStatus = iota
	InProgressCoreDumpStatus
	SolvedCoreDumpStatus
	RejectedCoreDumpStatus
)

type CoreDump struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	OsInfo    *OSInfo            `bson:"os_info,omitempty"`
	AppInfo   *AppInfo           `bson:"app_info,omitempty"`
	Status    CoreDumpStatus     `bson:"status,omitempty"`
	Data      string             `bson:"data,omitempty"`
	Timestamp time.Time          `bson:"timestamp,omitempty"`
}

func NewCoreDump() *CoreDump {
	return &CoreDump{}
}

func (c *CoreDump) GetOSInfo() *OSInfo {
	return c.OsInfo
}

func (c *CoreDump) GetAppInfo() *AppInfo {
	return c.AppInfo
}

func (c *CoreDump) GetStatus() CoreDumpStatus {
	return c.Status
}

func (c *CoreDump) GetData() string {
	return c.Data
}

func (c *CoreDump) GetTimestamp() time.Time {
	return c.Timestamp
}

func (c *CoreDump) SetOSInfo(info *OSInfo) {
	c.OsInfo = info
}

func (c *CoreDump) SetAppInfo(info *AppInfo) {
	c.AppInfo = info
}

func (c *CoreDump) SetStatus(status CoreDumpStatus) {
	c.Status = status
}

func (c *CoreDump) SetData(data string) {
	c.Data = data
}

func (c *CoreDump) SetTimestamp(timestamp time.Time) {
	c.Timestamp = timestamp
}
