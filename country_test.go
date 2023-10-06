package countries_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/pioz/countries"
	"github.com/stretchr/testify/assert"
)

func TestNotExistingCountry(t *testing.T) {
	assert.Nil(t, countries.Get("XX"))
}

func TestCountry(t *testing.T) {
	c := countries.Get("IT")
	assert.Equal(t, "{{recipient}}\n{{street}}\n{{postalcode}} {{city}} {{region_short}}\n{{country}}", c.AddressFormat)
	assert.Equal(t, "IT", c.Alpha2)
	assert.Equal(t, "ITA", c.Alpha3)
	assert.Equal(t, "Rome", c.Capital)
	assert.Equal(t, "Europe", c.Continent)
	assert.Equal(t, "39", c.CountryCode)
	assert.Equal(t, "EUR", c.CurrencyCode)
	assert.Equal(t, true, c.EEAMember)
	assert.Equal(t, true, c.EUMember)
	assert.Equal(t, true, c.G7Member)
	assert.Equal(t, true, c.G20Member)
	assert.Equal(t, false, c.ESMMember)
	assert.Equal(t, "IT", c.GEC)
	assert.Equal(t, 41.87194, c.Geo.Latitude)
	assert.Equal(t, 12.56738, c.Geo.Longitude)
	assert.Equal(t, 47.092, c.Geo.MaxLatitude)
	assert.Equal(t, 18.7975999, c.Geo.MaxLongitude)
	assert.Equal(t, 35.4897, c.Geo.MinLatitude)
	assert.Equal(t, 6.6267201, c.Geo.MinLongitude)
	assert.Equal(t, 47.092, c.Geo.Bounds.Northeast.Lat)
	assert.Equal(t, 18.7975999, c.Geo.Bounds.Northeast.Lng)
	assert.Equal(t, 35.4897, c.Geo.Bounds.Southwest.Lat)
	assert.Equal(t, 6.6267201, c.Geo.Bounds.Southwest.Lng)
	assert.Equal(t, "00", c.InternationalPrefix)
	assert.Equal(t, "ITA", c.IOC)
	assert.Equal(t, "The Italian Republic", c.ISOLongName)
	assert.Equal(t, "Italy", c.ISOShortName)
	assert.Equal(t, []string{"it"}, c.LanguagesOfficial)
	assert.Equal(t, []string{"it"}, c.LanguagesSpoken)
	assert.Equal(t, []int{3}, c.NationalDestinationCodeLengths)
	assert.Equal(t, []int{9, 11}, c.NationalNumberLengths)
	assert.Equal(t, "None", c.NationalPrefix)
	assert.Equal(t, "Italian", c.Nationality)
	assert.Equal(t, "380", c.Number)
	assert.Equal(t, "\\d{5}", c.PostalCodeFormat)
	assert.Equal(t, "Europe", c.Region)
	assert.Equal(t, "monday", c.StartOfWeek)
	assert.Equal(t, 126, len(c.Subdivisions))
	assert.Equal(t, "Southern Europe", c.Subregion)
	assert.Equal(t, 1, len(c.Timezones))
	assert.Equal(t, "Europe/Rome", c.Timezones[0])
	assert.Equal(t, "IT", c.UnLocode)
	assert.Equal(t, []string{"Italy", "Italien", "Italie", "Italia", "イタリア", "Italië"}, c.UnofficialNames)
	assert.Equal(t, 22, c.VatRates.Standard)
	assert.Equal(t, []int{10}, c.VatRates.Reduced)
	assert.Equal(t, 4, c.VatRates.SuperReduced)
	assert.Equal(t, 0, c.VatRates.Parking)
	assert.Equal(t, "EMEA", c.WorldRegion)
}

func ExampleGet() {
	c := countries.Get("US")
	fmt.Println(c.ISOShortName)
	// Output: United States of America
}

func ExampleGet_readmeIdentificationCodes() {
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
}

func ExampleGet_readmeNamesAndTranslations() {
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
}

func ExampleGet_readmeSubdivisions() {
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
}

func ExampleGet_readmeLocations() {
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
}

func ExampleGet_readmeBoundaryBoxes() {
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
}

