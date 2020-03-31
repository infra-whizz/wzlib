package wzlib_database_controller

type WzClient struct {
	ID     int `gorm:"primary_key"`
	Uid    string
	Fqdn   string
	RsaPk  string
	Status int
}
