package automation

import "time"

func (suite *AutomationTestSuite) TestMoneyRequest() {
	suite.GoToTransfer()

	if err := suite.page.FindByXPath("/html/body/div[3]/div[2]/section/div[1]/div/div/div/div[1]/a[2]").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("input[name='multi-autocomplete-input']").Fill("+12345678901"); err != nil {
		suite.Require().NoError(err)
	}

	time.Sleep(time.Second * 6)

	if err := suite.page.Find("#downshift-0-item-0 > div").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#react-transfer-container > div > div > div > div.css-1s5wvjt > div > div.css-a3b40m.transferPage.row > form > div.css-xcma15 > div > span.recipient-next > button").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#fn-amount").SendKeys("1"); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#react-transfer-container > div > div > div > form > div.preview-buttons > button").Click(); err != nil {
		suite.Require().NoError(err)
	}

	suite.page.CloseWindow()
}
