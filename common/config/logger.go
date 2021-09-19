package config

import (
	"fmt"
	"github.com/kingwel-xie/k2/core/logger"
	"os"
	"strings"

	logging "github.com/ipfs/go-log/v2"
	"github.com/mattn/go-isatty"
)

type Logger struct {
	Level     string `mapstructure:"level" json:"level" yaml:"level"`    // 级别
	Format    string `mapstructure:"format" json:"format" yaml:"format"` // 输出
	File      string `mapstructure:"file" json:"file"  yaml:"file"`      // 日志文件
	URL       string `mapstructure:"url" json:"url"  yaml:"url"`         // url that will be processed by sink in the zap
	Output    string `mapstructure:"output" json:"output" yaml:"output"` // possible values: stdout|stderr|file combine multiple values with '+'
	Labels    string `mapstructure:"labels" json:"labels" yaml:"labels"` // comma-separated key-value pairs, i.e. "app=example_app,dc=sjc-1"
	EnabledDB bool   `mapstructure:"log-db" json:"log-db" yaml:"log-db"` // log输出到数据库
}

var LoggerConfig = new(Logger)

// Setup 设置logger
func (e Logger) Setup() {
	cfg := toLoggerConfig(e)
	logging.SetupLogging(cfg)

	logger.DefaultLogger = logging.Logger("main")
}

func toLoggerConfig(log Logger) logging.Config {
	cfg := logging.Config{
		Format:          logging.ColorizedOutput,
		Stderr:          true,
		Level:           logging.LevelInfo,
		SubsystemLevels: map[string]logging.LogLevel{},
		Labels:          map[string]string{},
	}

	// Format
	var noExplicitFormat bool
	switch log.Format {
	case "color":
		cfg.Format = logging.ColorizedOutput
	case "nocolor":
		cfg.Format = logging.PlaintextOutput
	case "json":
		cfg.Format = logging.JSONOutput
	default:
		if log.Format != "" {
			fmt.Fprintf(os.Stderr, "ignoring unrecognized log format '%s'\n", log.Format)
		}
		noExplicitFormat = true
	}

	// Level
	if log.Level != "" {
		for _, kvs := range strings.Split(log.Level, ",") {
			kv := strings.SplitN(kvs, "=", 2)
			lvl, err := logging.LevelFromString(kv[len(kv)-1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "error setting log level %q: %s\n", kvs, err)
				continue
			}
			switch len(kv) {
			case 1:
				cfg.Level = lvl
			case 2:
				cfg.SubsystemLevels[kv[0]] = lvl
			}
		}
	}

	cfg.File = log.File
	// Disable stderr logging when a file is specified
	// https://github.com/ipfs/go-log/issues/83
	if cfg.File != "" {
		cfg.Stderr = false
	}

	cfg.URL = log.URL

	output := log.Output
	outputOptions := strings.Split(output, "+")
	for _, opt := range outputOptions {
		switch opt {
		case "stdout":
			cfg.Stdout = true
		case "stderr":
			cfg.Stderr = true
		case "file":
			if cfg.File == "" {
				fmt.Fprint(os.Stderr, "please specify a GOLOG_FILE value to write to")
			}
		case "url":
			fmt.Fprint(os.Stderr, "Url as log output is not supported any longer")
			//if cfg.URL == "" {
			//	fmt.Fprint(os.Stderr, "please specify a GOLOG_URL value to write to")
			//}
		}
	}

	if noExplicitFormat &&
		(!cfg.Stdout || !isTerm(os.Stdout)) &&
		(!cfg.Stderr || !isTerm(os.Stderr)) {
		cfg.Format = logging.PlaintextOutput
	}

	// labels
	if log.Labels != "" {
		labelKVs := strings.Split(log.Labels, ",")
		for _, label := range labelKVs {
			kv := strings.Split(label, "=")
			if len(kv) != 2 {
				fmt.Fprint(os.Stderr, "invalid label k=v: ", label)
				continue
			}
			cfg.Labels[kv[0]] = kv[1]
		}
	}

	return cfg
}

func isTerm(f *os.File) bool {
	return isatty.IsTerminal(f.Fd()) || isatty.IsCygwinTerminal(f.Fd())
}
