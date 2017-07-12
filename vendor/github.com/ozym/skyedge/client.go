package skyedge

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/taskcluster/httpbackoff"
)

type value struct {
	XMLName xml.Name `xml:"value"`
	Id      string   `xml:"id,attr"`
	Val     string   `xml:"val,attr"`
}

type values struct {
	XMLName xml.Name `xml:"values"`
	Values  []value  `xml:"value"`
}

var ErrClientEmptyResponse = errors.New("empty client response")
var ErrClientEmptyDocument = errors.New("empty client document")

func clientHandleDocument(client *http.Client, req *http.Request, handler func(d *goquery.Document) error) error {
	return clientHandleResponse(client, req, func(res *http.Response) error {
		var err error
		var doc *goquery.Document

		if res == nil {
			return ErrClientEmptyResponse
		}

		log.Println(res)
		if doc, err = goquery.NewDocumentFromResponse(res); doc == nil {
			return ErrClientEmptyDocument
		}
		if err != nil {
			return err
		}

		if handler != nil {
			return handler(doc)
		}

		return nil
	})
}

func clientHandleResponseBody(client *http.Client, req *http.Request, handler func(b []byte) error) error {
	return clientHandleResponse(client, req, func(res *http.Response) error {
		if res == nil {
			return ErrClientEmptyResponse
		}

		body := func(response *http.Response) []byte {
			if response.Body != nil {
				if b, err := ioutil.ReadAll(response.Body); err == nil {
					return b
				}
			}
			return []byte{}
		}(res)

		if handler != nil {
			return handler(body)
		}

		return nil
	})

	return nil
}

func clientHandleResponse(client *http.Client, req *http.Request, handler func(r *http.Response) error) error {
	var err error
	var res *http.Response

	if res, _, err = httpbackoff.ClientDo(client, req); res == nil {
		return ErrClientEmptyResponse
	}
	if err != nil {
		body := func(response *http.Response) []byte {
			if response.Body != nil {
				if b, err := ioutil.ReadAll(response.Body); err == nil {
					return b
				}
			}
			return []byte{}
		}(res)
		return fmt.Errorf("HTTP Response code %d: %s", res.StatusCode, strings.TrimSpace(string(body)))
	}
	defer func() {
		if res.Body != nil {
			io.Copy(ioutil.Discard, res.Body)
			res.Body.Close()
		}
	}()

	if handler != nil {
		return handler(res)
	}

	return nil
}
