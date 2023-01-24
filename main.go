package main

import (
	"errors"
	"flag"
	"fmt"
	packr "github.com/gobuffalo/packr/v2"
	"github.com/shammishailaj/goboiler/pkg"
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
			//fmt.Sprintf("%s%s", app.Path, "go.mod"):    "go.mod.tmpl",
			fmt.Sprintf("%s%s", app.Path, "Makefile"):   "Makefile.tmpl",
			fmt.Sprintf("%s%s", app.Path, "README.md"):  "README.md.tmpl",
			fmt.Sprintf("%s%s", app.Path, "VERSION"):    "VERSION.tmpl",
			fmt.Sprintf("%s%s", app.Path, ".gitignore"): "gitignore.tmpl",
			fmt.Sprintf("%s%s.go", app.Path, app.Name):  fmt.Sprintf("%s.go", app.Name),
		},
		Directories: make([]*dir, 0),
	}

	sep := string(os.PathSeparator)

	//s.Directories = append(s.Directories, &dir{
	//	Name: fmt.Sprintf("%s%s%s%s", app.Path, "cmd", sep, app),
	//	Files: map[string]string{
	//		fmt.Sprintf("%s%s%s%s%s%s", app.Path, "cmd", sep, app, sep, fmt.Sprintf("%s.go", app)): "app.go.tmpl",
	//	},
	//})

	s.Directories = append(s.Directories, &dir{
		Name: fmt.Sprintf("%s%s", app.Path, "configs"),
		Files: map[string]string{
			fmt.Sprintf("%s%s%s%s", app.Path, "configs", sep, "app.local.yml"): "config.yml.tmpl",
			fmt.Sprintf("%s%s%s%s", app.Path, "configs", sep, "app.dev.yml"):   "config.yml.tmpl",
			fmt.Sprintf("%s%s%s%s", app.Path, "configs", sep, "app.yml"):       "config.yml.tmpl",
		},
	})

	printMsg(fmt.Sprintf("app.DateTimeZone = %s", app.DateTimeZone))

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
		Name: fmt.Sprintf("%s%s%s%s%s%s%s%s", app.Path, "pkg", sep, "http", sep, "response", sep, "html"),
		Files: map[string]string{
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "pkg", sep, "http", sep, "response", sep, "html", sep, "doc.go"):       "response_html_doc.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "pkg", sep, "http", sep, "response", sep, "html", sep, "html.go"):      "response_html_html.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "pkg", sep, "http", sep, "response", sep, "html", sep, "html_test.go"): "response_html_html_test.go.tmpl",
		},
	})

	s.Directories = append(s.Directories, &dir{
		Name: fmt.Sprintf("%s%s%s%s%s%s%s%s", app.Path, "pkg", sep, "http", sep, "response", sep, "json"),
		Files: map[string]string{
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "pkg", sep, "http", sep, "response", sep, "json", sep, "doc.go"):       "response_json_doc.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "pkg", sep, "http", sep, "response", sep, "json", sep, "json.go"):      "response_json_json.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "pkg", sep, "http", sep, "response", sep, "json", sep, "json_test.go"): "response_json_test.go.tmpl",
		},
	})

	s.Directories = append(s.Directories, &dir{
		Name: fmt.Sprintf("%s%s%s%s%s%s%s%s", app.Path, "pkg", sep, "http", sep, "response", sep, "xml"),
		Files: map[string]string{
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "pkg", sep, "http", sep, "response", sep, "xml", sep, "doc.go"):      "response_xml_doc.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "pkg", sep, "http", sep, "response", sep, "xml", sep, "xml.go"):      "response_xml_xml.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "pkg", sep, "http", sep, "response", sep, "xml", sep, "xml_test.go"): "response_xml_test.go.tmpl",
		},
	})

	s.Directories = append(s.Directories, &dir{
		Name: fmt.Sprintf("%s%s%s%s%s%s", app.Path, "pkg", sep, "http", sep, "response"),
		Files: map[string]string{
			fmt.Sprintf("%s%s%s%s%s%s%s%s", app.Path, "pkg", sep, "http", sep, "response", sep, "doc.go"):          "response_doc.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s%s%s", app.Path, "pkg", sep, "http", sep, "response", sep, "builder.go"):      "response_builder.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s%s%s", app.Path, "pkg", sep, "http", sep, "response", sep, "builder_test.go"): "response_builder_test.go.tmpl",
		},
	})

	s.Directories = append(s.Directories, &dir{
		Name: fmt.Sprintf("%s%s%s%s", app.Path, "pkg", sep, "schemas"),
		Files: map[string]string{
			fmt.Sprintf("%s%s%s%s%s%s", app.Path, "pkg", sep, "schemas", sep, "home.go"):   "schemas_home.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s", app.Path, "pkg", sep, "schemas", sep, "semver.go"): "schemas_semver.go.tmpl",
		},
	})

	s.Directories = append(s.Directories, &dir{
		Name: fmt.Sprintf("%s%s%s%s", app.Path, "pkg", sep, "utils"),
		Files: map[string]string{
			fmt.Sprintf("%s%s%s%s%s%s", app.Path, "pkg", sep, "utils", sep, "array.go"):            "utils_array.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s", app.Path, "pkg", sep, "utils", sep, "aws_v2.go"):           "utils_aws_v2.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s", app.Path, "pkg", sep, "utils", sep, "constants.go"):        "utils_constants.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s", app.Path, "pkg", sep, "utils", sep, "couchbase_config.go"): "utils_couchbase_config.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s", app.Path, "pkg", sep, "utils", sep, "couchbasedb.go"):      "utils_couchbasedb.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s", app.Path, "pkg", sep, "utils", sep, "filesystem.go"):       "utils_filesystem.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s", app.Path, "pkg", sep, "utils", sep, "rdbms.go"):            "utils_rdbms.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s", app.Path, "pkg", sep, "utils", sep, "rdbms_utils.go"):      "utils_rdbms_utils.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s", app.Path, "pkg", sep, "utils", sep, "reflection.go"):       "utils_reflection.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s", app.Path, "pkg", sep, "utils", sep, "ses.go"):              "utils_ses.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s", app.Path, "pkg", sep, "utils", sep, "ses_old.go"):          "utils_ses_old.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s", app.Path, "pkg", sep, "utils", sep, "system.go"):           "utils_system.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s", app.Path, "pkg", sep, "utils", sep, "template.go"):         "utils_template.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s", app.Path, "pkg", sep, "utils", sep, "time.go"):             "utils_time.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s", app.Path, "pkg", sep, "utils", sep, "utils.go"):            "utils_utils.go.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s", app.Path, "pkg", sep, "utils", sep, "uuid.go"):             "utils_uuid.go.tmpl",
		},
	})

	s.Directories = append(s.Directories, &dir{
		Name: fmt.Sprintf("%s%s", app.Path, "vendor"),
	})

	s.Directories = append(s.Directories, &dir{
		Name: fmt.Sprintf("%s%s", app.Path, "storage"),
	})

	s.Directories = append(s.Directories, &dir{
		Name: fmt.Sprintf("%s%s%s%s", app.Path, "web", sep, "template"),
		Files: map[string]string{
			fmt.Sprintf("%s%s%s%s%s%s", app.Path, "web", sep, "template", sep, "base.gohtml"): "web_base.gohtml.tmpl",
		},
	})

	s.Directories = append(s.Directories, &dir{
		Name: fmt.Sprintf("%s%s%s%s%s%s%s%s", app.Path, "web", sep, "template", sep, "default", sep, "pages"),
		Files: map[string]string{
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "web", sep, "template", sep, "default", sep, "pages", sep, "home.gohtml"): "web_home.gohtml.tmpl",
		},
	})

	s.Directories = append(s.Directories, &dir{
		Name: fmt.Sprintf("%s%s%s%s%s%s%s%s", app.Path, "web", sep, "template", sep, "default", sep, "partials"),
		Files: map[string]string{
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "web", sep, "template", sep, "default", sep, "partials", sep, "footer.gohtml"):     "web_footer.gohtml.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "web", sep, "template", sep, "default", sep, "partials", sep, "header.gohtml"):     "web_header.gohtml.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "web", sep, "template", sep, "default", sep, "partials", sep, "javascript.gohtml"): "web_javascript.gohtml.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "web", sep, "template", sep, "default", sep, "partials", sep, "navbar.gohtml"):     "web_navbar.gohtml.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "web", sep, "template", sep, "default", sep, "partials", sep, "style.gohtml"):      "web_style.gohtml.tmpl",
		},
	})

	s.Directories = append(s.Directories, &dir{
		Name: fmt.Sprintf("%s%s%s%s%s%s%s%s", app.Path, "web", sep, "template", sep, "default", sep, "static"),
		Files: map[string]string{
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "web", sep, "template", sep, "default", sep, "static", sep, "datetime.js"): "web_datetime.js.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "web", sep, "template", sep, "default", sep, "static", sep, "default.css"): "web_default.css.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "web", sep, "template", sep, "default", sep, "static", sep, "domready.js"): "web_domready.js.tmpl",
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", app.Path, "web", sep, "template", sep, "default", sep, "static", sep, "lib.js"):      "web_lib.js.tmpl",
		},
	})

	s.create()
}

