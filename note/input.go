package note

type NoteInput struct {
	Title  string `json:"title" binding:"required"`
	Body   string `json:"body" binding:"required"`
	Secret string `json:"secret"`
	Type   string `json:"type" binding:"required"`
}

type NoteUpdateInput struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Secret string `json:"secret"`
	Type   string `json:"type"`
}

type NoteDeleteInput struct {
	Secret string `json:"secret"`
}
