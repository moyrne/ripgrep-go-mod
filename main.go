package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	var projectPath string
	flag.StringVar(&projectPath, "project-path", ".", "Path to directory where `go list -json -m all` should be executed")
	var outputPath string
	flag.StringVar(&outputPath, "output", ".ripgreprc", "Output ripgrep config path")

	flag.Parse()

	if err := ripgrep(projectPath, outputPath); err != nil {
		fmt.Fprintf(os.Stderr, "generate config failed: project-path:%v output-path:%v err:%v",
			projectPath, outputPath, err)
		os.Exit(1)
	}
}

func ripgrep(path, outputPath string) error {
	// Change to the specified directory
	if err := os.Chdir(path); err != nil {
		return err
	}

	var env Env
	envOutput, err := exec.Command("go", "env", "-json").Output()
	if err != nil {
		return err
	}
	if err := json.Unmarshal(envOutput, &env); err != nil {
		return err
	}

	// Execute go list -json -m all
	listOutput, err := exec.Command("go", "list", "-json", "-m", "all").Output()
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	data := fmt.Sprintf("--glob=%v\n", filepath.Join("**", "*"+env.GoRoot, "src*", "**"))
	fmt.Print(data)
	buf.WriteString(data)

	// Parse JSON output into a slice of Module
	decoder := json.NewDecoder(bytes.NewReader(listOutput))
	for decoder.More() {
		var mod Module
		if err := decoder.Decode(&mod); err != nil {
			return err
		}

		data := glob(mod)

		fmt.Print(data)
		buf.WriteString(data)
	}

	return os.WriteFile(outputPath, buf.Bytes(), 0666)
}

type Env struct {
	GoRoot     string `json:"GOROOT,omitempty"`
	GoModCache string `json:"GOMODCACHE,omitempty"`
}

// Module represents a Go module from `go list -json -m all` output
type Module struct {
	Path      string  `json:"Path"`
	Version   string  `json:"Version"`
	Replace   *Module `json:"Replace,omitempty"`
	Time      string  `json:"Time,omitempty"`
	Indirect  bool    `json:"Indirect,omitempty"`
	Dir       string  `json:"Dir,omitempty"`
	GoMod     string  `json:"GoMod,omitempty"`
	GoVersion string  `json:"GoVersion,omitempty"`
}

func glob(mod Module) string {
	if mod.Version == "" {
		return fmt.Sprintf("--glob=%v\n",
			filepath.Join("**", "*"+modPath(mod.Path)+"*", "**"),
		)
	}
	return fmt.Sprintf("--glob=%v\n",
		filepath.Join("**", "*"+modPath(mod.Path)+"@"+mod.Version+"*", "**"),
	)
}

func modPath(path string) string {
	b := &strings.Builder{}
	for _, v := range path {
		if v == '/' {
			v = filepath.Separator
		}
		if v >= 'A' && v <= 'Z' {
			v += 'a' - 'A'
			b.WriteRune('!')
		}
		b.WriteRune(v)
	}
	return b.String()
}
