package automation

import (
	"time"
)

func (suite *AutomationTestSuite) TestAddCurrency() {
	// Arrange
	suite.currency = suite.toCurrency("AUD")

	// Act
	suite.GoToWallet()
	suite.addCurrency()

	// Assert
}

// private

func (suite *AutomationTestSuite) addCurrency() {
	if err := suite.page.Find("#contents > main > section > div > div > div.currency-links > a.balanceDetails-manageCurrencies.test_mcm-addCurrency").Click(); err != nil {
		suite.Require().NoError(err)
	}

	time.Sleep(time.Second * 2)

	if err := suite.page.FindByClass(suite.currency).Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#mainModal > div > div > div > div > form > div > div > div > button").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#mainModal > div > div > div > div > a").Click(); err != nil {
		suite.Require().NoError(err)
	}
}
