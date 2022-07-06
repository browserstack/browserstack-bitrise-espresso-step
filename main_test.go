package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuild(t *testing.T) {
	t.Log("Build without app path")
	{
		build, err := build("", SAMPLE_TEST_SUITE, "username", "password")
		require.Equal(t, "", build)
		require.Error(t, err, FILE_NOT_AVAILABLE_ERROR)
	}

	t.Log("Build without test_suite_app path")
	{
		build, err := build(SAMPLE_APP, "", "", "")
		require.Equal(t, "", build)
		require.Error(t, err, FILE_NOT_AVAILABLE_ERROR)
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
