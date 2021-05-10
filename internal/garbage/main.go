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

type TransformedGarbage struct {
	Weight int
}






