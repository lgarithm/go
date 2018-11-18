package control

// Chain combines two functions of signature :: func() error
func Chain(f, g func() error) func() error {
	return func() error {
		if err := f(); err != nil {
			return err
		}
		return g()
	}
}
