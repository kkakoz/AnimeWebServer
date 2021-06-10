package usecase

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	countpb "red-bean-anime-server/api/count"
	"red-bean-anime-server/internal/app/anime/domain"
	userDomain "red-bean-anime-server/internal/app/user/domain"
	"red-bean-anime-server/internal/pkg/copier"
	"red-bean-anime-server/internal/pkg/msg"
	"red-bean-anime-server/internal/pkg/query"
	"red-bean-anime-server/pkg/auth"
	"red-bean-anime-server/pkg/db/mysqlx"
	"red-bean-anime-server/pkg/gerrors"
	"red-bean-anime-server/pkg/grpcx"
	"red-bean-anime-server/pkg/kafkax"
)

type AnimeUsecase struct {
	db           *gorm.DB
	animeRepo    domain.IAnimeRepo
	categoryRepo domain.ICategoryRepo
	countCli     countpb.CountServiceClient
	verifier     *auth.JwtTokenVerifier
	kafkaCli     sarama.SyncProducer
	videoRepo    domain.IVideoRepo
}

func (a *AnimeUsecase) UserUnLikeAnime(ctx context.Context, id int64, likeType bool) error {
	panic("implement me")
}

func (a *AnimeUsecase) UserLikeAnime(ctx context.Context, animeId int64, likeType bool) error {
	authorization := grpcx.GetAuthorization(ctx)
	user, err := a.verifier.Verifier(authorization)
	if err != nil {
		return err
	}
	countMsg := &msg.CountMsg{
		MsgType: msg.MsgTypeLike,
		UserId:  user.UserId,
		AnimeId: animeId,
	}
	countMsgData, err := json.Marshal(countMsg)
	err = kafkax.SendSyncMsgByte(a.kafkaCli, msg.MsgTypeCountTopic, countMsgData)
	return err
}

func (a *AnimeUsecase) GetAnimeInfo(ctx context.Context, animeId int64) (*domain.AnimeInfoRes, error) {
	authorization := grpcx.GetAuthorization(ctx)
	user, err := a.verifier.Verifier(authorization)
	if err != nil {
		return nil, err
	}
	count, err := a.countCli.GetAnimeCount(ctx, &countpb.AnimeCountReq{AnimeId: animeId, UserId: user.UserId})
	if err != nil {
		return nil, err
	}
	res, err := a.animeRepo.GetAnimeById(ctx, animeId)
	err = copier.Copy(res, count)
	if err != nil {
		return nil, err
	}
	videos, err := a.videoRepo.GetVideoByAnimeId(ctx, animeId)
	if err != nil {
		return nil, err
	}
	res.Videos = videos
	return res, nil
}

func (a *AnimeUsecase) GetAnimeList(ctx context.Context, page *query.Page, req *domain.AnimeListReq) ([]domain.Anime, error) {
	if req.CategoryId != 0 {
		animes, err := a.animeRepo.GetAnimeListByCategoryId(ctx, page, req)
		return animes, err
	}
	return a.animeRepo.GetAnimeList(ctx, page, req.Sort)
}

func (a *AnimeUsecase) AddAnime(ctx context.Context, addAnime *domain.AddAnime) error {
	ctx, tx := mysqlx.Begin(ctx, a.db)
	authorization := grpcx.GetAuthorization(ctx)
	user, err := a.verifier.Verifier(authorization)
	if err != nil {
		tx.Rollback()
		return err
	}
	if user.Auth != userDomain.UserAuthAdmin {
		return gerrors.NewBusErr("没有添加权限")
	}
	exist, err := a.animeRepo.GetAnimeExistByName(ctx, addAnime.Name)
	if err != nil {
		return err
	}
	if exist {
		return gerrors.NewBusErr("该动漫名已存在")
	}
	anime := &domain.Anime{
		Name:        addAnime.Name,
		Description: addAnime.Description,
		Year:        addAnime.Year,
		Quarter:     addAnime.Quarter,
	}
	err = a.animeRepo.AddAnime(ctx, anime)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = a.categoryRepo.AddAnimeCategory(ctx, int64(anime.ID), addAnime.CategoryIds)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = a.countCli.AddAnimeCount(ctx, &countpb.AnimeIdReq{AnimeId: anime.ID})
	if err != nil {
		tx.Rollback()
		return err
	}
	return errors.Wrap(tx.Commit().Error, "添加动漫失败")
}

func (a *AnimeUsecase) AddVideo(ctx context.Context, addVideo *domain.Video) error {
	exist, err := a.videoRepo.GetExistByAnimeIdEspisode(ctx, addVideo.AnimeId, addVideo.Episode)
	if exist {
		return gerrors.NewBusErr("上传集数已经存在")
	}
	video := &domain.Video{
		AnimeId: addVideo.AnimeId,
		Episode: addVideo.Episode,
		Name:    addVideo.Name,
		Url:     addVideo.Url,
	}
	err = a.animeRepo.AddVideo(ctx, video)
	return err
}

func NewAnimeUsecase(db *gorm.DB, animeRepo domain.IAnimeRepo, categoryRepo domain.ICategoryRepo,
	countCli countpb.CountServiceClient, produce sarama.SyncProducer, verifier *auth.JwtTokenVerifier, videoRepo domain.IVideoRepo) domain.IAnimeUsecase {
	return &AnimeUsecase{db: db, animeRepo: animeRepo, categoryRepo: categoryRepo,
		countCli: countCli, kafkaCli: produce, verifier: verifier, videoRepo: videoRepo}
}
