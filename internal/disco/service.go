package disco

import "github.com/BottoCarlos/MusicApp/internal/domain"

type Service struct {
	repo RepositoryInterface
}

func NewService(repo RepositoryInterface) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAll() []domain.Disco {
	return s.repo.GetAll()
}

func (s *Service) GetByID(id int) *domain.Disco {
	return s.repo.GetByID(id)
}

func (s *Service) Add(disco domain.Disco) {
	s.repo.Add(disco)
}

func (s *Service) Update(disco domain.Disco) {
	s.repo.Update(disco)
}

func (s *Service) Delete(id int) {
	s.repo.Delete(id)
}
