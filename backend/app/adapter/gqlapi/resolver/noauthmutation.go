package resolver

import(
	"errors"
	"fmt"

	"github.com/short-d/short/backend/app/adapter/gqlapi/input"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/adapter/gqlapi/scalar"
	"github.com/short-d/short/backend/app/usecase/changelog"
	"github.com/short-d/short/backend/app/usecase/shortlink"
	"github.com/short-d/short/backend/app/usecase/repository"
	"github.com/short-d/short/backend/app/usecase/keygen"
)

// NoAuthMutation represents GraphQL mutation resolver that does not require an authToken
type NoAuthMutation struct {
	changeLog        changelog.ChangeLog
	shortLinkCreator shortlink.Creator
	shortLinkUpdater shortlink.Updater
	userRepo 				 repository.User
	keyGen   				 keygen.KeyGenerator
}

// CreateShortLinkArgs represents the possible parameters for CreateShortLink endpoint
type NoAuthCreateShortLinkArgs struct {
	ShortLink input.ShortLinkInput
	IsPublic  bool
}

// CreateShortLink creates mapping between an alias and a long link for a given user
func (a NoAuthMutation) CreateShortLink(args *NoAuthCreateShortLinkArgs) (*ShortLink, error) {
	shortLink := args.ShortLink.CreateShortLinkInput()
	isPublic := args.IsPublic
	username := shortLink.GetUsername("")
	user := entity.User{
		ID: "dummyID",
	}

	exists, err := a.userRepo.IsEmailExist(username)
	if err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		return nil, ErrUnknown{}
	}
	
	if exists {
		user, err = a.userRepo.GetUserByEmail(username)
		if err != nil {
			fmt.Println("Error!")
			fmt.Println(err)
			return nil, ErrUnknown{}
		}
	} else {
		fmt.Println("Creating user")
		userID, err := a.createAccount(entity.SSOUser{
			ID: username,
			Email: username,
			Name: username,
		})
		if err != nil {
			fmt.Println("Error!")
			fmt.Println(err)
			return nil, ErrUnknown{}
		}
		fmt.Println("Getting created user")
		user, err = a.userRepo.GetUserByID(userID)
		if err != nil {
			fmt.Println("Error!")
			fmt.Println(err)
			return nil, ErrUnknown{}
		}

	}


	newShortLink, err := a.shortLinkCreator.CreateShortLink(shortLink, user, isPublic)
	if err == nil {
		return &ShortLink{shortLink: newShortLink}, nil
	}

	var (
		ae shortlink.ErrAliasExist
		l  shortlink.ErrInvalidLongLink
		c  shortlink.ErrInvalidCustomAlias
		m  shortlink.ErrMaliciousLongLink
	)
	if errors.As(err, &ae) {
		return nil, ErrAliasExist(shortLink.GetCustomAlias(""))
	}
	if errors.As(err, &l) {
		return nil, ErrInvalidLongLink{shortLink.GetLongLink(""), string(l.Violation)}
	}
	if errors.As(err, &c) {
		return nil, ErrInvalidCustomAlias{shortLink.GetCustomAlias(""), string(c.Violation)}
	}
	if errors.As(err, &m) {
		return nil, ErrMaliciousContent(shortLink.GetLongLink(""))
	}
	return nil, ErrUnknown{}
}

// UpdateShortLinkArgs represents the possible parameters for updateShortLink endpoint
type NoAuthUpdateShortLinkArgs struct {
	OldAlias  string
	ShortLink input.ShortLinkInput
}

// UpdateShortLink updates the relationship between the short link and the user
func (a NoAuthMutation) UpdateShortLink(args *NoAuthUpdateShortLinkArgs) (*ShortLink, error) {
	user := entity.User{
		Email: "alpha@example.com",
	}

	update := args.ShortLink.CreateShortLinkInput()

	newShortLink, err := a.shortLinkUpdater.UpdateShortLink(args.OldAlias, update, user)
	if err == nil {
		return &ShortLink{shortLink: newShortLink}, nil
	}

	var (
		ae shortlink.ErrAliasExist
		l  shortlink.ErrInvalidLongLink
		c  shortlink.ErrInvalidCustomAlias
		m  shortlink.ErrMaliciousLongLink
		nf shortlink.ErrShortLinkNotFound
		ns shortlink.ErrEmptyAlias
	)
	if errors.As(err, &ae) {
		return nil, ErrAliasExist(update.GetCustomAlias(""))
	}
	if errors.As(err, &l) {
		return nil, ErrInvalidLongLink{update.GetLongLink(""), string(l.Violation)}
	}
	if errors.As(err, &c) {
		return nil, ErrInvalidCustomAlias{update.GetCustomAlias(""), string(c.Violation)}
	}
	if errors.As(err, &m) {
		return nil, ErrMaliciousContent(update.GetLongLink(""))
	}
	if errors.As(err, &nf) {
		return nil, ErrShortLinkNotFound(args.OldAlias)
	}
	if errors.As(err, &ns) {
		return nil, ErrEmptyAlias{}
	}
	return nil, ErrUnknown{}
}

