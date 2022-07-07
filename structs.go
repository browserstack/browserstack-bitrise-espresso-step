package main

type TestMapping struct {
	Name     string      `json:"name,omitempty"`
	Strategy string      `json:"strategy,omitempty"`
	Values   interface{} `json:"values,omitempty"`
}

type TestSharding struct {
	NumberOfShards      int           `json:"numberOfShards,omitempty"`
	Mapping             []TestMapping `json:"mapping,omitempty"`
	AutoStrategyDevices []string      `json:"devices,omitempty"`
}

type BrowserStackPayload struct {
	App                    string      `json:"app"`
	TestSuite              string      `json:"testSuite"`
	Devices                []string    `json:"devices"`
	InstrumentationLogs    bool        `json:"instrumentationLogs"`
	NetworkLogs            bool        `json:"networkLogs"`
	DeviceLogs             bool        `json:"deviceLogs"`
	DebugScreenshots       bool        `json:"debugscreenshots,omitempty"`
	VideoRecording         bool        `json:"video"`
	Project                string      `json:"project,omitempty"`
	ProjectNotifyURL       string      `json:"projectNotifyURL,omitempty"`
	UseLocal               bool        `json:"local,omitempty"`
	ClearAppData           bool        `json:"clearPackageData,omitempty"`
	SingleRunnerInvocation bool        `json:"singleRunnerInvocation,omitempty"`
	Class                  []string    `json:"class,omitempty"`
	Package                []string    `json:"package,omitempty"`
	Annotation             []string    `json:"annotation,omitempty"`
	Size                   []string    `json:"size,omitempty"`
	UseMockServer          bool        `json:"allowDeviceMockServer,omitempty"`
	UseTestSharding        interface{} `json:"shards,omitempty"`

	// Apart from the inputs from UI, these are some more fields which we support.
	// We've mentioned the type and the json key for these field.
	// We don't have separate inputs field for each of them,
	// instead we have one field which can accept all these values,
	// which we dynamically add to our payload with the help of a function `appendExtraCapabilities`.

	// EnableSpoonFramework  bool     `json:"enableSpoonFramework,omitempty"`
	// GpsLocation           string   `json:"gpsLocation,omitempty"`
	// GeoLocation           string   `json:"geoLocation,omitempty"`
	// CallbackURL           string   `json:"callbackURL,omitempty"`
	// NetworkProfile        string   `json:"networkProfile,omitempty"`
	// CustomNetwork         string   `json:"customNetwork,omitempty"`
	// Language              string   `json:"language,omitempty"`
	// Locale                string   `json:"locale,omitempty"`
	// AppStoreConfiguration string   `json:"appStoreConfiguration,omitempty"`
	// DeviceOrientation     string   `json:"deviceOrientation,omitempty"`
	// AcceptInsecureCerts   bool     `json:"acceptInsecureCerts,omitempty"`
	// UploadMedia           []string `json:"UploadMedia,omitempty"`
	// LocalIdentifier       string   `json:"localIdentifier,omitempty"`
	// IdleTimeout           string   `json:"idleTimeout,omitempty"`
}
