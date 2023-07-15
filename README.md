# payment-calculator
Calculator of payment for housing services.

[![Build Status](http://51.250.69.53/badge/e6cf3c1b?branch=main)](http://51.250.69.53/repos/13)

## API Examples
### Taxes
```
curl -X GET -H "Accept: application/json" http://localhost:8080/api/taxes/
curl -X POST -H "Accept: application/json" http://localhost:8080/api/taxes/ -d '{"hot_price": 10.0, "cold_price": 5.0, "drenage_price": 3.0, "energy_price": 10.0}'
curl -X DELETE -H "Accept: application/json" http://localhost:8080/api/taxes/1
```
### Records
```
curl -X GET -H "Accept: application/json" http://localhost:8080/api/records/
curl -X POST -H "Accept: application/json" http://localhost:8080/api/records/ -d '{"hot_value": 15.0, "cold_value": 12.0, "drenage_value": 6.0, "energy_value": 20.0}'
curl -X DELETE -H "Accept: application/json" http://localhost:8080/api/records/1
```
## Development
### Run go application
```
go run ./
```

### Run angular application
```
cd web/routing-app
ng serve
```

## Build for Production
### Build Docker image
```
docker build -t payment-calculator ./
```
### Run Docker image for testing
#### Run container in background
```
docker run -v /home/<user name>:/root/ -d -p 127.0.0.1:8080:8080 payment-calculator
```
#### Run container in foreground
```
docker run -p 127.0.0.1:8080:8080 payment-calculator
```
### Build & run go application
```
go build ./
./payment-calculator
```
## Production
```
sudo docker run -v /home/<user name>:/var/payment-calculator/ -d -p 80:8080 payment-calculator
```

## Abstruse CI
