NAME = $(notdir $(PWD))

image: build
	docker build -t $(NAME):latest -f Dockerfile .

push@%: image
	$(eval VERSION = latest)
	$(eval TAG = $*/$(NAME):$(VERSION))
	docker tag $(NAME):$(VERSION) $(TAG)
	docker push $(TAG)

build:
	CGO_ENABLED=0 go build
