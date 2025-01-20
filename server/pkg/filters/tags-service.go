package filters

import (
	"database/sql"

	"github.com/Romma711/ozora_web_ecommerse/server/pkg/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateCategory(category *types.Category) error {
	_, err := s.db.Query(`INSERT INTO categories (name) VALUES (?)`, category.Name)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetCategories() ([]types.Category, error) {
	rows, err := s.db.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	tags, err := scanRowsIntoTags(rows)
	if err != nil {
		return nil, err
	}
	var categories []types.Category
	for i := 0; i < len(tags); i++ {
		var category types.Category
		category.ID = tags[i].ID
		category.Name = tags[i].Name
		categories = append(categories, category)
	}
	return categories, nil
}

func (s *Store) GetCategoryById(id int) (string, error) {
	row, err := s.db.Query("SELECT name FROM categories WHERE id = ?", id)
	if err != nil {
		return "", err
	}
	var name string
	for row.Next() {
		err = row.Scan(&name)
		if err != nil {
			return "", err
		}
	}
	return name, nil
}

func (s *Store) CreateType(type_ *types.Type) error {
	_, err := s.db.Exec("INSERT INTO types (name) VALUES (?)", type_.Name)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetTypes() ([]types.Type, error) {
	rows, err := s.db.Query("SELECT * FROM types")
	if err != nil {
		return nil, err
	}
	tags, err := scanRowsIntoTags(rows)
	if err != nil {
		return nil, err
	}
	var types_ []types.Type
	for i := 0; i < len(tags); i++ {
		var type_ types.Type
		type_.ID = tags[i].ID
		type_.Name = tags[i].Name
		types_ = append(types_, type_)
	}
	return types_, nil
}

func (s *Store) GetTypeById(id int) (string, error) {
	row, err := s.db.Query("SELECT name FROM types WHERE id = ?", id)
	if err != nil {
		return "", err
	}
	var name string
	for row.Next() {
		err = row.Scan(&name)
		if err != nil {
			return "", err
		}
	}
	return name, nil
}

func (s *Store) CreateArtWork(artWork *types.ArtWork) error {
	_, err := s.db.Exec("INSERT INTO artworks (title) VALUES (?)", artWork.Title)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetArtWorks() ([]types.ArtWork, error) {
	rows, err := s.db.Query("SELECT * FROM artworks")
	if err != nil {
		return nil, err
	}
	artWorks, err := scanRowsIntoArtWorks(rows)
	if err != nil {
		return nil, err
	}
	return artWorks, nil
}

func (s *Store) GetArtWorkById(id int) (string, error) {
	row, err := s.db.Query("SELECT title FROM artworks WHERE id = ?", id)
	if err != nil {
		return "", err
	}
	var name string
	for row.Next() {
		err = row.Scan(&name)
		if err != nil {
			return "", err
		}
	}
	return name, nil
}

func (s *Store) GetArtWorkRecomendation(number int) (*types.ArtWork, error) {
	row, err := s.db.Query("SELECT * FROM artworks WHERE id = ?", number)
	if err != nil {
		return nil ,err
	}
	
	artWork := new(types.ArtWork)
	for row.Next() {
		err = row.Scan(
			&artWork.ID,
			&artWork.Title,
			&artWork.Description,
			&artWork.Image,
			&artWork.Logo,
		)
		if err != nil {
			return nil, err
		}	
	}

	return artWork, nil
}
func scanRowsIntoTags(rows *sql.Rows) ([]types.Tag, error) {
	tags := make([]types.Tag, 0)
	tag := new(types.Tag)
	for rows.Next() {
		err := rows.Scan(
			&tag.ID,
			&tag.Name,
		)
		if err != nil {
			return nil, err
		}
		tags = append(tags, *tag)
	}
	return tags, nil
}

func scanRowsIntoArtWorks(rows *sql.Rows) ([]types.ArtWork, error) {
	artWorks := make([]types.ArtWork, 0)
	artWork := new(types.ArtWork)
	for rows.Next() {
		err := rows.Scan(
			&artWork.ID,
			&artWork.Title,
			&artWork.Description,
			&artWork.Image,
			&artWork.Logo,
		)
		if err != nil {
			return nil, err
		}
		artWorks = append(artWorks, *artWork)
	}
	return artWorks, nil
}