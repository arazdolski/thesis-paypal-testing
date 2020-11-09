package automation

import (
	"os"
	"testing"

	"github.com/subosito/gotenv"

	"github.com/sclevine/agouti"
	"github.com/stretchr/testify/suite"
)

type AutomationTestSuite struct {
	suite.Suite

	page *agouti.Page

	Driver *agouti.WebDriver
}

func (suite *AutomationTestSuite) SetupSuite() {

	gotenv.Load()

	suite.Driver = agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			// "--headless",
			"--allow-insecure-localhost",
			"--disable-gpu",
			"--homepage=about:blank",
			"--no-first-run",
			"--no-default-browser-check",
			"--no-sandbox",
		}),
		agouti.Debug,
	)

	if err := suite.Driver.Start(); err != nil {
		suite.FailNow(err.Error())
	}

	page, err := suite.Driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		suite.FailNow(err.Error())
	}

	suite.page = page

	if err := page.Navigate(os.Getenv("PAYPAL_URL")); err != nil {
		suite.FailNow(err.Error())
	}

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(AutomationTestSuite))
}
