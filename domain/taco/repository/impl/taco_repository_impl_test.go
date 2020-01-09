package impl

import (
	"github.com/fwchen/jellyfish/domain/taco"
	"github.com/fwchen/jellyfish/domain/taco/repository"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestTacoRepositoryImpl_buildListTacosSQL(t *testing.T) {
	sql, _, _ := buildListTacosSQL("u1", repository.ListTacoFilter{
		Statues: []taco.Status{taco.Done},
	})
	assert.Equal(t,
		`SELECT "id", "TRIM(content)", "TRIM(detail)", "type", "deadline", "status", "created_at", "updated_at" FROM "test" WHERE (("creator_id" = 'u1') AND ("status" IN ('Done')))`,
		sql)
}
