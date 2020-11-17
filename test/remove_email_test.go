package automation

func (suite *AutomationTestSuite) TestRemoveEmail() {

	// Arrange
	email := suite.randomString(10) + "@test.com"

	// Act
	suite.GoToSettings()
	suite.addEmail(email)

	if err := suite.page.FindByXPath("/html/body/div[3]/div[2]/div/div/div[2]/div[1]/div/div[2]/div[2]/div[2]/ul/li[2]/div[2]/span/a[2]").Click(); err != nil {
		suite.Require().NoError(err)
	}

	suite.checkHTMLContains(email)

	// Assert
	suite.Equal(suite.checkHTMLContains(email), false)

}
