package liveurlparser

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"github.com/wwq-2020/danmaku/pkg/common"
)

// Parser Parser
type Parser interface {
	Parse(context.Context, common.LiveSource, int64) (string, error)
}

type parser struct {
	httpClient *http.Client
}

// New New
func New(httpClient *http.Client) Parser {
	return &parser{
		httpClient: httpClient,
	}
}

func (p *parser) Parse(ctx context.Context, liveSource common.LiveSource, roomID int64) (string, error) {
	var err error
	var liveURL string
	switch liveSource {
	case common.LiveSourceZhanQi:
		liveURL, err = p.parseZhanQi(ctx, roomID)
	default:
		err = errors.New("unsupported")
	}
	if err != nil {
		return "", errors.WithMessagef(err, "fail to parse for livesource:%s, roomdID:%d", liveSource, roomID)
	}
	return liveURL, nil
}
