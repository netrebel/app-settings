# app-settings

Source: https://travix.io/making-your-go-app-configurable-bb5e5f4a9df9

* A settings file in your repository for local testing (appsettings.json)
* Environment variables that can be configured upon deployment
* Some logic in your Go code to combine these two into a settings struct that you can use in your application.

To test run it without ENV variable `GOOGLE_APPLICATION_CREDENTIALS`, you'll get:

```
panic: Failed to update with env vars: envconfig: keys GOOGLE_APPLICATION_CREDENTIALS, google_application_credentials not found
```

GOOGLE_APPLICATION_CREDENTIALS is required here, because it doesn’t contain the “optional” tag. And it’s also not in the settings file, because the credentials should only be stored on your local machine:

```
...
	Google struct {
      Application struct {
         Credentials string
      }
   }
...
```


and then add it with:

```
export GOOGLE_APPLICATION_CREDENTIALS="secret-password"
```

You'll get:
```
{Log:{MinFilter:Debug} Cors:{Origins:[*]} SomeService:{URL:https://some-service.com} SomePublisher:{Env:staging Project:staging-project Topic:some-pubsub-topic} Google:{Application:{Credentials:secret-password}} SomeXML:{Storage:{BucketID:some-xml-bucket} AuthBasicToken:LocalUse} DefaultProvider:}
```