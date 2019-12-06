# sangu-bca

## usage
1. There is a type named `Client` that should be instantiated through `NewClient` which hold any possible setting to the library.
2. There is a gateway classes which you will be using depending on whether you used. The gateway type need a Client instance.
3. Any activity is done in the gateway level (current: Get Token and Get Account Statement)

## example
see main.go