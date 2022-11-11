package dao

type Dao struct {
	// redis,mysql
}

func New() (*Dao, error) {
	dao := &Dao{}

	return dao, nil
}
