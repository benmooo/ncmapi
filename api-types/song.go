package apitypes

type Song struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Album    Album    `json:"al"`
	Artists  []Artist `json:"ar"`
	Duration int      `json:"dt"`
}
