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
```

```
touch Procfile
echo "web: go-vuejs-heroku" > Procfile
```
