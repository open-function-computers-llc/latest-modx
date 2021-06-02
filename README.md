# MODX Version API

This api was written so that we could easily request all the latest tagged releases from the MODX revolution github account.

This is a very simple API. It has 4 routes:
```
/      :list the most recent stable releases
/all   :list all the releases as they are provided by github
/rc    :list only release candidtates
/alpha :list only alpha releases
```

## Github rate limiting

Because we are requesting info directly from github, and we don't want to authenticate, we are caching the results for 15 minutes. So, if a new release goes out it might not be reflected in this API for up to 15 minutes.
