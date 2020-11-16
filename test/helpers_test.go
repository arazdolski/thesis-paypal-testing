package automation

import (
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
