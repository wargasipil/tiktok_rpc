package repo

import (
	core_concept "github.com/wargasipil/tiktok_rpc/lib/exec_ctx"
	"gorm.io/gorm"
)

type CollectionTask struct {
	Coll *Collection
	db   *gorm.DB
}

func NewCollectionTask(db *gorm.DB, name string) (*CollectionTask, error) {
	coll := Collection{
		CollName: name,
		Status:   CTaskRunning,
	}

	err := db.Save(&coll).Error

	return &CollectionTask{
		db:   db,
		Coll: &coll,
	}, err
}

func (task *CollectionTask) AddCount(c int) {
	task.Coll.Count += c
}

func (task *CollectionTask) UpdateProgress(curent int, total int) error {
	var prog float32 = (float32(curent) / float32(total)) * 100
	task.Coll.Progress = prog

	return task.db.Save(task.Coll).Error
}
func (task *CollectionTask) SetError(err error) {
	if err != nil {
		task.Coll.ErrMessage = err.Error()
	}
}
func (task *CollectionTask) Finish() error {
	task.Coll.Status = CTaskCompleted
	return task.db.Save(task.Coll).Error
}

type CollectionMap struct {
	CreatorItemCreatorID string
	CollectionCollName   string
}

func (CollectionMap) TableName() string {
	return "collection"
}

type IterateCollConfig struct {
	Offset   int    `json:"offset"`
	CollName string `json:"coll_name"`
	SortBy   string `json:"sort_by"`
}

func IterateCollection(db *gorm.DB, ctx *core_concept.TaskContext, config *IterateCollConfig) (<-chan *CreatorItem, error) {
	creatorChan := make(chan *CreatorItem, 30)

	query := db.Model(&CreatorItem{}).
		Joins("join collection on collection.creator_item_creator_id = creator_items.creator_id").
		Where(`collection.collection_coll_name = ?`, config.CollName)

	if config.SortBy != "" {
		query = query.Order("creator_items." + config.SortBy + " DESC")
	}

	var total int64

	err := query.Count(&total).Error
	if err != nil {
		return creatorChan, err
	}

	go func() {
		defer close(creatorChan)

		rows, err := query.Rows()
		if err != nil {
			ctx.SetError(err)
			return
		}
		defer rows.Close()

	Parent:
		for rows.Next() {
			var creator CreatorItem

			db.ScanRows(rows, &creator)
			if err != nil {
				ctx.SetError(err)
				continue Parent
			}

			select {
			case <-ctx.Ctx.Done():
				break Parent
			case creatorChan <- &creator:
				continue Parent
			}

		}
	}()

	return creatorChan, nil
}
