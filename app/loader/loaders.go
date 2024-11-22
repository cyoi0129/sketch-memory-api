package loader

import (
	"context"
	"net/http"

	"github.com/graph-gophers/dataloader"

	"database/sql"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

// 各DataLoaderを取りまとめるstruct
type Loaders struct {
	AuthorLoader *dataloader.Loader
	TagLoader    *dataloader.Loader
}

// Loadersの初期化メソッド
func NewLoaders(db *sql.DB) *Loaders {

	//ローダーの定義
	authorLoader := &AuthorLoader{
		DB: db,
	}
	tagLoader := &TagLoader{
		DB: db,
	}
	loaders := &Loaders{
		AuthorLoader: dataloader.NewBatchedLoader(authorLoader.BatchGetAuthors),
		TagLoader:    dataloader.NewBatchedLoader(tagLoader.BatchGetTags),
	}
	return loaders
}

// ミドルウェアはデータ ローダーをコンテキストに挿入します
func Middleware(loaders *Loaders, next http.Handler) http.Handler {
	loaders.AuthorLoader.ClearAll()
	loaders.TagLoader.ClearAll()
	// ローダーをリクエストコンテキストに挿入するミドルウェアを返す
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), loadersKey, loaders)
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}

// ContextからLoadersを取得する
func GetLoaders(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
