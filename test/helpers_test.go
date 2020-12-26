package automation

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/nouney/randomstring"
)

func (suite *AutomationTestSuite) randomString(num int) string {
	rsg, err := randomstring.NewGenerator("abcdefghijklmnopqrstuvwxyz")
	suite.Require().NoError(err)
	str := rsg.Generate(num)

	return str
}

func (suite *AutomationTestSuite) checkHTMLContains(str string) bool {
	time.Sleep(time.Second * 5)
	html, err := suite.page.HTML()
	suite.Require().NoError(err)

	return strings.Contains(html, str)
}

func (suite *AutomationTestSuite) checkSettingIsOn() bool {
	setting, err := suite.page.Find("#overpanel > div > div > div.overpanel-content > div.overpanel-body > form > div.personalizationButton > input").Text()
	suite.Require().NoError(err)

	if setting == "Turn It Off" {
		return false
	}
	
	return true
}

func (suite *AutomationTestSuite) toCurrency(currency string) string {
	return "test_" + currency
}

func (suite *AutomationTestSuite) paypalAmount(amount string) string {
	i, err := strconv.ParseFloat(amount, 64)
	suite.Require().NoError(err)

	j := strings.Replace(fmt.Sprintf("%.2f", i / 100), ".", ",", -1)

	return j
}
