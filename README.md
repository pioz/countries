# Countries

![Go](https://github.com/pioz/countries/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/pioz/countries)](https://goreportcard.com/report/github.com/pioz/countries)
[![codecov](https://codecov.io/gh/pioz/countries/branch/master/graph/badge.svg)](https://codecov.io/gh/pioz/countries)
[![GoReference](https://pkg.go.dev/badge/mod/github.com/pioz/countries)](https://pkg.go.dev/github.com/pioz/countries)

Countries is a port of [Ruby Countries](https://github.com/countries/countries) for GO.

- Standard ISO3166-1 (countries)
- Standard ISO3166-2 (states/subdivisions)
- Standard ISO4217 (currencies)
- Standard E.164 (phone numbers)
- Translation of country names
- VAT rates
- Address formats
- Timezones

## Installation

    go get github.com/pioz/countries

## Example

```go
us := countries.Get("US")
fmt.Println(us.ISOShortName)
fmt.Println(us.Alpha2)
fmt.Println(us.Region)
fmt.Println(us.PostalCodeFormat)
fmt.Println(us.StartOfWeek)
fmt.Println(us.Subdivision("CA").Name)
fmt.Println(us.EmojiFlag())
// Output: United States of America
// US
// Americas
// (\d{5})(?:[ \-](\d{4}))?
// sunday
// California
// 🇺🇸

// Get all countries in Europe
cc := countries.InRegion("Europe")
```

Please refer to the [godoc](https://godoc.org/github.com/pioz/countries) for all country fields, available functions and more.
Furthermore, tests are a good and helpful starting point.

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/pioz/countries/issues.

## License

The package is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
