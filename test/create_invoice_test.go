package automation

import "time"

func (suite *AutomationTestSuite) TestCreateInvoice() {
	// Arrange
	suite.email = "test@test.com"
	suite.item = "Books"
	suite.itemPrice = "100"
	successMessage := "We've sent your invoice"

	// Act
	suite.GoToTransfer()
	suite.createInvoice()

	// Assert
	suite.Equal(true, suite.checkHTMLContains(successMessage))
}

// private

func (suite *AutomationTestSuite) createInvoice() {
	err := suite.page.FindByXPath("/html/body/div[3]/div[2]/section/div[1]/div/div/div/div[1]/a[4]").Click()
	suite.Require().NoError(err)

	err = suite.page.FindByXPath("/html/body/div[3]/div[2]/section/div[1]/div/div/div/div[2]/div/div/div/div[3]/a[2]").Click()
	suite.Require().NoError(err)

	err = suite.page.Find("#createNewInvoice").Click()
	suite.Require().NoError(err)

	err = suite.page.Find("#billtotokenfield-tokenfield").Fill(suite.email)
	suite.Require().NoError(err)

	err = suite.page.Find("input[name='itemName']").Fill(suite.item)
	suite.Require().NoError(err)

	err = suite.page.Find("#itemPrice_0").Fill(suite.itemPrice)
	suite.Require().NoError(err)

	time.Sleep(time.Second * 2)

	err = suite.page.FindByXPath("/html/body/div[3]/div[2]/section/div/form[2]/div[13]/div/div/div/div/div[1]").Click()
	suite.Require().NoError(err)
}
