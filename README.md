# BrowserStack App Automate - Espresso

Run Espresso tests on BrowserStack

<details>
<summary>Description</summary>

Run your Espresso tests on BrowserStack App Automate. This step collects the built APK from `$BITRISE_APK_PATH` and test apk from `$BITRISE_TEST_APK_PATH` environment variables.

## Configure the Step

Before configuring this step, make sure you install [Bitrise CLI](https://github.com/bitrise-io/bitrise) in your machine.

Complete the following steps:

1. Clone the repository:
  ```bash
  git clone https://github.com/browserstack/browserstack-bitrise-espresso-step.git
  ```
2. Open the cloned repository and create a `.bitrise.secrets.yml` file at the same level of bitrise.yml to add your BrowserStack Username and Access Key.
  An example `.bitrise.secrets.yml` file is as follows:
  ```yml
    envs:
      - A_SECRET_PARAM_ONE: the value for secret one
      - A_SECRET_PARAM_TWO: the value for secret two
  ```
3. Go to the cloned repository and start the workflow editor:
  ```bash
  cd browserstack-bitrise-espresso-step
  bitrise :workflow-editor
  ```
4. Visit the workflow editor that starts on `http://localhost:50154/1.3.87/#!/workflows` by default.
5. On the workflow editor page, from the `WORKFLOW` drop-down, select `test`.
6. From the left navigation menu, click `Step Test`.
7. Provide values to the keys listed in the `Input variables` section. Check out the [configuration](#⚙️-configuration) section to learn about each key.
8. Save the configuration. You can now see the configuration in the `bitrise.yml` file created in the cloned repository.
9. Run the test using the following command:
  ```bash
  bitrise run test
  ```


## Troubleshooting

If you get the **Build already exists** error, it is because you have more than one instance of the Step in your Workflow. This doesn't work as Bitrise sends the build slug to Firebase and having the Step more than once in the same Workflow results in sending the same build slug multiple times.

## 🧩 Get started

Add this step directly to your workflow in the [Bitrise Workflow Editor](https://devcenter.bitrise.io/steps-and-workflows/steps-and-workflows-index/).

You can also run this step directly with [Bitrise CLI](https://github.com/bitrise-io/bitrise).

## ⚙️ Configuration

<details>
<summary>Inputs</summary>

| Key | Description | Flags | Default |
| --- | --- | --- | --- |
| `app_apk_path` | Path of the app (.apk) file. | required | `$BITRISE_APK_PATH` |
| `testsuite_apk_path` | Path of the test suite (.apk) file . | required | `$BITRISE_TEST_APK_PATH` |
| `devices` | Name of one or more device-OS combination in new line. For example: <br /> `Samsung Galaxy S9 Plus-9.0` <br />`Google Pixel 3a-9.0` | required | `Samsung Galaxy S9 Plus-9.0` |
| `instrumentation_logs` | Generate instrumentation logs of the test session  |  | `true` |
| `network_logs` | Generate network logs of your Espresso test sessions to capture network traffic, latency, etc. |  | `false` |
| `device_logs` | Generate device logs (Android logcat) |  | `false` |
| `debug_screenshots` | Capture the screenshots of the test execution|  | `false` |
| `video_recording` | Record video of the test execution  |  | `true` |
| `project` | Project name of the tests |  |  |
| `project_notify_url` | A callback URL to enable BrowserStack notify about completion of build under a given project.   |  |  |
| `use_local` | Enable local testing to retrieve app data hosted on local/private servers  |  | `false` |
| `use_test_sharding` | Enable test sharding to split tests cases into different groups instead of running them sequentially. <br />Add the sharding value json here. Examples: **Input for auto strategy**: <br /> ```{"numberOfShards": 2}, "devices": ["Google Pixel 3-9.0"]``` <br /> **Input for package strategy**:```{"numberOfShards": 2, "mapping": [{"name": "Shard 1", "strategy": "package", "values": ["com.foo.login", "com.foo.logout"]}, {"name": "Shard 2", "strategy": "package", "values": ["com.foo.dashboard"]}]}```  **Input for class strategy**: ```{"numberOfShards": 2, "mapping": [{"name": "Shard 1", "strategy": "class", "values": ["com.foo.login.user", "com.foo.login.admin"]}, {"name": "Shard 2", "strategy": "class", "values": ["com.foo.logout.user"]}]}```|  |  |
| `clear_app_data` | Enable to clear app data after every test run|  | `false`  |
| `filter_test` | "Key-value pairs of filters to run tests from supported test filtering strategies: class, package, annotation, size <br /> Examples: **For class filtering strategy**: `class com.android.foo.ClassA, class com.android.foo.ClassB,class com.android.foo.ClassC` <br /> **For package filtering strategy**: `package com.android.foo` <br /> **For annotation filtering strategy**: `size small`,`size medium`,`size large`  |  |  |
| `use_single_runner_invocation` | Enable to run all tests in a single instrumentation process to reduce overall build time.  |  | `false`  |
| `use_mock_server` | Enable to mock a web server in your espresso tests to mock your API responses. Learn more. |  | `false` |
| `check_build_status` | Wait for BrowserStack to complete the execution and get the test results  |  | `true` |
| `api_params` |"New line separated variables, key and value seperated by `=` For example: `coverage=true` <br />`geoLocation=CN"` |  |  |

</details>

<details>
<summary>Outputs</summary>

| Environment Variable | Description |
| --- | --- |
| `$BROWSERSTACK_BUILD_URL` |BrowserStack Dashboard url for the executed build|
| `$BROWSERSTACK_BUILD_STATUS`| Status of the executed build. Check out the [test results guide](https://www.browserstack.com/docs/app-automate/espresso/view-test-results) to learn about available status  |

</details>
