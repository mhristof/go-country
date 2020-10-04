package country

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindName(t *testing.T) {
	var cases = []struct {
		name          string
		country       string
		err           error
		expAlpha2Code string
	}{
		{
			name:          "find greece",
			country:       "greece",
			err:           nil,
			expAlpha2Code: "GR",
		},
		{
			name:    "error on multiple results",
			country: "united",
			err:     ErrorMultipleCountries,
		},
	}

	for _, test := range cases {
		country, err := FindName(test.country)
		assert.Equal(t, test.err, err, test.name)

		if test.expAlpha2Code != "" {
			assert.Equal(t, test.expAlpha2Code, country.Alpha2Code, test.name)
		}

	}
}
