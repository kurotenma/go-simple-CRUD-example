package gameDTO

type InsertGameRequest struct {
	Title       string `form:"title" validate:"required"`
	Url         string `form:"url"`
	Platform    string `form:"platform" validate:"required"`
	Description string `form:"description" validate:"required"`
}

type UpdateGameRequest struct {
	Title       string `form:"title" validate:"required"`
	Url         string `form:"url"`
	Platform    string `form:"platform" validate:"required"`
	Description string `form:"description" validate:"required"`
}

type GetGamesFilterRequest struct {
	Title         string   `form:"title"`
	Url           string   `form:"url"`
	Platform      []string `form:"platform"`
	Description   string   `form:"description"`
	Status        []string `form:"status"`
	DeletedStatus []string `form:"deleted_status"`
	PerPage       int      `form:"per_page"`
	Page          int      `form:"page"`
}

type GetGameResponse struct {
	CreatedAt   string `json:"created_at"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	Platform    string `json:"platform"`
	Description string `json:"description"`
	Status      string `json:"status"`
	IsDeleted   bool   `json:"is_deleted"`
}
type GetGamesResponse []GetGameResponse
