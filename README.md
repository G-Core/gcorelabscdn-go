# gcorelabscdn-go

![golangci-lint](https://github.com/G-Core/gcorelabscdn-go/actions/workflows/golangci-lint.yml/badge.svg)

`gcorelabscdn-go` is the G-Core CDN API SDK for the Go programming language.
 
* The GetCDNMetrics endoint returns varying responses which make unmarshalling it really difficult. For now we decided to just return the JSON strings instead. Because the Gcore SDK client also handles unmarshalling internally, we have to bypass the way Gcore SDK http layer, and build our own *http.Request instead.