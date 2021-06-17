package handler

import (
	"fmt"

	"github.com/afistapratama12/projectBackend/helper"
	"github.com/afistapratama12/projectBackend/note"
	"github.com/afistapratama12/projectBackend/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type noteHandler struct {
	service note.Service
}

func NewNoteHandler(service note.Service) *noteHandler {
	return &noteHandler{service}
}

func (h *noteHandler) GetAllNote(c *gin.Context) {
	notes, err := h.service.GetAll()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	noteFormat := helper.FormatNotes(notes)

	c.JSON(200, noteFormat)
}

func (h *noteHandler) GetAllNoteByUser(c *gin.Context) {
	userLogin := c.MustGet("currentUser").(user.User)

	// fmt.Println("line 36 notehandler ", userLogin)

	if userLogin.ID == "" || len(userLogin.ID) <= 1 {
		errResponse := gin.H{"error": "unauthorize user"}
		c.JSON(401, errResponse)
		return
	}

	notes, err := h.service.GetByUserLogin(userLogin.ID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	noteFormat := helper.FormatNotes(notes)

	c.JSON(200, noteFormat)
}

func (h *noteHandler) GetByIDNote(c *gin.Context) {
	noteID := c.Params.ByName("note_id")

	note, err := h.service.GetByID(noteID)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	formatter := helper.FormatNote(note)

	c.JSON(200, formatter)
}

func (h *noteHandler) SaveNewNote(c *gin.Context) {
	var inputNote note.NoteInput

	userLogin := c.MustGet("currentUser").(user.User)

	if userLogin.ID == "" || len(userLogin.ID) <= 1 {
		errResponse := gin.H{"error": "unauthorize user"}
		c.JSON(401, errResponse)
		return
	}

	if err := c.ShouldBindJSON(&inputNote); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}

	fmt.Println(inputNote)

	noteID := uuid.New()

	note, err := h.service.SaveNewNote(noteID.String(), userLogin.ID, inputNote)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	formatter := helper.FormatNote(note)

	c.JSON(201, formatter)
}

func (h *noteHandler) UpdateNote(c *gin.Context) {
	var noteID = c.Params.ByName("note_id")

	var updateInput note.NoteUpdateInput

	userLogin := c.MustGet("currentUser").(user.User)

	if userLogin.ID == "" || len(userLogin.ID) <= 1 {
		errResponse := gin.H{"error": "unauthorize user"}
		c.JSON(401, errResponse)
		return
	}

	if err := c.ShouldBindJSON(&updateInput); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}

	checkNote, _ := h.service.GetByID(noteID)

	if checkNote.ID != userLogin.ID || userLogin.Role != "admin" {
		errResponse := gin.H{"error": "unauthorize user"}
		c.JSON(401, errResponse)
		return
	}

	note, err := h.service.UpdateNote(noteID, updateInput)

	if err != nil {
		errResponse := gin.H{"error": err.Error()}
		c.JSON(500, errResponse)
		return
	}

	formatter := helper.FormatNote(note)

	c.JSON(200, formatter)
}

func (h *noteHandler) DeleteNote(c *gin.Context) {
	var delInput note.NoteDeleteInput

	var noteID = c.Params.ByName("note_id")

	userLogin := c.MustGet("currentUser").(user.User)

	if userLogin.ID == "" || len(userLogin.ID) <= 1 {
		errResponse := gin.H{"error": "unauthorize user"}
		c.JSON(401, errResponse)
		return
	}

	if err := c.ShouldBindJSON(&delInput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	checkNote, _ := h.service.GetByID(noteID)

	// masih terjadi error disini untuk verification
	if checkNote.UserID == userLogin.ID || userLogin.Role == "admin" {
		if checkNote.Secret != delInput.Secret {
			c.JSON(400, gin.H{"error": "secret note invalid, cannot delete note"})
			return
		}

		deleteMsg, err := h.service.DeleteNote(noteID)

		if err != nil {
			errResponse := gin.H{"error": err.Error()}
			c.JSON(500, errResponse)
			return
		}

		c.JSON(200, gin.H{"message": deleteMsg.(string)})
		return
	}

	errResponse := gin.H{"error": "unauthorize user"}
	c.JSON(401, errResponse)
}

func (h *noteHandler) UnDeleteNote(c *gin.Context) {
	var noteID = c.Params.ByName("note_id")

	userLogin := c.MustGet("currentUser").(user.User)

	if userLogin.Role != "admin" {
		errResponse := gin.H{"error": "unauthorize user not admin"}
		c.JSON(401, errResponse)
		return
	}

	undeleteMsg, err := h.service.UnDeleteNote(noteID)

	if err != nil {
		errResponse := gin.H{"error": err.Error()}
		c.JSON(500, errResponse)
		return
	}

	c.JSON(200, gin.H{"message": undeleteMsg.(string)})
}
