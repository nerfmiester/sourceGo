package Config

import "time"

// Cfg => config json
type Cfg struct {
	Enabled bool   `json:"enabled"`
	Path    string `json:"path"`
}

type Hoogie interface {
	Shootme() string
}

type Mawhat struct {
	State string
}

func NewMawhat() *Mawhat {
	m := new(Mawhat)
	m.State = "alive"
	return m
}

func (m Mawhat) Shootme() string {
	return "dead"
}

type Doohickie struct {
	State string
}

func NewDoohickie() *Doohickie {
	return &Doohickie{"passable"}
}

func (d Doohickie) Shootme() string {
	return "nearly dead"
}

type configuration struct {
	Enabled bool
	Path    string
}

type TomlConfig struct {
	Title   string
	Owner   ownerInfo
	DB      database `toml:"database"`
	Servers map[string]server
	Clients clients
}

type ownerInfo struct {
	Name string
	Org  string `toml:"organization"`
	Bio  string
	DOB  time.Time
}

type database struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type server struct {
	IP string
	DC string
}

type clients struct {
	Data  [][]interface{}
	Hosts []string
}
