package automation

import "time"

func (suite *AutomationTestSuite) TestMobilePayment() {
	suite.LoginToAccount()

	if err := suite.page.Find("#header-transfer").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("input[name='autocomplete-input']").Fill("+15417543010"); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#downshift-0-item-0 > div").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#fn-amount").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#fn-amount").SendKeys("1"); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#react-transfer-container > div > div > form > button.css-1mggxor.vx_btn").Click(); err != nil {
		suite.Require().NoError(err)
	}

	time.Sleep(time.Second * 2)

	if err := suite.page.Find("#personal").Click(); err != nil {
		suite.Require().NoError(err)
	}

	time.Sleep(time.Second * 5)

	if err := suite.page.Find("#react-transfer-container > div > div > form > button.css-1mggxor.vx_btn").Click(); err != nil {
		suite.Require().NoError(err)
	}

	suite.page.CloseWindow()
}
