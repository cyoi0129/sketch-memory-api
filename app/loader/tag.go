package loader

import (
	"context"
	"fmt"
	"log"
	"sketch/graph/model"

	"database/sql"

	"github.com/graph-gophers/dataloader"
	"github.com/lib/pq"
)

// TagLoader はデータベースからデータを読み取ります
type TagLoader struct {
	DB *sql.DB
}

// BatchGetTags は、ID によって多くのデータを取得できるバッチ関数を実装します。
func (u *TagLoader) BatchGetTags(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	// 単一のクエリで要求されたすべてのデータを読み取ります
	tagIDs := make([]string, len(keys))
	for ix, key := range keys {
		tagIDs[ix] = key.String()
	}

	tagsTemp := []*model.Tag{}

	rows, err := u.DB.Query("SELECT id, user_id, name FROM \"sketch_tags\" WHERE id = any($1)", pq.Array(tagIDs))
	if err != nil {
		err := fmt.Errorf("fail get tags, %w", err)
		log.Printf("%v\n", err)
		return nil
	}
	for rows.Next() {
		var tag model.Tag
		rows.Scan(&tag.ID, &tag.Name)
		tagsTemp = append(tagsTemp, &tag)
	}

	tagsByTagId := map[string]*model.Tag{}
	for _, tag := range tagsTemp {
		tagsByTagId[tag.ID] = tag
	}

	tags := make([]*model.Tag, len(tagIDs))

	for i, id := range tagIDs {
		tags[i] = tagsByTagId[id]
	}

	output := make([]*dataloader.Result, len(keys))
	for index := range tagIDs {
		tag := tags[index]
		output[index] = &dataloader.Result{Data: tag, Error: nil}
	}
	return output
}

// dataloader.Loadをwrapして型づけした実装
func LoadTag(ctx context.Context, tagID string) (*model.Tag, error) {
	loaders := GetLoaders(ctx)
	thunk := loaders.TagLoader.Load(ctx, dataloader.StringKey(tagID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*model.Tag), nil
}
