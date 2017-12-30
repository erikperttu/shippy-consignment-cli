build:
	docker build -t shippy-consignment-cli .
image:
	docker build -t shippy-consignment-cli .
run:
	docker run --net="host" \
		-e MICRO_REGISTRY=mdns \
		shippy-consignment-cli consignment.json \
  	    <insert_token>