package usecase_test

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/usecase"
	"OverflowBackend/mocks"
	"io/fs"
	"os"
	"time"

	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"

	"bou.ke/monkey"
)

func init() {
	log.SetLevel(log.PanicLevel)
}

func InitTestUseCase(ctrl *gomock.Controller) (*mocks.MockDatabaseRepository, usecase.UseCase) {
	mockDB := mocks.NewMockDatabaseRepository(ctrl)
	uc := usecase.UseCase{}
	uc.Init(mockDB, config.TestConfig())
	testTime := time.Now()
	monkey.Patch(time.Now, func() time.Time {return testTime})
	monkey.Patch(os.WriteFile, func(name string, data []byte, perm fs.FileMode) error {return nil})
	monkey.Patch(os.MkdirAll, func(path string, perm fs.FileMode) error {return nil})
	return mockDB, uc
}