func ExampleGet_readmeTelephoneRouting() {
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
}

func ExampleGet_readmeTimezones() {
	c := countries.Get("DE")
	fmt.Println(c.Timezones)
	// Output: [Europe/Berlin Europe/Busingen]
}

func ExampleGet_readmeFormattedAddresses() {
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
}

func ExampleGet_readmeVATRates() {
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
}

func ExampleGet_readmeEuropeanUnionMembership() {
	c := countries.Get("IT")
	fmt.Println(c.EUMember)
	// Output: true
}

func ExampleGet_readmeEuropeanEconomicAreaMembership() {
	c := countries.Get("FR")
	fmt.Println(c.EEAMember)
	// Output: true
}

func ExampleGet_readmeEuropeanSingleMarketMembership() {
	c := countries.Get("CH")
	fmt.Println(c.ESMMember)
	// Output: true
}

func ExampleGet_readmeCountryFinders() {
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
	// 34
}

func TestAll(t *testing.T) {
	assert.Equal(t, 249, len(countries.All))
	assert.Equal(t, "AD", countries.All[0].Alpha2)
}

func TestAlpha2(t *testing.T) {
	a := countries.Alpha2
	assert.Equal(t, 249, len(a))
	assert.Equal(t, "AD", a[0])
	assert.Equal(t, "ZW", a[248])
}

func TestRegions(t *testing.T) {
	regions := countries.Regions
	assert.Equal(t, 5, len(regions))
	assert.Equal(t, "Africa", regions[0])
	assert.Equal(t, "Americas", regions[1])
	assert.Equal(t, "Asia", regions[2])
	assert.Equal(t, "Europe", regions[3])
	assert.Equal(t, "Oceania", regions[4])
}

func TestSubregions(t *testing.T) {
	subregions := countries.Subregions
	assert.Equal(t, 22, len(subregions))
	assert.Equal(t, "Australia and New Zealand", subregions[0])
	assert.Equal(t, "Western Europe", subregions[21])
}

func TestInEU(t *testing.T) {
	cc := countries.InEU()
	assert.Equal(t, 34, len(cc))
}

func TestInRegion(t *testing.T) {
	cc := countries.InRegion("Europe")
	assert.Equal(t, 51, len(cc))
	cc = countries.InRegion("Africa")
	assert.Equal(t, 60, len(cc))
	cc = countries.InRegion("")
	assert.Equal(t, 3, len(cc))
}

func TestInSubregion(t *testing.T) {
	cc := countries.InSubregion("Southern Europe")
	assert.Equal(t, 16, len(cc))
	cc = countries.InSubregion("Western Asia")
	assert.Equal(t, 18, len(cc))
	cc = countries.InSubregion("")
	assert.Equal(t, 3, len(cc))
}

func TestSubdivision(t *testing.T) {
	it := countries.Get("IT")
	subdivision := it.Subdivision("RM")
	assert.Equal(t, "RM", subdivision.Code)
	assert.Equal(t, "Roma", subdivision.Name)
	assert.Equal(t, 41.9027835, subdivision.Geo.Latitude)
	assert.Equal(t, 12.4963655, subdivision.Geo.Longitude)
	assert.Equal(t, "Rome", subdivision.Translations["en"])
	assert.Equal(t, "metropolitan_city", subdivision.Type)
	assert.True(t, subdivision.Capital)
}

func ExampleCountry_Subdivision() {
	c := countries.Get("US")
	ca := c.Subdivision("CA")
	fmt.Println(ca.Name)
	// Output: California
}

func TestSubdivisionByName(t *testing.T) {
	c := countries.Get("IT")
	subdivision := c.SubdivisionByName("Veneto")
	assert.Equal(t, "34", subdivision.Code)
	assert.Equal(t, "Veneto", subdivision.Name)
	subdivision = c.SubdivisionByName("xx")
	assert.Equal(t, "", subdivision.Name)
}

func ExampleCountry_SubdivisionByName() {
	c := countries.Get("US")
	ca := c.SubdivisionByName("Texas")
	fmt.Println(ca.Code)
	// Output: TX
}

