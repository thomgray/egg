package egg

type AttributedString struct {
	str  string
	atts []stringAttribute
}

type stringAttribute struct {
	attribute Attribute
	rangeFrom int
	rangeTo   int
}

func MakeAttributedString(s string) AttributedString {
	return AttributedString{
		s,
		make([]stringAttribute, 0),
	}
}

func (as AttributedString) WithAttribute(att Attribute, from int, to int) AttributedString {
	newAtt := stringAttribute{
		att, from, to,
	}
	newAtts := append(as.atts, newAtt)
	return AttributedString{
		as.str,
		newAtts,
	}
}

// GetString - return the plain string without attributes
func (as AttributedString) GetString() string {
	return as.str
}

func (as AttributedString) Append(as2 AttributedString) AttributedString {
	newS := AttributedString{}
	newS.str = as.str + as2.str
	newS.atts = append(as.atts, as2.atts...)

	return newS
}

func (as AttributedString) GetAttributesAt(x int) Attribute {
	attOut := Attribute(0)
	for _, att := range as.atts {
		if x >= att.rangeFrom && x < att.rangeTo {
			attOut |= att.attribute
		}
	}

	return attOut
}
