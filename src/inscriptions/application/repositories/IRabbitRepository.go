package repositories

type IRabbitRepository interface {
    SendMessageToBroker(studentID, courseID, status string)
}

