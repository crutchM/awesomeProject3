package entity

type PostVideoRequest struct {
	Name       string `json:"name" binding:"required"`
	CustomerID string `json:"author"`
}

type VideoResponse struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	CreationDate string `json:"creationDate"`
	Description  string `json:"description"`
	CustomerID   string `json:"customerID"`
	Watches      int    `json:"watches"`
	Preview      string `json:"preview"`
}

type VideoMinimized struct {
	CustomerID   string `json:"customerID"`
	Name         string `json:"name"`
	CreationDate string `json:"creationDate"`
	Preview      string `json:"preview"`
	Watches      int    `json:"watches"`
}

type VideoDTO struct {
	Id           string
	Name         string
	CreationDate string
	Description  string
	CustomerID   string
	Watches      int
	Preview      string
}
