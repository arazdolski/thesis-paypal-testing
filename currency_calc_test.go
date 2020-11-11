package automation

import (
	"math"
	"strconv"
	"strings"
	"time"
)

func (suite *AutomationTestSuite) TestCurrencyCalc() {
	// Arrange
	convertFrom := 1.00

	// Act
	suite.GoToWallet()

	if err := suite.page.Find("#contents > main > section > div > div > div.currency-links > a.balanceDetails-currencyCalculator").Click(); err != nil {
		suite.Require().NoError(err)
	}

	time.Sleep(time.Second * 5)

	conversionRate, err := suite.page.Find("#body-content > div > div.vx_modal-glimpse.vx_modalPrepToOpen.vx_modalIsOpen > div > div > div > div.vx_modal-body.convert-inner-section > div:nth-child(2) > div > div > div.convert-rate.col-lg-12.col-md-12.col-xs-12 > span > span.exchange-data").Text()
	if err != nil {
		suite.Require().NoError(err)
	}

	convertTo, err := suite.page.Find("#body-content > div > div.vx_modal-glimpse.vx_modalPrepToOpen.vx_modalIsOpen > div > div > div > div.vx_modal-body.convert-inner-section > div:nth-child(2) > div > div > div.convert-to.padding-none.col-lg-6.col-md-6.col-xs-12 > div.col-lg-12.col-md-12.col-xs-12.convert-details.padding-none > span.convert-value").Text()
	if err != nil {
		suite.Require().NoError(err)
	}

	conversionRateDot := strings.Replace(conversionRate[8:13], ",", ".", -1)
	convertToDot := strings.Replace(convertTo, ",", ".", -1)

	conversionRateFloat, err := strconv.ParseFloat(conversionRateDot, 64)
	if err != nil {
		suite.Require().NoError(err)
	}

	convertToFloat, err := strconv.ParseFloat(convertToDot, 64)
	if err != nil {
		suite.Require().NoError(err)
	}

	// Assert
	suite.Equal(convertFrom*(math.Round(conversionRateFloat*100)/100), convertToFloat)
}
