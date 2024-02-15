# Sunsamago
## Your Sunsama Companion Cli, especially for Timesheets

### TODO
a lot of things
- Logging, to get the sessionID
- more output format
- test
- code quality
- ....

### Installation
- clone the repo
- cd into the repository directory
- run `go install ./cmd/sunsamago.go`
- run `sunsamago`

### Usage

set your environment variable `SUNSAMA_SESSION_ID` to your sunsama session id.
see [here](#how-to-get-your-sunsama-session-id) for how to get the session id

```shell
sunsamago timesheet --start "2024-02-15" --end "2024-02-15"  --round 15 --durationFormatter hour
```


#### How to get your sunsama session id
- logging to sunsama on your browser
- open devtool
- look for the cookie `sunsamaSession`