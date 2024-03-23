docker run -d -p 18123:8123 -p19000:9000 --name some-clickhouse-server --ulimit nofile=262144:262144 clickhouse/clickhouse-server

docker exec -it some-clickhouse-server clickhouse-client

CREATE DATABASE IF NOT EXISTS helloworld

CREATE TABLE helloworld.my_first_table
(
    user_id UInt32,
    message String,
    timestamp DateTime,
    metric Float32
)
ENGINE = MergeTree()
PRIMARY KEY (user_id, timestamp)

SELECT *
FROM helloworld.my_first_table
ORDER BY timestamp


go get github.com/ClickHouse/clickhouse-go
