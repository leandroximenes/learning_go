// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	resident := Resident{name, age, address}
	return &resident
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	if r.Address["street"] == "" {
		return false
	}

	if !(r.Age >= 0) {
		return false
	}

	if r.Name == "" {
		return false
	}

	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	*r = Resident{}
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	total := 0
	for _, v := range residents {
		if v.HasRequiredInfo() {
			total++
		}
	}
	return total
}
