
# OAuth 2.0 System with Google and GitHub in Go

This project demonstrates the implementation of an OAuth 2.0 system with Single Sign-On (SSO) support for Google and GitHub in Go.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
- [Usage](#usage)
  - [Google OAuth](#google-oauth)
  - [GitHub OAuth](#github-oauth)
- [Testing](#testing)


## Features

- Google OAuth login with SSO support.
- GitHub OAuth login with SSO support.
- Basic user information retrieval from respective identity providers.

## Getting Started

### Prerequisites

- Go installed on your machine.
- Google OAuth credentials (client ID and client secret).
- GitHub OAuth credentials (client ID and client secret).
### Installation

1. Clone the repository:

  ```bash
   git clone https://github.com/swag2716/OAuth-Go
   cd OAuth-Go
  ```
2. Install dependencies
  
  ```bash
    go get -u ./...
  ```
  
    
### Configuration

- For Google OAuth:
```bash
export GOOGLE_OAUTH_CLIENT_ID=your-google-client-id
export GOOGLE_OAUTH_CLIENT_SECRET=your-google-client-secret
```
- For GitHub OAuth:
```bash
export GITHUB_OAUTH_CLIENT_ID=your-github-client-id
export GITHUB_OAUTH_CLIENT_SECRET=your-github-client-secret
```
### Usage

#### Google OAuth

1. Start the server:
```bash
go run main.go
```
2. Open your browser and navigate to http://localhost:8000/auth/google/login.

3. Follow the prompts to log in with your Google account.

4. After successful authentication, you will be redirected back to the home page with user information displayed

#### Github OAuth

1. Start the server:
```bash
go run main.go
```
2. Open your browser and navigate to http://localhost:8000/auth/github/login.

3. Follow the prompts to log in with your Github account.

4. After successful authentication, you will be redirected back to the home page with user information displayed


### Testing

Run the tests to ensure the correct functionality of the OAuth system:

```bash
go test -v
```

