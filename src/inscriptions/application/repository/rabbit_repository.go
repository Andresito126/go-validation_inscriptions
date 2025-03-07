package repository

type IRabbitService interface {
    Consume(queue string, handler func([]byte)) error
}