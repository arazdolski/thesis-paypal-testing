package automation

import (
	"os"
	"time"
)

func (suite *AutomationTestSuite) LoginToAccount() {
	if err := suite.page.Find("#acceptAllButton").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#ul-btn").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("input[name='login_email']").Fill(os.Getenv("PAYPAL_USER_EMAIL")); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#btnNext").Click(); err != nil {
		suite.Require().NoError(err)
	}

	time.Sleep(time.Second)

	if err := suite.page.Find("input[name='login_password']").Fill(os.Getenv("PAYPAL_USER_PASSWORD")); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#btnLogin").Click(); err != nil {
		suite.Require().NoError(err)
	}
}

func (suite *AutomationTestSuite) GoToWallet() {
	suite.LoginToAccount()

	if err := suite.page.Find("#header-wallet").Click(); err != nil {
		suite.Require().NoError(err)
	}
}

func (suite *AutomationTestSuite) GoToSettings() {
	suite.LoginToAccount()

	if err := suite.page.Find("#header-settings").Click(); err != nil {
		suite.Require().NoError(err)
	}
}

func (suite *AutomationTestSuite) GoToTransfer() {
	suite.LoginToAccount()

	if err := suite.page.Find("#header-transfer").Click(); err != nil {
		suite.Require().NoError(err)
	}
}
