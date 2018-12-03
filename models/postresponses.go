package models

type UpdateProfileForm struct{
    PhoneNumber string      `json:"PhoneNumber" form:"PhoneNumber"`
    Location    string      `json:"Location" form:"Location"`
    Courses     []string    `json:"Courses[]" form:"Courses[]"`
    Languages   []string    `json:"Languages[]" form:"Languages[]"`
}
