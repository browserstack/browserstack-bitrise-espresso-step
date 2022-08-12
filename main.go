package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	// Exit after 2hours 30 mins
	time.AfterFunc(150*time.Minute, func() { failf("Session Timed Out") })

	username := os.Getenv("browserstack_username")
	access_key := os.Getenv("browserstack_accesskey")
	android_app := os.Getenv("android_app_under_test")
	test_suite := os.Getenv("espresso_test_suite")

	if username == "" || access_key == "" {
		failf(UPLOAD_APP_ERROR, "invalid credentials")
	}

	if android_app == "" {
		failf(AUT_NOT_FOUND)
	}

	if test_suite == "" {
		failf(TEST_SUITE_NOT_FOUND)
	}

	log.Print("Starting the build on BrowserStack App Automate")

	log.Print("Uploading app on BrowserStack App Automate")

	upload_app, err := upload(android_app, APP_UPLOAD_ENDPOINT, username, access_key)

	if err != nil {
		failf(err.Error())
	}

	upload_app_parsed_response := jsonParse(upload_app)

	if upload_app_parsed_response["app_url"] == "" {
		failf(err.Error())
	}

	app_url := upload_app_parsed_response["app_url"].(string)

	log.Print("Successfully uploaded the app")

	log.Print("Uploading test suite on BrowserStack App Automate")

	upload_test_suite, err := upload(test_suite, TEST_SUITE_UPLOAD_ENDPOINT, username, access_key)

	if err != nil {
		failf(err.Error())
	}

	test_suite_url := jsonParse(upload_test_suite)["test_suite_url"].(string)

	log.Print("Successfully uploaded the test suite")

	build_response, err := build(app_url, test_suite_url, username, access_key)

	if err != nil {
		failf(err.Error())
	}

	build_parsed_response := jsonParse(build_response)

	if build_parsed_response["message"] != "Success" {
		failf(BUILD_FAILED_ERROR, build_parsed_response["message"])
	}

	log.Print("Successfully started the build")

	check_build_status, _ := strconv.ParseBool(os.Getenv("check_build_status"))

	build_status := ""

	build_id := build_parsed_response["build_id"].(string)

	build_status, err = checkBuildStatus(build_id, username, access_key, check_build_status)

	if err != nil {
		failf(err.Error())
	}

	cmd_log_build_id, err_build_id := exec.Command("bitrise", "envman", "add", "--key", "BROWSERSTACK_BUILD_URL", "--value", APP_AUTOMATE_BUILD_DASHBOARD_URL+build_parsed_response["build_id"].(string)).CombinedOutput()
	cmd_log_build_status, err_build_status := exec.Command("bitrise", "envman", "add", "--key", "BROWSERSTACK_BUILD_STATUS", "--value", build_status).CombinedOutput()

	if err_build_id != nil {
		fmt.Printf("Failed to expose output with envman, error: %#v | output: %s", err, cmd_log_build_id)
		os.Exit(1)
	}

	if err_build_status != nil {
		fmt.Printf("Failed to expose output with envman, error: %#v | output: %s", err, cmd_log_build_status)
		os.Exit(1)
	}

	os.Exit(0)
}
