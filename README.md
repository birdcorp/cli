<div align="center">
  <a href="https://www.docuseal.com">
    <img 
      alt="Bird" 
      src="https://payments-webapp-assets-stage.s3.us-west-2.amazonaws.com/bird.png" 
      width="50">
  </a>
  <h1 style="border-bottom: none;">BirdPay CLI</h1>
  <h3>A Command-Line Interface for Managing BirdPay Merchant Services</h3>
  <p>
    Bird CLI is a command-line tool for managing BirdPay merchant operations. It provides functionality for handling orders, accounts, webhooks, and generating miniapp previews, all in one place to streamline your workflows.
  </p>
  <h2>
    <a href="https://docs-openapi.onrender.com">📚 Documentation</a>
    <span>|</span>
    <a href="https://docuseal.com/sign_up">🚀 Examples</a>
  </h2>
</div>



<br/>

## ✨ Features
- Create and test miniapp previews
- Manage merchant accounts and settings
- API key handling: Login, logout, and secure storage
- Create and manage orders with customizable line items
- Interactive prompts for intuitive command-line usage
- Preview links with browser integration
- JSON-formatted output for API responses
- Secure authentication workflows
- Cross-platform support

<br/>

## 🚀 Installation

Install Bird CLI using a single command:

### Script Installation

```sh
curl -sL https://link.birdwallet.xyz/cli | sh
```

<br/>

   
### Homebrew (macOS/Linux)

1. Add the Bird CLI Homebrew tap:
  ```sh
  brew tap birdcorp/homebrew-bird-cli
  ```

2. Install Bird CLI:
  ```sh
  brew install birdcli
  ```

3. Update Bird CLI:
  ```sh       
  brew update
  brew upgrade birdcli
  ```

<br/>


## 🔑 Authentication

### Login
Authenticate with your `API_KEY` to access Bird CLI features:

  ```bash
  birdcli login
  ```

### Logout

  ```bash
  birdcli logout
  ```

<br/>

## 🧾 Commands Overview

### Account Management

View account details:

  ```bash
  birdcli account
  ```

### Get a resource

Get a resource by ID: order, coupon, webhook, miniapp

  ```bash
  birdcli get <identifier>
  ```

### Orders

Create an order:

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

List all orders:
  ```bash
  birdcli orders list
  ```

<br/>

### Mini-Apps

**Development**

Create a miniapp preview to generate a scannable link. You can use ngrok for local testing or provide a staging URL:
  ```bash
  birdcli miniapps preview \
    --url <preview_url> \
    --name <name>
  ```


**Production**

Initialize a miniapp configuration to create a configurable `miniapp.config.json` file in your current directory.

  ```bash
  birdcli miniapp init
  ```
  
```
{
  "appInfo": {
    "appID": "180c9ec28117dc63",
    "name": "Minesweeper",
    "version": "1.0.0",
    "description": "Play minesweeper",
    "tags": []
  },
  "appearance": {
    "backgroundColor": "#000000",
    "foregroundColor": "#FFFFFF",
    "navBackgroundColor": "transparent",
    "navTextColor": "light"
  },
  "build": {
    "buildDirectory": "_build"
  },
  "assets": {
    "appIcon": "./app-icon.png"
  },
  "users": {
    "testers": []
  },
  "configuration": {
    "defaultLanguage": "en",
    "privacyPolicyUrl": "https://myapp.com/privacy",
    "termsOfServiceUrl": "https://myapp.com/terms"
  }
}
```

To publish a miniapp, run the following command in the directory containing the `miniapp.config.json` file. Note that you must update the version in the `miniapp.config.json` file before each publish:

  ```bash
  birdcli miniapp publish
  ```
  
<br/>

Get Miniprogram info
  ```bash
  birdcli get [appID]
  ```

List all miniapps:
  ```bash
  birdcli miniapp list
  ```


<br/>

### Webhooks

Create a webhook:
  ```bash
  birdcli webhook create --url <webhook_url>
  ```

List webhooks:
  ```bash
  birdcli webhook list
  ```

Delete a webhook:
  ```bash
  birdcli webhook delete --id <webhookID>
  ```

<br/>

### Events

Get event details by ID:
  ```bash
  birdcli events get <eventID>
  ```

Stream live events:
  ```bash
  birdcli events stream
  ```

<br/>



# 🌐 Example: Miniprogram Preview with `ngrok`

1. Start an `ngrok` proxy for your local server:
  ```bash
  ngrok http 3000
  ```
2. Use the `ngrok` URL to create a preview:
  ```bash
  birdcli miniapps create-preview --url <ngrok_url> --name "MyApp"
  ```

<br/>

# 🛠️ Development and Releases

To create a new release:

1. Authenticate with GitHub:
  ```bash
  export GITHUB_TOKEN=<your_github_token>
  ```
2. Tag the release version:
  ```bash
  git tag v1.0.x
  git push origin v1.0.x
  ```
3. Build and publish the release:
  ```bash
  goreleaser release --config .goreleaser/mac.yml --clean
  ```

<br/>

# 💬 Feedback and Contributions
We welcome your feedback and contributions! Create an issue or submit a pull request on GitHub.

