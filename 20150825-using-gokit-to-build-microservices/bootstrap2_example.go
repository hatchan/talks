func main() {
	ctx := context.Background()
	logger := log.NewLogfmtLogger(os.Stderr)
	svc := NewMongoAuthService(...)

	e := makeAuthenticateEndpoint(svc)

	e = loggingMiddleware(log.NewContext(logger).With("method", "authenticate"))(e)

	authenticateHandler := httptransport.Server{
		Context: ctx, 
		Endpoint: e, 
		DecodeFunc: decodeAuthenticateRequest, 
		EncodeFunc: encodeAuthenticateResponse,
	}
	
	http.Handle("/authenticate", authenticateHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

