package main

import (
	"errors"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/urfave/cli.v1"
	"gopkg.in/urfave/cli.v1/altsrc"
)

const VERSION = "0.8.13"

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)
	cli.AppHelpTemplate = fmt.Sprintf(`%s

WEBSITE: https://github.com/docker-commands/enforce

`, cli.AppHelpTemplate)
}

func main() {
	app := cli.NewApp()
	settingApplication(app)
	app.Action = action
	app.Run(os.Args)
	os.Exit(1)
}

func action(c *cli.Context) {
	error := validate(c)
	if error != nil {
		log.Error("invalid Arguments")
		os.Exit(1)
	}
}

func validate(c *cli.Context) error {
	log.WithFields(log.Fields{
		"len(args)": len(c.Args()),
		"args":      c.Args(),
	}).Debug("validation Arguments")

	if len(c.Args()) != 2 {
		return errors.New("args")
	}
	return nil
}
func settingApplication(app *cli.App) {
	app.Version = VERSION
	app.Name = "enforcer"
	app.Usage = "docker continaer rancher"
	app.Description = "enforcer launch docker container , when reciept message from queue."
	app.Copyright = "(c) 2017 Ryuichi Tokugami"
	flags := []cli.Flag{
		cli.StringFlag{
			Name:   "access-key, a",
			Usage:  "AWS access key for s3 access",
			EnvVar: "AWS_ACCESS_KEY",
		},
		cli.StringFlag{
			Name:   "secret-key, s",
			Usage:  "AWS secret access key for s3 access",
			EnvVar: "AWS_SECRET_ACCESS_KEY",
		},
		cli.StringFlag{
			Name:  "config, f",
			Usage: "Load configuration from `FILE`",
		},
		cli.UintFlag{
			Name:  "concurrent, c",
			Usage: "concurrent count",
		},
	}
	app.Before = altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("load"))

	app.Flags = flags
}
