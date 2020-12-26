package automation

import (
	"time"
)

func (suite *AutomationTestSuite) TestMoneyRequest() {
	// Arrange
	suite.mobile = "+12345678901"
	suite.amount = "1"
	successMessage := "You requested " + suite.paypalAmount(suite.amount) + "&nbsp;EUR from " + suite.mobile

	// Act
	suite.GoToTransfer()
	suite.moneyRequest()

	// Assert
	suite.Equal(true, suite.checkHTMLContains(successMessage))
}

// private

func (suite *AutomationTestSuite) moneyRequest() {
	err := suite.page.FindByXPath("/html/body/div[3]/div[2]/section/div[1]/div/div/div/div[1]/a[2]").Click()
	suite.Require().NoError(err)

	err = suite.page.Find("input[name='multi-autocomplete-input']").Fill(suite.mobile)
	suite.Require().NoError(err)

	time.Sleep(time.Second * 6)

	err = suite.page.Find("#downshift-1-item-0 > div").Click()
	suite.Require().NoError(err)

	err = suite.page.Find("#react-transfer-container > div > div > div > div.css-1s5wvjt > div > div.css-a3b40m.transferPage.row > form > div.css-xcma15 > div > span.recipient-next > button").Click()
	suite.Require().NoError(err)

	err = suite.page.Find("#fn-amount").SendKeys(suite.amount)
	suite.Require().NoError(err)

	err = suite.page.Find("#react-transfer-container > div > div > div > form > div.preview-buttons > button").Click()
	suite.Require().NoError(err)
}
