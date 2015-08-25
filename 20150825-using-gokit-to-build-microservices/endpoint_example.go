package main

type authenticateRequest struct {
	Credentials credentials.Credentials
}

type authenticateResponse struct {
	User  user.User
	Error error
}

func makeAuthenticateEndpoint(svc AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(authenticateRequest)
		u, err := svc.Authenticate(req.Credentials)
		return authenticateResponse{u, err}, nil
	}
}
