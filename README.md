# Find-Pairs API
Returns all pairs of indices
where the sum of the elements at those indices equals the target value.

### /find-pairs [POST] `apllication/json`
-**Endpoint** - `/find-pairs`
-**Method** - `POST`
-**Request Body** - `application/json`

#### Run Application 
```go

go run main.go

```

#### Send request in `Postman`
Example 
```json

{
"numbers":[1,2,3,4,5],
"target":6
}

```

#### Response 
```json 
{
    "solutions": [
        [ 0, 4],
        [1,3]
    ]
}

```
#### Check the Test Cases

```go

go test -v

```