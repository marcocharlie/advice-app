package config

import "time"

const (
	//AdviceSlipBaseURL indicate the url of the main external datasource for advices

	AdviceSlipBaseURL = "https://api.adviceslip.com/advice/search/"

	//CacheExpirationTime indicate how often data is refreshed from AdviceSlipBaseURL

	CacheExpirationTime time.Duration = 5 * time.Minute
)
