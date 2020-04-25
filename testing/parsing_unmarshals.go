package ldetesting

func (p *CustomStr) unmarshalData(rest []byte) (string, error) {
	return string(rest), nil
}
