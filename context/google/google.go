package google

import (
	"context"
	"learn/context/userip"
	"net/http"
)

type Results []Result
type Result struct {
	Title, URL string
}

func Search(ctx context.Context, query string) (Results, error) {
	req, err := http.NewRequest("GET",
		"https://ajax.googleapis.com/ajax/services/search/web?v=1.0", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Set("q", query)

	if userIP, ok := userip.FromContext(ctx); ok {
		q.Set("userip", userIP.String())
	}
	req.URL.RawQuery = q.Encode()
}
