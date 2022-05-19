FILE = main.go
image_name = alfredvalderrama/metapod:v1

start:
	export POD_NAMESPACE="prod"
	export POD_IP="127.0.0.1"
	export NODE_NAME="example.com"
	export POD_NAME="index-html"
	export POD_CPU_REQUEST="120m"
	export POD_CPU_LIMIT="200m"
	export POD_MEM_REQUEST="120m"
	export POD_MEM_LIMIT="200m"
	export POD_SERVICE_ACCOUNT="test"
	export CLUSTER_NAME="Cluster1"
	export GIN_MODE=debug

	go run $(FILE)

build: $(FILE)
	echo "[INFO]: Building package $(FILE)"
	go build $(FILE)

build-image:
	echo "[INFO]: Docker image build process starting ..."
	docker build -t $(image_name) .

deploy-image:
	docker push $(image_name)
