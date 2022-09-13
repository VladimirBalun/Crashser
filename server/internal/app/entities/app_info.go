package entities

type ProgrammingLanguage int

const (
	UnknownProgrammingLanguage ProgrammingLanguage = iota
	GoProgrammingLanguage
	CppProgrammingLanguage
	JavaProgrammingLanguage
)

type AppInfo struct {
	Name                string              `bson:"name,omitempty"`
	Version             string              `bson:"version,omitempty"`
	ProgrammingLanguage ProgrammingLanguage `bson:"programming_language,omitempty"`
}

func NewAppInfo() *AppInfo {
	return &AppInfo{}
}

func (app *AppInfo) GetName() string {
	return app.Name
}

func (app *AppInfo) GetVersion() string {
	return app.Version
}

func (app *AppInfo) GetProgrammingLanguage() ProgrammingLanguage {
	return app.ProgrammingLanguage
}

func (app *AppInfo) SetName(name string) {
	app.Name = name
}

func (app *AppInfo) SetVersion(version string) {
	app.Version = version
}

func (app *AppInfo) SetProgrammingLanguage(language ProgrammingLanguage) {
	app.ProgrammingLanguage = language
}
