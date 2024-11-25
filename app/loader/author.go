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

// AuthorLoader はデータベースからデータを読み取ります
type AuthorLoader struct {
	DB *sql.DB
}

// BatchGetAuthors は、ID によって多くのデータを取得できるバッチ関数を実装します。
func (u *AuthorLoader) BatchGetAuthors(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	// 単一のクエリで要求されたすべてのデータを読み取ります
	authorIDs := make([]string, len(keys))
	for ix, key := range keys {
		authorIDs[ix] = key.String()
	}

	authorsTemp := []*model.Author{}

	rows, err := u.DB.Query("SELECT id, name, birth FROM \"sketch_authors\" WHERE id = any($1)", pq.Array(authorIDs))
	if err != nil {
		err := fmt.Errorf("fail get authors, %w", err)
		log.Printf("%v\n", err)
		return nil
	}
	for rows.Next() {
		var author model.Author
		rows.Scan(&author.ID, &author.Name, &author.Birth)
		authorsTemp = append(authorsTemp, &author)
	}

	authorsByAuthorId := map[string]*model.Author{}
	for _, author := range authorsTemp {
		authorsByAuthorId[author.ID] = author
	}

	authors := make([]*model.Author, len(authorIDs))

	for i, id := range authorIDs {
		authors[i] = authorsByAuthorId[id]
	}

	output := make([]*dataloader.Result, len(keys))
	for index := range authorIDs {
		author := authors[index]
		output[index] = &dataloader.Result{Data: author, Error: nil}
	}
	return output
}

// dataloader.Loadをwrapして型づけした実装
func LoadAuthor(ctx context.Context, authorID string) (*model.Author, error) {
	loaders := GetLoaders(ctx)
	thunk := loaders.AuthorLoader.Load(ctx, dataloader.StringKey(authorID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*model.Author), nil
}
