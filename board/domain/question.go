package domain

type Question struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	UserName  string `json:"userName"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type QuestionRepository interface {
	GetByID(id int64) (*Question, error)
	Fetch(offset, limit int) ([]*Question, error)
	Create(question *Question) (*Question, error)
	Update(id int64, question *Question) (*Question, error)
	Delete(id int64) error
}

type QuestionUseCase interface {
	GetByID(id int64) (res *Question, err error)
	Fetch(offset, limit int) (res []*Question, err error)
	Create(question *Question) (*Question, error)
	Update(id int64, question *Question) (*Question, error)
	Delete(id int64) error
}
