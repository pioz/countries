# Countries

![Go](https://github.com/pioz/countries/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/pioz/countries)](https://goreportcard.com/report/github.com/pioz/countries)
[![codecov](https://codecov.io/gh/pioz/countries/branch/master/graph/badge.svg)](https://codecov.io/gh/pioz/countries)
[![awesome-go](https://camo.githubusercontent.com/2727609d8bfde9ba1a95be1449eb878bfafa4d76789ba05661857e2c8ac70fa1/68747470733a2f2f63646e2e7261776769742e636f6d2f73696e647265736f726875732f617765736f6d652f643733303566333864323966656437386661383536353265336136336531353464643865383832392f6d656469612f62616467652e737667)
[![GoReference](https://pkg.go.dev/badge/mod/github.com/pioz/countries)](https://pkg.go.dev/github.com/pioz/countries)

Countries is a port of [Ruby Countries](https://github.com/countries/countries) for Go.

- Standard ISO3166-1 (countries)
- Standard ISO3166-2 (states/subdivisions)
- Standard ISO4217 (currencies)
- Standard E.164 (phone numbers)
- Country Name Translations
- VAT Rates
- Address Formats
- Timezones

## Installation

    go get github.com/pioz/countries

## Usage

### Identification Codes

```go
c := countries.Get("US")
fmt.Println(c.Number)
fmt.Println(c.Alpha2)
fmt.Println(c.Alpha3)
fmt.Println(c.GEC)
// Output:
// 840
// US
// USA
// US
```

### Names & Translations

```go
c := countries.Get("US")
fmt.Println(c.ISOLongName)
fmt.Println(c.ISOShortName)
fmt.Println(c.UnofficialNames)
fmt.Println(c.Translations["en"])
fmt.Println(c.Translations["it"])
fmt.Println(c.Translations["de"])
fmt.Println(c.Nationality)
fmt.Println(c.Capital)
fmt.Println(c.EmojiFlag())
// Output:
// The United States of America
// United States of America
// [United States USA Vereinigte Staaten von Amerika États-Unis Estados Unidos アメリカ合衆国 Verenigde Staten Соединенные Штаты Америки]
// United States
// Stati Uniti
// Vereinigte Staaten
// American
// Washington
// 🇺🇸
```

### Subdivisions

```go
c := countries.Get("US")
ca := c.Subdivision("CA")
tx := c.SubdivisionByName("Texas")
fmt.Println(len(c.Subdivisions))
fmt.Println(ca.Name)
fmt.Println(ca.Type)
fmt.Println(ca.Translations["de"])
fmt.Println(ca.Geo.Latitude)
fmt.Println(tx.Code)
// Output:
// 57
// California
// state
// Kalifornien
// 36.778261
// TX
```

### Locations

```go
c := countries.Get("US")
fmt.Println(c.Geo.Latitude)
fmt.Println(c.Geo.Longitude)
fmt.Println(c.Region)
fmt.Println(c.Subregion)
fmt.Println(c.Continent)
fmt.Println(c.WorldRegion)
// Output:
// 37.09024
// -95.712891
// Americas
// Northern America
// North America
// AMER
```

### Boundary Boxes

```go
c := countries.Get("US")
fmt.Println(c.Geo.MinLatitude)
fmt.Println(c.Geo.MaxLatitude)
fmt.Println(c.Geo.MinLongitude)
fmt.Println(c.Geo.MaxLongitude)
fmt.Println(c.Geo.Bounds.Northeast.Lat)
fmt.Println(c.Geo.Bounds.Northeast.Lng)
fmt.Println(c.Geo.Bounds.Southwest.Lat)
fmt.Println(c.Geo.Bounds.Southwest.Lng)
// Output:
// 18.91619
// 71.3577635769
// -171.791110603
// -66.96466
// 71.3577635769
// -66.96466
// 18.91619
// -171.791110603
```

### Telephone Routing (E164)

```go
c := countries.Get("US")
fmt.Println(c.CountryCode)
fmt.Println(c.NationalDestinationCodeLengths)
fmt.Println(c.NationalNumberLengths)
fmt.Println(c.InternationalPrefix)
fmt.Println(c.NationalPrefix)
// Output:
// 1
// [3]
// [10]
// 011
// 1
```

### Timezones

```go
c := countries.Get("DE")
fmt.Println(c.Timezones)
// Output: [Europe/Berlin Europe/Busingen]
```

### Formatted Addresses

```go
c := countries.Get("US")
fmt.Println(c.AddressFormat)
fmt.Println("---")
fmt.Println(c.FormatAddress("John Smith", "1084 Nuzum Court", "14214", "Buffalo", "New York"))
// Output:
// {{recipient}}
// {{street}}
// {{city}} {{region_short}} {{postalcode}}
// {{country}}
// ---
// John Smith
// 1084 Nuzum Court
// Buffalo NY 14214
// United States of America
```

### VAT Rates

```go
c := countries.Get("IE")
fmt.Println(c.VatRates.Standard)
fmt.Println(c.VatRates.Reduced)
fmt.Println(c.VatRates.SuperReduced)
fmt.Println(c.VatRates.Parking)
// Output:
// 23
// [9 13]
// 4
// 13
```

### European Union Membership

```go
c := countries.Get("IT")
fmt.Println(c.EUMember)
// Output: true
```

### European Economic Area Membership

```go
c := countries.Get("FR")
fmt.Println(c.EEAMember)
// Output: true
```

### European Single Market Membership

```go
c := countries.Get("CH")
fmt.Println(c.ESMMember)
// Output: true
```

### GDPR Compliant

```go
c := countries.Get("IT")
fmt.Println(c.GDPRCompliant())
// Output: true
```

### Country Finders

```go
allCountries := countries.All
countriesInEurope := countries.InRegion("Europe")
countriesInSouthernAsia := countries.InSubregion("Southern Asia")
countriesInEU := countries.InEU()
fmt.Println(len(allCountries))
fmt.Println(len(countriesInEurope))
fmt.Println(len(countriesInSouthernAsia))
fmt.Println(len(countriesInEU))
// Output:
// 249
// 51
// 9
// 27
```

Please refer to the [godoc](https://godoc.org/github.com/pioz/countries) for all country fields, available functions and more.
Furthermore, tests are a good and helpful starting point.

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/pioz/countries/issues.

## License

The package is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
