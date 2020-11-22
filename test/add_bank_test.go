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
	if err := suite.page.Find("#contents > main > div > section.fiList-icon_container > div:nth-child(1) > a > span").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#contents > main > div > section.fiList-icon_container > div:nth-child(1) > a > span").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("input[name='accountNumber']").Fill(suite.IBAN); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#mainModal > div > div > div > form:nth-child(1) > button").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#mainModal > div > div > div > div.stickyButtonFooter-wrapper.vx_btn-group_stacked > div > div > form > button.btn.vx_btn.test_accept-mandate.mandate_lg-btn").Click(); err != nil {
		suite.Require().NoError(err)
	}

	time.Sleep(time.Second * 2)

	if err := suite.page.Find("#mainModal > div > div > a").Click(); err != nil {
		suite.Require().NoError(err)
	}
}

func (suite *AutomationTestSuite) addAndRemoveBank() {
	bankName := suite.checkHTMLContains(suite.currency)

	if bankName != true {
		suite.addBank()
	} else {
		suite.removeBank()
		suite.addBank()
	}
}
