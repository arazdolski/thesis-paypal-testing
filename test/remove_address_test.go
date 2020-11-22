package automation

import (
	"time"
)

func (suite *AutomationTestSuite) TestRemoveAddress() {
	// Arrange
	suite.address = "Address 1"
	suite.city = "City"
	suite.postalCode = "12321321"
	suite.state = "City"
	successMessage := "Your address has been removed."

	// Act
	suite.GoToSettings()
	suite.addAddress()
	suite.removeAddress()

	// Assert
	suite.Equal(true, suite.checkHTMLContains(successMessage))
}

func (suite *AutomationTestSuite) removeAddress() {
	if err := suite.page.Find("#overpanel > div > div > div.overpanel-content > div.overpanel-body > div > ul > li:nth-child(2) > p.manageAddresses-actions > a:nth-child(2)").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#overpanel > div > div > div.overpanel-content > div.overpanel-body > form > div.button.center > input").Click(); err != nil {
		suite.Require().NoError(err)
	}

	time.Sleep(time.Second * 5)
}
