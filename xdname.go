package xmldoc

// XDName wraps local name and name space.
type XDName struct {
	LocalName string
	NameSpace string
}

// String convert the name to a string either NameSpace:LocalName or LocalName
func (n XDName) String() string {
	if n.NameSpace != "" {
		return n.NameSpace + ":" + n.LocalName
	}
	return n.LocalName
}
