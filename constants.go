package main

const (
	POOLING_INTERVAL_IN_MS             = 30000 // 30 secs
	BROWSERSTACK_DOMAIN                = "https://api-cloud.browserstack.com"
	APP_UPLOAD_ENDPOINT                = "/app-automate/espresso/v2/app"
	TEST_SUITE_UPLOAD_ENDPOINT         = "/app-automate/espresso/v2/test-suite"
	APP_AUTOMATE_BUILD_ENDPOINT        = "/app-automate/espresso/v2/build"
	APP_AUTOMATE_BUILD_STATUS_ENDPOINT = "/app-automate/espresso/v2/builds/"
	APP_AUTOMATE_BUILD_DASHBOARD_URL   = "https://app-automate.browserstack.com/dashboard/v2/builds/"

	SAMPLE_APP        = "bs://b91841adbf33515fef7a1cca869a9526a86f9a0e"
	SAMPLE_TEST_SUITE = "bs://535a0932c8a785384b8470ec6166e093cd3b2c5f"
	SAMPLE_BUILD_ID   = "56fec97937b22c785a6c5e08c13f629d505f5cd9"

	UPLOAD_APP_ERROR         = "Failed to upload app on BrowserStack, error : %s"
	FILE_NOT_AVAILABLE_ERROR = "Failed to upload test suite on BrowserStack, error: file not available"
	INVALID_FILE_TYPE_ERROR  = "Failed to upload test suite on BrowserStack, error: invalid file type"
	BUILD_FAILED_ERROR       = "Failed to execute build on BrowserStack, error: %s"
	FETCH_BUILD_STATUS_ERROR = "Failed to fetch test results, error: %s"
	HTTP_ERROR               = "Something went wrong while processing your request, error: %s"
)
