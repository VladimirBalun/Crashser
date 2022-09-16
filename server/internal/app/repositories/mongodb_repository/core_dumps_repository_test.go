package mongodb_repository

import (
	"errors"
	"math/rand"
	"server/internal/app/entities"
	mock_repositories "server/internal/app/repositories/mock"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_AddCoreDump(t *testing.T) {
	t.Parallel()
	c := gomock.NewController(t)
	defer c.Finish()

	randomDump := generateRandomSliceOfCoreDumps(1)
	tests := []struct {
		name  string
		dump  entities.CoreDump
		stubs func(store *mock_repositories.MockCoreDumpsRepository, dump entities.CoreDump)
		error error
	}{
		{
			name: "ok",
			dump: randomDump[0],
			stubs: func(r *mock_repositories.MockCoreDumpsRepository, dump entities.CoreDump) {
				r.EXPECT().AddCoreDump(dump).Return(nil)
			},
			error: nil,
		},
		{
			name: "error",
			dump: entities.CoreDump{},
			stubs: func(r *mock_repositories.MockCoreDumpsRepository, dump entities.CoreDump) {
				r.EXPECT().AddCoreDump(dump).Return(errors.New("error adding core dump"))

			},
			error: errors.New("error adding core dump"),
		},
		{
			name: "timeout",
			dump: entities.CoreDump{},
			stubs: func(r *mock_repositories.MockCoreDumpsRepository, dump entities.CoreDump) {
				r.EXPECT().AddCoreDump(dump).Return(errors.New("timeout"))

			},
			error: errors.New("timeout"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := mock_repositories.NewMockCoreDumpsRepository(c)

			test.stubs(r, test.dump)
			err := r.AddCoreDump(test.dump)
			require.Equal(t, test.error, err)
		})
	}
}

func Test_GetCoreDump(t *testing.T) {
	t.Parallel()
	c := gomock.NewController(t)
	defer c.Finish()

	sliceOfDumps := generateRandomSliceOfCoreDumps(5)

	tests := []struct {
		name  string
		stubs func(store *mock_repositories.MockCoreDumpsRepository)
		dumps []entities.CoreDump
		error error
	}{
		{
			name: "ok",
			stubs: func(r *mock_repositories.MockCoreDumpsRepository) {
				r.EXPECT().GetCoreDumps().Return(sliceOfDumps, nil)
			},
			dumps: sliceOfDumps,
			error: nil,
		},
		{
			name: "error",
			stubs: func(r *mock_repositories.MockCoreDumpsRepository) {
				r.EXPECT().GetCoreDumps().Return(nil, errors.New("error getting dumps"))
			},
			dumps: nil,
			error: errors.New("error getting dumps"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := mock_repositories.NewMockCoreDumpsRepository(c)
			test.stubs(r)

			res, err := r.GetCoreDumps()
			require.Equal(t, test.error, err)
			require.Equal(t, test.dumps, res)
		})
	}
}

func Test_DeleteCoreDump(t *testing.T) {
	t.Parallel()
	c := gomock.NewController(t)
	defer c.Finish()

	tests := []struct {
		name   string
		stubs  func(store *mock_repositories.MockCoreDumpsRepository, dumpID string)
		dumpID string
		error  error
	}{
		{
			name: "ok",
			stubs: func(r *mock_repositories.MockCoreDumpsRepository, dumpID string) {
				r.EXPECT().DeleteCoreDump(dumpID).Return(nil)
			},
			dumpID: "dsfdfds454",
			error:  nil,
		},
		{
			name: "error",
			stubs: func(r *mock_repositories.MockCoreDumpsRepository, dumpID string) {
				r.EXPECT().DeleteCoreDump(dumpID).Return(errors.New("error enter dump id"))
			},
			dumpID: "",
			error:  errors.New("error enter dump id"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := mock_repositories.NewMockCoreDumpsRepository(c)
			test.stubs(r, test.dumpID)

			err := r.DeleteCoreDump(test.dumpID)
			require.Equal(t, test.error, err)
		})
	}
}

func init() {
	rand.Seed(time.Now().Unix())
}

func generateRandomSliceOfCoreDumps(quantity int) []entities.CoreDump {
	type asInfo struct {
		Name    string
		Arch    string
		Version string
	}

	type appInfo struct {
		Name    string
		Version string
	}
	osArr := []asInfo{
		{
			Name:    "linux",
			Arch:    "amd64",
			Version: "ubuntu 22.04",
		},
		{
			Name:    "linux",
			Arch:    "amd64",
			Version: "ubuntu 18.06",
		},
		{
			Name:    "windows",
			Arch:    "amd64",
			Version: "10",
		},
		{
			Name:    "darwin",
			Arch:    "amd64",
			Version: "10.0.3",
		},
		{
			Name:    "darwin",
			Arch:    "amd64",
			Version: "10.0.1",
		},
	}
	appArr := []appInfo{
		{
			Name:    "financial",
			Version: "v0.0.1",
		},
		{
			Name:    "financial",
			Version: "v1.0.1",
		},
		{
			Name:    "sports",
			Version: "v2.1.1",
		},
		{
			Name:    "sports",
			Version: "v1.1.1",
		},
		{
			Name:    "educational",
			Version: "v0.1.1",
		},
		{
			Name:    "educational",
			Version: "v3.1.1",
		},
	}
	var result []entities.CoreDump
	for i := 0; i < quantity; i++ {
		coreDump := entities.NewCoreDump()
		randomIdxOs := rand.Intn(len(osArr))
		randomIdxApp := rand.Intn(len(appArr))

		osInfo := entities.NewOSInfo()
		osInfo.SetName(osArr[randomIdxOs].Name)
		osInfo.SetArchitecture(osArr[randomIdxOs].Arch)
		osInfo.SetVersion(osArr[randomIdxOs].Version)
		coreDump.SetOSInfo(osInfo)

		appInfo := entities.NewAppInfo()
		appInfo.SetName(appArr[randomIdxApp].Name)
		appInfo.SetProgrammingLanguage(entities.ProgrammingLanguage(rand.Intn(4)))
		appInfo.SetVersion(appArr[randomIdxApp].Version)
		coreDump.SetAppInfo(appInfo)

		coreDump.SetStatus(1)
		coreDump.SetData(time.Now().Format("2006-01-02"))

		randomTime := rand.Int63n(time.Now().Unix()-94608000) + 94608000
		randomNow := time.Unix(randomTime, 0)
		coreDump.SetTimestamp(randomNow)

		result = append(result, *coreDump)
	}

	return result
}
