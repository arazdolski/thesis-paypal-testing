package automation

import (
	"time"
)

func (suite *AutomationTestSuite) TestAddBank() {
	// Arrange
	suite.bankName = "Rabobank Nederland"
	suite.IBAN = "EE123456789012345678"

	// Act
	suite.GoToWallet()
	suite.addAndRemoveBank()

	// Assert
	suite.Equal(true, suite.checkHTMLContains(suite.bankName))
}

// private

func (suite *AutomationTestSuite) addBank() {
	err := suite.page.Find("#contents > main > div > section.fiList-icon_container > div:nth-child(1) > a > span").Click()
	suite.Require().NoError(err)

	err = suite.page.Find("#contents > main > div > section.fiList-icon_container > div:nth-child(1) > a > span").Click()
	suite.Require().NoError(err)

	err = suite.page.Find("input[name='accountNumber']").Fill(suite.IBAN)
	suite.Require().NoError(err)

	err = suite.page.Find("#mainModal > div > div > div > form:nth-child(1) > button").Click()
	suite.Require().NoError(err)

	err = suite.page.Find("#mainModal > div > div > div > div.stickyButtonFooter-wrapper.vx_btn-group_stacked > div > div > form > button.btn.vx_btn.test_accept-mandate.mandate_lg-btn").Click()
	suite.Require().NoError(err)

	time.Sleep(time.Second * 2)

	err = suite.page.Find("#mainModal > div > div > a").Click()
	suite.Require().NoError(err)
}

func (suite *AutomationTestSuite) addAndRemoveBank() {
	bankName := suite.checkHTMLContains(suite.bankName)

	if bankName != true {
		suite.addBank()
	} else {
		suite.removeBank()
		suite.addBank()
	}
}
