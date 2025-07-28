package postgres

type Client struct {
	Firstn string `gorm:"column:first_name"`
	Lastn  string `gorm:"column:last_name"`
	NID    string `gorm:"primaryKey;column:nid"`
}

type Account struct {
	Balance uint32 `gorm:"column:balance"`
	AID     string `gorm:"primaryKey;column:aid"`
	Client  Client `gorm:"foreignKey:NID"`
	NID     string `gorm:"column:nid"`
}

type Transaction struct {
	AID     string  `gorm:"column:aid"`
	Sum     uint32  `gorm:"column:amount"`
	TID     string  `gorm:"primaryKey;column:tid"`
	Account Account `gorm:"foreignKey:AID"`
}

type File struct {
	Filename string  `gorm:"column:filename"`
	Hash     string  `gorm:"column:hash;unique"`
	AID      string  `gorm:"column:aid"`
	Bytes    []byte  `gorm:"column:bytes"`
	Account  Account `gorm:"foreignKey:AID"`
}
