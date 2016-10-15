package env

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"syscall"
)

func mainEnvDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/.gip", usr.HomeDir), nil
}

func envDir(name string) (string, error) {
	mainEnvDir, err := mainEnvDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", mainEnvDir, name), nil
}

func ensureDir(path string) error {
	f, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			if mErr := os.Mkdir(path, 0777); mErr != nil {
				return mErr
			}
		} else {
			return err

		}
	}

	if f != nil && !f.IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	}

	return nil
}

func Init() error {
	envDir, err := mainEnvDir()
	if err != nil {
		return err
	}

	if err := ensureDir(envDir); err != nil {
		return err
	}

	return nil
}

func ActivateOrCreateEnv(name string) error {
	if err := Init(); err != nil {
		return err
	}

	envDir, err := envDir(name)
	if err != nil {
		return err
	}

	f, err := os.Stat(envDir)
	if err != nil {
		if os.IsNotExist(err) {
			if mErr := os.Mkdir(envDir, 0777); mErr != nil {
				return mErr
			}
		} else {
			return err

		}
	}

	if f != nil && !f.IsDir() {
		return fmt.Errorf("%s is not a directory", envDir)
	}

	shell := os.Getenv("SHELL")

	ps1 := os.Getenv("PS1")
	if ps1 == "" {
		ps1 = fmt.Sprintf("(%s) $ ", name)
	} else {
		ps1 = fmt.Sprintf("(%s) %s", name, ps1)
	}

	if err := os.Setenv("GOPATH", envDir); err != nil {
		return err
	}

	if err := os.Setenv("GIP", "true"); err != nil {
		return err
	}

	if err := os.Setenv("PS1", ps1); err != nil {
		return err
	}

	if err := syscall.Exec(shell, []string{shell, "--noprofile", "--norc"}, os.Environ()); err != nil {
		return err
	}

	return nil
}

func DeleteEnv(name string) error {
	envDir, err := envDir(name)
	if err != nil {
		return err
	}

	if err := os.RemoveAll(envDir); err != nil {
		return err
	}

	return nil
}

func ListEnvs() ([]string, error) {
	if err := Init(); err != nil {
		return nil, err
	}

	envDir, err := mainEnvDir()
	if err != nil {
		return nil, err
	}

	files, err := ioutil.ReadDir(envDir)
	if err != nil {
		return nil, err
	}

	envs := make([]string, 0, len(files))

	for _, f := range files {
		if f.IsDir() {
			envs = append(envs, f.Name())
		}
	}

	return envs, nil
}
