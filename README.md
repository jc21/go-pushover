# Pushover CLI

Yes, it's another super simple command line interface for [Pushover](https://pushover.net/).

```bash
Send a single Pushover notification to a User
Usage: pushover-cli [--apptoken APPTOKEN] [--usertoken USERTOKEN] [--title TITLE] [--sound SOUND] [--device DEVICE] [--priority PRIORITY] [--url URL] [--urltitle URLTITLE] MESSAGE

Positional arguments:
  MESSAGE                message text

Options:
  --apptoken APPTOKEN, -a APPTOKEN
                         pushover app token can also be read from environment PUSHOVER_APP_TOKEN
  --usertoken USERTOKEN, -u USERTOKEN
                         pushover user token can also be read from environment PUSHOVER_USER_TOKEN
  --title TITLE, -t TITLE
                         title text
  --sound SOUND, -s SOUND
                         sound
  --device DEVICE, -d DEVICE
                         device
  --priority PRIORITY, -p PRIORITY
                         priority 0-3
  --url URL              url
  --urltitle URLTITLE    url title
  --help, -h             display this help and exit
  --version              display version and exit
```

### Installing

```bash
go get github.com/jc21/pushover-cli
```

or build with:

```bash
git clone https://github.com/jc21/pushover-cli && cd pushover-cli
make
./pushover-cli -h
```

