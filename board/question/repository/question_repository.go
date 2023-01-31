package question

import (
	"database/sql"
	"fmt"

	"board/datasource"
	model "board/domain"
)

func NewQuestionRepository(ds datasource.DataSource) (*QuestionRepository, error) {
	db, err := ds.GetConnection()
	if err != nil {
		return nil, err
	}
	return &QuestionRepository{DBEngine: db}, nil
}

type QuestionRepository struct {
	DBEngine *sql.DB
}

func (r *QuestionRepository) GetByID(id int64) (*model.Question, error) {
	fmt.Println(">> ", r.DBEngine)
	return &model.Question{
		Title:    "MySQL 사용법",
		Body:     "MySQL은 쉽게 사용 할 수 있습니다.\n정말로",
		UserName: "yundream",
	}, nil
}

func (r *QuestionRepository) Fetch(offset, limit int) ([]*model.Question, error) {
	fmt.Println(">> ", r.DBEngine)
	return []*model.Question{
		{
			Title:    "MySQL 사용법",
			Body:     "MySQL은 쉽게 사용 할 수 있습니다.\n정말로",
			UserName: "yundream",
		},
		{
			Title:    "GoLang의 미래",
			Body:     "GoLang의 미래는 밝아보입니다....",
			UserName: "yundream",
		},
	}, nil
}

func (r *QuestionRepository) Create(question *model.Question) (*model.Question, error) {
	return question, nil
}

func (r *QuestionRepository) Update(id int64, question *model.Question) (*model.Question, error) {
	return &model.Question{Title: "Update"}, nil
}

func (r *QuestionRepository) Delete(id int64) error {
	fmt.Println("삭제")
	return nil
}
