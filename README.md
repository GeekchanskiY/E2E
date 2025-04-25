# E2E—eye to eye

Service for my own usage, monolith app with lots of features. Probably I'll split it to
microservices in the future, but definitely not now.

features:
 - advanced users and roles management
 - financial operations management allows analyzing consumption better
 - currency scraper (myfin.by)
 - notes management
 - work time management
 - relationship score tracking
 - activity calendar
 - all services are synchronized and create a nice ecosystem

## Dev setup
```bash
  npm i -g @redocly/cli@latest
  go install github.com/swaggo/swag/cmd/swag@latest
```

```.env
APP_SECRET=secret
APP_HOST=127.0.0.1
APP_PORT=8080
DB_HOST=secret
DB_POST=5432
DB_USER=secret
DB_PASSWORD=secret
DB_NAME=e2e
```
