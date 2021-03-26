package dto

//CVUpdateDTO is a model that client use when updating a CV
type CVUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}

//CVCreateDTO is is a model that clinet use when create a new CV
type CVCreateDTO struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}
