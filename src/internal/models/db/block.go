package db

import (
	"database/sql"
)

type Block struct {
	Hash      string         `gorm:"primaryKey"`
	Number    int32          `gorm:"column:number"`
	TxCount   int16          `gorm:"column:tx_count"`
	PrevHash  sql.NullString `gorm:"column:prev_block_hash"`
	Timestamp string         `gorm:"column:timestamp"`
}
