package mobile

type ArticleOverview struct {
	ID        *string `json:"id"`
	Source    *string `json:"source"`
	Starred   *bool   `json:"starred"`
	Timestamp *string `json:"timestamp"`
	Title     *string `json:"title"`
}

type Article struct {
	Body     *string          `json:"body"`
	ImageURL *string          `json:"imageUrl"`
	Overview *ArticleOverview `json:"overview"`
}

type Category struct {
	ID       string `json:"id"`
	ImageURL string `json:"imageUrl"`
	Title    string `json:"title"`
}

type ListArticlesResponse struct {
	Articles []*ArticleOverview `json:"articles"`
}

type ListCategoriesResponse struct {
	Categories []*Category `json:"categories"`
}
