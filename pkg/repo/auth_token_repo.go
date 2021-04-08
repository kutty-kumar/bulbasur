package repo

type AuthTokenRepo interface {
	logout(userId string, token string) error
}
