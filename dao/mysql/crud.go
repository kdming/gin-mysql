package mysql

import (
	"gorm.io/gorm"
)

func Insert(data interface{}) error {
	return db.Create(data).Error
}

func FindOne(model interface{}, sql string, params []interface{}, out interface{}) error {
	err := db.Model(model).Where(sql, params...).First(out).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return err
}

func FindOneDesc(model interface{}, sql string, params []interface{}, out interface{}) error {
	err := db.Model(model).Where(sql, params...).Order("id desc").First(out).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return err
}

func Find(model interface{}, sql string, params []interface{}, out interface{}) error {
	return db.Model(model).Where(sql, params...).Find(out).Error
}

func FindByPage(model interface{}, sql string, params []interface{}, page int, limit int, out interface{}) error {
	return db.Model(model).Where(sql, params...).Offset((page - 1) * limit).Limit(limit).Find(out).Error
}

func FindByPageDesc(model interface{}, sql string, params []interface{}, page int, limit int, out interface{}) error {
	return db.Model(model).Where(sql, params...).Order("id desc").Offset((page - 1) * limit).Limit(limit).Find(out).Error
}

func Count(model interface{}, sql string, params []interface{}) (int, error) {
	var total int64
	err := db.Model(model).Where(sql, params...).Count(&total).Error
	return int(total), err
}

func Update(model interface{}, sql string, params []interface{}, update interface{}) error {
	return db.Model(model).Where(sql, params).Updates(update).Error
}

func Delete(table, sql string, params []interface{}) error {
	return db.Table(table).Delete(sql, params).Error
}

func Exec(sql string, params []interface{}, out interface{}) error {
	return db.Raw(sql, params...).Scan(out).Error
}

func InsertMany(model interface{}, data interface{}, batchSize int) error {
	return db.Model(model).CreateInBatches(data, batchSize).Error
}

func BulkCopy(model interface{}, data interface{}) error {
	// return db.Table(model.TableName()).Create(data).Error
	return nil
}
