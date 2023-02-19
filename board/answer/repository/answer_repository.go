package answer

import (
	"database/sql"

	"board/datasource"
	model "board/domain"
)

func NewanswerRepository(ds datasource.DataSource) (*answerRepository, error) {
	db, err := ds.GetConnection()
	if err != nil {
		return nil, err
	}
	return &answerRepository{DBEngine: db}, nil
}

type answerRepository struct {
	DBEngine *sql.DB
}

func (r *answerRepository) FindAllByQuestionId(id uint) ([]*model.Answer, error) {
	// var q model.Answer

	// row := r.DBEngine.QueryRow("SELECT id, title, content, writerId, images, createdAt,updatedAt FROM tbanswer WHERE id = ?", id)
	// err := row.Scan(&q.Id, &q.Title, &q.Content, &q.WriterId, &q.Images, &q.CreatedAt, &q.UpdatedAt)
	// switch {
	// case err == sql.ErrNoRows:
	// 	return nil, errors.New("no data")
	// case err != nil:
	// 	return nil, errors.New("query error")
	// }
	// return &q, nil
	return nil, nil
}

func (r *answerRepository) FindAllByWriterId(limit, offset int) ([]*model.Answer, error) {
	// var qList []*model.Answer

	// rows, err := r.DBEngine.Query("SELECT id, title, content, writerId, images, createdAt,updatedAt FROM tbanswer LIMIT ? OFFSET ?", limit, offset)
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var q model.Answer
	// 	rows.Scan(&q.Id, &q.Title, &q.Content, &q.WriterId, &q.Images, &q.CreatedAt, &q.UpdatedAt)
	// 	qList = append(qList, &q)
	// }

	// return qList, nil
	return nil, nil
}

func (r *answerRepository) Create(answerInput *model.AnswerInput) (*model.Answer, error) {
	// var answer *model.Answer
	// result, err := r.DBEngine.Exec("INSERT INTO `tbanswer`(`Title`,`Content`,`WriterId`,`Images`,`CreatedAt`) VALUES (?, ?, ?, ?, now())",
	// 	answerInput.Title, answerInput.Content, answerInput.WriterId, answerInput.Images)
	// if err != nil {
	// 	return answer, err
	// }

	// id, err := result.LastInsertId()
	// answer.Id = uint(id)
	// if err != nil {
	// 	return answer, err
	// }
	// return answer, nil
	return nil, nil
}

func (r *answerRepository) Update(id uint, answerUpdate map[string]interface{}) (*model.Answer, error) {
	// // TODO: map key-value check & make query logic
	// var answer *model.Answer
	// _, err := r.DBEngine.Exec("UPDATE tbanswer SET Title = ?, Content = ?, Images = ?, updatedAt = now() WHERE id = ?",
	// 	answerUpdate["Title"], answerUpdate["Content"], answerUpdate["Images"])
	// if err != nil {
	// 	return answer, err
	// }
	// return answer, nil
	return nil, nil
}

func (r *answerRepository) Delete(id uint) error {
	// result, err := r.DBEngine.Exec("DELETE FROM `tbanswer` WHERE id = ?", id)
	// if err != nil {
	// 	return err
	// }

	// rows, err := result.RowsAffected()
	// if err != nil {
	// 	return err
	// }
	// if rows != 1 {
	// 	return fmt.Errorf("expected to affect 1 row, affected %d", rows)
	// }
	// return err
	return nil
}
