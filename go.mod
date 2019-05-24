module github.com/sylr/prometheus-azure-exporter

require (
	contrib.go.opencensus.io/exporter/ocagent v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go v24.1.0+incompatible
	github.com/Azure/azure-storage-blob-go v0.6.0
	github.com/Azure/go-autorest v11.9.0+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dimchansky/utfbom v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.4.7
	github.com/jessevdk/go-flags v1.4.0
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/patrickmn/go-cache v0.0.0-20180815053127-5633e0862627
	github.com/prometheus/client_golang v0.9.3
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	golang.org/x/crypto v0.0.0-20190513172903-22d7a77e9e5f // indirect
	gopkg.in/yaml.v2 v2.2.2
)

replace (
	github.com/patrickmn/go-cache => github.com/sylr/go-cache v2.1.1-0.20190112150453-7f6fb256aaca+incompatible
	github.com/prometheus/client_golang => github.com/sylr/prometheus-client-golang v0.0.0-20190106175946-16e6956cdb08
)
