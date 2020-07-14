package liveurlparser

import (
	"context"
	"net/http"
	"testing"
)

func TestParseZhanQi(t *testing.T) {
	p := New(http.DefaultClient)
	realP := p.(*parser)
	ctx := context.Background()
	roomIDs := []int64{
		1,
		2,
	}
	for _, roomID := range roomIDs {
		if _, err := realP.parseZhanQi(ctx, roomID); err != nil {
			t.Fatalf("fail to parseZhanQi for roomID:%d, err:%#v", roomID, err)
		}
	}
}
