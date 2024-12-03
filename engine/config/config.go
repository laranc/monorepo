package config

import "github.com/yuin/gopher-lua"

type ConfigError struct{}

func (e *ConfigError) Error() string {
	return "Error parsing config file"
}

func LoadConfig(path string) (*lua.LTable, error) {
	l := lua.NewState()
	defer l.Close()
	if err := l.DoFile(path); err != nil {
		return nil, err
	}
	table := l.Get(-1)
	if tbl, ok := table.(*lua.LTable); ok {
		return tbl, nil
	}
	return nil, new(ConfigError)
}

func Get[T lua.LValue](table *lua.LTable, value string) T {
	return table.RawGetString(value).(T)
}
