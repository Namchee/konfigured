package entity

// Validaton represents config validation result
type Validation struct {
	Filename string
	Valid    bool
}

// GetInvalidValidations returns validations that are deemed invalid
func GetInvalidValidations(validations []Validation) []Validation {
	filtered := []Validation{}

	for _, v := range validations {
		if !v.Valid {
			filtered = append(filtered, v)
		}
	}

	return filtered
}
