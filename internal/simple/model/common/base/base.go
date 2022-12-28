package baseModel

import "time"

type Base struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	IsDeleted bool
}

func (b *Base) CreatedData() {
	b.CreatedAt = time.Now()
	b.IsDeleted = false
}

func (b *Base) UpdatedData() {
	b.UpdatedAt = time.Now()
}
func (b *Base) DeletedData() {
	b.DeletedAt = time.Now()
	b.IsDeleted = true
}

type BaseCreate struct {
	CreatedAt time.Time
	IsDeleted bool
}

func (b *BaseCreate) CreatedBase() {
	b.CreatedAt = time.Now()
	b.IsDeleted = true
}

type BaseUpdate struct {
	UpdatedAt time.Time
}

func (b *BaseUpdate) UpdatedBase() {
	b.UpdatedAt = time.Now()
}

type BaseDelete struct {
	DeletedAt time.Time
	IsDeleted bool
}

func (b *BaseDelete) DeletedBase() {
	b.DeletedAt = time.Now()
	b.IsDeleted = true
}