func main() {
	var (
		pathToApp string
		version   bool
	)

	const VERSION = "1.0.0"

	flag.BoolVar(&version, "version", false, "Prints version and exits")
	flag.StringVar(&pathToApp, "path", "", "Specify absolute path to app")
	flag.Parse()

	if version {
		printMsg(fmt.Sprintf("Version %s", VERSION))
		os.Exit(0)
	}

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
			Validate: func(answer *string) error {
				a := strings.ToLower(*answer)
				if a != "y" && a != "n" {
					return errors.New("invalid option. Only [y/N] available")
				}

				return nil
			},
			Action: func(answer *string) error {
				if strings.ToLower(*answer) == "n" {
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
		Validate: func(answer *string) error {
			err := errors.New("app name should be in lower snake case [a-z0-9_]")
			if len(*answer) == 0 {
				return err
			}
			r := regexp.MustCompile("^[a-z0-9_]*$")
			if !r.MatchString(*answer) {
				return err
			}

			return nil
		},
		Action: func(answer *string) error {
			app.Name = *answer

			return nil
		},
	})
	if err != nil {
		printErr(err.Error())
		os.Exit(1)
	}

	err = userAction(&action{
		Question: "Enter App Description:",
		Validate: func(answer *string) error {
			err := errors.New("illegal description")
			if len(*answer) == 0 {
				return err
			}
			return nil
		},
		Action: func(answer *string) error {
			app.Description = *answer
			return nil
		},
	})
	if err != nil {
		printErr(err.Error())
		os.Exit(1)
	}

	err = userAction(&action{
		Question: "Enter App CLI Description:",
		Validate: func(answer *string) error {
			err := errors.New("illegal CLI description")
			if len(*answer) == 0 {
				return err
			}
			return nil
		},
		Action: func(answer *string) error {
			app.DescriptionCLI = *answer
			return nil
		},
	})
	if err != nil {
		printErr(err.Error())
		os.Exit(1)
	}

	rv := new(pkg.RuntimeVersion).ReadVersion()
	app.RuntimeVersion = fmt.Sprintf("%s.%s", rv.Major, rv.Minor)
	t := time.Now()
	app.DateYear = fmt.Sprintf("%d", t.Year())
	app.DateYYYYMMDD = fmt.Sprintf("%d%d%d", t.Year(), t.Month(), t.Day())

	err = userAction(&action{
		Question: "Enter App Licence URL: ",
		Validate: func(answer *string) error {
			if len(*answer) == 0 {
				printMsg("illegal Licence URL")
				printMsg("Using MIT Licence by default")
				app.SoftwareLicenceURL = "https://opensource.org/licenses/MIT"
				return nil
			}
			return nil
		},
		Action: func(answer *string) error {
			app.SoftwareLicenceURL = *answer
			return nil
		},
	})

	err = userAction(&action{
		Question: fmt.Sprintf("Enter the default timezone your app will use[%s]:", time.Now().Location().String()),
		Validate: func(answer *string) error {
			if len(*answer) == 0 {
				*answer = time.Now().Location().String()
				printMsg(fmt.Sprintf("Can not use empty value for timezone. Using Local time zone by default [%s]", *answer))
				return nil
			}
			loc, locErr := time.LoadLocation(*answer)
			if locErr != nil {
				*answer = time.Now().Location().String()
				printMsg(fmt.Sprintf("Error validating timezone:  %s Using Local time zone by default [%s]", locErr.Error(), *answer))
				return nil
			} else {
				printSuccess("Successfully set timezone to: " + loc.String())
			}
			return nil
		},
		Action: func(answer *string) error {
			app.DateTimeZone = *answer

			return nil
		},
	})
	if err != nil {
		printErr(err.Error())
		app.DateTimeZone = time.Now().Location().String()
	}

	err = userAction(&action{
		Question: "Enter port number the app will run on [6000-10000]:",
		Validate: func(answer *string) error {
			err := errors.New("invalid port number")
			if len(*answer) == 0 {
				return err
			}
			r := regexp.MustCompile("^[0-9]*$")
			if !r.MatchString(*answer) {
				return err
			}

			return nil
		},
		Action: func(answer *string) error {
			app.Port = *answer

			return nil
		},
	})
	if err != nil {
		printErr(err.Error())
		os.Exit(1)
	}

	err = userAction(&action{
		Question: "Select logger:\r\n[1]: github.com/Sirupsen/logrus\r\n[2]: github.com/uber-go/zap",
		Validate: func(answer *string) error {
			i, err := strconv.Atoi(*answer)
			if err != nil {
				return err
			}

			if i != 1 && i != 2 {
				return errors.New("invalid choice")
			}

			return nil
		},
		Action: func(answer *string) error {
			i, err := strconv.Atoi(*answer)
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
		Validate: func(answer *string) error {
			i, err := strconv.Atoi(*answer)
			if err != nil {
				return err
			}

			if i != 1 && i != 2 {
				return errors.New("invalid choice")
			}

			return nil
		},
		Action: func(answer *string) error {
			i, err := strconv.Atoi(*answer)
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
