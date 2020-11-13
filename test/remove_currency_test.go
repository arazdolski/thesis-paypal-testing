package automation

import (
	"time"
)

func (suite *AutomationTestSuite) TestRemoveCurrency() {
	suite.GoToWallet()

	currency, err := suite.page.Find("#contents > main > section > div > div > ul > li:nth-child(2) > div > span.flex-item.multiCurrency-label_left > span.multiCurrency-label_left.multiCurrency-label_alignMiddle").Text()
	if err != nil {
		suite.Require().NoError(err)
	}

	if currency != "AUD" {
		suite.addCurrency()
		suite.removeCurrency()

	} else {
		suite.removeCurrency()
	}
}

// private

func (suite *AutomationTestSuite) removeCurrency() {
	if err := suite.page.Find("#contents > main > section > div > div > ul > li:nth-child(2) > div > div > div > button > span").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#contents > main > section > div > div > ul > li:nth-child(2) > div > div > div > div > ul > li:nth-child(2) > a").Click(); err != nil {
		suite.Require().NoError(err)
	}

	time.Sleep(time.Second * 5)

	if err := suite.page.Find("#mainModal > div > div > div > div > form > button").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#mainModal > div > div > div > div > a").Click(); err != nil {
		suite.Require().NoError(err)
	}
}
