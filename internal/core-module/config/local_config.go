package config

type TarantoolConfig struct {
	Address  string `json:"addr"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type LocalConfig struct {
	GrpcAddr        string          `json:"core-addr"`
	TarantoolConfig TarantoolConfig `json:"tarantool"`
}
