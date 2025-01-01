package enums

type Integrity int

// Integrity States
const (
	INTACT                 = iota
	VIOLATED               = iota
	NOT_VERIFIED           = iota
	INTEGRITY_INITIAL_SCAN = iota
)

var integrityNames = map[Integrity]string{
	INTACT:                 "INTACT",
	VIOLATED:               "VIOLATED",
	NOT_VERIFIED:           "NOT_VERIFIED",
	INTEGRITY_INITIAL_SCAN: "INITIAL_SCAN",
}

func (integrity Integrity) String() string {
	return integrityNames[integrity]
}
