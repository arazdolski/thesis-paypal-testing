package automation

import "time"

func (suite *AutomationTestSuite) TestChangeName() {
	// Arrange
	suite.firstname = "John"
	suite.lastname = "Doe"

	// Act
	suite.GoToSettings()
	suite.changeName()

	// Assert
	suite.Equal(true, suite.checkHTMLContains(suite.firstname))
	suite.Equal(true, suite.checkHTMLContains(suite.lastname))
}

// private

func (suite *AutomationTestSuite) changeName() {
	err := suite.page.Find("#accountTab > div > div:nth-child(1) > div.row.vx_panel.vx_panel-profile > div:nth-child(2) > div.col-sm-8.col-md-9.profile-detailColumn.js_name > div.row.profileDetail-container > a").Click()
	suite.Require().NoError(err)

	err = suite.page.Find("#overpanel > div > div > div.overpanel-content > div.overpanel-body > div > a:nth-child(3)").Click()
	suite.Require().NoError(err)

	err = suite.page.Find("input[name='firstName']").Fill(suite.firstname)
	suite.Require().NoError(err)

	err = suite.page.Find("input[name='lastName']").Fill(suite.lastname)
	suite.Require().NoError(err)

	err = suite.page.Find("#legalDisclaimer").Click()
	suite.Require().NoError(err)

	time.Sleep(time.Second * 5)

	err = suite.page.FindByXPath("/html/body/section[2]/div/div/div[2]/div[2]/form/div[5]/input").Click()
	suite.Require().NoError(err)

	time.Sleep(time.Second * 10)
}
