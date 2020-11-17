package automation

import "time"

func (suite *AutomationTestSuite) TestRemoveAddress() {
	suite.GoToSettings()

	suite.addAddress()

	if err := suite.page.Find("#overpanel > div > div > div.overpanel-content > div.overpanel-body > div > ul > li:nth-child(2) > p.manageAddresses-actions > a:nth-child(2)").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#overpanel > div > div > div.overpanel-content > div.overpanel-body > form > div.button.center > input").Click(); err != nil {
		suite.Require().NoError(err)
	}

	time.Sleep(time.Second * 5)
}
