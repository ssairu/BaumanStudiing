package main

import (
	"bytes"
	"database/sql"
	"encoding/xml"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"os"
)

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	FullText    string `xml:"rambler:fulltext"`
}

func getRSS() []Item {
	response, err := http.Get("https://kinolexx.ru/rss")

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

	fmt.Println(rss.Channel.Items)
	return rss.Channel.Items
}

func main() {
	fmt.Println("start connect to bd")
	db, err := sql.Open("mysql", "iu9networkslabs:Je2dTYr6@tcp(students.yss.su)/iu9networkslabs")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("end connect to bd")
	fmt.Println("start get from bd")
	rows, err := db.Query("select * from iu9networkslabs.ArtyomPenkinRSS")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	fmt.Println("end get from bd")

	oldItems := []Item{}

	for rows.Next() {
		p := Item{}
		var id int32
		err := rows.Scan(&id, &p.Title, &p.Link, &p.Description, &p.FullText)
		if err != nil {
			fmt.Println(err)
			continue
		}
		oldItems = append(oldItems, p)
	}

	newItems := getRSS()

	for _, newItem := range newItems {
		fmt.Println(newItem)
		flag := true
		for _, old := range oldItems {
			if newItem.Description == old.Description && newItem.Link == old.Link &&
				newItem.Title == old.Title && newItem.FullText == old.FullText {
				flag = false
			}
		}
		if flag {
			result, err := db.Exec("insert into iu9networkslabs.ArtyomPenkinRSS (title, link, description, full_text) values (?, ?, ?, ?)",
				newItem.Title, newItem.Link, newItem.Description, newItem.FullText)
			if err != nil {
				panic(err)
			}
			fmt.Println(result.LastInsertId()) // id добавленного объекта
			fmt.Println(result.RowsAffected()) // количество затронутых строк
		}
	}
}
