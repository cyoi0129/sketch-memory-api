package db

import (
	"fmt"
	"sketch/graph/model"
)

func FetchReviewList() ([]*model.Review, error) {
	var reviews []*model.Review
	rows, err := DB.Query("SELECT id, reviewer, item_id, score, comment FROM \"sketch_reviews\"")
	if err != nil {
		fmt.Println(err)
		return reviews, err
	}
	for rows.Next() {
		var db_review model.DbReview
		rows.Scan(&db_review.ID, &db_review.Reviewer, &db_review.ItemID, &db_review.Score, &db_review.Comment)
		review := convertReview(&db_review)
		reviews = append(reviews, review)
	}

	return reviews, nil
}

func FetchReviewByItem(item_id string) ([]*model.Review, error) {
	var reviews []*model.Review
	rows, err := DB.Query("SELECT id, reviewer, item_id, score, comment FROM \"sketch_reviews\" WHERE item_id = $1", item_id)
	if err != nil {
		fmt.Println(err)
		return reviews, err
	}
	for rows.Next() {
		var db_review model.DbReview
		rows.Scan(&db_review.ID, &db_review.Reviewer, &db_review.ItemID, &db_review.Score, &db_review.Comment)
		review := convertReview(&db_review)
		reviews = append(reviews, review)
	}

	return reviews, nil
}

func CreateReview(input model.NewReview) (*model.Review, error) {
	var review *model.Review
	db_review := model.DbReview{
		Reviewer: input.Reviewer,
		ItemID:   input.Item,
		Score:    input.Score,
		Comment:  input.Comment,
	}
	err := DB.QueryRow("INSERT INTO sketch_reviews(reviewer, item_id, score, comment) VALUES($1,$2,$3,$4) RETURNING id", input.Reviewer, input.Item, input.Score, input.Comment).Scan(&db_review.ID)
	if err != nil {
		fmt.Println(err)
		return review, err
	}
	review = convertReview(&db_review)
	return review, nil
}
