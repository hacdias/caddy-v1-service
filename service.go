package service

import (
	"flag"
	"fmt"
	"os"

	"github.com/NetWatcher/service"
	"github.com/mholt/caddy"
)

var (
	logger       service.Logger
	name, action string
	instance     *caddy.Instance
)

func init() {
	flag.StringVar(&name, "name", "Caddy", "Caddy's service name")
	flag.StringVar(&action, "service", "", "Install, uninstall, start, stop, restart")

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
	instance, err = caddy.Start(caddyfile)
	if err != nil {
		return err
	}

	return nil
}

func (p *program) Stop(s service.Service) error {
	instance.ShutdownCallbacks()
	err := instance.Stop()
	instance = nil
	return err
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
		exit(err)
	}

	if action != "" {
		exit(actionHandler(action, s))
	}

	exit(s.Run())
	return nil
}

func exit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}

func actionHandler(action string, s service.Service) error {
	if action != "status" {
		return service.Control(s, action)
	}

	code, _ := s.Status()

	switch code {
	case 0:
		fmt.Println("Caddy service is not installed.")
	case 1:
		fmt.Println("Caddy service is not running.")
	case 4:
		fmt.Println("Caddy service is running.")
	default:
		fmt.Println("Error: ", code)
	}

	return nil
}
