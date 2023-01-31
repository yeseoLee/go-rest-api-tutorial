package question

import (
	"board/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type QuestionHandler struct {
	QUseCase domain.QuestionUseCase
}

func NewQuestionHandler(e *echo.Echo, us domain.QuestionUseCase) {
	handler := QuestionHandler{
		QUseCase: us,
	}
	e.GET("/questions", handler.GetQuestions)
	e.GET("/question/{id}", handler.GetQuestion)
	e.POST("/question", handler.AddQuestion)
	e.PUT("/question", handler.EditQuestion)
	e.DELETE("/question", handler.RemoveQuestion)
}

func (h *QuestionHandler) GetQuestions(c echo.Context) error {
	return c.String(http.StatusOK, "getquestions")
}

func (h *QuestionHandler) GetQuestion(c echo.Context) error {
	/*
		func (a *ArticleHandler) GetArticle(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			id, _ := strconv.ParseInt(vars["id"], 10, 16)
			article, _ := a.AUseCase.GetByID(id)
			response, _ := json.Marshal(article)
			fmt.Fprint(w, string(response))
		}
	*/
	return c.String(http.StatusOK, "getquestion")
}

func (h *QuestionHandler) AddQuestion(c echo.Context) error {
	return c.String(http.StatusOK, "addquestion")
}

func (h *QuestionHandler) EditQuestion(c echo.Context) error {
	return c.String(http.StatusOK, "editquestion")
}

func (h *QuestionHandler) RemoveQuestion(c echo.Context) error {
	return c.String(http.StatusOK, "removequestion")
}
