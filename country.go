package country

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var (
	// ErrorMultipleCountries Error thrown when there are multiple countries matched
	ErrorMultipleCountries = errors.New("Multiple countries found for the search term")
)

// RestCountryResp Response struct from https://restcountries.eu/rest/v2/name/name
type RestCountryResp struct {
	Countries []RestCountry
}

// RestCountry A country representation from https://restcountries.eu/rest/v2/name/name
type RestCountry struct {
	Alpha2Code   string   `json:"alpha2Code"`
	Alpha3Code   string   `json:"alpha3Code"`
	AltSpellings []string `json:"altSpellings"`
	Area         float64  `json:"area"`
	Borders      []string `json:"borders"`
	CallingCodes []string `json:"callingCodes"`
	Capital      string   `json:"capital"`
	Cioc         string   `json:"cioc"`
	Currencies   []struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Demonym   string  `json:"demonym"`
	Flag      string  `json:"flag"`
	Gini      float64 `json:"gini"`
	Languages []struct {
		Iso639_1   string `json:"iso639_1"`
		Iso639_2   string `json:"iso639_2"`
		Name       string `json:"name"`
		NativeName string `json:"nativeName"`
	} `json:"languages"`
	Latlng        []float64 `json:"latlng"`
	Name          string    `json:"name"`
	NativeName    string    `json:"nativeName"`
	NumericCode   string    `json:"numericCode"`
	Population    int64     `json:"population"`
	Region        string    `json:"region"`
	RegionalBlocs []struct {
		Acronym       string   `json:"acronym"`
		Name          string   `json:"name"`
		OtherAcronyms []string `json:"otherAcronyms"`
		OtherNames    []string `json:"otherNames"`
	} `json:"regionalBlocs"`
	Subregion      string   `json:"subregion"`
	Timezones      []string `json:"timezones"`
	TopLevelDomain []string `json:"topLevelDomain"`
	Translations   struct {
		Br string `json:"br"`
		De string `json:"de"`
		Es string `json:"es"`
		Fa string `json:"fa"`
		Fr string `json:"fr"`
		Hr string `json:"hr"`
		It string `json:"it"`
		Ja string `json:"ja"`
		Nl string `json:"nl"`
		Pt string `json:"pt"`
	} `json:"translations"`
}

// FindName Returns one result from the URL https://restcountries.eu/rest/v2/name/name
// This function will error if there are multiple matches
func FindName(name string) (*RestCountry, error) {

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	resp, err := http.Get("https://restcountries.eu/rest/v2/name/" + name)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var results RestCountryResp
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyBytes, &results.Countries)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if len(results.Countries) > 1 {
		return nil, ErrorMultipleCountries
	}

	return &results.Countries[0], nil
}
