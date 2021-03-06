package models

import (
    _ "github.com/jinzhu/gorm"
)

type Nonce struct {
    Address string `gorm:"column:address;primary_key" json:"address"`
    NetworkID string `gorm:"column:network_id;primary_key" json:"network_id"`
    Nonce int `gorm:"column:nonce;default:1;" json:"nonce"`
    Timestamp int64 `gorm:"column:timestmap;default:0;" json:"timestamp"`
}

func (u Nonce) TableName() string {
    return "nonce"
}
