# https://hub.docker.com/r/otel/opentelemetry-collector-contrib
IMAGE := otel/opentelemetry-collector-contrib
TAG   := latest

CONFIG := config.yaml
OUTPUT := output.json

all:
	@cat $(MAKEFILE_LIST)

run: clean _req_DRIVER _req_DSN
	@touch $(OUTPUT)
	docker run -v ./$(CONFIG):/etc/otelcol-contrib/config.yaml -v ./output.json:/output.json $(IMAGE):$(TAG)

_req_%: $(CONFIG)
	@grep $* .env | cut -d "=" -f2 | xargs -I{} sed -i 's#%$*%#{}#g' $<

$(CONFIG):
	@cp $@.template $@

clean:
	@rm -f $(CONFIG) $(OUTPUT)

tail:
	tail -F $(OUTPUT)
