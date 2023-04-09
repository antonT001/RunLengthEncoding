package services

const LEN_CHUNK = 3 // TODO вынести в конст или в конфиг

type RleService interface {
	Encode(msg []string) []string
	Decode(msg []string) []string
}

type rleService struct {
}

func NewRleService() *rleService {
	return &rleService{}
}

func (s rleService) Decode(msg []string) []string {
	return RunLengthDecode(msg)
}

func (s rleService) Encode(msg []string) []string {
	return RunLengthEncode(msg)
}
