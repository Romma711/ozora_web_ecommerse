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


//CATEGORY FUNCTIONS//

//Esta es una funcion que crea una categoria en la base de datos
func (s *Store) CreateCategory(category *types.Category) error {
	_, err := s.db.Query(`INSERT INTO categories (name) VALUES (?)`, category.Name)
	if err != nil {
		return err
	}
	return nil
}

//Esta funcion recupera todas las categorias de la base de datos
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

//Esta funcion recupera una categoria en especifico
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

func (s *Store) UpdateCategory(category types.Category) error{
	_, err := s.db.Query("UPDATE categories SET name = ? WHERE id = ?", category.Name, category.ID)
	if err != nil {
		return err
	}
	return nil
}

//TYPES FUNCTIONS//

//Esta funcion crea un tipo de producto en la base de datos
func (s *Store) CreateType(type_ *types.Type) error {
	_, err := s.db.Exec("INSERT INTO types (name) VALUES (?)", type_.Name)
	if err != nil {
		return err
	}
	return nil
}

//esta funcion recupera todos los tipos de productos de la base de datos
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

//Esta funcion recupera un tipo de producto en especifico de la base de datos
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


func (s *Store) UpdateType(type_ types.Type) error{
	_, err := s.db.Query("UPDATE categories SET name = ? WHERE id = ?", type_.Name, type_.ID)
	if err != nil {
		return err
	}
	return nil
}

//ARTWORK FUNTIONS//

//ESta funcion que crea una obra en la base de datos
func (s *Store) CreateArtWork(artWork *types.ArtWork) error {
	_, err := s.db.Exec("INSERT INTO artworks (title) VALUES (?)", artWork.Title)
	if err != nil {
		return err
	}
	return nil
}

//Esta funcion recupera todas las obras de la base de datos
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

//Esta funcion recupera una obra en especifico de la base de datos
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

//Esta funcion recupera una obra aleatoria de la base de datos
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

func (s *Store) GetNotedArtWork() ([]types.ArtWork, error){
	rows, err := s.db.Query("SELECT * FROM artworks WHERE noted = 1")
	if err != nil{
		return nil, err
	}

	artWorks, err := scanRowsIntoArtWorks(rows)

	return artWorks, err
}

func (s *Store) UpdateArtWork(artWork types.ArtWork) error{
	_, err := s.db.Query("UPDATE artworks SET title = ?, description = ?, image = ?, logo = ? WHERE id = ?", artWork.Title, artWork.Description, artWork.Image, artWork.Logo, artWork.ID)
	if err != nil {
		return err
	}
	return nil
}

//SCAN ROWS DB FUNCTIONS//

//Esta funcion escanea los datos traidos de la base de datos y los convierte en un tag
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

//Esta funcion escanea los datos traidos de la base de datos y los convierte en una obra
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