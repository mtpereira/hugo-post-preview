package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/chromedp/chromedp"
)

// elementScreenshot takes a screenshot of a specific element.
func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(sel, chromedp.ByQuery),
		chromedp.Screenshot(sel, res, chromedp.ByQuery),
	}
}

func screenshotPost(ctx context.Context, post string) error {
	// capture screenshot of an element
	var buf []byte
	if err := chromedp.Run(ctx, elementScreenshot(fmt.Sprintf("http://localhost:1313/post/%s/", post), `article.post`, &buf)); err != nil {
		return err
	}
	if err := ioutil.WriteFile(fmt.Sprintf("%s.png", post), buf, 0644); err != nil {
		return err
	}
	return nil
}

func main() {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	defer cancel()

	if err := screenshotPost(ctx, "inato"); err != nil {
		log.Fatal(err)
	}
}