func TestTranslations(t *testing.T) {
	c := countries.Get("IT")
	assert.Equal(t, "Italy", c.Translations["en"])
	assert.Equal(t, "Italia", c.Translations["it"])
	assert.Equal(t, "Italie", c.Translations["fr"])
	assert.Equal(t, "", c.Translations["xx"])
}

func TestHasPostalCode(t *testing.T) {
	it := countries.Get("IT")
	assert.True(t, it.HasPostalCode())

	jm := countries.Get("JM")
	assert.False(t, jm.HasPostalCode())
}

func TestMatchPostalCode(t *testing.T) {
	it := countries.Get("IT")
	assert.True(t, it.MatchPostalCode("35018"))
	assert.False(t, it.MatchPostalCode("3501F"))

	jm := countries.Get("JM")
	assert.False(t, jm.MatchPostalCode("35018"))

	for _, c := range countries.All {
		assert.NotPanics(t, func() {
			c.MatchPostalCode("35018")
		})
	}
}

func TestFormatAddress(t *testing.T) {
	it := countries.Get("IT")
	address := it.FormatAddress("Enrico Pilotto", "via Garibaldi 15", "97011", "Acate", "Ragusa")
	assert.Equal(t, "Enrico Pilotto\nvia Garibaldi 15\n97011 Acate RG\nItaly", address)
	address = it.FormatAddress("Enrico Pilotto", "via Garibaldi 15", "97011", "Acate", "RG")
	assert.Equal(t, "Enrico Pilotto\nvia Garibaldi 15\n97011 Acate RG\nItaly", address)
	address = it.FormatAddress("Enrico Pilotto", "via Garibaldi 15", "97011", "Acate", "xx")
	assert.Equal(t, "Enrico Pilotto\nvia Garibaldi 15\n97011 Acate xx\nItaly", address)

	es := countries.Get("ES")
	address = es.FormatAddress("Enrico Pilotto", "via Garibaldi 15", "97011", "Acate", "Andalucía")
	assert.Equal(t, "Enrico Pilotto\nvia Garibaldi 15\n97011 Acate\nAndalucía\nSpain", address)
	address = es.FormatAddress("Enrico Pilotto", "via Garibaldi 15", "97011", "Acate", "AN")
	assert.Equal(t, "Enrico Pilotto\nvia Garibaldi 15\n97011 Acate\nAndalucía\nSpain", address)
	address = es.FormatAddress("Enrico Pilotto", "via Garibaldi 15", "97011", "Acate", "xx")
	assert.Equal(t, "Enrico Pilotto\nvia Garibaldi 15\n97011 Acate\nxx\nSpain", address)

	for _, c := range countries.All {
		address := c.FormatAddress("Enrico Pilotto", "via Garibaldi 15", "97011", "Acate", "RG")
		assert.False(t, strings.Contains(address, "{{"), fmt.Sprintf("Invalid formatted address for country %s", c.Alpha2))
	}
}

func ExampleCountry_FormatAddress() {
	c := countries.Get("US")
	fmt.Println(c.FormatAddress("John Smith", "1084 Nuzum Court", "14214", "Buffalo", "New York"))
	// Output:
	// John Smith
	// 1084 Nuzum Court
	// Buffalo NY 14214
	// United States of America
}

func ExampleCountry_GDPRCompliant() {
	c := countries.Get("IT")
	fmt.Println(c.GDPRCompliant())
	// Output: true
}

func TestGDPRCompliant(t *testing.T) {
	c := countries.Get("IT")
	assert.True(t, c.GDPRCompliant())

	c = countries.Get("GB")
	assert.True(t, c.GDPRCompliant())

	c = countries.Get("US")
	assert.False(t, c.GDPRCompliant())
}

func ExampleCountry_EmojiFlag() {
	c := countries.Get("US")
	fmt.Println(c.EmojiFlag())
	// Output: 🇺🇸
}

func TestTimezones(t *testing.T) {
	c := countries.Get("DE")
	zones := c.Timezones
	assert.Equal(t, 2, len(zones))
	assert.Equal(t, "Europe/Berlin", zones[0])
	assert.Equal(t, "Europe/Busingen", zones[1])
}
