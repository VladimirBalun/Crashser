package entities

type OSInfo struct {
	name         string
	version      string
	architecture string
}

func NewOSInfo() *OSInfo {
	return &OSInfo{}
}

func (os *OSInfo) Name() string {
	return os.name
}

func (os *OSInfo) Version() string {
	return os.version
}

func (os *OSInfo) Architecture() string {
	return os.architecture
}

func (os *OSInfo) SetName(name string) {
	os.name = name
}

func (os *OSInfo) SetVersion(version string) {
	os.version = version
}

func (os *OSInfo) SetArchitecture(architecture string) {
	os.architecture = architecture
}
