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

	// Address
	address, city, postalCode, state string

	// Bank
	IBAN, currency, bankName string
	
	// Data
	email, firstname, lastname, item, itemPrice, mobile, amount string
	conversionRateFloat, convertToFloat float64
}

func (suite *AutomationTestSuite) SetupSuite() {
	gotenv.Load()

	suite.Driver = agouti.ChromeDriver()

	suite.Require().NoError(
		suite.Driver.Start(),
	)

	page, err := suite.Driver.NewPage(agouti.Browser("chrome"))
	suite.Require().NoError(err)

	suite.page = page

	suite.page.SetImplicitWait(5000)
}

func (suite *AutomationTestSuite) SetupTest() {
	err := suite.page.Navigate(os.Getenv("PAYPAL_URL"))
	suite.Require().NoError(err)

	text, err := suite.page.Find("#header-logout").Text()
	if err != nil && text != "Log Out" {
		suite.LoginToAccount()
	}
}

func (suite *AutomationTestSuite) TearDownSuite() {
	suite.page.CloseWindow()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(AutomationTestSuite))
}
