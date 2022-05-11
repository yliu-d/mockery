package test

import (
	"github.com/yliu-d/mockery/v2/pkg/fixtures/http"
)

type HasConflictingNestedImports interface {
	RequesterNS
	Z() http.MyStruct
}
