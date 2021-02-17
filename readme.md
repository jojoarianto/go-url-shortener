# go-url-shortener

Service that imitate TinyURL

## getting started

This project requires go to be installed. on OS X with homebrew you can just run  ```brew install go```

## how to play
clone this project/repository
```bash
git clone https://github.com/jojoarianto/go-url-shortener
```

migrate database
```
make migrate-schema
```

run build to download and build binari
```bash
make build
```

run app
```bash
./go-url-shortener
```

## run with docker

Build script
```
docker build --pull --rm -f "Dockerfile" -t gourlshortener:latest .
```

Run docker container
```
docker run --rm -it  -p 3000:3000/tcp gourlshortener:latest
```

## development mode

to run dev mode run
```bash
make run
```

to run test 
```bash
make test
```

see another make command on makefile

## api documentation

### post /shorten

```
POST /shorten
Content-Type: "application/json"

{
  "url": "https://blog.trello.com/navigate-communication-styles-difficult-times",
  "shortcode": "example"
}
```

Attribute | Description
--------- | -----------
**url**   | url to shorten
shortcode | preferential shortcode

##### Returns:

```
201 Created
Content-Type: "application/json"

{
  "shortcode": :shortcode
}
```

A random shortcode is generated if none is requested, the generated short code has exactly 6 alpahnumeric characters and passes the following regexp: ```^[0-9a-zA-Z_]{6}$```.

##### Errors:

Error | Description
----- | ------------
400   | ```url``` is not present
409   | The the desired shortcode is already in use. **Shortcodes are case-sensitive**.
422   | The shortcode fails to meet the following regexp: ```^[0-9a-zA-Z_]{6}$```.


### get /:shortcode

```
GET /:shortcode
Content-Type: "application/json"
```

Attribute      | Description
-------------- | -----------
**shortcode**  | url encoded shortcode

##### Returns

**302** response with the location header pointing to the shortened URL

```
HTTP/1.1 302 Found
Location: https://blog.trello.com/navigate-communication-styles-difficult-times
```

##### Errors

Error | Description
----- | ------------
404   | The ```shortcode``` cannot be found in the system

### get /:shortcode/stats

```
GET /:shortcode/stats
Content-Type: "application/json"
```

Attribute      | Description
-------------- | -----------
**shortcode**  | url encoded shortcode

##### Returns

```
200 OK
Content-Type: "application/json"

{
  "startDate": "2012-04-23T18:25:43.511Z",
  "lastSeenDate": "2012-04-23T18:25:43.511Z",
  "redirectCount": 1
}
```

Attribute         | Description
--------------    | -----------
**startDate**     | date when the url was encoded, conformant to [ISO8601](http://en.wikipedia.org/wiki/ISO_8601)
**redirectCount** | number of times the endpoint ```GET /shortcode``` was called
lastSeenDate      | date of the last time the a redirect was issued, not present if ```redirectCount == 0```

##### Errors

Error | Description
----- | ------------
404   | The ```shortcode``` cannot be found in the system
