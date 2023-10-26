package model

type Document struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Doc_file    string `json:"doc_file"`
	Img_name    string `json:"img_name"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

func GetAllDocs() ([]Document, error) {
	var documents []Document

	query := `SELECT id, title, category, description, doc_file, img_name, created_at, updated_at;`

	rows, err := db.Query(query)
	if err != nil {
		return documents, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uint64
		var title, category, description, doc_file, img_name, created_at, updated_at string

		err := rows.Scan(&id, &title, &category, &description, &doc_file, &img_name, &created_at, &updated_at)
		if err != nil {
			return documents, err
		}

		document := Document{
			ID:          id,
			Title:       title,
			Description: description,
			Category:    category,
			Doc_file:    doc_file,
			Img_name:    img_name,
			Created_at:  created_at,
			Updated_at:  updated_at,
		}

		documents = append(documents, document)
	}

	return documents, nil
}

func GetDoc(id uint64) (Document, error) {
	var document Document

	query := `SELECT id, title, category, description, doc_file, img_name, created_at, updated_at FROM docs WHERE id=$1;`

	row, err := db.Query(query, id)
	if err != nil {
		return document, err
	}

	defer row.Close()

	if row.Next() {
		var title, category, description, doc_file, img_name, created_at, updated_at string

		err := row.Scan(&title, &category, &description, &doc_file, &img_name, &created_at, &updated_at)
		if err != nil {
			return document, err
		}

		document = Document{
			ID:          id,
			Title:       title,
			Description: description,
			Category:    category,
			Doc_file:    doc_file,
			Img_name:    img_name,
			Created_at:  created_at,
			Updated_at:  updated_at,
		}
	}

	return document, nil

}

func CreateDoc(doc Document) error {
	query := `INSERT INTO docs (title, category, description, doc_file, img_name, created_at, updated_at) VALUES ($1, $2, $3, $4, $5);`
	_, err := db.Exec(query, doc.Title, doc.Category, doc.Description, doc.Doc_file, doc.Img_name, doc.Created_at, doc.Updated_at)

	if err != nil {
		return err
	}

	return nil
}

func UpdateDoc(doc Document) error {
	query := `UPDATE docs SET title=$1, category=$2, description=$3, doc_file=$4, img_name=$5, created_at=$6, updated_at=$7 WHERE id=$6;`
	_, err := db.Exec(query, doc.Title, doc.Category, doc.Description, doc.Doc_file, doc.Img_name, doc.Created_at, doc.Updated_at, doc.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteDoc(id uint64) error {
	query := `DELETE FROM docs WHERE id=$1;`
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
