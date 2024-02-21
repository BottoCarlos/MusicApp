package disco

import "github.com/BottoCarlos/MusicApp/internal/domain"

type ServiceInterface interface {
	GetAll() []domain.Disco
	GetByID(id int) *domain.Disco
	Add(disco domain.Disco)
	Update(disco domain.Disco)
	Delete(id int)
}
