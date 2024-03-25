package main

func loginService(userId int64) (AuthToken, error) {
	token, err := makeToken(userId)
	if err != nil {
		return AuthToken{}, err
	}

	return AuthToken{
		Token:       token,
		ExpiredTime: 1000,
	}, nil
}
