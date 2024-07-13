package meshmodel

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/layer5io/meshery/server/models"
	reg "github.com/layer5io/meshery/server/models/meshmodel/registration"
	"github.com/layer5io/meshkit/logger"
	meshmodel "github.com/layer5io/meshkit/models/meshmodel/registry"
)

var ModelsPath = "../meshmodel"

func GetModelDirectoryPaths() ([]string, error) {
	dirEntries := make([]string, 0)
	modelsDirs, err := os.ReadDir(ModelsPath)
	if(err != nil){
		return dirEntries, err
	}
	for _, modelDir := range modelsDirs {
		modelVersionsDirPath := filepath.Join(ModelsPath, modelDir.Name())
		modelVersionsDir, err := os.ReadDir(modelVersionsDirPath)
		if(err != nil){
			return dirEntries, err
		}
		for _, version := range modelVersionsDir {
			modelDefVersionsDirPath := filepath.Join(modelVersionsDirPath, version.Name())
			modelDefVersionsDir, err := os.ReadDir(modelDefVersionsDirPath)
			if(err != nil){
				return dirEntries, err
			}
			for _, defVersion := range modelDefVersionsDir {
				defPath := filepath.Join(modelDefVersionsDirPath, defVersion.Name())
				dirEntries = append(dirEntries, defPath)
			}
		}
	}
	return dirEntries, nil
}

// seed the local meshmodel components
func SeedComponents(log logger.Handler, hc *models.HandlerConfig, regm *meshmodel.RegistryManager) {
	regHelper := reg.NewRegistrationHelper(log, hc, regm)
	modelDirPaths, err := GetModelDirectoryPaths()
	if(err != nil){
		// handle error
		fmt.Println(err)
	}
	for _, dirPath := range modelDirPaths {
		dir := 	reg.NewDir(dirPath)
		err := regHelper.Register(dir)
		if err != nil {
			// handle error
			fmt.Println(err)
		}
	}
	regHelper.RegistryLog()
}