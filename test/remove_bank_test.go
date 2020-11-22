package automation

import "time"

func (suite *AutomationTestSuite) TestRemoveBank() {
	// Arrange
	suite.bankName = "Rabobank Nederland"
	suite.IBAN = "EE123456789012345678"
	
	// Act
	suite.GoToWallet()
	suite.removeAndCheckBank()

	// Assert
	suite.Equal(false, suite.checkHTMLContains(suite.bankName))
}

// private

func (suite *AutomationTestSuite) removeBank() {
	if err := suite.page.Find("#contents > main > div > section.fiListGroup_testTreatment.nemo_fiListGroup > ul > li:nth-child(2) > a > span > span:nth-child(3) > span.fiListItem-identifier").Click(); err != nil {
		suite.Require().NoError(err)
	}

	time.Sleep(time.Second * 2)

	if err := suite.page.Find("#contents > main > section > div > div > div.fiDetails-body > div.fiDetails-actionLinks > a").Click(); err != nil {
		suite.Require().NoError(err)
	}

	time.Sleep(time.Second * 2)

	if err := suite.page.Find("#mainModal > div > div > div > div > button").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#mainModal > div > div > div > div > a").Click(); err != nil {
		suite.Require().NoError(err)
	}

	time.Sleep(time.Second * 2)
}

func (suite *AutomationTestSuite) removeAndCheckBank() {
	bankName := suite.checkHTMLContains(suite.bankName)

	if bankName != true {
		suite.addBank()
		suite.removeBank()
	} else {
		suite.removeBank()
	}
}
