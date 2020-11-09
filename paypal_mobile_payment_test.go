package automation

import (
	"os"
	"time"

	"github.com/juju/errors"
	"github.com/subosito/gotenv"
)

func (suite *AutomationTestSuite) TestMobilePayment() error {

	gotenv.Load()

	if err := suite.page.Find("#ul-btn").Click(); err != nil {
		return errors.Annotate(err, "clicking login button failed")
	}

	if err := suite.page.Find("input[name='login_email']").Fill(os.Getenv("PAYPAL_USER_EMAIL")); err != nil {
		return errors.Annotate(err, "filling email field failed")
	}

	if err := suite.page.Find("#btnNext").Click(); err != nil {
		return errors.Annotate(err, "clicking Next button failed")
	}

	time.Sleep(time.Second)

	if err := suite.page.Find("input[name='login_password']").Fill(os.Getenv("PAYPAL_USER_PASSWORD")); err != nil {
		return errors.Annotate(err, "filling password field failed")
	}

	if err := suite.page.Find("#btnLogin").Click(); err != nil {
		return errors.Annotate(err, "clicking Next button failed")
	}

	if err := suite.page.Find("#header-transfer").Click(); err != nil {
		return errors.Annotate(err, "clicking Next button failed")
	}

	time.Sleep(time.Second)

	if err := suite.page.Find("input[name='autocomplete-input']").Fill("+15417543010"); err != nil {
		return errors.Annotate(err, "clicking Next button failed")
	}

	time.Sleep(time.Second * 3)

	if err := suite.page.Find("#downshift-0-item-0 > div").Click(); err != nil {
		return errors.Annotate(err, "clicking on Send to failed")
	}

	time.Sleep(time.Second * 3)

	if err := suite.page.Find("#fn-amount").Click(); err != nil {
		return errors.Annotate(err, "clicking Next button failed")
	}

	if err := suite.page.Find("#fn-amount").SendKeys("1"); err != nil {
		return errors.Annotate(err, "clicking Next button failed")
	}

	suite.page.CloseWindow()

	return nil
}
