package domain

// input(DTO)

// output(VO)

// DBStruct(DAO)
type Question struct {
	Id         int64    `json:"id"`
	Title      string   `json:"title"`
	Content    string   `json:"body"`
	Writer     string   `json:"writer"`
	WriterId   string   `json:"writerId"`
	Images     []string `json:"images"`
	CreateTime string   `json:"createTime"`
	UpdateTime string   `json:"updateTime"`
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
