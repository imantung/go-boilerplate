package dbkit

import (
	"testing"
)

func TestPgConfig(t *testing.T) {
	testcases := []struct {
		config     PgConfig
		connString string
	}{
		{
			PgConfig{DbName: "some-dbname", User: "some-user", Password: "with space"},
			"dbname=some-dbname user=some-user password=with\\ space ",
		},
		{
			PgConfig{Host: "some-host", Port: 99999, ConnectTimeout: 1000},
			"host=some-host port=99999 connect_timeout=1000 ",
		},
		{
			PgConfig{FallbackApplicationName: "some-application"},
			"fallback_application_name=some-application ",
		},
		{
			PgConfig{SslMode: "some-mode", SslKey: "some-key", SslCert: "some-cert", SslRootCert: "some-root-cert"},
			"sslmode=some-mode sslcert=some-cert sslkey=some-key sslrootcert=some-root-cert ",
		},
	}

	for _, tt := range testcases {
		connString := tt.config.ConnectionString()
		if connString != tt.connString {
			t.Fatalf("want '%s' got '%s'", tt.connString, connString)
		}
	}

}
