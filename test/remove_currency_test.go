package automation

import (
	"time"
)

func (suite *AutomationTestSuite) TestRemoveCurrency() {
	// Arrange
	suite.currency = suite.toCurrency("AUD")

	// Act
	suite.GoToWallet()
	suite.addAndRemoveCurrency()

	// Assert
	suite.Equal(false, suite.checkHTMLContains(suite.currency))
}

// private

func (suite *AutomationTestSuite) removeCurrency() {
	err := suite.page.Find("#contents > main > section > div > div > ul > li:nth-child(2) > div > div > div > button > span").Click()
	suite.Require().NoError(err)

	err = suite.page.Find("#contents > main > section > div > div > ul > li:nth-child(2) > div > div > div > div > ul > li:nth-child(2) > a").Click()
	suite.Require().NoError(err)

	time.Sleep(time.Second * 5)

	err = suite.page.Find("#mainModal > div > div > div > div > form > button").Click()
	suite.Require().NoError(err)

	err = suite.page.Find("#mainModal > div > div > div > div > a").Click()
	suite.Require().NoError(err)
}

func (suite *AutomationTestSuite) addAndRemoveCurrency() {
	currency := suite.checkHTMLContains(suite.currency)

	if currency != true {
		suite.addCurrency()
		suite.removeCurrency()

	} else {
		suite.removeCurrency()
	}
}
