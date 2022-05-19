# Company Legal Form

`go-company-legal-form` provides you with a simple, easy to use library to
remove or extract the legal form of a company.

## Usage

Install using `go get github.com/tilotech/go-company-legal-form`.

And then either use the default list:

```go
package main

import (
  "fmt"

  legalform "github.com/tilotech/go-company-legal-form"
)

func main() {
  companyName, legalForm := legalform.Default.Strip("Example Inc.")
  fmt.Println(companyName, legalForm)
}
```

Or create your own by creating a custom instance of `LegalForms`. Please note,
that the index values of `LegalForms` must be all lower case and
[some special characters](https://github.com/tilotech/go-company-legal-form/blob/4756e4973476350012a60f9b4facfee226266821/strip.go#L42)
must be removed.
As an example, using a custom instance gives you the possibility to limit the
recognition only to specific countries.

## Known Issues

* only recognizes legal forms that are at the end of the full company name,
  e.g. the legal form ("GmbH") of "Example GmbH Textilien + Angelware" is not
  recognized
* despite the huge list of supported legal forms, some legal forms might still
  be missing - please raise an issue with a company example for such cases