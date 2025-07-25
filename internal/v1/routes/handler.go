package routes

// Handler provides the dependencies for any endpoint, and is the reciever of the endpoint handling functions
type Handler struct {
	// db        *sql.DB
	// validator *validators.RequestValidator
}

func New() *Handler {
	return &Handler{
		// TODO: insert http Client for mapbox and NWS services
		// db:        db,  (*sql.DB)
		// validator: validators.New(),
	}
}
