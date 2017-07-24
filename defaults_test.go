package diamonds_test

import (
	. "github.com/harryganz/diamonds"

	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DefaultPageGetter", func() {
	table := []struct {
		params       Parameters
		rowStart     int
		responseCode int
		responseBody io.Reader
		output       io.Reader
		errors       error
		context      string
		outcome      string
	}{
		{
			Parameters{}, 0, 200, bytes.NewBufferString(""),
			bytes.NewBufferString("<html><body><table><tbody></tbody></table></body></html>"),
			nil, "when a successful empty response is sent",
			"returns a reader wrapped in html, body, table, and tbody tags",
		},
	}

	for _, v := range table {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(v.responseCode)
			io.Copy(w, v.responseBody)
		}))
		defer ts.Close()

		res, err := DefaultPageGetter(ts.URL, v.params, v.rowStart)
		Context(v.context, func() {
			It(v.outcome, func() {
				obs, _ := ioutil.ReadAll(res)
				exp, _ := ioutil.ReadAll(v.output)
				Expect(obs).To(Equal(exp))
				Expect(err == v.errors).To(BeTrue())
			})
		})
	}
})
