### CLI

### Account

- **Set API Key**
  ```bash
  go run main.go account set-api-key [API_KEY]
  ```

- **Get Account Info**
  ```bash
  go run main.go account
  ```

- **Delete API Key**
  ```bash
  go run main.go account delete-api-key
  ```

### Orders

- **Create Order**
  ```bash
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

- **List Orders**
  ```bash
  go run main.go orders list
  ```

- **Get Order by ID**
  ```bash
  go run main.go orders get [orderID]
  ```

- **Delete Order**
  ```bash
  go run main.go orders delete [orderID]
  ```

### Miniprogram

- **Initialize Miniprogram (creates config file)**
  ```bash
  go run main.go miniprogram init
  ```

- **Create Miniprogram Preview**
  ```bash
  go run main.go miniprogram create-preview \
    --url https://miniprogram-developer.onrender.com/ \
    --name "Miniprogram Developer"
  ```

- **Create Miniprogram**
  ```bash
  go run main.go miniprogram create
  ```

- **Get Miniprogram by ID**
  ```bash
  go run main.go miniprogram get [appID]
  ```

- **List Miniprograms**
  ```bash
  go run main.go miniprogram list
  ```

- **Delete Miniprogram**
  ```bash
  go run main.go miniprogram delete <appID>
  ```

- **Release Miniprogram**
  ```bash
  go run main.go miniprogram release <appID>
  ```

### Webhooks

- **Create Webhook**
  ```bash
  birdcli webhook create --url https://www.example.com
  ```

- **List Webhooks**
  ```bash
  go run main.go webhook list
  ```

- **Delete Webhook**
  ```bash
  go run main.go webhook delete --id 1234567890
  ```

### Events

- **Get Event by ID**
  ```bash
  go run main.go events get <eventID>
  ```

- **Stream Events**
  ```bash
  go run main.go events stream
  ```

### Release
