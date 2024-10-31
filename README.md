### CLI

### Account

- **Set API Key**
  ```bash
  birdcli account set-api-key [API_KEY]
  ```

- **Get Account Info**
  ```bash
  birdcli account
  ```

- **Delete API Key**
  ```bash
  birdcli account delete-api-key
  ```

### Orders

- **Create Order**
  ```bash
  birdcli orders create \
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
  birdcli orders list
  ```

- **Get Order by ID**
  ```bash
  birdcli orders get [orderID]
  ```

- **Delete Order**
  ```bash
  birdcli orders delete [orderID]
  ```

### Miniprogram

- **Initialize Miniprogram (creates config file)**
  ```bash
  birdcli miniprogram init
  ```

- **Create Miniprogram Preview**

This command generates a preview of the miniprogram using a specified URL. For local testing, you can use ngrok to create a secure tunnel. The command will produce a URL that includes a QR code, allowing you to easily scan and develop the miniprogram on mobile.

  ```bash
  birdcli miniprogram create-preview \
    --url https://miniprogram-developer.onrender.com/ \
    --name "Miniprogram Developer"
  ```

- **Create Miniprogram**
  ```bash
  birdcli miniprogram create
  ```

- **Get Miniprogram by ID**
  ```bash
  birdcli miniprogram get [appID]
  ```

- **List Miniprograms**
  ```bash
  birdcli miniprogram list
  ```

- **Delete Miniprogram**
  ```bash
  birdcli miniprogram delete <appID>
  ```

- **Release Miniprogram**
  ```bash
  birdcli miniprogram release <appID>
  ```

### Webhooks

- **Create Webhook**
  ```bash
  birdcli webhook create --url https://www.example.com
  ```

- **List Webhooks**
  ```bash
  birdcli webhook list
  ```

- **Delete Webhook**
  ```bash
  birdcli webhook delete --id 1234567890
  ```

### Events

- **Get Event by ID**
  ```bash
  birdcli events get <eventID>
  ```

- **Stream Events**
  ```bash
  birdcli events stream
  ```

### Release
