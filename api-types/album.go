package apitypes

type Album struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Cover string `json:"picUrl"`
	Size  int    `json:"size"`
	Tags  string `json:"tags"`
	Type  string `json:"type"`
}
