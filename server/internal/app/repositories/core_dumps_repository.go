package repositories

import (
	"server/internal/app/entities"
)

type CoreDumpsRepository interface {
	GetCoreDumps(setters ...entities.OptionsMongo) ([]entities.CoreDump, error)
	AddCoreDump(coreDump entities.CoreDump) error
	DeleteCoreDump(id string) error
	DeleteAllCoreDumps() error
}
