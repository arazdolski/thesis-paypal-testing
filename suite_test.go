package automation

import (
	"os"
	"testing"
	"time"

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

// private

func (suite *AutomationTestSuite) LoginToAccount() {
	if err := suite.page.Find("#ul-btn").Click(); err != nil {
		suite.Require().NoError(err)
	}
	
	if err := suite.page.Find("input[name='login_email']").Fill(os.Getenv("PAYPAL_USER_EMAIL")); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#btnNext").Click(); err != nil {
		suite.Require().NoError(err)
	}

	time.Sleep(time.Second)

	if err := suite.page.Find("input[name='login_password']").Fill(os.Getenv("PAYPAL_USER_PASSWORD")); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#btnLogin").Click(); err != nil {
		suite.Require().NoError(err)
	}
}
