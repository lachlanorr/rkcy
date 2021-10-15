#!/bin/bash
cd "$(dirname "$0")"

./init_db_aws.sh

./rpg platform replace --config_file_path=platform_aws.json --otelcol_endpoint=${otelcol_endpoint} --admin_brokers=${kafka_hosts[0]}
./rpg run --otelcol_endpoint=${otelcol_endpoint} --admin_brokers=${kafka_hosts[0]}
