package automation

import (
	"strings"
	"time"

	"github.com/nouney/randomstring"
)

func (suite *AutomationTestSuite) randomString(num int) string {
	rsg, err := randomstring.NewGenerator("abcdefghijklmnopqrstuvwxyz")
	if err != nil {
		suite.Require().NoError(err)
	}
	str := rsg.Generate(num)

	return str
}

func (suite *AutomationTestSuite) checkHTMLContains(str string) bool {
	time.Sleep(time.Second * 5)
	html, err := suite.page.HTML()
	if err != nil {
		suite.Require().NoError(err)
	}

	return strings.Contains(html, str)
}
