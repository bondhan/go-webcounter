package application

import (
	"github.com/adjust/redismq"
	"github.com/bondhan/go-webcounter/domain"
	"github.com/bondhan/go-webcounter/domain/repository"
	"github.com/bondhan/go-webcounter/infrastructure/utils/redisclient"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"strconv"
)

// VisitorApp ...
type VisitorApp interface {
	IncrementCounter() (string, error)
	GetLastCounterFromDB() (domain.Visitor, error)
	IncrementCounterDirectDB() (domain.Visitor, error)
}
type visitorApp struct {
	visitorRepo  repository.VisitorRepository
	visitorQueue *redismq.BufferedQueue
	redisClient  *redis.Client
}

// NewVisitorApp ...
func NewVisitorApp(visitorRepo repository.VisitorRepository, visitorQueue *redismq.BufferedQueue, redisClient *redis.Client) VisitorApp {
	return &visitorApp{
		visitorRepo:  visitorRepo,
		visitorQueue: visitorQueue,
		redisClient:  redisClient,
	}
}

func (va *visitorApp) IncrementCounter() (string, error) {

	var err error
	var number uint64

	content, ok := redisclient.GetContentByKey("counter")
	if !ok {
		// get last counte from DB
		lastVisitor, err := va.visitorRepo.GetVisitor()
		if err != nil {
			logrus.Fatalf("Fail get last visitor %f", err)
		}

		// convert to string as we store in redis
		lastVisitorCounterStr := strconv.FormatUint(lastVisitor.Counter, 10)

		// create new redis key
		redisclient.CreateNewKey("counter", lastVisitorCounterStr, -1)

		number = lastVisitor.Counter

	} else {
		number, err = strconv.ParseUint(content.(string), 10, 64)
		if err != nil {
			return "", err
		}
	}

	//increment the counter
	number = number + 1

	numberStr := strconv.FormatUint(number, 10)

	redisclient.CreateNewKey("counter", numberStr, -1)
	if err != nil {
		return "", err
	}

	err = va.visitorQueue.Put(numberStr)

	return numberStr, err
}

func (va *visitorApp) GetLastCounterFromDB() (domain.Visitor, error) {

	visitorCounter, err := va.visitorRepo.GetVisitor()

	return visitorCounter, err
}

func (va *visitorApp) IncrementCounterDirectDB() (domain.Visitor, error) {

	writtenCounter, err := va.visitorRepo.IncrementVisitorNoParam()

	return writtenCounter, err
}
