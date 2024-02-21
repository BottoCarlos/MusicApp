package disco

import (
	"encoding/json"
	"io"
	"os"

	"github.com/BottoCarlos/MusicApp/internal/domain"
)

type Repository struct {
	discos []domain.Disco
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Load() error {
	jsonFile, err := os.Open("db/data.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &r.discos)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetAll() []domain.Disco {
	return r.discos
}

func (r *Repository) GetByID(id int) *domain.Disco {
	for i := range r.discos {
		if r.discos[i].ID == id {
			return &r.discos[i]
		}
	}
	return nil
}

func (r *Repository) Add(disco domain.Disco) {
	r.discos = append(r.discos, disco)
}

func (r *Repository) Update(disco domain.Disco) {
	for i := range r.discos {
		if r.discos[i].ID == disco.ID {
			r.discos[i] = disco
			return
		}
	}
}

func (r *Repository) Delete(id int) {
	for i := range r.discos {
		if r.discos[i].ID == id {
			r.discos = append(r.discos[:i], r.discos[i+1:]...)
			return
		}
	}
}
