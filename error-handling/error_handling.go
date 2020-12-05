package erratum

func Use(f func() (Resource,error), s string) (err error) {
	var r Resource

	for {
		r, err = f()

		if err == nil {
			break
		}

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