package config

type Entra struct {
	Enable   bool   `mapstructure:"enable"`
	TenantId string `mapstructure:"tenant-id"`
	ClientId string `mapstructure:"client-id"`
	Realm    string `mapstructure:"realm"`
	Mgmt     struct {
		SecretKey string `mapstructure:"secret-key"`
	} `mapstructure:"mgmt"`
}

var EntraConfig = new(Entra)
