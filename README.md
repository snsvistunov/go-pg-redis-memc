# Golang REST API application with Redis, PostgreSQL and in-app cache

## Stack

 - Go 1.18 
 - Redis 
 - PostgreSQL 
 - Docker

## Development

1. The Web Service reached under the path /film/:title

2. The Application has to realize the following logic:

	a. An HTTP request is sent against the endpoint /film/:title.
    
    b. The application requests data stored in a database (PostgreSQL).
	
    b. The request-payload contains a JSON object with movie title information.

    c. The application caches data: 
        - 15 seconds in application memory, 
        - 30 seconds in Redis store. 

## General description

Docker, docker-compose, curl should be preinstalled.

Used Docker version 20.10.17 and docker-compose version 1.29.0.

Golang REST API application, stores data in PostgreSQL and two level cache (application memory, Redis) - on get requests application checks the application memory and then the Redis before request the data in PostgreSQL.

Application returns JSON in format: 

```
{film:{"film_id" : integer, "title" : "character varying", "description": "text", "release_year" : integer, "language_id":smallint, "rental_rate":numeric, "replacement_cost":numeric, "rating":"mpaa_rating", "last_update": "timestamp without time zone", "special-features":"text[]", "full-text":"txvector"}}
```

#### Directories:

`controllers` - endpoints and acts on models.

`models` - PostgreSQL, Redis, application cache data structures and acts on data.

## How to deploy

The Docker image can be rebuilt by

```
docker build -t go-pg-redis-memc .
```

Deploy application:

```
docker-compose up --build 
```
 
It will run go, redis, containers and available as: 

http://127.0.0.1:8080/film/:title

Stop application:

```
docker-compose down
```

### Testing

`Makefile` - contains curl commands to send.

`make test-get` - send GET requests to retrieve data.

Example:
```
$ make test-get

```

## Additional

1. The Web Service has response time log reached under the path /log .

2. The time log is stored in the database PostgreSQL.

3. The log records the last response time for a specific request in milliseconds.

4. Application response time log returns JSON in format: 

```
{ResponseTimeLog:[{"Id" : integer, "Request" : "text", "TimeDb" : numeric, "TimeRedis" : numeric, "TimeMemory" : numeric,}]}
```

## Configuration

1. Demo DB deployed on the cloud services https://www.elephantsql.com/ .

2. Demo dataset https://www.postgresqltutorial.com/postgresql-sample-database/ .

3. Http-port listening "8080" by default.

4. You can change cofiguration params in models/config.go file.