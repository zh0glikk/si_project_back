package ventilation

type Ventilation struct {
	IsWorking bool
}

func New() *Ventilation {
	return &Ventilation{
		IsWorking: false,
	}
}


