package impl

import (
	"github.com/fwchen/jellyfish/domain/taco"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestTacoRepositoryImpl_buildListTacosSQL(t *testing.T) {
	sql, _, _ := buildListTacosSQL("u1", taco.ListTacoFilter{
		Statues: []taco.Status{taco.Done},
	})
	assert.Equal(t,
		sql,
		`SELECT "id", "creator_id", TRIM("content"), TRIM("detail"), TRIM("type"), "deadline", TRIM("status"), "box_id", "created_at", "updated_at" FROM "taco" WHERE (("creator_id" = 'u1') AND ("status" IN ('Done')))`,
	)
}
