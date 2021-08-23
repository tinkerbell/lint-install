package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/karrick/godirwalk"
	"k8s.io/klog/v2"
)

var (
	dryRunFlag     = flag.Bool("dry-run", false, "Display changes to make")
	goFlag         = flag.String("go", "error", "Level to lint Go with: [ignore, warn, error]")
	shellFlag      = flag.String("shell", "error", "Level to lint Shell with: [ignore, warn, error]")
	dockerfileFlag = flag.String("dockerfile", "error", "Level to lint Dockerfile with: [ignore, warn, error]")

	//go:embed .golangci.yml
	goLintConfig []byte

	//go:embed Makefile.tmpl
	makeTmpl string
)

type Language int

const (
	Go Language = iota
	Shell
	Dockerfile
)

type Config struct {
	Args       string
	Go         string
	Dockerfile string
	Shell      string
}

// applicableLinters returns a list of languages with known linters within a given directory.
func applicableLinters(root string) (map[Language]bool, error) {
	klog.Infof("Searching for linters to use for %s ...", root)
	found := map[Language]bool{}

	err := godirwalk.Walk(root, &godirwalk.Options{
		Callback: func(path string, de *godirwalk.Dirent) error {
			if strings.HasSuffix(path, ".go") {
				found[Go] = true
			}
			if strings.HasSuffix(path, "Dockerfile") {
				found[Dockerfile] = true
			}
			if strings.HasSuffix(path, ".sh") {
				found[Shell] = true
			}
			return nil
		},
		Unsorted: true,
	})

	return found, err
}

// updateMakefile updates the Makefile within a project with lint rules.
func updateMakefile(root string, cfg Config, dryRun bool) (string, error) {
	dest := filepath.Join(root, "Makefile")
	var existing []byte
	var err error

	if _, err = os.Stat(dest); err == nil {
		klog.Infof("Found existing %s", dest)
		existing, err = os.ReadFile(dest)
		if err != nil {
			return "", err
		}
	}

	var newRules bytes.Buffer
	t := template.Must(template.New("Makefile").Parse(makeTmpl))
	if err = t.Execute(&newRules, cfg); err != nil {
		return "", fmt.Errorf("execute: %w", err)
	}

	ignore := false
	inserted := false
	proposed := []byte{}
	for x, line := range bytes.Split(existing, []byte("\n")) {
		if bytes.HasPrefix(line, []byte("# BEGIN: lint-install")) {
			ignore = true
			inserted = true
			proposed = append(proposed, newRules.Bytes()...)
			continue
		}

		if bytes.HasPrefix(line, []byte("# END: lint-install")) {
			ignore = false
			continue
		}

		if ignore {
			continue
		}

		if x > 0 {
			proposed = append(proposed, []byte("\n")...)
		}
		proposed = append(proposed, line...)
	}

	if !inserted {
		proposed = append(proposed, newRules.Bytes()...)
	}

	edits := myers.ComputeEdits("Makefile", string(existing), string(proposed))
	change := gotextdiff.ToUnified(filepath.Base(dest), filepath.Base(dest), string(existing), edits)
	if !dryRun {
		if err := os.WriteFile(dest, proposed, 0755); err != nil {
			return "", err
		}
	}
	return fmt.Sprint(change), nil
}

// updateGoLint updates the golangci-lint configuration for a project.
func updateGoLint(root string, dryRun bool) (string, error) {
	dest := filepath.Join(root, ".golangci.yml")
	var existing []byte
	var err error

	if _, err = os.Stat(dest); err == nil {
		klog.Infof("Found existing %s", dest)
		existing, err = os.ReadFile(dest)
		if err != nil {
			return "", err
		}
	}

	proposed := string(goLintConfig)
	edits := myers.ComputeEdits(".golangci.yml", string(existing), proposed)
	change := gotextdiff.ToUnified(filepath.Base(dest), filepath.Base(dest), string(existing), edits)

	if !dryRun {
		if err := os.WriteFile(dest, goLintConfig, 0755); err != nil {
			return "", err
		}
	}

	return fmt.Sprint(change), nil
}

// main creates peanut butter & jelly sandwiches with utmost precision.
func main() {
	klog.InitFlags(nil)
	flag.Parse()

	if len(flag.Args()) == 0 {
		klog.Exitf("usage: lint-install [directory..]")
	}

	for _, root := range flag.Args() {
		needs, err := applicableLinters(root)
		if err != nil {
			klog.Exitf("failed to find linters: %v", err)
		}
		if len(needs) == 0 {
			continue
		}

		if needs[Go] {
			diff, err := updateGoLint(root, *dryRunFlag)
			if err != nil {
				klog.Exitf("update go lint config failed: %v", err)
			}
			if diff != "" {
				klog.Infof("go lint config changes:\n%s", diff)
			} else {
				klog.Infof("go lint config has no changes")
			}
		}

		cfg := Config{Args: strings.Join(os.Args[1:], " ")}

		if needs[Go] {
			cfg.Go = *goFlag
		}
		if needs[Dockerfile] {
			cfg.Dockerfile = *dockerfileFlag
		}
		if needs[Shell] {
			cfg.Shell = *shellFlag
		}

		diff, err := updateMakefile(root, cfg, *dryRunFlag)
		if err != nil {
			klog.Exitf("update Makefile failed: %v", err)
		}
		if diff != "" {
			klog.Infof("Makefile changes:\n%s", diff)
		} else {
			klog.Infof("Makefile has no changes")
		}
	}
}
