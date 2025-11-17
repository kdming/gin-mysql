package mysql

import "gorm.io/gorm"

func NewTax() *tax {
	return &tax{
		db:        db.Begin(),
		committed: false,
	}
}

type tax struct {
	db        *gorm.DB
	committed bool
}

func (t *tax) InsertMany(data interface{}, batchSize int) error {
	return t.db.CreateInBatches(data, batchSize).Error
}

func (t *tax) RawSql(sql string, params []interface{}, out interface{}) error {
	return t.db.Raw(sql, params...).Scan(out).Error
}

func (t *tax) Commit() error {
	if t.committed {
		return nil
	} else {
		t.committed = true
		return t.db.Commit().Error
	}
}

func (t *tax) Rollback() error {
	if t.committed {
		return nil
	} else {
		return t.db.Rollback().Error
	}
}
