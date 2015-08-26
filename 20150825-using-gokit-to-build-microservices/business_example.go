type AuthService interface {
	Authenticate(credentials.Credentials) (user.User, error)
}

type FakeAuthService struct {}

func (s *FakeAuthService) Authenticate(creds credentials.Credentials) (user.User, error) {
	if creds.Username == "bvdberg" { 
		return user.User{
			ID:       1,
			Username: "bvdberg",
			Role:     "Admin",
		}, nil
	}
	return nil, nil
}

type MongoAuthService struct {} ...
