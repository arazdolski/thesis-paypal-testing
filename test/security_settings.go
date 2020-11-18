package automation

import "time"

func (suite *AutomationTestSuite) TestTurnOnInterestBasedMarketingInOtherAccounts() {
	suite.GoToSettings()

	if err := suite.page.Find("#securityLink").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#securityTab > div > ul > li.otherPersonalization > div > div.col-xs-2.modify > span > a").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if suite.checkSettingIsOn() == true {
		return
	} else {
		if err := suite.page.Find("#overpanel > div > div > div.overpanel-content > div.overpanel-body > form > div.personalizationButton > input").Click(); err != nil {
			suite.Require().NoError(err)
		}
	}
}

func (suite *AutomationTestSuite) TestTurnOffInterestBasedMarketingInOtherAccounts() {
	suite.GoToSettings()

	if err := suite.page.Find("#securityLink").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#securityTab > div > ul > li.otherPersonalization > div > div.col-xs-2.modify > span > a").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if suite.checkSettingIsOn() == false {
		return
	} else {
		if err := suite.page.Find("#overpanel > div > div > div.overpanel-content > div.overpanel-body > form > div.personalizationButton > input").Click(); err != nil {
			suite.Require().NoError(err)
		}
	}
}
