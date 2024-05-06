package models

import (
	"errors"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("UNIQUE constraint failed: users.email")
	ErrUnknownComment     = errors.New("comment does not exist")
	ErrUnknownPost        = errors.New("post does not exist")
)

var (
	JSDuplicateEmail     = "user/signup/$2a$12$.HjibCBIzegvrQ/dwDRzkO8DzBVEnFU6PlGlqhtCYu4hLMCgF.F3G"
	JSInvalidCredentials = "user/login/$2a$12$qSexwUFOrmwkXucsSGdZu.1d0.YCu3/4gaxfVwdrRdb0b4ZnjP2de"
	JSSuccesfulSignup    = "user/login/$2a$12$B1q8dQJdtB/chK3XMVZmFeWsVKds1t0Oyw8KmMsyDdkVpxEEz.zmS"
	JSLogin              = "user/login/$2a$12$TV0unRiB6mg0U.4FhoznvuyD8mR5wG9v5SZWi8.T8lMY0bvc5CiCy"
	JSSignup             = "user/signup/$2a$12$0Oun5XW9qwWIj1sKSfktn.w2ldBuRh9Btbjz.i3Kgnp8SVt.KC12K"
)
