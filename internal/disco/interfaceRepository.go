package disco

import "github.com/BottoCarlos/MusicApp/internal/domain"

type RepositoryInterface interface {
	Load() error
	GetAll() []domain.Disco
	GetByID(id int) *domain.Disco
	Add(disco domain.Disco)
	Update(disco domain.Disco)
	Delete(id int)
}
