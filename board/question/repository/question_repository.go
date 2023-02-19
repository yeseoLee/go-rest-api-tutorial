package question

import (
	"database/sql"
	"errors"
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

func (r *QuestionRepository) FindById(id uint) (*model.Question, error) {
	var q model.Question

	row := r.DBEngine.QueryRow("SELECT id, title, content, writerId, images, createdAt,updatedAt FROM tbQuestion WHERE id = ?", id)
	err := row.Scan(&q.Id, &q.Title, &q.Content, &q.WriterId, &q.Images, &q.CreatedAt, &q.UpdatedAt)
	switch {
	case err == sql.ErrNoRows:
		return nil, errors.New("no data")
	case err != nil:
		return nil, errors.New("query error")
	}
	return &q, nil
}

func (r *QuestionRepository) FindAll(limit, offset int) ([]*model.Question, error) {
	var qList []*model.Question

	rows, err := r.DBEngine.Query("SELECT id, title, content, writerId, images, createdAt,updatedAt FROM tbQuestion LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var q model.Question
		rows.Scan(&q.Id, &q.Title, &q.Content, &q.WriterId, &q.Images, &q.CreatedAt, &q.UpdatedAt)
		qList = append(qList, &q)
	}

	return qList, nil
}

func (r *QuestionRepository) Create(questionInput *model.QuestionInput) (*model.Question, error) {
	var question *model.Question
	result, err := r.DBEngine.Exec("INSERT INTO `tbQuestion`(`Title`,`Content`,`WriterId`,`Images`,`CreatedAt`) VALUES (?, ?, ?, ?, now())",
		questionInput.Title, questionInput.Content, questionInput.WriterId, questionInput.Images)
	if err != nil {
		return question, err
	}

	id, err := result.LastInsertId()
	question.Id = uint(id)
	if err != nil {
		return question, err
	}
	return question, nil
}

func (r *QuestionRepository) Update(id uint, questionUpdate map[string]interface{}) (*model.Question, error) {
	// TODO: map key-value check & make query logic
	var question *model.Question
	_, err := r.DBEngine.Exec("UPDATE tbQuestion SET Title = ?, Content = ?, Images = ?, updatedAt = now() WHERE id = ?",
		questionUpdate["Title"], questionUpdate["Content"], questionUpdate["Images"])
	if err != nil {
		return question, err
	}
	return question, nil
}

func (r *QuestionRepository) Delete(id uint) error {
	result, err := r.DBEngine.Exec("DELETE FROM `tbQuestion` WHERE id = ?", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return fmt.Errorf("expected to affect 1 row, affected %d", rows)
	}
	return err
}
