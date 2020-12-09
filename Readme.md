# <img src="https://uploads-ssl.webflow.com/5ea5d3315186cf5ec60c3ee4/5edf1c94ce4c859f2b188094_logo.svg" alt="Pip.Services Logo" width="200"> <br/> Tokenizers, parsers and expression calculators Golang

This module is a part of the [Pip.Services](http://pip.services.org) polyglot microservices toolkit.
It provides syntax and lexical analyzers and expression calculator optimized for repeated calculations.

The module contains the following packages:
- [**Calculator**](https://godoc.org/github.com/pip-services3-go/pip-services3-expressions-go/calculator) - Expression calculator
- [**CSV**](https://godoc.org/github.com/pip-services3-go/pip-services3-expressions-go/csv) - CSV tokenizer
- [**IO**](https://godoc.org/github.com/pip-services3-go/pip-services3-expressions-go/io) - input/output utility classes to support lexical analysis
- [**Tokenizers**](https://godoc.org/github.com/pip-services3-go/pip-services3-expressions-go/tokenizers) - lexical analyzers to break incoming character streams into tokens
- [**Variants**](https://godoc.org/github.com/pip-services3-go/pip-services3-expressions-go/variants) - dynamic objects that can hold any values and operators for them

<a name="links"></a> Quick links:

* [API Reference](https://godoc.org/github.com/pip-services3-go/pip-services3-expressions-go/)
* [Change Log](CHANGELOG.md)
* [Get Help](https://www.pipservices.org/community/help)
* [Contribute](https://www.pipservices.org/community/contribute)

## Use

Get the package from the Github repository:
```bash
go get -u github.com/pip-services3-go/pip-services3-expressions-go@latest
```

## Develop

For development you shall install the following prerequisites:
* Golang v1.12+
* Visual Studio Code or another IDE of your choice
* Docker
* Git

Run automated tests:
```bash
go test -v ./test/...
```

Generate API documentation:
```bash
./docgen.ps1
```

Before committing changes run dockerized test as:
```bash
./test.ps1
./clear.ps1
```

## Contacts

The Golang version of Pip.Services is created and maintained by:
- **Sergey Seroukhov**
