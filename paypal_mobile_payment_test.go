package automation

import (
	"os"
	"testing"
	"time"

	"github.com/subosito/gotenv"
)

func (suite *AutomationTestSuite) TestMobilePayment(t *testing.T) {

	gotenv.Load()

	if err := suite.page.Find("#ul-btn").Click(); err != nil {
		t.Fatal(err, "clicking login button failed")
	}

	if err := suite.page.Find("input[name='login_email']").Fill(os.Getenv("PAYPAL_USER_EMAIL")); err != nil {
		t.Fatal(err, "filling email field failed")
	}

	if err := suite.page.Find("#btnNext").Click(); err != nil {
		t.Fatal(err, "clicking Next button failed")
	}

	time.Sleep(time.Second)

	if err := suite.page.Find("input[name='login_password']").Fill(os.Getenv("PAYPAL_USER_PASSWORD")); err != nil {
		t.Fatal(err, "filling password field failed")
	}

	if err := suite.page.Find("#btnLogin").Click(); err != nil {
		t.Fatal(err, "clicking Next button failed")
	}

	if err := suite.page.Find("#header-transfer").Click(); err != nil {
		t.Fatal(err, "clicking Next button failed")
	}

	time.Sleep(time.Second)

	if err := suite.page.Find("input[name='autocomplete-input']").Fill("+15417543010"); err != nil {
		t.Fatal(err, "clicking Next button failed")
	}

	time.Sleep(time.Second * 3)

	if err := suite.page.Find("#downshift-0-item-0 > div").Click(); err != nil {
		t.Fatal(err, "clicking on Send to failed")
	}

	time.Sleep(time.Second * 3)

	if err := suite.page.Find("#fn-amount").Click(); err != nil {
		t.Fatal(err, "clicking Next button failed")
	}

	if err := suite.page.Find("#fn-amount").SendKeys("1"); err != nil {
		t.Fatal(err, "clicking Next button failed")
	}

	suite.page.CloseWindow()

}
