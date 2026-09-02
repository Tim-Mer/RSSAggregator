package main

type command struct {
	Name string
	Args []string
}

type commands struct {
	cmdMap map[string]func(*state, command) error
}

func (c *commands) Run(s *state, cmd command) error {
	return c.cmdMap[cmd.Name](s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmdMap[name] = f
}

func (c *commands) Initialise() error {
	c.cmdMap = make(map[string]func(*state, command) error)
	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("users", handlerListUsers)
	c.register("agg", handlerAgg)
	c.register("addfeed", handlerAddFeed)
	return nil
}
