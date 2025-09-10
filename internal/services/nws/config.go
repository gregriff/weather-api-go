package nws

const (
	TestCoordsLat  = 35
	TestCoordsLong = -78

	GridpointURLParams = "%s/%d,%d"
	NWSBaseURL         = "https://api.weather.gov"
	PointsURL          = "/points/%s,%s"
	ForecastURL        = "/gridpoints/%s/forecast"
	HourlyForecastURL  = "/gridpoints/%s/forecast/hourly"
)

// python reference
// NWS_AUTH_HEADERS = {
//     "User-Agent": f"{CONFIG.nws.user_agent_identifier}, {CONFIG.nws.user_agent_email}"
// }
// TEST_COORDS = (CONFIG.nws.test_lat, CONFIG.nws.test_long)
