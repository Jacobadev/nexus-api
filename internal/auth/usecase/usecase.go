package usecase

import (
	"github.com/gateway-address/config"
	"github.com/gateway-address/internal/auth"
	"github.com/gateway-address/internal/auth/repository"
	model "github.com/gateway-address/internal/models"
	"github.com/gateway-address/pkg/httpErrors"
	"github.com/gateway-address/pkg/logger"
	"github.com/gateway-address/pkg/utils"
	"github.com/pkg/errors"
)

const (
	basePrefix    = "api-auth:"
	cacheDuration = 3600
)

// Auth UseCase
type authUC struct {
	cfg      *config.Config
	authRepo *repository.RepositorySqlite
	logger   logger.Logger
}

// Auth UseCase constructor
func NewAuthUseCase(cfg *config.Config, authRepo *repository.RepositorySqlite, log logger.Logger) auth.UseCase {
	return &authUC{cfg: cfg, authRepo: authRepo, logger: log}
}

func (u *authUC) Register(user *model.User) (*model.UserWithToken, error) {
	existsUser, err := u.authRepo.FindByEmail(user)
	if existsUser != nil || err == nil {
		u.logger.Error("%s, err: %v", httpErrors.ErrEmailAlreadyExists, err)
		return nil, errors.New(httpErrors.ErrEmailAlreadyExists)
	}

	if err := user.PrepareCreate(); err != nil {
		u.logger.Error("%s, err: %v", httpErrors.ErrPreparing, err)
		return nil, errors.WithMessage(err, httpErrors.ErrPreparing)
	}

	createdUser, err := u.authRepo.Register(user)
	if err != nil {
		u.logger.Error("%s, err: %v", httpErrors.ErrRegistering, err)
		return nil, err
	}
	createdUser.SanitizePassword()
	token, err := utils.GenerateJWTToken(createdUser, u.cfg)
	if err != nil {
		u.logger.Infof("%s err: %v", httpErrors.ErrGeneratingJWT, err)
		return nil, err
	}

	return &model.UserWithToken{
		User:  createdUser,
		Token: token,
	}, nil
}

func (u *authUC) Login(user *model.User) (*model.UserWithToken, error) {
	foundUser, err := u.authRepo.FindByEmail(user)
	if err != nil {
		u.logger.Errorf("%s, err: %v", httpErrors.ErrNoSuchUser, err)
		return nil, err
	}
	if err = foundUser.ComparePasswords(user.Password); err != nil {
		u.logger.Error("%s, err: %v", httpErrors.ErrWrongPassword, err)
		return nil, err
	}
	u.logger.Infof("user logged sucessfully: %d", foundUser.ID)
	foundUser.SanitizePassword()

	token, err := utils.GenerateJWTToken(foundUser, u.cfg)
	if err != nil {
		u.logger.Infof("%s err: %v", httpErrors.ErrGeneratingJWT, err)
		return nil, err
	}

	return &model.UserWithToken{
		User:  foundUser,
		Token: token,
	}, nil
}

// Update existing user
// func (u *authUC) Update(ctx context.Context, user *models.User) (*models.User, error) {
// 	if err := user.PrepareUpdate(); err != nil {
// 		return nil, httpErrors.NewBadRequestError(errors.Wrap(err, "authUC.Register.PrepareUpdate"))
// 	}
//
// 	updatedUser, err := u.authRepo.Update(ctx, user)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	updatedUser.SanitizePassword()
//
// 	if err = u.redisRepo.DeleteUserCtx(ctx, u.GenerateUserKey(user.UserID.String())); err != nil {
// 		u.logger.Errorf("AuthUC.Update.DeleteUserCtx: %s", err)
// 	}
//
// 	updatedUser.SanitizePassword()
//
// 	return updatedUser, nil
// }
//
// // Delete new user
// func (u *authUC) Delete(ctx context.Context, userID uuid.UUID) error {
// 	if err := u.authRepo.Delete(ctx, userID); err != nil {
// 		return err
// 	}
//
// 	return nil
// }
//
// // Get user by id
// func (u *authUC) GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
// 	return u.authRepo.GetByID(ctx, userID)
// }
//
// // Find users by name
// func (u *authUC) FindByName(ctx context.Context, name string, query *utils.PaginationQuery) (*models.UsersList, error) {
// 	return u.authRepo.FindByName(ctx, name, query)
// }
//
// // Get users with pagination
// func (u *authUC) GetUsers(ctx context.Context, pq *utils.PaginationQuery) (*models.UsersList, error) {
// 	return u.authRepo.GetUsers(ctx, pq)
// }
//
// // Login user, returns user model with jwt token
// func (u *authUC) Login(ctx context.Context, user *models.User) (*models.UserWithToken, error) {
// 	return &models.UserWithToken{}, nil
// }
//
// // Upload user avatar
// func (u *authUC) UploadAvatar(ctx context.Context, userID uuid.UUID, file models.UploadInput) (*models.User, error) {
// 	return nil, nil
// }
//
// func (u *authUC) GenerateUserKey(userID string) string {
// 	return fmt.Sprintf("%s: %s", basePrefix, userID)
// }
//
// func (u *authUC) generateAWSMinioURL(bucket string, key string) string {
// 	return fmt.Sprintf("%s/minio/%s/%s", u.cfg.AWS.MinioEndpoint, bucket, key)
//
