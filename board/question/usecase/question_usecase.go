package question

import "board/domain"

type questionUsecase struct {
	questionRepo domain.QuestionRepository
}

func NewQuestionUseCase(u domain.QuestionRepository) domain.QuestionUseCase {
	return &questionUsecase{
		questionRepo: u,
	}
}

func (u *questionUsecase) GetByID(id int64) (*domain.Question, error) {
	res, err := u.questionRepo.GetByID(id)
	return res, err
}
func (u *questionUsecase) Fetch(offset, limit int) ([]*domain.Question, error) {
	return u.questionRepo.Fetch(offset, limit)
}

func (u *questionUsecase) Create(question *domain.Question) (*domain.Question, error) {
	return u.questionRepo.Create(question)
}

func (u *questionUsecase) Update(id int64, question *domain.Question) (*domain.Question, error) {
	return u.questionRepo.Update(id, question)
}

func (u *questionUsecase) Delete(id int64) error {
	return u.questionRepo.Delete(id)
}
