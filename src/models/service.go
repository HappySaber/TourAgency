package models

// Service представляет структуру для таблицы services
type Service struct {
	ID    uint   `json:"id" gorm:"primary_key;column:id"`                   // Идентификатор услуги (SERIAL)
	Name  string `json:"name" gorm:"type:varchar(32);not null;column:name"` // Название услуги
	Price string `json:"price" gorm:"type:varchar(64);column:price"`        // Цена услуги
}
