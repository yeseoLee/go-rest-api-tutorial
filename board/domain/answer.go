package domain

type Answer struct {
	Id         int64    `json:"id"`
	QuestionId int64    `json:"questionId"`
	Content    string   `json:"content"`
	Writer     string   `json:"writer"`
	WriterId   string   `json:"writerId"`
	Images     []string `json:"images"`
	CreateTime string   `json:"createTime"`
	UpdateTime string   `json:"updateTime"`
}

type AnswerRepository interface {
	GetByID(id int64) (*Answer, error)
	Fetch(offset, limit int) ([]*Answer, error)
	Create(answer *Answer) (*Answer, error)
	Update(id int64, answer *Answer) (*Answer, error)
	Delete(id int64) error
}

type AnswerUseCase interface {
	GetByID(id int64) (res *Answer, err error)
	Fetch(offset, limit int) (res []*Answer, err error)
	Create(answer *Answer) (*Answer, error)
	Update(id int64, answer *Answer) (*Answer, error)
	Delete(id int64) error
}
