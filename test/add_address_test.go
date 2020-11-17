package automation

import "time"

func (suite *AutomationTestSuite) TestAddAddress() {
	suite.GoToSettings()

	suite.addAddress()
}

// private

func (suite *AutomationTestSuite) addAddress() {
	
	if err := suite.page.Find("#accountTab > div > div.col-md-6.section.top.profile-panel > div.row.vx_panel.vx_panel-address.panel-overflow > a").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("input[name='line1']").Fill("Address 1"); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("input[name='city']").Fill("City"); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("input[name='postalCode']").Fill("12321321"); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("input[name='state']").Fill("City"); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#overpanel > div > div > div.overpanel-content > div.overpanel-body > form > div:nth-child(6) > input").Click(); err != nil {
		suite.Require().NoError(err)
	}

	time.Sleep(time.Second * 10)

	if err := suite.page.Find("#overpanel > div > div > div.overpanel-content > div.overpanel-body > div > input").Click(); err != nil {
		suite.Require().NoError(err)
	}
}
