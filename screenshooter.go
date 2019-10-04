package screenshooter

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"time"

	"github.com/chromedp/chromedp"
)

// Screenshooter takes screenshots of given webpages.
type Screenshooter struct {
	ctx      context.Context
	cancel   context.CancelFunc
	tasks    chromedp.Tasks
	Timeout  time.Duration
	SavePath string
}

// New initializes and returns a Screenshooter.
func New(config ...func(*Screenshooter)) *Screenshooter {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	ss := Screenshooter{
		ctx:    ctx,
		cancel: cancel,
	}

	for _, c := range config {
		c(&ss)
	}

	return &ss
}

// elementScreenshot takes a screenshot of a specific element.
func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(sel, chromedp.ByQuery),
		chromedp.Screenshot(sel, res, chromedp.ByQuery),
	}
}

// Take a screenshot of the given post.
func (ss *Screenshooter) Take(postURL url.URL, element string, filename string) error {
	defer ss.cancel()
	if ss.Timeout > 0 {
		ss.ctx, ss.cancel = context.WithTimeout(ss.ctx, ss.Timeout)
	}

	var buf []byte
	err := chromedp.Run(ss.ctx, elementScreenshot(postURL.String(), element, &buf))
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/%s.png", ss.SavePath, filename), buf, 0644)
	if err != nil {
		return err
	}

	return nil
}
