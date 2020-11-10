package automation

func (suite *AutomationTestSuite) TestPayPalLogin(){
	suite.LoginToAccount()

	suite.page.CloseWindow()

}
