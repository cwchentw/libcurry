package libcurry

import (
	"github.com/pelletier/go-toml"
	"os"
)

func LoadConfigFile() (*toml.TomlTree, error) {
	var tree *toml.TomlTree
	rcPath := getRcPath()
	isInit := IsInit()
	if isInit {
		tree = toml.TreeFromMap(make(map[string]interface{}))
	} else {
		_tree, err := toml.LoadFile(rcPath)

		if err != nil {
			return _tree, err
		}

		tree = _tree
	}
	return tree, nil
}

func WriteConfigFile(tree *toml.TomlTree) error {
	rcPath := getRcPath()
	f, err := os.Create(rcPath)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(tree.ToString())

	return nil
}
