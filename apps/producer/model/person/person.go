package person

type Person struct {
	name           string
	age            int64
	identification map[string]string
}

func (p Person) GetName() string {
	return p.name
}

func (p Person) GetAge() int64 {
	return p.age
}

func (p Person) GetIdentification(ids ...string) []string {
	var evaluatedIds []string
	for _, value := range ids {
		evaluatedIds = append(evaluatedIds, value)
	}
	return evaluatedIds
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) SetAge(age int64) {
	p.age = age
}
