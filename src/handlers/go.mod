module handlers

require (
	general v1.0.0
	db v1.0.0
)

replace general v1.0.0 => ../general
replace db v1.0.0 => ../db

require (
	github.com/gorilla/mux v1.8.0
)

go 1.17
