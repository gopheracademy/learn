# Setting Up and Deploying Learn

## Set Up Databases

Setup the users and passwords listed in `database.yml`.

## Setup Auth Keys

Currently we support 5 different login methods: `GitHub`, `LinkedIn`, `Twitter`, `FaceBook`, and `Google+`. In production OAuth applications need to be created and set up to point at the production URL of the application.

If we don't want to support these five logins, we need to remove them from `actions/auth.go` and `templates/_nav_bar.html`.

```text
GITHUB_KEY
GITHUB_SECRET
```

```text
LINKEDIN_KEY
LINKEDIN_SECRET
```

```text
TWITTER_KEY
TWITTER_SECRET
```

```text
FACEBOOK_KEY
FACEBOOK_SECRET
```

```text
GPLUS_KEY
GPLUS_SECRET
```

#### Development

Here is a development key for GitHub. It only works on `127.0.0.1`.

```text
GITHUB_KEY="308545eb0909a67d1581"
GITHUB_SECRET="c8ed04e6dcb9a37dd3711de0991e2fc51e04303e"
```

## GitHub Callback

The ENV var `GITHUB_WEBHOOK_SECRET` needs to be set up with the production secret. This doesn't need to be set up in development.

## Training Repo

All that is left is to have access to the `github.com/gopheracademy/training` repo locally from the web server. The repo should be set up at `$GOPATH/src/github.com/gopheracademy/training`.

The application needs both `go` and `git` tooling in order to run properly.
