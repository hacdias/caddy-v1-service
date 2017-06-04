package service

import (
	"flag"
	"fmt"
	"os"

	"github.com/kardianos/service"
	"github.com/mholt/caddy"
)

var (
	logger       service.Logger
	name, action string
)

func init() {
	flag.StringVar(&name, "name", "Caddy", "Caddy's service name")
	flag.StringVar(&action, "service", "", "install, uninstall, start, stop, restart")

	caddy.RegisterEventHook("service", hook)
}

type program struct{}

func (p *program) Start(s service.Service) error {
	// Get Caddyfile input
	caddyfile, err := caddy.LoadCaddyfile(flag.Lookup("type").Value.String())
	if err != nil {
		return err
	}

	// Start your engines
	_, err = caddy.Start(caddyfile)
	if err != nil {
		return err
	}

	return nil
}

func (p *program) Stop(s service.Service) error {
	return caddy.Stop()
}

func hook(event caddy.EventName, info interface{}) error {
	if event != caddy.StartupEvent {
		return nil
	}

	config := &service.Config{
		Name:        name,
		DisplayName: name,
		Description: "Caddy's service",
		Arguments:   []string{},
	}

	flag.VisitAll(func(f *flag.Flag) {
		// ignore our own flags
		if f.Name == "service" || f.Name == "name" {
			return
		}

		// ignore flags with default value
		if f.Value.String() == f.DefValue {
			return
		}

		config.Arguments = append(config.Arguments, "-"+f.Name, f.Value.String())
	})

	s, err := service.New(&program{}, config)
	if err != nil {
		return err
	}

	if action != "" {
		err = service.Control(s, action)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	err = s.Run()
	if err != nil {
		return err
	}

	os.Exit(0)
	return nil
}
