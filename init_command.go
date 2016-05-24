package libcurry

import (
	"bufio"
	"errors"
	"fmt"
	"path"
	"os"
)

func InitSetting() error {
	configTree, err := loadConfigFile()
	if err != nil {
		return err
	}

	var appID string
	var dataPath string
	var _appID interface{} = ""
	var _dataPath interface{} = path.Join(getHomeDir(), ".currency")

	_appID = configTree.GetDefault("app_id", _appID).(string)
	_dataPath = configTree.GetDefault("data_dir", _dataPath).(string)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Input your app id [%s]: ", _appID.(string))
	if scanner.Scan() {
		appID = scanner.Text()

		if appID == "" && _appID.(string) == "" {
			return errors.New("Empty app id")
		}
	}

	fmt.Printf("Input your data directory [%s]: ", _dataPath.(string))
	if scanner.Scan() {
		dataPath = scanner.Text()
	}

	if appID != "" {
		_appID = appID
	}

	if dataPath != "" {
		_dataPath = dataPath
	}

	configTree.Set("app_id", _appID)
	configTree.Set("data_dir", _dataPath)

	err = writeConfigFile(configTree)
	if err != nil {
		return err
	}

	return nil
}
