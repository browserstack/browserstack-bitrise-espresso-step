package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuild(t *testing.T) {
	t.Log("Build without app path")
	{
		build, err := build("", SAMPLE_TEST_SUITE, "username", "password")
		require.Equal(t, "", build)
		require.Error(t, err, AUT_NOT_FOUND)
	}

	t.Log("Build without test_suite_app path")
	{
		build, err := build(SAMPLE_APP, "", "", "")
		require.Equal(t, "", build)
		require.Error(t, err, TEST_SUITE_NOT_FOUND)
	}

	t.Log("Build with invalid credentials")
	{
		build, err := build(SAMPLE_APP, SAMPLE_TEST_SUITE, "a", "a")

		require.Equal(t, build, "{\"message\":\"Unauthorized\"}")

		require.NoError(t, err)
	}
}

func TestUpload(t *testing.T) {
	t.Log("It should throw file not found error with empty path")
	{
		build, err := upload("", APP_UPLOAD_ENDPOINT, "username", "password")
		t.Log(build, err)
		require.Equal(t, "", build)
		require.Error(t, err)
	}

	t.Log("It should throw file not found error with invalid path")
	{
		build, err := upload("invalidpath", APP_UPLOAD_ENDPOINT, "username", "password")
		t.Log(build, err)
		require.Equal(t, "", build)
		require.Error(t, err)
	}
}

func TestCheckBuildStatus(t *testing.T) {
	t.Log("It should throw FETCH_BUILD_STATUS_ERROR if build_id is not passed")
	{
		build, err := checkBuildStatus("", "username", "password", false)
		t.Log(build, err)
		require.Equal(t, "", build)
		require.Error(t, err)
	}
	t.Log("It should throw unauthorized error while checking build status")
	{
		expected := map[string]interface{}{"message": "unautorized"}
		build, err := checkBuildStatus(SAMPLE_BUILD_ID, "username", "password", false)

		require.Equal(t, "", build)
		require.Error(t, err, expected)
	}
}

func TestGetDevices(t *testing.T) {
	t.Log("It should return devices list")
	{
		expected := []string{"Samsung Galaxy S9 Plus-9.0", "Samsung Galaxy S10 Plus-10.0"}
		t.Setenv("devices_list", "Samsung Galaxy S9 Plus-9.0\nSamsung Galaxy S10 Plus-10.0")

		devices, _ := getDevices()

		require.Equal(t, expected, devices)
	}
	t.Log("It should throw error if devices not found in env")
	{
		t.Setenv("devices_list", "")
		_, err := getDevices()

		t.Log(err)
		require.Error(t, err)
	}
}

func TestParseNetworkLogsOptions(t *testing.T) {
	t.Log("It should treat an empty value as not set")
	{
		options, err := parseNetworkLogsOptions("")
		require.NoError(t, err)
		require.Nil(t, options)
	}
	t.Log("It should parse captureContent true")
	{
		options, err := parseNetworkLogsOptions(`{"captureContent": true}`)
		require.NoError(t, err)
		require.NotNil(t, options)
		require.True(t, options.CaptureContent)
	}
	t.Log("It should parse captureContent false")
	{
		options, err := parseNetworkLogsOptions(`{"captureContent": false}`)
		require.NoError(t, err)
		require.NotNil(t, options)
		require.False(t, options.CaptureContent)
	}
	t.Log("It should return an actionable error on malformed JSON")
	{
		_, err := parseNetworkLogsOptions(`{captureContent: true}`)
		require.Error(t, err)
		require.Contains(t, err.Error(), "network_logs_options")
	}
	t.Log("It should return an actionable error on a wrong value type")
	{
		_, err := parseNetworkLogsOptions(`{"captureContent": "true"}`)
		require.Error(t, err)
		require.Contains(t, err.Error(), "network_logs_options")
	}
}

func TestCreateBuildPayloadNetworkLogsOptions(t *testing.T) {
	t.Setenv("devices_list", "Google Pixel 6 Pro-15.0")

	t.Log("captureContent true -> networkLogs forced true + networkLogsOptions.captureContent true")
	{
		t.Setenv("network_logs", "false")
		t.Setenv("network_logs_options", `{"captureContent": true}`)

		payload := createBuildPayload()
		require.True(t, payload.NetworkLogs)
		require.NotNil(t, payload.NetworkLogsOptions)
		require.True(t, payload.NetworkLogsOptions.CaptureContent)

		marshalled, err := json.Marshal(payload)
		require.NoError(t, err)
		require.Contains(t, string(marshalled), `"networkLogs":true`)
		require.Contains(t, string(marshalled), `"networkLogsOptions":{"captureContent":true}`)
	}

	t.Log("options unset -> networkLogsOptions omitted (backward compatible)")
	{
		t.Setenv("network_logs", "true")
		t.Setenv("network_logs_options", "")

		payload := createBuildPayload()
		require.Nil(t, payload.NetworkLogsOptions)

		marshalled, err := json.Marshal(payload)
		require.NoError(t, err)
		require.NotContains(t, string(marshalled), "networkLogsOptions")
	}
}
