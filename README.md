# dockerhub-statsd

Serves `/metrics` endpoint with prometheus output:

```
dockerhub_image_pull_count{image="myimage"} 12345
```

where `myimage` is a docker image and `12345` is number of pulls according to docker hub API

# Usage

* Listen address specified by `LISTEN` environment variable, example: `:80`.
* List of images must be passed as arguments.
