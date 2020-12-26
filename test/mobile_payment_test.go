package automation

import (
	"time"
)

func (suite *AutomationTestSuite) TestMobilePayment() {
	// Arrange
	suite.mobile = "+12345678901"
	suite.amount = "1"
	successMessage := "You've sent " + suite.paypalAmount(suite.amount) + "&nbsp;EUR to " + suite.mobile

	// Act
	suite.GoToTransfer()
	suite.mobilePayment()

	// Assert
	suite.Equal(true, suite.checkHTMLContains(successMessage))
}

// private

func (suite *AutomationTestSuite) mobilePayment() {
	err := suite.page.Find("input[name='autocomplete-input']").Fill(suite.mobile)
	suite.Require().NoError(err)

	err = suite.page.Find("#downshift-0-item-0 > div").Click()
	suite.Require().NoError(err)

	err = suite.page.Find("#fn-amount").Click()
	suite.Require().NoError(err)

	err = suite.page.Find("#fn-amount").SendKeys("1")
	suite.Require().NoError(err)

	err = suite.page.Find("#react-transfer-container > div > div > form > button.css-1mggxor.vx_btn").Click()
	suite.Require().NoError(err)

	time.Sleep(time.Second * 2)

	err = suite.page.Find("#personal").Click()
	suite.Require().NoError(err)

	time.Sleep(time.Second * 15)

	err = suite.page.Find("#react-transfer-container > div > div > form > button.css-1mggxor.vx_btn").Click()
	suite.Require().NoError(err)

	time.Sleep(time.Second * 2)
}
