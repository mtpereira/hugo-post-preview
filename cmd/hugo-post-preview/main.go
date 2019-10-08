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
	filename := flag.String("filename", "./post.png", "File where the screenshot will be saved to.")
	element := flag.String("element", "article.post", "Name of the CSS ID to select and capture.")
	siteURL := flag.String("url", "http://localhost:1313", "URL for the website. Defaults to the hugo server default address.")
	postsPath := flag.String("postsPath", "post", "Path for the posts.")
	timeout := flag.Duration("timeout", time.Duration(time.Second*10), "Time to wait before giving up on the request to the site. Disabled by setting a negative value.")
	debug := flag.Bool("debug", false, "Enables debug information.")
	flag.Parse()

	if *post == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	postURL, err := url.Parse(fmt.Sprintf("%s/%s/%s", *siteURL, *postsPath, *post))
	if err != nil {
		log.Fatalln(err)
	}

	ss := s.New(s.Timeout(*timeout), s.Debug(*debug))

	file, err := os.Create(fmt.Sprintf("%s", *filename))
	if err != nil {
		log.Fatalln(err)
	}

	err = ss.Take(*postURL, *element, file)
	if err != nil {
		log.Fatalln(err)
	}
}
