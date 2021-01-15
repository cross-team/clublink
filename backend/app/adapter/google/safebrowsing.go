package google

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/short-d/app/fw/webreq"
	"github.com/short-d/short/backend/app/usecase/risk"
)

var _ risk.BlackList = (*SafeBrowsing)(nil)

const safeBrowsingLookupAPI = "https://safebrowsing.googleapis.com/v4/threatMatches:find"

type threatType string

const (
	unknown               threatType = "THREAT_TYPE_UNSPECIFIED"
	malware               threatType = "MALWARE"
	socialEngineering     threatType = "SOCIAL_ENGINEERING"
	potentiallyHarmfulApp threatType = "POTENTIALLY_HARMFUL_APPLICATION"
	unwantedSoftware      threatType = "UNWANTED_SOFTWARE"
)

type platformType string

const (
	allPlatforms platformType = "ALL_PLATFORMS"
	anyPlatform  platformType = "ANY_PLATFORM"
	window       platformType = "WINDOWS"
	linux        platformType = "LINUX"
	osx          platformType = "OSX"
	chrome       platformType = "CHROME"
	ios          platformType = "IOS"
	android      platformType = "Android"
)

type threatEntryType string

const (
	unspecified  threatEntryType = "THREAT_ENTRY_TYPE_UNSPECIFIED"
	maliciousURL threatEntryType = "URL"
	executable   threatEntryType = "EXECUTABLE"
)

type threat struct {
	URL string `json:"url"`
}

type lookupAPIRequest struct {
	ThreatInfo threatInfo `json:"threatInfo"`
}

type threatInfo struct {
	ThreatTypes   []threatType   `json:"threatTypes"`
	PlatformTypes []platformType `json:"platformTypes"`
	ThreatEntries []threat       `json:"threatEntries"`
}

type lookupAPIResponse struct {
	Matches []match `json:"matches"`
}

type match struct {
	ThreatType      threatType      `json:"threatType"`
	PlatformType    platformType    `json:"platformType"`
	Threat          threat          `json:"threat"`
	CacheDuration   string          `json:"cacheDuration"`
	ThreatEntryType threatEntryType `json:"threatEntryType"`
}

// SafeBrowsing represents Googles SafeBrowsing APIs.
type SafeBrowsing struct {
	apiKey      string
	httpRequest webreq.HTTP
}

// HasURL checks whether a given URL is blacklisted by Google.
func (s SafeBrowsing) HasURL(url string) (bool, error) {
	api := s.auth(safeBrowsingLookupAPI)
	body := lookupAPIRequest{
		ThreatInfo: threatInfo{
			ThreatTypes: []threatType{
				malware,
				socialEngineering,
				potentiallyHarmfulApp,
				unwantedSoftware,
			},
			PlatformTypes: []platformType{
				allPlatforms,
			},
			ThreatEntries: []threat{
				{URL: url},
			},
		},
	}

	buf, err := json.Marshal(body)
	if err != nil {
		return false, err
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	res := lookupAPIResponse{}
	err = s.httpRequest.JSON(http.MethodPost, api, headers, string(buf), &res)
	if err != nil {
		return false, err
	}
	return len(res.Matches) > 0, nil
}

func (s SafeBrowsing) auth(baseURL string) string {
	return fmt.Sprintf("%s/?key=%s", baseURL, s.apiKey)
}

// NewSafeBrowsing initializes Google's SafeBrowsing API client.
func NewSafeBrowsing(apiKey string, req webreq.HTTP) SafeBrowsing {
	return SafeBrowsing{
		apiKey:      apiKey,
		httpRequest: req,
	}
}
