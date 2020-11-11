package automation

func (suite *AutomationTestSuite) TestPayPalLogin() {
	suite.loginToAccount()

	suite.page.CloseWindow()

}
