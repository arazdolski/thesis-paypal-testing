package automation

import "time"

func (suite *AutomationTestSuite) TestMobilePayment() {
	suite.GoToTransfer()

	if err := suite.page.Find("input[name='autocomplete-input']").Fill("+12345678901"); err != nil {
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

	time.Sleep(time.Second * 15)

	if err := suite.page.Find("#react-transfer-container > div > div > form > button.css-1mggxor.vx_btn").Click(); err != nil {
		suite.Require().NoError(err)
	}
}
