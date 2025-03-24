package resources

type Resources struct {
	Env *Env
}

func Get() (*Resources, error) {
	env := NewEnv()
	err := env.Load()
	if err != nil {
		return nil, err
	}

	return &Resources{
		Env: env,
	}, nil
}
