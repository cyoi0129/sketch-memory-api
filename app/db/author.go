package db

import (
	"fmt"
	"sketch/graph/model"

	"github.com/lib/pq"
)

func FetchAuthorList() ([]*model.Author, error) {
	var authors []*model.Author
	rows, err := DB.Query("SELECT id, user_id, name, birth FROM \"sketch_authors\"")
	if err != nil {
		fmt.Println(err)
		return authors, err
	}
	for rows.Next() {
		var db_author model.DbAuthor
		rows.Scan(&db_author.ID, &db_author.UserID, &db_author.Name, &db_author.Birth)
		author := convertAuthor(&db_author)
		authors = append(authors, author)
	}

	return authors, nil
}

func FetchAuthorById(author_id string) (*model.Author, error) {
	var author model.DbAuthor
	row := DB.QueryRow("SELECT id, user_id, name, birth FROM \"sketch_authors\" WHERE id = $1", author_id)
	err := row.Scan(&author.ID, &author.UserID, &author.Name, &author.Birth)
	if err != nil {
		fmt.Println(err)
		return convertAuthor(&author), err
	}
	return convertAuthor(&author), nil
}

func FetchAuthors(author_ids []string) ([]*model.Author, error) {
	var authors []*model.Author
	rows, err := DB.Query("SELECT id, user_id, name, birth FROM \"sketch_authors\" WHERE id = any($1)", pq.Array(author_ids))
	if err != nil {
		fmt.Println(err)
		return authors, err
	}
	for rows.Next() {
		var db_author model.DbAuthor
		rows.Scan(&db_author.ID, &db_author.UserID, &db_author.Name, &db_author.Birth)
		author := convertAuthor(&db_author)
		authors = append(authors, author)
	}

	return authors, nil
}
