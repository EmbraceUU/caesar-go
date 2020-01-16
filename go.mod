module caesar-go

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/net => github.com/golang/net v0.0.0-20190813141303-74dc4d7220e7
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190813064441-fde4db37ae7a
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190816200558-6889da9d5479
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20190717185122-a985d3407aa7
)

require (
	github.com/go-errors/errors v1.0.1 // indirect
	github.com/lunixbochs/vtclean v1.0.0 // indirect
	github.com/pkg/errors v0.9.0 // indirect
	github.com/shopspring/decimal v0.0.0-20191009025716-f1972eb1d1f5
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/viant/afs v0.12.0 // indirect
	github.com/viant/assertly v0.5.1 // indirect
	github.com/viant/toolbox v0.29.1
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20180412165947-fbb02b2291d2
	gopkg.in/yaml.v2 v2.2.7 // indirect
)

go 1.12
