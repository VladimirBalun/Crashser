package repositories

import (
	"server/sources/models"
)

func addCoreDump(coreDump models.CoreDump) bool {
	return true
}

func removeCoreDumpByUUID(uuid string) bool {
	return true
}

func changeCoreDumpStatusByUUID(uuid string, newStatus models.CoreDumpStatus) bool {
	return true
}