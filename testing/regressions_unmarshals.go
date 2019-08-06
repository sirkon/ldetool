package ldetesting

func (p *Regression1) unmarshalState(s string) (uint8, error) {
	return s[0], nil
}
