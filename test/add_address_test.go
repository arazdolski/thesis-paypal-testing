package automation

import "time"

func (suite *AutomationTestSuite) TestAddAddress() {
	// Arrange
	suite.address = "Address 123"
	suite.city = "City"
	suite.postalCode = "123213121"
	suite.state = "City"

	// Act
	suite.GoToSettings()
	suite.addAddress()

	// Assert
	suite.Equal(true, suite.checkHTMLContains(suite.address))
}

// private

func (suite *AutomationTestSuite) addAddress() {
	suite.page.Find("#accountTab > div > div.col-md-6.section.top.profile-panel > div.row.vx_panel.vx_panel-address.panel-overflow > a").Click()
	suite.Require().NoError(err)

	suite.page.Find("input[name='line1']").Fill(suite.address)
	suite.Require().NoError(err)

	suite.page.Find("input[name='city']").Fill(suite.city)
	suite.Require().NoError(err)

	suite.page.Find("input[name='postalCode']").Fill(suite.postalCode)
	suite.Require().NoError(err)

	suite.page.Find("input[name='state']").Fill(suite.state)
	suite.Require().NoError(err)

	suite.page.Find("#overpanel > div > div > div.overpanel-content > div.overpanel-body > form > div:nth-child(6) > input").Click()
	suite.Require().NoError(err)

	time.Sleep(time.Second * 10)

	suite.page.Find("#overpanel > div > div > div.overpanel-content > div.overpanel-body > div > input").Click()
	suite.Require().NoError(err)
}
