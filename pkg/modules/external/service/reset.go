package external_service

func (serv *service) Reset() {
	serv.repository.Reset()
}
