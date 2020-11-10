package automation

func (suite *AutomationTestSuite) TestAddBank() {
	suite.LoginToAccount()

	if err := suite.page.Find("#header-wallet").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#contents > main > div > section.fiList-icon_container > div:nth-child(1) > a > span").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#contents > main > div > section.fiList-icon_container > div:nth-child(1) > a > span").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("input[name='accountNumber']").Fill("EE123456789012345678"); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#mainModal > div > div > div > form:nth-child(1) > button").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#mainModal > div > div > div > div.stickyButtonFooter-wrapper.vx_btn-group_stacked > div > div > form > button.btn.vx_btn.test_accept-mandate.mandate_lg-btn").Click(); err != nil {
		suite.Require().NoError(err)
	}
}
