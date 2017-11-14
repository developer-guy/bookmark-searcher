package bookmark

type Bookmarks struct {
	Checksum string `json:"checksum"`
	Roots    roots  `json:"roots"`
	Version  int    `json:"version"`
}

type roots struct {
	BookmarkBar              bookmark `json:"bookmark_bar"`
	Other                    others   `json:"other"`
	Sync_Transaction_Version string   `json:"sync_transaction_version"`
	Synced                   others   `json:"synced"`
}

type bookmark struct {
	Children []Children `json:"children"`
}

type Children struct {
	Children                 []bookmarkItem `json:"children"`
	Date_Added               string         `json:"date_added"`
	Date_Modified            string         `json:"date_modified"`
	Id                       string         `json:"id"`
	Name                     string         `json:"name"`
	Sync_Transaction_Version string         `json:"sync_transaction_version"`
	Type                     string         `json:"type"`
}

type others struct {
	Children      []bookmarkItem `json:"children"`
	Date_Added    string         `json:"date_added"`
	Date_Modified string         `json:"date_modified"`
	Id            string         `json:"id"`
	Name          string         `json:"name"`
	Type          string         `json:"type"`
}

type bookmarkItem struct {
	Date_Added               string   `json:"date_added"`
	Id                       string   `json:"id"`
	Meta_Info                metaInfo `json:"meta_info"`
	Name                     string   `json:"name"`
	Sync_Transaction_Version string   `json:"sync_transaction_version"`
	Type                     string   `json:"type"`
	Url                      string   `json:"url"`
}

type metaInfo struct {
	Last_Visited_Desktop string `json:"last_visited_desktop"`
}
