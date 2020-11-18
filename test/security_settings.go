package automation

func (suite *AutomationTestSuite) TestTurnOnInterestBasedMarketingInOtherAccounts() {
	suite.GoToSettings()

	suite.GoToSecurity()

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

	suite.GoToSecurity()

	if suite.checkSettingIsOn() == false {
		return
	} else {
		if err := suite.page.Find("#overpanel > div > div > div.overpanel-content > div.overpanel-body > form > div.personalizationButton > input").Click(); err != nil {
			suite.Require().NoError(err)
		}
	}
}
