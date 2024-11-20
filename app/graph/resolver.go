package graph

import "sketch/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	users       []*model.User
	items       []*model.Item
	authors     []*model.Author
	tags        []*model.Tag
	reviews     []*model.Review
	item        *model.Item
	tag         *model.Tag
	author      *model.Author
	tagItems    []*model.Item
	authorItems []*model.Item
	itemReviews []*model.Review
}
