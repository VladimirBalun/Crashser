package entities

type ProgrammingLanguage int

const (
	UnknownProgrammingLanguage ProgrammingLanguage = iota
	GoProgrammingLanguage
	CppProgrammingLanguage
	JavaProgrammingLanguage
)

type AppInfo struct {
	name                string
	version             string
	programmingLanguage ProgrammingLanguage
}

func NewAppInfo() *AppInfo {
	return &AppInfo{}
}

func (app *AppInfo) Name() string {
	return app.name
}

func (app *AppInfo) Version() string {
	return app.version
}

func (app *AppInfo) ProgrammingLanguage() ProgrammingLanguage {
	return app.programmingLanguage
}

func (app *AppInfo) SetName(name string) {
	app.name = name
}

func (app *AppInfo) SetVersion(version string) {
	app.version = version
}

func (app *AppInfo) SetProgrammingLanguage(language ProgrammingLanguage) {
	app.programmingLanguage = language
}
