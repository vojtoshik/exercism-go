package erratum

// Use opens uses resource returned by f() handling all the errors
// as requested
func Use(f func() (Resource, error), s string) (err error) {
	var r Resource

	for r, err = f(); err != nil; r, err = f() {
		if _, ok := err.(TransientError); ok {
			continue
		}

		return err
	}

	defer r.Close()

	defer func(res Resource) {
		if r := recover(); r != nil {

			if fe, ok := r.(FrobError); ok {
				res.Defrob(fe.defrobTag)
			}

			err = r.(error)
		}
	}(r)

	r.Frob(s)

	return nil
}
