package usecase

type FSConfigure interface {
	faq(domainName string) error
}

type FSConfiguration struct {
	fsType string
}


/// returns a  faq file
// TODO update the bill board
func (fs FSConfiguration) faq(domainName string) error  {
	return nil
}
