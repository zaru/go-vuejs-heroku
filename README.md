### Vue.js init

```
vue init webpack vue-app
cd vue-app
npm install
```

### Go packages

```
glide init
glide get github.com/pressly/goose/cmd/goose
glide install
```

### DB migration

```
goose sqlite3 ./sample.db create init sql
goose sqlite3 ./sample.db up
```

### Heroku

```
heroku create go-vuejs-heroku
heroku buildpacks:add heroku/go --app go-vuejs-heroku
heroku buildpacks:add heroku/nodejs --app go-vuejs-heroku

heroku config:set NPM_CONFIG_PRODUCTION=false
heroku config:set GOVERSION=go1.9
```

```
touch Procfile
```

in `Procfile`

```
release: goose sqlite3 ./sample.db up
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


in `server.go`

```
func main() {
	e := echo.New()
	e.Static("/vue", "vue-app/dist")
}
```

```
git push heroku master
```
