# payment-calculator
Calculator of payment for housing services.

```
curl -X GET -H "Accept: application/json" http://localhost:8080/api/taxes/
curl -X POST -H "Accept: application/json" http://localhost:8080/api/taxes/ -d '{"hot_price": 10.0, "cold_price": 5.0, "drenage_price": 3.0, "energy_price": 10.0}'
curl -X DELETE -H "Accept: application/json" http://localhost:8080/api/taxes/1
```