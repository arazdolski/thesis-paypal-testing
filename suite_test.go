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

	suite.Driver = agouti.ChromeDriver()

	if err := suite.Driver.Start(); err != nil {
		suite.Require().NoError(err)
	}

	page, err := suite.Driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		suite.Require().NoError(err)
	}

	suite.page = page

	suite.page.SetImplicitWait(15000)

	if err := page.Navigate(os.Getenv("PAYPAL_URL")); err != nil {
		suite.Require().NoError(err)
	}
}

func (suite *AutomationTestSuite) TearDownTest() {
	suite.page.CloseWindow()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(AutomationTestSuite))
}
