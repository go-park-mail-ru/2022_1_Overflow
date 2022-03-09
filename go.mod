module OverflowBackend

require (
	general v1.0.0
	handlers v1.0.0
)

require (
	db v1.0.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.11.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.2.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.10.0 // indirect
	github.com/jackc/pgx/v4 v4.15.0 // indirect
	github.com/jackc/puddle v1.2.1 // indirect
	github.com/rs/cors v1.8.2 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/text v0.3.6 // indirect
)

replace db v1.0.0 => ./src/db

replace general v1.0.0 => ./src/general

replace handlers v1.0.0 => ./src/handlers

go 1.17
