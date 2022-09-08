package repositories

import "server/internal/app/entities"

type CoreDumpsRepository interface {
	GetCoreDumps() ([]entities.CoreDump, error)
	AddCoreDump(coreDump entities.CoreDump) error
	DeleteCoreDump(id string) error
}
