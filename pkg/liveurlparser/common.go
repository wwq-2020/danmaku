package liveurlparser

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

func (p *parser) doGet(ctx context.Context, url string, resp interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return errors.WithMessagef(err, "failed to new http request for url:%s", url)
	}

	httpResp, err := p.httpClient.Do(req)
	if err != nil {
		return errors.WithMessagef(err, "failed to do http request for url:%s", url)

	}
	defer httpResp.Body.Close()
	contentType := httpResp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		return errors.Errorf("got unexpected contentType:%s for url:%s", contentType, url)
	}

	if err := json.NewDecoder(httpResp.Body).Decode(resp); err != nil {
		return errors.WithMessagef(err, "failed to decode resp body for url:%s", url)
	}
	return nil
}
