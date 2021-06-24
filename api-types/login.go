package apitypes

type LoginResult struct {
	Code    int     `json:"code"`
	Account Account `json:"account"`
	Token   string  `json:"token"`
	Profile Profile `json:"profile"`

	Bindings []Binding `json:"bindings"`
	Cookie   string    `json:"cookie"`
}

type Account struct {
	ID                 int    `json:"id"`
	Username           string `json:"userName"`
	Type               int    `json:"type"`
	Status             int    `json:"status"`
	WhitelistAuthority int    `json:"whitelistAuthority"`
	CreateTime         int    `json:"createTime"`
	Salt               string `json:"salt"`
	TokenVersion       int    `json:"tokenVersion"`
	Ban                int    `json:"ban"`
	BaoyueVersion      int    `json:"baoyueVersion"`
	DonateVersion      int    `json:"donateVersion"`
	VipType            int    `json:"vipType"`
	ViptypeVersion     int    `json:"viptypeVersion"`
	AnonimousUser      bool   `json:"anonimousUser"`
}

type Profile struct {
	City                      int    `json:"city"`
	Description               string `json:"description"`
	DjStatus                  int    `json:"djStatus"`
	AvatarImgId               int    `json:"avatarImgId"`
	AvatarUrl                 string `json:"avatarUrl"`
	BackgroundImgId           int    `json:"backgroundImgId"`
	VipType                   int    `json:"vipType"`
	Province                  int    `json:"province"`
	AccountStatus             int    `json:"accountStatus"`
	Gender                    int    `json:"gender"`
	Mutual                    bool   `json:"mutual"`
	AuthStatus                int    `json:"authStatus"`
	Birthday                  int    `json:"birthday"`
	DefaultAvatar             bool   `json:"defaultAvatar"`
	Nickname                  string `json:"nickname"`
	Followed                  bool   `json:"followed"`
	BackgroundUrl             string `json:"backgroundUrl"`
	DetailDescription         string `json:"detailDescription"`
	UserId                    int    `json:"userId"`
	UserType                  int    `json:"userType"`
	AvatarImgIdStr            string `json:"avatarImgIdStr"`
	BackgroundImgIdStr        string `json:"backgroundImgIdStr"`
	Signature                 string `json:"signature"`
	Authority                 int    `json:"authority"`
	AvatarImgId_str           string `json:"avatarImgId_str"`
	Followeds                 int    `json:"followeds"`
	Follows                   int    `json:"follows"`
	EventCount                int    `json:"eventCount"`
	PlaylistCount             int    `json:"playlistCount"`
	PlaylistBeSubscribedCount int    `json:"playlistBeSubscribedCount"`
}

type Binding struct {
	ExpiresIn    int    `json:"expiresIn"`
	TokenJsonStr string `json:"tokenJsonStr"`
	RefreshTime  int    `json:"refreshTime"`
	BindingTime  int    `json:"bindingTime"`
	Expired      bool   `json:"expired"`
	UserId       int    `json:"userId"`
	Url          string `json:"url"`
	Id           int    `json:"id"`
	Type         int    `json:"type"`
}
