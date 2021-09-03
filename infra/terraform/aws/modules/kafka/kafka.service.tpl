[Unit]
Description=Apache Kafka server (broker)
Documentation=http://kafka.apache.org/documentation.html
Requires=network.target remote-fs.target
After=network.target remote-fs.target kafka.service

[Service]
Type=simple
User=kafka
Group=kafka
Environment=JAVA_HOME=/usr/lib/jvm/java-11-openjdk-amd64
ExecStart=/opt/kafka/bin/kafka-server-start.sh /etc/kafka.properties
ExecStop=/opt/kafka/bin/kafka-server-stop.sh

[Install]
WantedBy=multi-user.target
