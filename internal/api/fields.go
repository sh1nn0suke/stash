package api

import (
	"context"
	"slices"

	"github.com/99designs/gqlgen/graphql"
)

type queryFields []string

func collectQueryFields(ctx context.Context) queryFields {
	fields := graphql.CollectAllFields(ctx)
	return queryFields(fields)
}

func (f queryFields) Has(field string) bool {
	return slices.Contains(f, field)
}
