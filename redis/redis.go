package redis

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/iris-contrib/go.uuid"
	"github.com/ivahaev/go-logger"
	"github.com/orcaman/concurrent-map"
	"os"
)

var redisClient *redis.Client

const REQUEST_CHANNEL = "requests"
const RESPONSE_CHANNEL = "response"

var channels *redis.PubSub

var callbacks = cmap.New()

type CallBack func(payload interface{})

type Request struct {
	Command  string
	CallBack CallBack
	Id       string
	PayLoad  []interface{}
}

type Response struct {
	Status  string      `json:"status"`
	Content interface{} `json:"content"`
	Id      string      `json:"id"`
}

func InitRedis() {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	pass := os.Getenv("REDIS_PASSWORD")

	if len(host) < 1 {
		host = "localhost"
	}

	if len(port) < 1 {
		port = "6379"
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: pass,
		DB:       0,
	})

	channels = redisClient.Subscribe(REQUEST_CHANNEL, RESPONSE_CHANNEL)

	go listen()
}

func listen() {
	for {
		message := <-channels.Channel()
		if message.Channel == RESPONSE_CHANNEL {
			response := Response{}

			err := json.Unmarshal([]byte(message.Payload), &response)

			if err != nil {
				logger.Warnf("Can't unmarshal JSON: %s", err.Error())
				return
			}

			result, success := callbacks.Get(response.Id)

			if result != nil && success {
				req, _ := result.(Request)
				req.CallBack(response.Content)
				callbacks.Remove(req.Id)
			}
		}
	}
}

func GetRedis() *redis.Client {
	return redisClient
}

func PublishRequest(request Request) {
	bytes, err := json.Marshal(request)

	if err != nil {
		logger.Error(err)
		panic(err)
	}

	redisClient.Publish(REQUEST_CHANNEL, string(bytes))
	callbacks.Set(request.Id, request)
}

func NewRequest(command string, callback CallBack, payLoad ...interface{}) *Request {
	id := uuid.Must(uuid.NewV4())
	return &Request{Command: command, CallBack: callback, PayLoad: payLoad, Id: id.String()}
}
