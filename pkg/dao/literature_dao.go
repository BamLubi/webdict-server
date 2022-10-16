package dao

import (
	"context"
	"golang.org/x/text/message/catalog"
	"webdict-server/pkg/entity"
)

type LiteratureDao interface {
	Insert(ctx context.Context, literature entity.Literature) bool
	FindAll(ctx context.Context) []entity.Literature
	FindById(ctx context.Context, id string) entity.Literature
	Update(ctx context.Context, literature entity.Literature) bool
	FuzzyFind(ctx catalog.Context, key string) []entity.Literature
}