// ChangeInput represents possible properties for Change
type NoAuthChangeInput struct {
	Title           string
	SummaryMarkdown *string
}

// CreateChangeArgs represents the possible parameters for CreateChange endpoint
type NoAuthCreateChangeArgs struct {
	Change ChangeInput
}

func (a NoAuthMutation) CreateChange(args *NoAuthCreateChangeArgs) (*Change, error) {
	user := entity.User{
		Email: "alpha@example.com",
	}
	
	change, err := a.changeLog.CreateChange(args.Change.Title, args.Change.SummaryMarkdown, user)
	if err == nil {
		change := newChange(change)
		return &change, nil
	}

	var (
		u changelog.ErrUnauthorizedAction
	)
	if errors.As(err, &u) {
		return nil, ErrUnauthorizedAction(fmt.Sprintf("user %s is not allowed to create a change", user.ID))
	}
	return nil, ErrUnknown{}
}

// DeleteChangeArgs represents the possible parameters for DeleteChange endpoint
type NoAuthDeleteChangeArgs struct {
	ID string
}

// DeleteChange removes a Change with given ID from change log
func (a NoAuthMutation) DeleteChange(args *NoAuthDeleteChangeArgs) (*string, error) {
	user := entity.User{
		Email: "alpha@example.com",
	}

	err := a.changeLog.DeleteChange(args.ID, user)
	if err == nil {
		return &args.ID, nil
	}

	var (
		u changelog.ErrUnauthorizedAction
	)
	if errors.As(err, &u) {
		return nil, ErrUnauthorizedAction(fmt.Sprintf("user %s is not allowed to delete the change %s", user.ID, args.ID))
	}
	return nil, ErrUnknown{}
}

// UpdateChangeArgs represents the possible parameters for UpdateChange endpoint.
type NoAuthUpdateChangeArgs struct {
	ID     string
	Change ChangeInput
}

func (a NoAuthMutation) UpdateChange(args *NoAuthUpdateChangeArgs) (*Change, error) {
	user := entity.User{
		Email: "alpha@example.com",
	}

	change, err := a.changeLog.UpdateChange(
		args.ID,
		args.Change.Title,
		args.Change.SummaryMarkdown,
		user,
	)
	if err == nil {
		change := newChange(change)
		return &change, nil
	}

	var (
		u changelog.ErrUnauthorizedAction
	)
	if errors.As(err, &u) {
		return nil, ErrUnauthorizedAction(fmt.Sprintf("user %s is not allowed to update the change %s", user.ID, args.ID))
	}
	return nil, ErrUnknown{}
}

// ViewChangeLog records the time when the user viewed the change log
func (a NoAuthMutation) ViewChangeLog() (scalar.Time, error) {
	user := entity.User{
		Email: "alpha@example.com",
	}

	lastViewedAt, err := a.changeLog.ViewChangeLog(user)
	return scalar.Time{Time: lastViewedAt}, err
}

func (a NoAuthMutation) createAccount(ssoUser entity.SSOUser) (string, error) {
	userID, err := a.generateUnassignedUserID()
	if err != nil {
		return "", err
	}
	err = a.createUser(userID, ssoUser.Name, ssoUser.Email)
	return userID, err
}

func (a NoAuthMutation) generateUnassignedUserID() (string, error) {
	newKey, err := a.keyGen.NewKey()
	return string(newKey), err
}

func (a NoAuthMutation) createUser(id string, name string, email string) error {
	user := entity.User{
		ID:    id,
		Name:  name,
		Email: email,
	}
	return a.userRepo.CreateUser(user)
}

func newNoAuthMutation(
	changeLog changelog.ChangeLog,
	shortLinkCreator shortlink.Creator,
	shortLinkUpdater shortlink.Updater,
	userRepo repository.User,
	keyGen keygen.KeyGenerator,
) NoAuthMutation {
	return NoAuthMutation{
		changeLog:        changeLog,
		shortLinkCreator: shortLinkCreator,
		shortLinkUpdater: shortLinkUpdater,
	  userRepo:         userRepo,
		keyGen:						keyGen,
	}
}