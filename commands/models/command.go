package models

type Command struct {
	CmdType string
	Args    []string
	Flags   map[string]string
}
