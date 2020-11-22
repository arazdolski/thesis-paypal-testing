package automation

func (suite *AutomationTestSuite) TestRemoveEmail() {
	// Arrange
	suite.email = suite.randomString(10) + "@test.com"

	// Act
	suite.GoToSettings()
	suite.addEmail(suite.email)
	suite.clickRemoveEmail()
	suite.checkHTMLContains(suite.email)

	// Assert
	suite.Equal(suite.checkHTMLContains(suite.email), false)

}

func (suite *AutomationTestSuite) clickRemoveEmail() {
	if err := suite.page.FindByXPath("/html/body/div[3]/div[2]/div/div/div[2]/div[1]/div/div[2]/div[2]/div[2]/ul/li[2]/div[2]/span/a[2]").Click(); err != nil {
		suite.Require().NoError(err)
	}
}
