package models

type Branch_Offices []Branch_Office

type Branch_Office struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Location  string   `json:"location"`
	Telephone string   `json:"telephone"`
	Articles  Articles `json:"articles"`
}
