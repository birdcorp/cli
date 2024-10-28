### Cli


### Account


Set API key
```
go run main.go account set-api-key [API_KEY]
```

Get account info
```
go run main.go account me
```

Delete API key
```
go run main.go account delete-api-key
```

### Orders

Create order 

```
go run main.go orders create \
  --total-value "10.99" \
  --currency "USD" \
  --line-items '[
    {
      "label": "Item1",
      "value": "5.99",
      "status": "final",
      "type": "item"
    },
    {
      "label": "Item2",
      "value": "5.00",
      "status": "pending",
      "type": "tax"
    },
    {
      "label": "Shipping",
      "value": "0.00",
      "status": "pending",
      "type": "shipping"
    }
  ]' \
  --required-shipping-fields "name,postalAddress,phone,email" \
  --required-billing-fields "name,postalAddress,phone,email"
```

List orders

```
go run main.go orders list
```

Get order by ID

```
go run main.go orders get [orderID]
```

Delete order

```
go run main.go orders delete [orderID]
```




### Miniprogram

  
Initialize miniprogram, creates config file
```
go run main.go miniprogram init
```

Create miniprogram preview

```
go run main.go miniprogram create-preview \
 --url https://miniprogram-developer.onrender.com/ \
 --name "Miniprogram Developer"
```


Create miniprogram 

```
go run main.go miniprogram create
```


Get Miniprogram by ID
```
go run main.go miniprogram get [appID]
```

List miniprograms

```
go run main.go miniprogram list
```


Delete miniprogram
```
go run main.go miniprogram delete <appID>
```


```
go run main.go miniprogram release <appID>
```



