package question

import (
	"database/sql"
	"time"

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
	//r.DBEngine
	return &model.Question{}, nil
}

func (r *QuestionRepository) Fetch(offset, limit int) ([]*model.Question, error) {
	return []*model.Question{}, nil
}

func (r *QuestionRepository) Create(question *model.Question) (*model.Question, error) {
	question.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	result, err := r.DBEngine.Exec("INSERT INTO `tbQuestion`(`Title`,`Content`,`Writer`,`WriterId`,`Images`,`CreateTime`) VALUES (?, ?, ?, ?, ?, ?)",
		question.Title, question.Content, question.Writer, question.WriterId, question.Images, question.CreateTime)
	if err != nil {
		return question, err
	}
	question.Id, err = result.LastInsertId()
	if err != nil {
		return question, err
	}
	return question, nil
}

func (r *QuestionRepository) Update(id int64, question *model.Question) (*model.Question, error) {
	return &model.Question{}, nil
}

func (r *QuestionRepository) Delete(id int64) error {
	return nil
}
