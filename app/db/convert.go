package db

import (
	"sketch/graph/model"
)

func convertTag(tag *model.DbTag) *model.Tag {
	return &model.Tag{
		ID:   tag.ID,
		Name: tag.Name,
	}
}

func convertAuthor(author *model.DbAuthor) *model.Author {
	return &model.Author{
		ID:    author.ID,
		Name:  author.Name,
		Birth: author.Birth,
	}
}

func convertReview(review *model.DbReview) *model.Review {
	return &model.Review{
		ID:       review.ID,
		Score:    review.Score,
		Comment:  review.Comment,
		Reviewer: review.Reviewer,
	}
}

func convertItem(item *model.DbItem) *model.Item {
	var tags []*model.Tag
	for _, tag_id := range item.Tags {
		tag, _ := FetchTagById(tag_id)
		tags = append(tags, tag)
	}
	author, _ := FetchAuthorById(item.AuthorID)
	return &model.Item{
		ID:     item.ID,
		Title:  item.Title,
		Image:  item.Image,
		Date:   item.Date,
		Status: item.Status,
		Author: author,
		Tags:   tags,
	}
}
