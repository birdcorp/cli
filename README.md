

<h1 align="center" style="border-bottom: none">
  <div>
    <a href="https://www.docuseal.com">
      <img  alt="DocuSeal" src="https://github.com/docusealco/docuseal/assets/5418788/c12cd051-81cd-4402-bc3a-92f2cfdc1b06" width="80" />
      <br>
    </a>
    Cli
  </div>
</h1>
<h3 align="center">
  Bird CLI - A command line interface for managing BirdPay merchant services
</h3>
<p>
Bird CLI is a command line interface for managing BirdPay merchant services, including miniprogram previews, account management, and API key operations.
</p>

<h2 align="center">
  <a href="https://demo.docuseal.tech">✨ Docs</a>
  <span>|</span>
  <a href="https://docuseal.com/sign_up">☁️ Examples</a>
</h2>

## Features
- Create miniprogram previews for testing and development
- Manage merchant account information and settings
- Handle API key operations (login, logout)
- Create and manage orders with line items
- Interactive command-line prompts for easy input
- Browser integration for preview links
- JSON output formatting for API responses
- Secure authentication handling
- Cross-platform support


## Install

### Homebrew
```sh
brew update

brew tap birdcorp/homebrew-bird-cli

brew install birdcli
```

#### Curl
```sh
curl -sL https://gist.githubusercontent.com/andyfen/64296525a465dd1f9cab7528f236c6b3/raw/be96fe8ab3c5fc962777c6192bceb0ee0eafd646/gistfile1.txt | sh
```


### Update
```sh       
brew update

brew upgrade birdcli
```





## Authentication

You need to obtain a merchant `API_KEY` first to use the cli.

- **Login**
  ```bash
  birdcli login
  ```

### Account

- **Get Account Info**
  ```bash
  birdcli account
  ```

- **Delete API Key**
  ```bash
  birdcli logout
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
  birdcli orders retrieve [orderID]
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
  birdcli miniprograms create-preview \
    --url https://miniprogram-developer.onrender.com/ \
    --name "Miniprogram Developer"
  ```

Example: w/ngrok url for localhost proxy

```bash
ngrok http 3000
```

```bash
  birdcli miniprograms create-preview \
    --url https://56fc-32-133-145-153.ngrok-free.app \
    --name "Miniprogram Developer"
```


- **Create Miniprogram**
  ```bash
  birdcli miniprogram init
  ```

- **Release Miniprogram**
  ```bash
  birdcli miniprogram publish
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





#### Release
export GITHUB_TOKEN=<token>

git tag v1.0.xxx

git push origin v1.0.xxx

goreleaser release --config .goreleaser/mac.yml --clean


