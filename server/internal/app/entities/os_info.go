package entities

type OSInfo struct {
	Name         string `bson:"name,omitempty"`
	Version      string `bson:"version,omitempty"`
	Architecture string `bson:"architecture,omitempty"`
}

func NewOSInfo() *OSInfo {
	return &OSInfo{}
}

func (os *OSInfo) GetName() string {
	return os.Name
}

func (os *OSInfo) GetVersion() string {
	return os.Version
}

func (os *OSInfo) GetArchitecture() string {
	return os.Architecture
}

func (os *OSInfo) SetName(name string) {
	os.Name = name
}

func (os *OSInfo) SetVersion(version string) {
	os.Version = version
}

func (os *OSInfo) SetArchitecture(architecture string) {
	os.Architecture = architecture
}
