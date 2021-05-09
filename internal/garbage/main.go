package garbage

type GarbageType string

const (
	Glass 	GarbageType = "glass"
	Plastic GarbageType = "plastic"
	Organic	GarbageType = "organic"
)

type Garbage struct {
	Weight 		int
	GarbageType GarbageType
}

func NewGarbage(weight int, garbageType GarbageType) *Garbage {
	return &Garbage{Weight: weight, GarbageType: garbageType}
}

type TransformedGarbage struct {
	Weight int
}

func NewTransformedGarbage(weight int) *TransformedGarbage {
	return &TransformedGarbage{Weight: weight}
}





