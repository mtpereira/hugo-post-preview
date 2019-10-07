package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	s "github.com/mtpereira/hugo-post-preview"
)

func main() {
	post := flag.String("post", "", "Name of the post")
	filename := flag.String("filename", "", "Filename for the screenshot. Defaults to the value of -post.")
	element := flag.String("element", "article.post", "Name of the CSS ID to select and capture.")
	siteURL := flag.String("url", "http://localhost:1313", "URL for the website. Defaults to the hugo server default address.")
	postsPath := flag.String("postsPath", "post", "Path for the posts.")
	timeout := flag.Duration("timeout", time.Duration(time.Second*10), "Time to wait before giving up on the request to the site. Disabled by setting a negative value.")
	savePath := flag.String("savePath", ".", "Path for saving the taken screenshots.")
	debug := flag.Bool("debug", false, "Enables debug information.")
	flag.Parse()

	if *post == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *filename == "" {
		*filename = *post
	}

	postURL, err := url.Parse(fmt.Sprintf("%s/%s/%s", *siteURL, *postsPath, *post))
	if err != nil {
		log.Fatalln(err)
	}

	ss := s.New(s.Timeout(*timeout), s.StoragePath(*savePath), s.Debug(*debug))

	err = ss.Take(*postURL, *element, *filename)
	if err != nil {
		log.Fatalln(err)
	}
}
