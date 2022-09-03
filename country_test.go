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
	assert.Equal(t, "Europe", c.Continent)
	assert.Equal(t, "39", c.CountryCode)
	assert.Equal(t, "EUR", c.CurrencyCode)
	assert.Equal(t, true, c.EEAMember)
	assert.Equal(t, true, c.EUMember)
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
	assert.Equal(t, true, c.PostalCode)
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
	us := countries.Get("US")
	fmt.Println(us.ISOShortName)
	// Output: United States of America
}

func ExampleGet_fields() {
	us := countries.Get("US")
	fmt.Println(us.ISOShortName)
	fmt.Println(us.Alpha2)
	fmt.Println(us.Region)
	fmt.Println(us.PostalCodeFormat)
	fmt.Println(us.StartOfWeek)
	fmt.Println(us.Subdivision("CA").Name)
	fmt.Println(us.Timezones[0])
	fmt.Println(us.EmojiFlag())
	// Output: United States of America
	// US
	// Americas
	// (\d{5})(?:[ \-](\d{4}))?
	// sunday
	// California
	// America/New_York
	// 🇺🇸
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
}

func TestInEU(t *testing.T) {
	cc := countries.InEU()
	assert.Equal(t, 27, len(cc))
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
	subdivision := it.Subdivision("VE")
	assert.Equal(t, "VE", subdivision.Code)
	assert.Equal(t, "Venezia", subdivision.Name)
	assert.Equal(t, 45.4408474, subdivision.Geo.Latitude)
	assert.Equal(t, "Venice", subdivision.Translations["en"])
	assert.Equal(t, "metropolitan_city", subdivision.Type)
}

func ExampleCountry_Subdivision() {
	us := countries.Get("US")
	ca := us.Subdivision("CA")
	fmt.Println(ca.Name)
	// Output: California
}

func TestSubdivisionByName(t *testing.T) {
	it := countries.Get("IT")
	subdivision := it.SubdivisionByName("Veneto")
	assert.Equal(t, "34", subdivision.Code)
	assert.Equal(t, "Veneto", subdivision.Name)
	subdivision = it.SubdivisionByName("xx")
	assert.Equal(t, "", subdivision.Name)
}

func ExampleCountry_SubdivisionByName() {
	us := countries.Get("US")
	ca := us.SubdivisionByName("Texas")
	fmt.Println(ca.Code)
	// Output: TX
}

func TestTranslation(t *testing.T) {
	it := countries.Get("IT")
	assert.Equal(t, "Italy", it.Translation("en"))
	assert.Equal(t, "Italia", it.Translation("it"))
	assert.Equal(t, "Italie", it.Translation("fr"))
	assert.Equal(t, "", it.Translation("xx"))
}

func ExampleCountry_Translation() {
	us := countries.Get("US")
	fmt.Println(us.Translation("fr"))
	fmt.Println(us.Translation("zh_CN"))
	// Output: États-Unis
	// 美国
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
	us := countries.Get("US")
	fmt.Println(us.FormatAddress("John Smith", "1084 Nuzum Court", "14214", "Buffalo", "New York"))
	// Output: John Smith
	// 1084 Nuzum Court
	// Buffalo NY 14214
	// United States of America
}

func ExampleCountry_EmojiFlag() {
	us := countries.Get("US")
	fmt.Println(us.EmojiFlag())
	// Output: 🇺🇸
}

func TestTimezones(t *testing.T) {
	us := countries.Get("DE")
	zones := us.Timezones
	assert.Equal(t, 2, len(zones))
	assert.Equal(t, "Europe/Berlin", zones[0])
	assert.Equal(t, "Europe/Busingen", zones[1])
}
