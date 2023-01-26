package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

type TextLang struct {
	XMLName  xml.Name `xml:"TextLang"`
	Chardata string   `xml:",chardata"`
	Text     []struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
	} `xml:"Text"`
}

func main() {
	data, err := ioutil.ReadFile("explore_plugin_full.xml")
	if err != nil {
		log.Fatal(err)
	}
	tlb := &TextLang{}
	err = xml.Unmarshal([]byte(data), &tlb)
	if err != nil {
		log.Fatal(err)
	}
	data, err = ioutil.ReadFile("explore_plugin_full_new.xml")
	if err != nil {
		log.Fatal(err)
	}
	tla := &TextLang{}
	err = xml.Unmarshal([]byte(data), &tla)
	if err != nil {
		log.Fatal(err)
	}

	var diff []string
	m := make(map[string]bool)

	for _, item := range tlb.Text {
		m[item.Name] = true
	}

	for _, item := range tla.Text {
		if _, ok := m[item.Name]; !ok {
			diff = append(diff, item.Name)
		}
	}

	for _, item := range diff {
		fmt.Println(item)
	}
}
