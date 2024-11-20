package db

import (
	"fmt"
	"sketch/graph/model"

	"github.com/lib/pq"
)

func FetchItemList() ([]*model.Item, error) {
	var items []*model.Item
	rows, err := DB.Query("SELECT id, user_id, title, image, status, author_id, tags, date FROM \"sketch_items\"")
	if err != nil {
		fmt.Println(err)
		return items, err
	}
	for rows.Next() {
		var db_item model.DbItem
		rows.Scan(&db_item.ID, &db_item.UserID, &db_item.Title, &db_item.Image, &db_item.Status, &db_item.AuthorID, pq.Array(&db_item.Tags), &db_item.Date)
		item := convertItem(&db_item)
		items = append(items, item)
	}

	return items, nil
}

func FetchItemById(item_id string) (*model.Item, error) {
	var item *model.Item
	var db_item model.DbItem
	row := DB.QueryRow("SELECT id, user_id, title, image, status, author_id, tags, date FROM \"sketch_items\" WHERE id = $1", item_id)
	err := row.Scan(&db_item.ID, &db_item.UserID, &db_item.Title, &db_item.Image, &db_item.Status, &db_item.AuthorID, pq.Array(&db_item.Tags), &db_item.Date)
	if err != nil {
		fmt.Println(err)
		return item, err
	}
	item = convertItem(&db_item)
	return item, nil
}

func FetchItemByAuthor(author_id string) ([]*model.Item, error) {
	var items []*model.Item
	rows, err := DB.Query("SELECT id, user_id, title, image, status, author_id, tags, date FROM \"sketch_items\" WHERE author_id = $1", author_id)
	if err != nil {
		fmt.Println(err)
		return items, err
	}
	for rows.Next() {
		var db_item model.DbItem
		rows.Scan(&db_item.ID, &db_item.UserID, &db_item.Title, &db_item.Image, &db_item.Status, &db_item.AuthorID, pq.Array(&db_item.Tags), &db_item.Date)
		item := convertItem(&db_item)
		items = append(items, item)
	}

	return items, nil
}

func FetchItemByTag(tag_id string) ([]*model.Item, error) {
	var items []*model.Item
	rows, err := DB.Query("SELECT id, user_id, title, image, status, author_id, tags, date FROM \"sketch_items\" WHERE $1 = any(tags)", tag_id)
	if err != nil {
		fmt.Println(err)
		return items, err
	}
	for rows.Next() {
		var db_item model.DbItem
		rows.Scan(&db_item.ID, &db_item.UserID, &db_item.Title, &db_item.Image, &db_item.Status, &db_item.AuthorID, pq.Array(&db_item.Tags), &db_item.Date)
		item := convertItem(&db_item)
		items = append(items, item)
	}

	return items, nil
}

func FetchItems(item_ids []string) ([]*model.Item, error) {
	var items []*model.Item
	rows, err := DB.Query("SELECT id, user_id, title, image, status, author_id, tags, date FROM \"sketch_items\" WHERE id = any($1)", pq.Array(item_ids))
	if err != nil {
		fmt.Println(err)
		return items, err
	}
	for rows.Next() {
		var db_item model.DbItem
		rows.Scan(&db_item.ID, &db_item.UserID, &db_item.Title, &db_item.Image, &db_item.Status, &db_item.AuthorID, pq.Array(&db_item.Tags), &db_item.Date)
		item := convertItem(&db_item)
		items = append(items, item)
	}

	return items, nil
}
