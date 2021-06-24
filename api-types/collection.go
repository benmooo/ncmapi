package apitypes

type Collection struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserID int    `json:"userId"`

	Cover       string `json:"coverImgUrl"`
	Subscribers int    `json:"bookCount"`
	Description string `json:"description"`

	Subscribed  bool `json:"subscribed"`
	SongCount   int  `json:"trackCount"`
	SpecialType int  `json:"specialType"`

	HighQuality bool `json:"highQuality"`
	PlayCount   int  `json:"playCount"`
	// Creator struct {nickname: "", userId: 249810913, userType: 4, authStatus: 1, expertTags: null, experts: null}
	// OfficialTags string `json:"officialTags"`
}
