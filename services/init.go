package services

func Init() error {
	if err := initTask(); err != nil {
		return err
	}
	if err := initArea(); err != nil {
		return err
	}
	return nil
}
