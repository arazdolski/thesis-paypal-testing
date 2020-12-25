package automation

import (
	"time"
)

func (suite *AutomationTestSuite) TestAddEmailInCaseOfValidEmail() {
	// Arrange
	suite.email = suite.randomString(10) + "@test.com"

	// Act
	suite.GoToSettings()
	suite.addEmail(suite.email)

	// Assert
	suite.Equal(true, suite.checkHTMLContains(suite.email))
}

func (suite *AutomationTestSuite) TestAddEmailInCaseOfInvalidEmail() {
	// Arrange
	suite.email = "abcabc"

	// Act
	suite.GoToSettings()
	suite.addEmail(suite.email)

	// Assert
	suite.Equal(true, suite.findEmailError())
}

// private

func (suite *AutomationTestSuite) addEmail(email string) {
	suite.page.Find("#accountTab > div > div.col-md-6.section.top.profile-panel > div.row.vx_panel.vx_panel-email.panel-overflow > a").Click()
	suite.Require().NoError(err)

	suite.page.Find("input[name='email']").Fill(suite.email)
	suite.Require().NoError(err)

	time.Sleep(time.Second * 5)

	suite.page.Find("#overpanel > div > div > div.overpanel-content > div.overpanel-body > form > div.button > input").Click()
	suite.Require().NoError(err)

	time.Sleep(time.Second * 5)
}

func (suite *AutomationTestSuite) findEmailError() bool {
	_, err := suite.page.Find("#overpanel > div > div > div.overpanel-content > div.overpanel-body > form > div.textInput.email.email.email.lap.hasError").Text()
	if err == nil {
		return true
	}
	return false
}
