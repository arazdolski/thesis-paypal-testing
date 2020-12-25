package automation

import (
	"os"
	"time"
)

func (suite *AutomationTestSuite) LoginToAccount() {
	// suite.page.Find("#acceptAllButton").Click()
	// 	suite.Require().NoError(err)
	// }

	suite.page.Find("#ul-btn").Click()
	suite.Require().NoError(err)

	suite.page.Find("input[name='login_email']").Fill(os.Getenv("PAYPAL_USER_EMAIL"))
	suite.Require().NoError(err)

	suite.page.Find("#btnNext").Click()
	suite.Require().NoError(err)

	time.Sleep(time.Second)

	suite.page.Find("input[name='login_password']").Fill(os.Getenv("PAYPAL_USER_PASSWORD"))
	suite.Require().NoError(err)

	suite.page.Find("#btnLogin").Click()
	suite.Require().NoError(err)
}

func (suite *AutomationTestSuite) GoToWallet() {
	time.Sleep(time.Second)
	suite.page.Find("#header-wallet").Click()
	suite.Require().NoError(err)
}

func (suite *AutomationTestSuite) GoToSettings() {
	suite.page.Find("#header-settings").Click()
	suite.Require().NoError(err)
}

func (suite *AutomationTestSuite) GoToTransfer() {
	time.Sleep(time.Second)
	suite.page.Find("#header-transfer").Click()
	suite.Require().NoError(err)
}

func (suite *AutomationTestSuite) GoToSecurity() {
	suite.page.Find("#securityLink").Click()
	suite.Require().NoError(err)
}
