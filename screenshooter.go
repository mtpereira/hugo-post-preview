package screenshotter

import (
	"context"
	"io"
	"log"
	"net/url"
	"time"

	"github.com/chromedp/chromedp"
)

// Screenshotter takes screenshots of given webpages.
type Screenshotter struct {
	ctx     context.Context
	cancel  context.CancelFunc
	tasks   chromedp.Tasks
	timeout time.Duration
	options []chromedp.ContextOption
}

// Debug enables debug output on the Screenshooter.
func Debug(enabled bool) func(*Screenshotter) {
	return func(ss *Screenshotter) {
		if enabled {
			ss.options = append(ss.options, chromedp.WithDebugf(log.Printf))
		}
	}
}

// Timeout configures a timeout for taking screenshots.
func Timeout(duration time.Duration) func(*Screenshotter) {
	return func(ss *Screenshotter) {
		ss.timeout = duration
	}
}

// New initializes and returns a Screenshooter.
func New(config ...func(*Screenshotter)) *Screenshotter {
	var ss Screenshotter

	for _, c := range config {
		c(&ss)
	}

	ctx, cancel := chromedp.NewContext(context.Background(), ss.options...)
	ss.ctx = ctx
	ss.cancel = cancel

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
func (ss *Screenshotter) Take(postURL url.URL, element string, w io.Writer) error {
	defer ss.cancel()
	if ss.timeout > 0 {
		ss.ctx, ss.cancel = context.WithTimeout(ss.ctx, ss.timeout)
	}

	var buf []byte
	err := chromedp.Run(ss.ctx, elementScreenshot(postURL.String(), element, &buf))
	if err != nil {
		return err
	}

	_, err = w.Write(buf)
	if err != nil {
		return err
	}

	return nil
}
