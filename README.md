```
vue init webpack vue-app
cd vue-app
npm install
```

```
dep init
dep ensure
```

```
heroku create go-vuejs-heroku
heroku buildpacks:add heroku/go --app go-vuejs-heroku
heroku buildpacks:add heroku/nodejs --app go-vuejs-heroku

heroku config:set NPM_CONFIG_PRODUCTION=false
```

```
touch Procfile
echo "web: go-vuejs-heroku" > Procfile
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
