### Cli



```
birdcli miniprogram init
```

```
birdcli miniprogram create
```



### Account

```
go run main.go auth set-api-key [API_KEY]
```

```
go run main.go auth me
```


```
go run main.go auth delete-api-key
```


### Miniprogram

```
go run main.go miniprogram init
```

```
go run main.go miniprogram create
```

```
go run main.go miniprogram create-preview <appID> --url <url>
```

```
go run main.go miniprogram get [appID]
```

```
go run main.go miniprogram list
```

```
go run main.go miniprogram delete <appID>
```


```
go run main.go miniprogram release <appID>
```



### Orders

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

```
go run main.go orders list
```

```
go run main.go orders get [orderID]
```

```
go run main.go orders delete [orderID]
```

```
go run main.go orders create
```
