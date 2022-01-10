## Criar tópico
kafka-topics --create --topic=teste --bootstrap-server=localhost:9092 --partitions=3

## Descrever tópico
kafka-topics --bootstrap-server=localhost:9092 --topic=teste --describe

## Criar consumer
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste
## Criar Producer
kafka-console-producer --bootstrap-server=localhost:9092 --topic=teste
## Cria consumer groups
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste  --group=x
## Ver grupos e offsets de consumer
kafka-consumer-groups --bootstrap-server=localhost:9092 --group=x --describe
