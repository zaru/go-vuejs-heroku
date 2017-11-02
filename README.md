# Go, VueJS, PostgreSQL on Heroku demo

- Go 1.9
  - glide
- VueJS 2.5
  - vue-cli

### Vue.js init

```
vue init webpack vue-app
cd vue-app
npm install
```

### Go packages

#### glide

```
glide init
glide get github.com/pressly/goose/cmd/goose
glide install
```

#### dep

```
dep init
dep ensure
```

### DB migration

```
goose postgres $DATABASE_URL create init sql
goose postgres $DATABASE_URL up
```

### Heroku

```
heroku create go-vuejs-heroku
heroku buildpacks:add heroku/go --app go-vuejs-heroku
heroku buildpacks:add heroku/nodejs --app go-vuejs-heroku
heroku addons:create heroku-postgresql:hobby-dev --app go-vuejs-heroku

heroku config:set NPM_CONFIG_PRODUCTION=false
heroku config:set GOVERSION=go1.9
heroku config:set GO_INSTALL_PACKAGE_SPEC="./cmd/... ."
```

```
touch Procfile
```

in `Procfile`

```
release: goose postgres $DATABASE_URL up
web: go-vuejs-heroku
```

```
touch package.json
```

in `/package.json`

```
{
  "name": "go-vuejs-heroku",
  "version": "0.0.1",
  "engines": {
    "node": "7.10.0",
    "npm": "4.2.0"
  },
  "scripts": {
    "postinstall": "npm --prefix ./vue-app install ./vue-app && cd ./vue-app && npm run build"
  }
}
```

```
git push heroku master
```
