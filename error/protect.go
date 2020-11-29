package error

func Go(x func()) {
	defer func() {
		if err := recover(); err != nil {
			//XX
		}
	}()
	go x()
}


