package automation

import "time"

func (suite *AutomationTestSuite) TestCreateInvoice() {
	suite.loginToAccount()

	if err := suite.page.Find("#header-transfer").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.FindByXPath("/html/body/div[3]/div[2]/section/div[1]/div/div/div/div[1]/a[4]").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.FindByXPath("/html/body/div[3]/div[2]/section/div[1]/div/div/div/div[2]/div/div/div/div[3]/a[2]").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#createNewInvoice").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#billtotokenfield-tokenfield").Fill("test@test.com"); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("input[name='itemName']").Fill("Books"); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#itemPrice_0").Fill("100"); err != nil {
		suite.Require().NoError(err)
	}

	time.Sleep(time.Second * 2)

	if err := suite.page.FindByXPath("/html/body/div[3]/div[2]/section/div/form[2]/div[13]/div/div/div/div/div[1]").Click(); err != nil {
		suite.Require().NoError(err)
	}

	//We've sent your invoice {InvoiceNum} appear

	suite.page.CloseWindow()
}
