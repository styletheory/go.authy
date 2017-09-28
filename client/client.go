package client

import (
  "github.com/styletheory/go.authy/pkg/verification"
)

type AuthyClient struct {
  ApiKey                string
  BaseUrl               string
  Verification          *verification.VerificationService
}

func NewAuthyClient(apiKey string) (*AuthyClient, error) {
  return &AuthyClient {
    ApiKey:         apiKey,
    BaseUrl:        "https://api.authy.com",
    Verification:   verification.NewVerificationService(apiKey, "https://api.authy.com"),
  }, nil
}