package automation

func (suite *AutomationTestSuite) TestTurnOnInterestBasedMarketingInOtherAccounts() {
	suite.GoToSettings()

	suite.GoToSecurity()
	suite.clickInterestBasedMarketingInOtherAccountsSetting()

	if suite.checkSettingIsOn() == true {
		return
	}

	err := suite.page.Find("#overpanel > div > div > div.overpanel-content > div.overpanel-body > form > div.personalizationButton > input").Click()
	suite.Require().NoError(err)
}

func (suite *AutomationTestSuite) TestTurnOffInterestBasedMarketingInOtherAccounts() {
	suite.GoToSettings()

	suite.GoToSecurity()
	suite.clickInterestBasedMarketingInOtherAccountsSetting()

	if suite.checkSettingIsOn() == false {
		return
	}

	err := suite.page.Find("#overpanel > div > div > div.overpanel-content > div.overpanel-body > form > div.personalizationButton > input").Click()
	suite.Require().NoError(err)
}

func (suite *AutomationTestSuite) TestTurnOnInterestBasedMarketingOnPaypal() {
	suite.GoToSettings()

	suite.GoToSecurity()
	suite.clickInterestBasedMarketingOnPaypalAccountsSetting()

	if suite.checkSettingIsOn() == true {
		return
	}

	err := suite.page.Find("#overpanel > div > div > div.overpanel-content > div.overpanel-body > form > div.personalizationButton > input").Click()
	suite.Require().NoError(err)
}

func (suite *AutomationTestSuite) TestTurnOffInterestBasedMarketingOnPaypal() {
	suite.GoToSettings()

	suite.GoToSecurity()
	suite.clickInterestBasedMarketingOnPaypalAccountsSetting()

	if suite.checkSettingIsOn() == false {
		return
	}

	err := suite.page.Find("#overpanel > div > div > div.overpanel-content > div.overpanel-body > form > div.personalizationButton > input").Click()
	suite.Require().NoError(err)
}

func (suite *AutomationTestSuite) TestTurnOffCookieSettings() {

	suite.GoToSettings()
	suite.GoToSecurity()
	suite.clickCookieSettings()
	suite.uncheckAllCheckboxes()
	suite.clickSaveCookieSetting()
}

// private

func (suite *AutomationTestSuite) clickInterestBasedMarketingInOtherAccountsSetting() {
	err := suite.page.Find("#securityTab > div > ul > li.otherPersonalization > div > div.col-xs-2.modify > span > a").Click()
	suite.Require().NoError(err)
}

func (suite *AutomationTestSuite) clickInterestBasedMarketingOnPaypalAccountsSetting() {
	err := suite.page.Find("#securityTab > div > ul > li.paypalPersonalization > div > div.col-xs-2.modify > span > a").Click()
	suite.Require().NoError(err)
}

func (suite *AutomationTestSuite) clickCookieSettings() {
	err := suite.page.Find("#securityTab > div > ul > li:nth-child(7) > div > div.col-xs-2.modify > span > a").Click()
	suite.Require().NoError(err)
}

func (suite *AutomationTestSuite) uncheckAllCheckboxes() {
	err := suite.page.Find("#mainModal > div > div.vx_modal-content.padding-20 > div > div > form > div:nth-child(4) > div:nth-child(1) > div > label").Click()
	suite.Require().NoError(err)

	err = suite.page.Find("#mainModal > div > div.vx_modal-content.padding-20 > div > div > form > div:nth-child(6) > div:nth-child(1) > div > label").Click()
	suite.Require().NoError(err)

	err = suite.page.Find("#mainModal > div > div.vx_modal-content.padding-20 > div > div > form > div:nth-child(8) > div:nth-child(1) > div > label").Click()
	suite.Require().NoError(err)
}

func (suite *AutomationTestSuite) clickSaveCookieSetting() {
	err := suite.page.Find("#mainModal > div > div.vx_modal-content.padding-20 > div > div > form > div.cookieAction > button").Click()
	suite.Require().NoError(err)
}
