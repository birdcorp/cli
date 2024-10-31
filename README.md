### CLI

### Install

```bash
make install
```

### Authentication

You need to obtain a merchant `API_KEY` first to use the cli.

- **Set API Key**
  ```bash
  birdcli account set-api-key [API_KEY]
  ```

### Account

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
        "label": "Sun Hat",
        "type": "item",
        "value": "5.99",
        "status": "final",
        "thumbnail_url": "https://placehold.co/60x60.png"
      },
      {
        "label": "Sales Tax",
        "type": "tax",
        "value": "5.00",
        "status": "pending"
      },
      {
        "label": "Delivery",
        "type": "shipping",
        "value": "0.00",
        "status": "pending"
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

This command opens a screen to scan a qrcode of the miniprogram using a given URL.

  ```bash
  birdcli miniprogram create-preview \
    --url https://miniprogram-developer.onrender.com/ \
    --name "Miniprogram Developer"
  ```

Example: w/ngrok url for localhost proxy

```bash
ngrok http 3000
```

```bash
  birdcli miniprogram create-preview \
    --url https://56fc-32-133-145-153.ngrok-free.app \
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
