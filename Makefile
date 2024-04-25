PWD=/home/qy/project/logger-monitor

run-minio:
	docker run -dt                         \
	-p 9000:9000 -p 9001:9001                     \
	-v $(PWD)/minio/volume:/mnt/data                             \
	-v $(PWD)/minio/config:/etc/config       \
	-e "MINIO_CONFIG_ENV_FILE=/etc/config/config.env"    \
	--name "minio"                          \
	minio/minio server /mnt/data --console-address ":9001"

run-loki:
	docker run --name loki -d -v "$(PWD)/loki/loki-volume:/mnt/config" -p 3100:3100 grafana/loki:2.9.2 -config.file=/mnt/config/loki-config.yaml

run-promtail:
	docker run --name promtail -d -v $(PWD)/loki/promtail-volume:/mnt/config -v /var/log:/var/log --link loki grafana/promtail:2.9.2 -config.file=/mnt/config/promtail-config.yaml