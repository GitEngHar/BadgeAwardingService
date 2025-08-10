package get

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/labstack/echo/v4"
	"hello-world/domain"
	"hello-world/domain/management"
	"hello-world/domain/notification"
	"hello-world/infra/db/dynamo"
	"hello-world/infra/queue"
	user "hello-world/usecase/badge"
	"hello-world/usecase/badgeRank"
	pushUseCase "hello-world/usecase/push"
	badgeAwardUsecase "hello-world/usecase/userBadgeAward"
	"net/http"
)

type Handler struct{}

func NewPublisherHandler() *Handler {
	return &Handler{}
}

// Do 配信したいメッセージをキューに入れておく
func (h Handler) Do(ctx context.Context) error {
	sqsConfig := queue.NewConfig(ctx)
	dbConfig := dynamo.NewConnectionDynamoDBForLocal()
	queuePublisherRepo := queue.NewPublisher(*sqsConfig)
	dynamodbRepo := dynamo.NewUserRepository(dbConfig)
	//TODO: テーブルが存在しない場合はエラーを返す
	publishMessageUseCase := pushUseCase.NewPublishMessageUseCase(queuePublisherRepo)
	badgeAwardGetUseCase := badgeAwardUsecase.NewGetUseCase(dynamodbRepo)
	pushTargetUsers, err := badgeAwardGetUseCase.GetAwardTargetUserByDate(ctx)
	if err != nil {
		return err
	}
	//TODO: 通知情報から通知情報を作成
	badgeAwardPublishers, err := createPublishers(ctx, pushTargetUsers, dynamodbRepo)
	if err != nil {
		return err
	}
	if len(badgeAwardPublishers) == 0 {
		return errors.New("no publishers found")
	}
	err = publishMessageUseCase.Do(ctx, badgeAwardPublishers)
	//err := publishMessageUseCase.Do(ctx, publisher.MessageBody, publisher.UserName, publisher.Address, publisher.Message)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func createPublishers(ctx context.Context, pushTargetUsers []map[string]types.AttributeValue, dynamodbRepo management.Repository) ([]notification.Publisher, error) {
	//TODO: imgUrlをちゃんとした内容にする
	var imgUrl string
	var publishers []notification.Publisher
	var userDomain *management.User
	var badgeRankDomain *management.BadgeDetailsByRank
	var publisher *notification.Publisher
	var awardComment *notification.AwardComment
	// userIDからuser情報を取得する
	userGetUseCase := user.NewGetUseCase(dynamodbRepo)
	// badgeRankIDからbadgeとbadgeRank情報を取得する
	badgeRankGetUseCase := badgeRank.NewGetUseCase(dynamodbRepo)

	for _, pushTargetUser := range pushTargetUsers {
		userID, isExistsKey := domain.GetStringAttr(pushTargetUser, "user_id")
		if !isExistsKey {
			return nil, errors.New("user_id does not exist key")
		}
		userMap, err := userGetUseCase.Do(ctx, userID)
		if err != nil {
			return nil, err
		}
		badgeRankID, isExistsKey := domain.GetStringAttr(pushTargetUser, "badge_rank_id")
		if !isExistsKey {
			return nil, errors.New("badge_rank_id does not exist key")
		}
		badgeRankMap, err := badgeRankGetUseCase.Do(ctx, badgeRankID)
		if err != nil {
			return nil, err
		}
		// 不正な値で処理をしないように、Domainに変換する
		userDomain, err = management.MapUserToDomain(userMap)
		if err != nil {
			return nil, err
		}
		badgeRankDomain, err = management.MapBadgeRankToDomain(badgeRankMap)
		if err != nil {
			return nil, err
		}
		awardComment, err = notification.NewAwardComment(badgeRankDomain.BadgeName, badgeRankDomain.Message, badgeRankDomain.Reason, badgeRankDomain.Effect)
		if err != nil {
			return nil, err
		}
		publisher, err = notification.NewPublisher(awardComment, imgUrl, string(userDomain.MailAddress))
		if publisher != nil {
			publishers = append(publishers, *publisher)
		} else {
			return nil, errors.New("failed to create the publisher userid :  " + userDomain.ID)
		}

	}
	return publishers, nil
}

func (h Handler) Hub(ctx echo.Context) error {
	return h.Do(ctx.Request().Context())
}
