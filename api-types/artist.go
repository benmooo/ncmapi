package apitypes

type Artist struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AccountID int    `json:"accountId"`
	Cover     string `json:"picUrl"`
	AlbumSize int    `json:"albumSize"`
	MvSize    int    `json:"mvSize"`
	Followed  bool   `json:"followed"`
}
