module github.com/AlonaBot/alonabot-plugin-help

go 1.27.1

tool google.golang.org/protobuf/cmd/protoc-gen-go

require github.com/AlonaBot/AlonaBot v0.0.0-20261120143159-a9c9e4d5812e

require (
	github.com/google/jsonschema-go v0.4.3 // indirect
	golang.org/x/sys v0.49.0 // indirect
	google.golang.org/protobuf v1.36.12 // indirect
)

replace gorm.io/gorm => github.com/AlonaBot/gorm v1.31.3-0.20261120022453-34b98340007d
