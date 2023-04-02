package album

type Album struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Author      string `json:"author,omitempty"`
	Year        int    `json:"year,omitempty"`
	CoverURL    string `json:"coverURL,omitempty"`
	Description string `json:"description,omitempty"`
	Duration    string `json:"duration,omitempty"`
	Songs       []Song `json:"songs,omitempty"`
	GeniousLink string `json:"geniousLink,omitempty"`
}

type Song struct {
	Title    string `json:"title,omitempty"`
	Duration string `json:"duration,omitempty"`
}
