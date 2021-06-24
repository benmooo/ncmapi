package apitypes

type SearchResult struct {
	Code   int `json:"code"`
	Result H   `json:"result"`
}

type SearchSongResult struct {
	Code   int `json:"code"`
	Result struct {
		Songs []Song `json:"songs"`
	} `json:"result"`
}

type SearchCollectionResult struct {
	Code   int `json:"code"`
	Result struct {
		Collection []Song `json:"playlists"`
	} `json:"result"`
}

type SearchArtistResult struct {
	Code   int `json:"code"`
	Result struct {
		Artists []Artist `json:"artists"`
	} `json:"result"`
}

type SearchAlbumResult struct {
	Code   int `json:"code"`
	Result struct {
		Albums []Album `json:"albums"`
	} `json:"result"`
}

// 1: 单曲, 10: 专辑, 100: 歌手, 1000: 歌单, 1002: 用户, 1004: MV, 1006: 歌词, 1009: 电台, 1014: 视频
const (
	SearchTypeSong       = 1
	SearchTypeAlbum      = 10
	SearchTypeArtist     = 100
	SearchTypeCollection = 1000
	SearchTypeUser       = 1002
	SearchTypeMV         = 1004
	SearchTypeLyric      = 1006
	SearchTypePodcast    = 1009
	SearchTypeVideo      = 1014
)

// type SearchPayload struct {
// 	Keyword string `json:"s"`
// 	Type    int    `json:"type"`
// 	Limit   int    `json:"limit"`
// 	Offset  int    `json:"offset"`
// }
