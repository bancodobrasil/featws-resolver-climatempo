module github.com/bancodobrasil/featws-resolver-climatempo

go 1.16

require (
	github.com/bancodobrasil/featws-resolver-adapter-go v0.0.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.10.1
)

replace github.com/bancodobrasil/featws-resolver-adapter-go => ../featws-resolver-adapter-go
