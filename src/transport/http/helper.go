package http

func unmarshalJSON(reader io.Reader, obj interface{}) error {
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return api.ErrInvalidJSON
	}

	if err = json.Unmarshal(body, obj); err != nil {
		if _, ok := err.(base64.CorruptInputError); ok {
			return api.ErrInvalidBase64
		}
		return api.ErrInvalidJSON
	}

	return nil
}


var httpErr *errors.HTTPError
func handleError(ctx context.Context, err error, w http.ResponseWriter) {
	if err != nil && xerrors.As(err, httpErr) {
		w.SetStatus(err.StatusCode)
			
	} else {
		w.SetStatus(http.StatusInternalServerError)
	}

	w.Write(err.Error())
}
