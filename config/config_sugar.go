package config

import (
	"github.com/stack-labs/stack-rpc/pkg/config/reader"
)

var (
	_sugar Config
)

func Get(path ...string) reader.Value {
	return _sugar.Get(path...)
}

// todo, doest work
func ServerName() string {
	return _defaultCfg.Stack.Server.Name
}
