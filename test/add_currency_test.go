package automation

import (
	"time"
)

func (suite *AutomationTestSuite) TestAddCurrency() {
	// Arrange
	suite.currency = suite.toCurrency("AUD")

	// Act
	suite.GoToWallet()
	suite.removeAndAddCurrency()

	// Assert
	suite.Equal(true, suite.checkHTMLContains(suite.currency))
}

// private

func (suite *AutomationTestSuite) addCurrency() {
	suite.page.Find("#contents > main > section > div > div > div.currency-links > a.balanceDetails-manageCurrencies.test_mcm-addCurrency").Click()
	suite.Require().NoError(err)

	time.Sleep(time.Second * 2)

	suite.page.FindByClass(suite.currency).Click()
	suite.Require().NoError(err)

	suite.page.Find("#mainModal > div > div > div > div > form > div > div > div > button").Click()
	suite.Require().NoError(err)

	suite.page.Find("#mainModal > div > div > div > div > a").Click()
	suite.Require().NoError(err)
}

func (suite *AutomationTestSuite) removeAndAddCurrency() {
	currency := suite.checkHTMLContains(suite.currency)

	if currency == true {
		suite.removeCurrency()
		suite.addCurrency()
	} else {
		suite.addCurrency()
	}
}
