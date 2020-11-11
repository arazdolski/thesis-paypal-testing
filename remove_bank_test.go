package automation

import "time"

func (suite *AutomationTestSuite) TestRemoveBank() {
	suite.loginToAccount()

	if err := suite.page.Find("#header-wallet").Click(); err != nil {
		suite.Require().NoError(err)
	}

	bank, err := suite.page.Find("#contents > main > div > section.fiListGroup_testTreatment.nemo_fiListGroup > ul > li:nth-child(2) > a > span > span:nth-child(3) > span.fiListItem-identifier").Text()
	if err != nil {
		suite.Require().NoError(err)
	}

	if bank != "Rabobank Nederland" {
		suite.addBank()
		suite.removeBank()
	} else {
		suite.removeBank()
	}
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
}