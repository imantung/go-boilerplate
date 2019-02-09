# Go Helper

Helper library collection for go

### DBKit

Clean way to generate postgres connection string
```go
pgConfig := dbkit.PgConfig{
  User: "pqgotest",
  DbName: "pqgotest",
  SslMode: "verify-full",
}

pgConfig.ConnectionString() // return "user=pqgotest dbname=pqgotest sslmode=verify-full "
```
