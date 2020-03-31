package wzlib_database_controller

type WzClient struct {
	ID     int    `gorm:"primary_key"`
	Uid    string `gorm:"unique; not null"`
	Fqdn   string `gorm:"unique; not null"`
	RsaPk  string `gorm:"unique; not null"`
	Status int    `gorm:"not null"`
}
