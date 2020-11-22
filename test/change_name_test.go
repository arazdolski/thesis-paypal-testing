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
}

// private

func (suite *AutomationTestSuite) changeName() {
	if err := suite.page.Find("#accountTab > div > div:nth-child(1) > div.row.vx_panel.vx_panel-profile > div:nth-child(2) > div.col-sm-8.col-md-9.profile-detailColumn.js_name > div.row.profileDetail-container > a").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#overpanel > div > div > div.overpanel-content > div.overpanel-body > div > a:nth-child(3)").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("input[name='firstName']").Fill(suite.firstname); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("input[name='lastName']").Fill(suite.lastname); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#legalDisclaimer").Click(); err != nil {
		suite.Require().NoError(err)
	}

	time.Sleep(time.Second * 5)

	if err := suite.page.FindByXPath("/html/body/section[2]/div/div/div[2]/div[2]/form/div[5]/input").Click(); err != nil {
		suite.Require().NoError(err)
	}

	time.Sleep(time.Second * 10)
}
