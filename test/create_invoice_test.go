package automation

import "time"

func (suite *AutomationTestSuite) TestCreateInvoice() {
	// Arrange
	suite.email = "test@test.com"
	suite.item = "Books"
	suite.itemPrice = "100"

	// Act
	suite.GoToTransfer()
	suite.createInvoice()

	// Assert
	//We've sent your invoice {InvoiceNum} appear
}

// private

func (suite *AutomationTestSuite) createInvoice() {
	if err := suite.page.FindByXPath("/html/body/div[3]/div[2]/section/div[1]/div/div/div/div[1]/a[4]").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.FindByXPath("/html/body/div[3]/div[2]/section/div[1]/div/div/div/div[2]/div/div/div/div[3]/a[2]").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#createNewInvoice").Click(); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#billtotokenfield-tokenfield").Fill(suite.email); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("input[name='itemName']").Fill(suite.item); err != nil {
		suite.Require().NoError(err)
	}

	if err := suite.page.Find("#itemPrice_0").Fill(suite.itemPrice); err != nil {
		suite.Require().NoError(err)
	}

	time.Sleep(time.Second * 2)

	if err := suite.page.FindByXPath("/html/body/div[3]/div[2]/section/div/form[2]/div[13]/div/div/div/div/div[1]").Click(); err != nil {
		suite.Require().NoError(err)
	}
}
