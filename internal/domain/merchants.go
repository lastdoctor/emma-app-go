package domain

type Merchants struct {
	ID          int64  `json:"id"`
	DisplayName string `json:"display_name"`
	IconUrl     string `json:"icon_url"`
	FunnyGifUrl string `json:"funny_gif_url"`
}
