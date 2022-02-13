package main

import (
	"errors"
	"flag"
	"fmt"
	packr "github.com/gobuffalo/packr/v2"
	"goboiler/pkg"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func createStructure(app *application) {
	s := structure{
		Box: packr.New("tplBox", "./templates"),
		App: app,
		Files: map[string]string{
			fmt.Sprintf("%s%s", app.Path, "go.mod"):   "go.mod.tmpl",
			fmt.Sprintf("%s%s", app.Path, "Makefile"): "Makefile.tmpl",
		},
		Directories: make([]*dir, 0),
	}

	sep := string(os.PathSeparator)

	s.Directories = append(s.Directories, &dir{
		Name: fmt.Sprintf("%s%s%s%s", app.Path, "cmd", sep, app),
		Files: map[string]string{
			fmt.Sprintf("%s%s%s%s%s%s", app.Path, "cmd", sep, app, sep, fmt.Sprintf("%s.go", app)): "app.go.tmpl",
		},
	})

	s.Directories = append(s.Directories, &dir{
		Name: fmt.Sprintf("%s%s", app.Path, "configs"),
		Files: map[string]string{
			fmt.Sprintf("%s%s%s%s", app.Path, "configs", sep, "app.local.yml"): "config.yml.tmpl",
			fmt.Sprintf("%s%s%s%s", app.Path, "configs", sep, "app.dev.yml"):   "config.yml.tmpl",
			fmt.Sprintf("%s%s%s%s", app.Path, "configs", sep, "app.yml"):       "config.yml.tmpl",
		},
	})

	s.Directories = append(s.Directories, &dir{
		Name: fmt.Sprintf("%s%s", app.Path, "deployments"),
		Files: map[string]string{
			fmt.Sprintf("%s%s%s%s", app.Path, "deployments", sep, "build_cli.sh"):       "build_cli.sh.tmpl",
			fmt.Sprintf("%s%s%s%s", app.Path, "deployments", sep, "Dockerfile"):         "Dockerfile.tmpl",
			fmt.Sprintf("%s%s%s%s", app.Path, "deployments", sep, "docker-compose.yml"): "docker-compose.yml.tmpl",
		},
	})

	s.Directories = append(s.Directories, &dir{
		Name: fmt.Sprintf("%s%s%s%s%s%s%s%s", app.Path, "internal", sep, "app", sep, app.Name, sep, "cmd"),
		Files: map[string]string{
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "internal", sep, "app", sep, app.Name, sep, "cmd", sep, "docs.go"):    "docs.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "internal", sep, "app", sep, app.Name, sep, "cmd", sep, "root.go"):    "root.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "internal", sep, "app", sep, app.Name, sep, "cmd", sep, "serve.go"):   "serve.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "internal", sep, "app", sep, app.Name, sep, "cmd", sep, "version.go"): "version.go.tmpl",
		},
	})

	s.Directories = append(s.Directories, &dir{
		Name: fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "internal", sep, "app", sep, app.Name, sep, "server", sep, "handlers"),
		Files: map[string]string{
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s", app.Path, "internal", sep, "app", sep, app.Name, sep, "server", sep, "handlers", sep, "ping.go"): "ping.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s", app.Path, "internal", sep, "app", sep, app.Name, sep, "server", sep, "handlers", sep, "home.go"): "home.go.tmpl",
		},
	})

	s.Directories = append(s.Directories, &dir{
		Name: fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "internal", sep, "app", sep, app.Name, sep, "server", sep, "middleware"),
		Files: map[string]string{
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s", app.Path, "internal", sep, "app", sep, app.Name, sep, "server", sep, "middleware", sep, "log.go"): "log.go.tmpl",
		},
	})

	s.Directories = append(s.Directories, &dir{
		Name: fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "internal", sep, "app", sep, app.Name, sep, "server", sep, "routes"),
		Files: map[string]string{
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s", app.Path, "internal", sep, "app", sep, app.Name, sep, "server", sep, "routes", sep, "list.go"):   "list.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s", app.Path, "internal", sep, "app", sep, app.Name, sep, "server", sep, "routes", sep, "routes.go"): "routes.go.tmpl",
		},
	})

	s.Directories = append(s.Directories, &dir{
		Name: fmt.Sprintf("%s%s", app.Path, "vendor"),
	})

	s.Directories = append(s.Directories, &dir{
		Name: fmt.Sprintf("%s%s", app.Path, "storage"),
	})

	s.create()
}

