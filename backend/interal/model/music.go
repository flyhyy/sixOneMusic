package model

type MusicAllResponse struct {
	PageSize int         `json:"page_size"`
	Page     int         `json:"page"`
	Data     []MusicBase `json:"data"`
}

type MusicBase struct {
	Title     string `json:"title"`
	Singer    string `json:"singer"`
	Duration  int    `json:"duration"`
	Size      int64  `json:"size"`
	ID        uint   `json:"id"`
	CoverPath string `json:"cover_path"`
	IsCollect bool   `json:"is_collect"`
	AlbumName string `json:"album_name"`
}

type MusicCollectRequest struct {
	SongID    uint `json:"song_id"`
	IsCollect bool `json:"is_Collect"`
}
