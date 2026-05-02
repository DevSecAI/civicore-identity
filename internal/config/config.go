// CIV-SAST-003: hardcoded OIDC client secret.
package config

type Config struct {
	Listen           string
	DBDsn            string
	OIDCClientID     string
	OIDCClientSecret string
}

func Load() Config {
	return Config{
		Listen:           ":8080",
		DBDsn:            "postgres://civicore:Civicore2024!@db:5432/civicore?sslmode=disable",
		OIDCClientID:     "civicore-prod",
		OIDCClientSecret: "oidc_secret_civicore_4M9XzQ1RyEXAMPLE",
	}
}
