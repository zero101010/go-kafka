## O que são os tópicos
- Os tópicos são como se fossem as filas do rabbitMq, no entanto eles tem divisões e estruturas diferentes
- Os tópicos podem ser consumidos por mais de um serviço sem retirar a informacao e possuem uma camada de banco de dados, caso seja necessário reprocessar um dado por algum motivo.
- A estrutura do tópico é divido em Partições e offsets como mostra a imagem abaixo

![Kafka estrutura](https://atitudereflexiva.files.wordpress.com/2019/11/kafka-architecture.png)