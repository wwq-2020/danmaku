package liveurlparser

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type zhanQiAPIResp struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    *zhanQiAPIRespData `json:"data"`
}

type zhanQiAPIRespDataTagsCommon struct {
	PcIcon     string `json:"pcIcon"`
	PcIconSize string `json:"pcIconSize"`
}
type zhanQiAPIRespDataTagsSystem struct {
	PcIcon     string `json:"pcIcon"`
	PcIconSize string `json:"pcIconSize"`
}

type zhanQiAPIRespDataTags struct {
	Common *zhanQiAPIRespDataTagsCommon `json:"common"`
	System *zhanQiAPIRespDataTagsSystem `json:"system"`
}
type zhanQiAPIRespDataFlashvars struct {
	Servers     string `json:"Servers"`
	VideoLevels string `json:"VideoLevels"`
	Cdns        string `json:"cdns"`
	H5Cdns      string `json:"h5Cdns"`
	Status      int    `json:"Status"`
	RoomID      int    `json:"RoomId"`
	ComLayer    bool   `json:"ComLayer"`
	ComDef      bool   `json:"ComDef"`
	VideoTitle  string `json:"VideoTitle"`
	WebHost     string `json:"WebHost"`
	VideoType   string `json:"VideoType"`
	GameID      int    `json:"GameId"`
	Online      int    `json:"Online"`
	Pv          string `json:"pv"`
	TuristRate  int    `json:"TuristRate"`
	UseDlIP     int    `json:"UseDlIp"`
	Gpu         int    `json:"Gpu"`
	StarRoom    int    `json:"StarRoom"`
	GadObj      string `json:"GadObj"`
}
type zhanQiAPIRespDataPermission struct {
	Fans     bool `json:"fans"`
	Guess    bool `json:"guess"`
	Replay   bool `json:"replay"`
	Multi    bool `json:"multi"`
	Shift    bool `json:"shift"`
	Yf       bool `json:"yf"`
	Flycan   bool `json:"flycan"`
	ComDef   bool `json:"ComDef"`
	Firework bool `json:"firework"`
}
type zhanQiAPIRespDataAnchorAttrHots struct {
	Level          string `json:"level"`
	Fight          string `json:"fight"`
	NowLevelStart  string `json:"nowLevelStart"`
	NextLevelFight string `json:"nextLevelFight"`
}
type zhanQiAPIRespDataAnchorAttr struct {
	Hots       *zhanQiAPIRespDataAnchorAttrHots `json:"hots"`
	AnchorCard string                           `json:"anchorCard"`
}
type zhanQiAPIRespDataIsStar struct {
	Start   int `json:"start"`
	IsMonth int `json:"isMonth"`
}
type zhanQiAPIRespDataRoomFootWeight struct {
	InfoDisplayOrder int `json:"infoDisplayOrder"`
	NewsGames        int `json:"newsGames"`
	Recommend        int `json:"recommend"`
	Info             int `json:"info"`
	Extend           int `json:"extend"`
	Video            int `json:"video"`
}

type zhanQiAPIRespData struct {
	ID             string                           `json:"id"`
	UID            string                           `json:"uid"`
	Nickname       string                           `json:"nickname"`
	Gender         string                           `json:"gender"`
	Avatar         string                           `json:"avatar"`
	Code           string                           `json:"code"`
	URL            string                           `json:"url"`
	Title          string                           `json:"title"`
	GameID         string                           `json:"gameId"`
	Spic           string                           `json:"spic"`
	Bpic           string                           `json:"bpic"`
	Online         string                           `json:"online"`
	Status         string                           `json:"status"`
	Level          string                           `json:"level"`
	Type           string                           `json:"type"`
	LiveTime       string                           `json:"liveTime"`
	HotsLevel      string                           `json:"hotsLevel"`
	VideoID        string                           `json:"videoId"`
	Verscr         string                           `json:"verscr"`
	PublishURL     string                           `json:"publishUrl"`
	AnchorCoverImg string                           `json:"anchorCoverImg"`
	AnchorNotice   string                           `json:"anchorNotice"`
	RoomDesc       string                           `json:"roomDesc"`
	ChatStatus     string                           `json:"chatStatus"`
	Tags           *zhanQiAPIRespDataTags           `json:"tags"`
	Style          int                              `json:"style"`
	Legal          int                              `json:"legal"`
	Flashvars      *zhanQiAPIRespDataFlashvars      `json:"flashvars"`
	ClassID        string                           `json:"classId"`
	ClassName      string                           `json:"className"`
	ClassURL       string                           `json:"classUrl"`
	NewGameName    string                           `json:"newGameName"`
	FatherGameID   string                           `json:"fatherGameId"`
	FatherGameName string                           `json:"fatherGameName"`
	FatherGameURL  string                           `json:"fatherGameUrl"`
	Fid            string                           `json:"fid"`
	GameName       string                           `json:"gameName"`
	GameURL        string                           `json:"gameUrl"`
	GameIcon       string                           `json:"gameIcon"`
	GameBpic       string                           `json:"gameBpic"`
	Permission     *zhanQiAPIRespDataPermission     `json:"permission"`
	FansTitle      string                           `json:"fansTitle"`
	AnchorAttr     *zhanQiAPIRespDataAnchorAttr     `json:"anchorAttr"`
	Follows        string                           `json:"follows"`
	IsStar         *zhanQiAPIRespDataIsStar         `json:"isStar"`
	Bonus          bool                             `json:"bonus"`
	Fall           []interface{}                    `json:"fall"`
	Firepower      int                              `json:"firepower"`
	RoomFootWeight *zhanQiAPIRespDataRoomFootWeight `json:"roomFootWeight"`
}

func (p *parser) parseZhanQi(ctx context.Context, roomID int64) (string, error) {
	url := fmt.Sprintf("https://m.zhanqi.tv/api/static/v2.1/room/domain/%d.json", roomID)
	resp := &zhanQiAPIResp{}
	if err := p.doGet(ctx, url, resp); err != nil {
		return "", errors.WithStack(err)
	}
	if resp.Code != 0 {
		return "", errors.Errorf("got unexpected code:%d, message:%s", resp.Code, resp.Message)
	}
	if resp.Data == nil {
		return "", errors.New("got unexpected nil data")
	}
	return fmt.Sprintf("https://dlhdl-cdn.zhanqi.tv/zqlive/%s.flv", resp.Data.VideoID), nil
}
