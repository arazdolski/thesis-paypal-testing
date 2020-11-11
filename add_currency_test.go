package automation

import (
	"time"
)

func (suite *AutomationTestSuite) TestAddCurrency() {
	suite.GoToWallet()

	suite.addCurrency()
}

// private

func (suite *AutomationTestSuite) addCurrency() {
	if err := suite.page.Find("#contents > main > section > div > div > div.currency-links > a.balanceDetails-manageCurrencies.test_mcm-addCurrency").Click(); err != nil {
		suite.Require().NoError(err)
	}

	time.Sleep(time.Second * 2)

	if err := suite.page.Find("#AUDRadioBtn").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#mainModal > div > div > div > div > form > div > div > div > button").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#mainModal > div > div > div > div > a").Click(); err != nil {
		suite.Require().NoError(err)
	}
}
