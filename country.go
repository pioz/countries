//go:generate go run generator/main.go data

package countries

import (
	"regexp"
	"strings"
)

// Coord represents a geographic coordinate.
type Coord struct {
	Lat float64 `yaml:"lat"`
	Lng float64 `yaml:"lng"`
}

// Bounds represents a country bounds: the northeast and the southwest
// geographic coordinates.
type Bounds struct {
	Northeast Coord `yaml:"northeast"`
	Southwest Coord `yaml:"southwest"`
}

// Bounds store geographic informations of a country.
type Geo struct {
	Latitude     float64 `yaml:"latitude"`
	Longitude    float64 `yaml:"longitude"`
	MaxLatitude  float64 `yaml:"max_latitude"`
	MaxLongitude float64 `yaml:"max_longitude"`
	MinLatitude  float64 `yaml:"min_latitude"`
	MinLongitude float64 `yaml:"min_longitude"`
	Bounds       Bounds  `yaml:"bounds"`
}

// VatRates store the VAT (Value Added Tax) rates of a country.
type VatRates struct {
	Standard     int   `yaml:"standard"`
	Reduced      []int `yaml:"reduced"`
	SuperReduced int   `yaml:"super_reduced"`
	Parking      int   `yaml:"parking"`
}

// Country store all information about a country.
type Country struct {
	AddressFormat                  string                 `yaml:"address_format"`
	Alpha2                         string                 `yaml:"alpha2"`
	Alpha3                         string                 `yaml:"alpha3"`
	Capital                        string                 `yaml:"capital"`
	Continent                      string                 `yaml:"continent"`
	CountryCode                    string                 `yaml:"country_code"`
	CurrencyCode                   string                 `yaml:"currency_code"`
	EEAMember                      bool                   `yaml:"eea_member"`
	EUMember                       bool                   `yaml:"eu_member"`
	G7Member                       bool                   `yaml:"g7_member"`
	G20Member                      bool                   `yaml:"g20_member"`
	ESMMember                      bool                   `yaml:"esm_member"`
	GEC                            string                 `yaml:"gec"`
	Geo                            Geo                    `yaml:"geo"`
	InternationalPrefix            string                 `yaml:"international_prefix"`
	IOC                            string                 `yaml:"ioc"`
	ISOLongName                    string                 `yaml:"iso_long_name"`
	ISOShortName                   string                 `yaml:"iso_short_name"`
	LanguagesOfficial              []string               `yaml:"languages_official"`
	LanguagesSpoken                []string               `yaml:"languages_spoken"`
	NationalDestinationCodeLengths []int                  `yaml:"national_destination_code_lengths"`
	NationalNumberLengths          []int                  `yaml:"national_number_lengths"`
	NationalPrefix                 string                 `yaml:"national_prefix"`
	Nationality                    string                 `yaml:"nationality"`
	Number                         string                 `yaml:"number"`
	PostalCodeFormat               string                 `yaml:"postal_code_format"`
	Region                         string                 `yaml:"region"`
	StartOfWeek                    string                 `yaml:"start_of_week"`
	Subdivisions                   map[string]Subdivision `yaml:"-"`
	Subregion                      string                 `yaml:"subregion"`
	Timezones                      []string               `yaml:"-"`
	Translations                   map[string]string      `yaml:"-"`
	UnLocode                       string                 `yaml:"un_locode"`
	UnofficialNames                []string               `yaml:"unofficial_names"`
	VatRates                       VatRates               `yaml:"vat_rates"`
	WorldRegion                    string                 `yaml:"world_region"`
}

// Subdivision store information about a subdivision like a region or a province
// or a state or a metropolitan city of a country.
type Subdivision struct {
	Name         string            `yaml:"name"`
	Code         string            `yaml:"code"`
	Type         string            `yaml:"type"`
	Capital      bool              `yaml:"capital"`
	Geo          Geo               `yaml:"geo"`
	Translations map[string]string `yaml:"translations"`
}

// InEU returns all countries that are members of the European Union.
func InEU() []Country {
	result := make([]Country, 0)
	for _, c := range All {
		if c.EUMember {
			result = append(result, c)
		}
	}
	return result
}

// InRegion returns all countries that are part of the region.
func InRegion(region string) []Country {
	result := make([]Country, 0)
	for _, c := range All {
		if c.Region == region {
			result = append(result, c)
		}
	}
	return result
}

// InSubregion returns all countries that are part of the subregion.
func InSubregion(subregion string) []Country {
	result := make([]Country, 0)
	for _, c := range All {
		if c.Subregion == subregion {
			result = append(result, c)
		}
	}
	return result
}

// Subdivision returns the country's subdivision identified by code. If the code
// is not valid or not found returns a zero value Subdivision.
func (c *Country) Subdivision(code string) Subdivision {
	return c.Subdivisions[code]
}

// SubdivisionByName returns the country's subdivision with name name. If the
// name is not valid or not found returns a zero value Subdivision.
func (c *Country) SubdivisionByName(name string) Subdivision {
	for _, s := range c.Subdivisions {
		if s.Name == name {
			return s
		}
	}
	return Subdivision{}
}

// HasPostalCode determines whether the country has postal codes. It returns
// true if the country has postal codes, and false if it does not.
func (c *Country) HasPostalCode() bool {
	return c.PostalCodeFormat != ""
}

// MatchPostalCode returns true if postalCode has a valid format for the
// country. If the country does not have a postal code, returns false.
func (c *Country) MatchPostalCode(postalCode string) bool {
	if !c.HasPostalCode() {
		return false
	}
	r := regexp.MustCompile(c.PostalCodeFormat)
	return r.Match([]byte(postalCode))
}

// FormatAddress returns the formatted address based on country.AddressFormat
// template.
func (c *Country) FormatAddress(recipient, street, postalCode, city, region string) string {
	subdivision := c.Subdivision(region)
	if subdivision.Name == "" {
		subdivision = c.SubdivisionByName(region)
	}
	regionName := subdivision.Name
	regionShortName := subdivision.Code
	if regionName == "" {
		regionName = region
	}
	if regionShortName == "" {
		regionShortName = region
	}
	a := c.AddressFormat
	a = strings.ReplaceAll(a, "{{recipient}}", recipient)
	a = strings.ReplaceAll(a, "{{street}}", street)
	a = strings.ReplaceAll(a, "{{postalcode}}", postalCode)
	a = strings.ReplaceAll(a, "{{city}}", city)
	a = strings.ReplaceAll(a, "{{region}}", regionName)
	a = strings.ReplaceAll(a, "{{region_short}}", regionShortName)
	a = strings.ReplaceAll(a, "{{country}}", c.ISOShortName)
	return a
}

var flagsCodePoints = map[rune]rune{
	'a': '🇦',
	'b': '🇧',
	'c': '🇨',
	'd': '🇩',
	'e': '🇪',
	'f': '🇫',
	'g': '🇬',
	'h': '🇭',
	'i': '🇮',
	'j': '🇯',
	'k': '🇰',
	'l': '🇱',
	'm': '🇲',
	'n': '🇳',
	'o': '🇴',
	'p': '🇵',
	'q': '🇶',
	'r': '🇷',
	's': '🇸',
	't': '🇹',
	'u': '🇺',
	'v': '🇻',
	'w': '🇼',
	'x': '🇽',
	'y': '🇾',
	'z': '🇿',
}

// EmojiFlag returns the country Emoji flag.
func (c *Country) EmojiFlag() string {
	var flag []rune
	for _, r := range strings.ToLower(c.Alpha2) {
		flag = append(flag, flagsCodePoints[r])
	}
	return string(flag)
}
