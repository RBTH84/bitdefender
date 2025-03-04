module github.com/RBTH84/bitdefender

go 1.20

require (
	github.com/elastic/go-elasticsearch/v8 v8.7.0 // Elasticsearch client
	github.com/gorilla/mux v1.8.1
	github.com/sirupsen/logrus v1.9.3
	github.com/urfave/cli v1.22.16
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.5 // indirect
	github.com/elastic/elastic-transport-go/v8 v8.2.0 // indirect
	github.com/malice-plugins/pkgs v1.1.7
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)

replace golang.org/x/sys => golang.org/x/sys v0.7.0
