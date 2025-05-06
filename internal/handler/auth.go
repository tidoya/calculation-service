package handler

import (
	Calculation "calculation-service/internal/entity/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signIn(w *gin.Context) {

}

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body todo.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input Calculation.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Autorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

	// err := r.ParseForm()
	// if err != nil {
	// 	http.Error(w, "Error parsing form", http.StatusBadRequest)
	// 	return
	// }
	// username := r.FormValue("username")
	// password := r.FormValue("password")
	// fmt.Print(username, password)
	// w.Header().Set("Content-Type", "text/plain")
	// fmt.Fprintf(w, "Received username: %s have %s pass", username, password)
}

func (h *Handler) logout(w *gin.Context) {

}
