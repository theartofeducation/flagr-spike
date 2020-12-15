# flagr-spike

Spiking the feature flag service [Flagr](https://github.com/checkr/flagr).

## Getting Started

### Setup Environment

1. Copy `.env.sample` to `.env`
    * For Heroku app change `FLAGR_URL` to the app's URL from 1Password
    * Added `FLAGR_BASIC_AUTH_USERNAME` and `FLAGR_BASIC_AUTH_PASSWORD` from 1Password

### Using Docker

1. `docker-compose build`
1. `docker-compose up`

### Using Heroku

1. Ensure the dyno has been turned on