func main() {
	var pathToApp string

	flag.StringVar(&pathToApp, "path", "", "Specify absolute path to app")
	flag.Parse()

	if len(pathToApp) == 0 {
		printErr("--path is required")
		os.Exit(1)
	}

	if !path.IsAbs(pathToApp) {
		printErr("--path should be absolute")
		os.Exit(1)
	}

	pathToApp = path.Clean(pathToApp)

	stat, err := os.Stat(pathToApp)
	if err != nil && !os.IsNotExist(err) {
		printErr(err.Error())
		os.Exit(1)
	} else if os.IsNotExist(err) {
		err := userAction(&action{
			Question: fmt.Sprintf("Dir %s doesn't exist. Create? [y/N]:", pathToApp),
			Validate: func(answer string) error {
				a := strings.ToLower(answer)
				if a != "y" && a != "n" {
					return errors.New("invalid option. Only [y/N] available")
				}

				return nil
			},
			Action: func(answer string) error {
				if strings.ToLower(answer) == "n" {
					return errors.New("can't continue without app dir")
				}

				return os.Mkdir(pathToApp, permMode)
			},
		})

		if err != nil {
			printErr(err.Error())
			os.Exit(0)
		}

		printSuccess(fmt.Sprintf("Created %s", pathToApp))

		stat, err = os.Stat(pathToApp)
		if err != nil {
			printErr(err.Error())
			os.Exit(1)
		}
	}

	if !stat.IsDir() {
		printErr("--path should be a directory")
		os.Exit(1)
	}

	app := new(application)
	app.Path = fmt.Sprintf("%s%s", pathToApp, string(os.PathSeparator))

	err = userAction(&action{
		Question: "Enter app name [a-z0-9_]:",
		Validate: func(answer string) error {
			err := errors.New("app name should be in lower snake case [a-z0-9_]")
			if len(answer) == 0 {
				return err
			}
			r := regexp.MustCompile("^[a-z0-9_]*$")
			if !r.MatchString(answer) {
				return err
			}

			return nil
		},
		Action: func(answer string) error {
			app.Name = answer

			return nil
		},
	})
	if err != nil {
		printErr(err.Error())
		os.Exit(1)
	}

	err = userAction(&action{
		Question: "Enter App Description:",
		Validate: func(answer string) error {
			err := errors.New("illegal description")
			if len(answer) == 0 {
				return err
			}
			return nil
		},
		Action: func(answer string) error {
			app.Description = answer
			return nil
		},
	})
	if err != nil {
		printErr(err.Error())
		os.Exit(1)
	}

	err = userAction(&action{
		Question: "Enter App CLI Description:",
		Validate: func(answer string) error {
			err := errors.New("illegal CLI description")
			if len(answer) == 0 {
				return err
			}
			return nil
		},
		Action: func(answer string) error {
			app.DescriptionCLI = answer
			return nil
		},
	})
	if err != nil {
		printErr(err.Error())
		os.Exit(1)
	}

	app.RuntimeVersion = pkg.RuntimeVersionMap()["string"]
	t := time.Now()
	app.DateYear = fmt.Sprintf("%d", t.Year())
	app.DateYYYYMMDD = fmt.Sprintf("%d%d%d", t.Year(), t.Month(), t.Day())

	err = userAction(&action{
		Question: "Enter App Licence URL: ",
		Validate: func(answer string) error {
			if len(answer) == 0 {
				printMsg("illegal Licence URL")
				printMsg("Using MIT Licence by default")
				app.SoftwareLicenceURL = "https://opensource.org/licenses/MIT"
				return nil
			}
			return nil
		},
		Action: func(answer string) error {
			app.SoftwareLicenceURL = answer
			return nil
		},
	})

	err = userAction(&action{
		Question: fmt.Sprintf("Enter the default timezone your app will use[%s]:", time.Now().Location().String()),
		Validate: func(answer string) error {
			if len(answer) == 0 {
				app.DateTimeZone = time.Now().Location().String()
				printMsg(fmt.Sprintf("Can not use empty value for timezone. Using Local time zone by default [%s]", app.DateTimeZone))
				return nil
			}
			loc, locErr := time.LoadLocation(answer)
			if locErr != nil {
				app.DateTimeZone = time.Now().Location().String()
				printMsg(fmt.Sprintf("Error validating timezone:  %s Using Local time zone by default [%s]", locErr.Error(), app.DateTimeZone))
				return nil
			} else {
				printSuccess("Successfully set timezone to: " + loc.String())
			}
			return nil
		},
		Action: func(answer string) error {
			app.DateTimeZone = answer

			return nil
		},
	})
	if err != nil {
		printErr(err.Error())
		app.DateTimeZone = time.Now().Location().String()
	}

	err = userAction(&action{
		Question: "Enter port number the app will run on [6000-10000]:",
		Validate: func(answer string) error {
			err := errors.New("invalid port number")
			if len(answer) == 0 {
				return err
			}
			r := regexp.MustCompile("^[0-9]*$")
			if !r.MatchString(answer) {
				return err
			}

			return nil
		},
		Action: func(answer string) error {
			app.Port = answer

			return nil
		},
	})
	if err != nil {
		printErr(err.Error())
		os.Exit(1)
	}

	err = userAction(&action{
		Question: "Select logger:\r\n[1]: github.com/Sirupsen/logrus\r\n[2]: github.com/uber-go/zap",
		Validate: func(answer string) error {
			i, err := strconv.Atoi(answer)
			if err != nil {
				return err
			}

			if i != 1 && i != 2 {
				return errors.New("invalid choice")
			}

			return nil
		},
		Action: func(answer string) error {
			i, err := strconv.Atoi(answer)
			if err != nil {
				return err
			}

			switch i {
			case 1:
				app.Logger = "logrus"
				app.LoggerPackage = "github.com/sirupsen/logrus"
			case 2:
				app.Logger = "zap"
				app.LoggerPackage = "go.uber.org/zap"
			}

			return nil
		},
	})

	err = userAction(&action{
		Question: "Select router:\r\n[1]: github.com/gorilla/mux\r\n[2]: github.com/go-chi/chi",
		Validate: func(answer string) error {
			i, err := strconv.Atoi(answer)
			if err != nil {
				return err
			}

			if i != 1 && i != 2 {
				return errors.New("invalid choice")
			}

			return nil
		},
		Action: func(answer string) error {
			i, err := strconv.Atoi(answer)
			if err != nil {
				return err
			}

			switch i {
			case 1:
				app.Router = "mux"
				app.RouterPackage = "github.com/gorilla/mux"
			case 2:
				app.Router = "chi"
				app.RouterPackage = "github.com/go-chi/chi"
			}

			return nil
		},
	})

	printMsg(fmt.Sprintf("Creating %s application..", app))

	createStructure(app)

	printSuccess("Success!")
}
