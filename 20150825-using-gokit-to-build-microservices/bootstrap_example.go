func main() {
	ctx := context.Background()
	svc := NewMongoAuthService(...)

	authenticateHandler := httptransport.Server{Context: ctx, Endpoint: makeAuthenticateEndpoint(svc), DecodeFunc: decodeAuthenticateRequest, EncodeFunc: encodeAuthenticateResponse }
	
	http.Handle("/authenticate", authenticateHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodeAuthenticateRequest(r *http.Request) (interface{}, error) {
	var request authenticateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	r.Body.Close()
	return request, nil
}

func encodeAuthenticateResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
