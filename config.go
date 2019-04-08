package main

import (
    "github.com/JeremyLoy/config"
    "github.com/alexflint/go-arg"
)

// Populated at build time:
var GitCommit string
var AppVersion string

type AppConfig struct {
    AppToken  string `config:"PUSHOVER_APP_TOKEN" arg:"-a" help:"pushover app token can also be read from environment PUSHOVER_APP_TOKEN"`
    UserToken string `config:"PUSHOVER_USER_TOKEN" arg:"-u" help:"pushover user token can also be read from environment PUSHOVER_USER_TOKEN"`
    Title     string `arg:"-t" help:"title text"`
    Sound    string `arg:"-s" help:"sound"`
    Device   string `arg:"-d" help:"device"`
    Priority int    `arg:"-p" help:"priority 1-3"`
    Url      string `help:"url"`
    UrlTitle string `help:"url title"`
    Message  string `arg:"positional" help:"message text"`
}

func (AppConfig) Version() string {
    return "v" + AppVersion + " (" + GitCommit + ")"
}

func (AppConfig) Description() string {
    return "Send a single Pushover notification to a User"
}

func GetConfig() AppConfig {
    var AppArguments AppConfig

    config.FromEnv().To(&AppArguments)
    arg.MustParse(&AppArguments)

    return AppArguments
}
