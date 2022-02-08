package types

import (
	"encoding/xml"
)

type DBLP struct {
	XMLName       xml.Name       `xml:"dblp" json:"-"`
	Articles      []Article      `xml:"article,omitempty" json:"articles,omitempty"`
	Inproceedings []Inproceeding `xml:"inproceedings,omitempty" json:"inproceedings,omitempty"`
	Proceedings   []Proceedings  `xml:"proceedings,omitempty" json:"proceedings,omitempty"`
	Books         []Book         `xml:"book,omitempty" json:"book,omitempty"`
	Incollections []Incollection `xml:"incollection,omitempty" json:"incollection,omitempty"`
	PHDThesis     []PHDThesis    `xml:"phdthesis,omitempty" json:"phdthesis,omitempty"`
	MasterThesis  []MasterThesis `xml:"mastersthesis,omitempty" json:"masterthesis,omitempty"`
	WWW           []WWW          `xml:"www,omitempty" json:"www,omitempty"`
	Data          []Data         `xml:"data,omitempty" json:"data,omitempty"`
}

type Fields struct {
	Author    string `xml:"author,omitempty" json:"author,omitempty"`
	Editor    string `xml:"editor,omitempty" json:"editor,omitempty"`
	Title     string `xml:"title,omitempty" json:"title,omitempty"`
	BookTitle string `xml:"booktitle,omitempty" json:"booktitle,omitempty"`
	Pages     string `xml:"pages,omitempty" json:"pages,omitempty"`
	Year      string `xml:"year,omitempty" json:"year,omitempty"`
	Address   string `xml:"address,omitempty" json:"address,omitempty"`
	Journal   string `xml:"journal,omitempty" json:"journal,omitempty"`
	Volume    string `xml:"volume,omitempty" json:"volume,omitempty"`
	Number    string `xml:"number,omitempty" json:"number,omitempty"`
	Month     string `xml:"month,omitempty" json:"month,omitempty"`
	Url       string `xml:"url,omitempty" json:"url,omitempty"`
	EE        string `xml:"ee,omitempty" json:"ee,omitempty"`
	CDRom     string `xml:"cdrom,omitempty" json:"cdrom,omitempty"`
	Cite      string `xml:"cite,omitempty" json:"cite,omitempty"`
	Publisher string `xml:"publisher,omitempty" json:"publisher,omitempty"`
	Note      string `xml:"note,omitempty" json:"note,omitempty"`
	Crossref  string `xml:"crossref,omitempty" json:"crossref,omitempty"`
	ISBN      string `xml:"isbn,omitempty" json:"isbn,omitempty"`
	Series    string `xml:"series,omitempty" json:"series,omitempty"`
	School    string `xml:"school,omitempty" json:"school,omitempty"`
	PublNr    string `xml:"publnr,omitempty" json:"publnr,omitempty"`
}

type CommonAttributes struct {
	Key      string `xml:"key,attr" json:"key"`
	MDate    string `xml:"mdate,attr,omitempty" json:"mdate,omitempty"`
	PublType string `xml:"publtype,attr,omitempty" json:"publtype,omitempty"`
	CDate    string `xml:"cdate,attr,omitempty" json:"cdate,omitempty"`
}

type Article struct {
	Fields
	CommonAttributes
	ReviewId string `xml:"reviewid,attr,omitempty" json:"reviewid,omitempty"`
	Rating   string `xml:"rating,attr,omitempty" json:"rating,omitempty"`
}

type Inproceeding struct {
	Fields
	CommonAttributes
}

type Proceedings struct {
	Fields
	CommonAttributes
}

type Book struct {
	Fields
	CommonAttributes
}

type Incollection struct {
	Fields
	CommonAttributes
}

type PHDThesis struct {
	Fields
	CommonAttributes
}

type MasterThesis struct {
	Fields
	CommonAttributes
}

type WWW struct {
	Fields
	CommonAttributes
}

type Data struct {
	Fields
	CommonAttributes
}
