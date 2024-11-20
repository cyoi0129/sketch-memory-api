package db

import (
	"fmt"
	"sketch/graph/model"

	"github.com/lib/pq"
)

func FetchTagList() ([]*model.Tag, error) {
	var tags []*model.Tag
	rows, err := DB.Query("SELECT id, user_id, name FROM \"sketch_tags\"")
	if err != nil {
		fmt.Println(err)
		return tags, err
	}
	for rows.Next() {
		var db_tag model.DbTag
		rows.Scan(&db_tag.ID, &db_tag.UserID, &db_tag.Name)
		tag := convertTag(&db_tag)
		tags = append(tags, tag)
	}

	return tags, nil
}

func FetchTagById(tag_id string) (*model.Tag, error) {
	var tag *model.Tag
	var db_tag model.DbTag
	row := DB.QueryRow("SELECT id, user_id, name FROM \"sketch_tags\" WHERE id = $1", tag_id)
	err := row.Scan(&db_tag.ID, &db_tag.UserID, &db_tag.Name)
	if err != nil {
		fmt.Println(err)
		return tag, err
	}
	tag = convertTag(&db_tag)
	return tag, nil
}

func FetchTags(tag_ids []string) ([]*model.Tag, error) {
	var tags []*model.Tag
	rows, err := DB.Query("SELECT id, user_id, name FROM \"sketch_tags\" WHERE id = any($1)", pq.Array(tag_ids))
	if err != nil {
		fmt.Println(err)
		return tags, err
	}
	for rows.Next() {
		var db_tag model.DbTag
		rows.Scan(&db_tag.ID, &db_tag.UserID, &db_tag.Name)
		tag := convertTag(&db_tag)
		tags = append(tags, tag)
	}

	return tags, nil
}
