package automation

import (
	"os"
	"testing"
	"time"

	"github.com/sclevine/agouti"
	"github.com/subosito/gotenv"
)

func TestPayPalLogin(t *testing.T) {

	gotenv.Load()

	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			// "--headless",
			"--allow-insecure-localhost",
			"--disable-gpu",
			"--homepage=about:blank",
			"--no-first-run",
			"--no-default-browser-check",
			"--no-sandbox",
		}),
		agouti.Debug,
	)

	if err := driver.Start(); err != nil {
		t.Fatal(err, "starting Chromedriver failed")
	}

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		t.Fatal(err, "opening new page failed")
	}

	if err := page.Navigate(os.Getenv("PAYPAL_URL")); err != nil {
		t.Fatal(err, "navigating to URL failed")
	}

	if err := page.Find("#ul-btn").Click(); err != nil {
		t.Fatal(err, "clicking login button failed")
	}

	if err := page.Find("input[name='login_email']").Fill(os.Getenv("PAYPAL_USER_EMAIL")); err != nil {
		t.Fatal(err, "filling email field failed")
	}

	if err := page.Find("#btnNext").Click(); err != nil {
		t.Fatal(err, "clicking Next button failed")
	}

	time.Sleep(time.Second)

	if err := page.Find("input[name='login_password']").Fill(os.Getenv("PAYPAL_USER_PASSWORD")); err != nil {
		t.Fatal(err, "filling password field failed")
	}

	if err := page.Find("#btnLogin").Click(); err != nil {
		t.Fatal(err, "clicking Next button failed")
	}

	page.CloseWindow()

}
