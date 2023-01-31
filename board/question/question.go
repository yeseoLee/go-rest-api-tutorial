package question

import (
	"board/datasource"
	handler "board/question/handler"
	repository "board/question/repository"
	usecase "board/question/usecase"

	"github.com/labstack/echo/v4"
)

func RegistQuestionRoute(ds datasource.DataSource, e *echo.Echo) {
	// Question
	qr, err := repository.NewQuestionRepository(ds)
	if err != nil {
		panic(err)
	}
	qu := usecase.NewQuestionUseCase(qr)
	handler.NewQuestionHandler(e, qu)

}
