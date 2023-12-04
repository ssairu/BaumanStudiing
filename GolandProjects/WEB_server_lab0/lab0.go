package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

func main() {

	http.HandleFunc("/about", func(w http.ResponseWriter, _ *http.Request) {

		fmt.Fprintln(w, "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n "+
			"   <meta charset=\"UTF-8\">\n  "+
			"  <title>Home page</title>\n</head>\n<body>\n"+
			"<p>\n    About page\n</p>\n"+
			"<p>\n    <a href=\"/\">Home page</a>\n</p>"+
			"<p><img src=\"https://sun9-57.userapi.com/impg/nE1FRMhYAXIu9i2pd9CroVQZq9suPxblhVhF6Q/8axZixRzwwY.jpg?"+
			"size=1080x1015&quality=96&sign=ec889be9796c55bc883e6a401fda124b&type=album\" align=”middle”>"+
			"<p/>\n</body>\n</html>")
	})

	http.HandleFunc("/news", func(w http.ResponseWriter, _ *http.Request) {
		response, err := http.Get("https://www.mos.ru/rss")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer response.Body.Close()

		XMLdata, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		rss := new(Rss)

		buffer := bytes.NewBuffer(XMLdata)

		decoded := xml.NewDecoder(buffer)

		err = decoded.Decode(rss)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Title : %s\n", rss.Channel.Title)
		fmt.Printf("Description : %s\n", rss.Channel.Description)
		fmt.Printf("Link : %s\n", rss.Channel.Link)

		total := len(rss.Channel.Items)

		fmt.Printf("Total items : %v\n", total)

		fmt.Fprintln(w, "<!DOCTYPE html>\n<html lang=\"en\">\n"+
			"<head>\n<meta charset=\"UTF-8\">\n<title>Home page</title>\n"+
			"</head>\n<body>\n"+
			"<p> <a href=\"/\">Home page</a><p/>\n<p>")
		for i := 0; i < total; i++ {
			fmt.Fprintf(w, "<p>[%d] item title : %s\n <br>", i, rss.Channel.Items[i].Title)
			fmt.Fprintf(w, ". item description : %s\n<br>", rss.Channel.Items[i].Description)
			fmt.Fprintf(w, ". item link : <a href=\"%s\">\"%s\"</a>\n\n</p>",
				rss.Channel.Items[i].Link, rss.Channel.Items[i].Link)
		}
		fmt.Fprintln(w, "\n</p>\n</body>\n</html>")

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {
			w.WriteHeader(404)
			w.Write([]byte("404 - not found\n"))
			return
		}

		fmt.Fprintln(w, "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n "+
			"   <meta charset=\"UTF-8\">\n"+
			"    <title>Home page</title>\n</head>\n<body>\n<p>\n"+
			"    Home page\n</p>\n<p>\n    <a href=\"/about\">About page</a>\n</p>\n<p>\n"+
			"    <a href=\"/news\">RSS page</a>\n</p>\n</body>\n</html>")
	})

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":80", nil))
}
