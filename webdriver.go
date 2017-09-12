package visitor

import (
  "sourcegraph.com/sourcegraph/go-selenium"
  "log"
)

func createWebDriver() *selenium.WebDriver {
  caps := selenium.Capabilities(map[string]interface{}{
    "browserName": "chrome",
    "chromeOptions": map[string][]string {
      "args": {"--headless"},
    },
  })

  webDriver, err := selenium.NewRemote(caps, "http://localhost:9515")
  if err != nil {
    log.Fatal(err)
  }

  return &webDriver
}
