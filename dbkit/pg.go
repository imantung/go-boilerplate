package dbkit

import (
	"strconv"
	"strings"
)

// PgConfig is clean way to generate connection string for postgres
type PgConfig struct {
	Host                    string
	Port                    int
	DbName                  string
	User                    string
	Password                string
	SslMode                 string
	FallbackApplicationName string
	ConnectTimeout          int
	SslCert                 string
	SslKey                  string
	SslRootCert             string
}

// ConnectionString return connection string based on https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters
func (c PgConfig) ConnectionString() string {
	var builder strings.Builder

	c.appendString(&builder, "host", c.Host)
	c.appendInt(&builder, "port", c.Port)
	c.appendString(&builder, "dbname", c.DbName)
	c.appendString(&builder, "user", c.User)
	c.appendString(&builder, "password", c.Password)
	c.appendString(&builder, "fallback_application_name", c.FallbackApplicationName)
	c.appendInt(&builder, "connect_timeout", c.ConnectTimeout)
	c.appendString(&builder, "sslmode", c.SslMode)
	c.appendString(&builder, "sslcert", c.SslCert)
	c.appendString(&builder, "sslkey", c.SslKey)
	c.appendString(&builder, "sslrootcert", c.SslRootCert)

	return builder.String()
}

func (c PgConfig) appendString(builder *strings.Builder, key, value string) bool {
	// NOTE: it will set to default value when empty by driver library
	if value != "" {
		builder.WriteString(key)
		builder.WriteString("=")
		builder.WriteString(strings.Replace(value, " ", "\\ ", -1))
		builder.WriteString(" ")
		return true
	}

	return false
}

func (c PgConfig) appendInt(builder *strings.Builder, key string, value int) bool {
	// NOTE: it will set to default value when 0 by driver library
	if value > 0 {
		builder.WriteString(key)
		builder.WriteString("=")
		builder.WriteString(strconv.Itoa(value))
		builder.WriteString(" ")
		return true
	}

	return false
}
