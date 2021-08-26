package constants

type TenantConfig struct {
	CollectionPrefix string
}

var TenantsConfig map[string]*TenantConfig

func init() {

	TenantsConfig = map[string]*TenantConfig{
		"localhost:8080": {
			CollectionPrefix: "", //local_
		},
	}
}
