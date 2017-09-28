package errors

import (
  "fmt"
) 

type InvalidVerificationMethodError struct {
  ErrorCode int
  Reason    string
}

type InvalidVerificationCodeError struct {
    ErrorCode int
    Reason string
}

type NoVerificationFoundError struct {
    ErrorCode int
    Reason string
}

type UnknownError struct {
    ErrorCode int
    Reason string
}

func (e *InvalidVerificationMethodError) Error() string {
  return fmt.Sprintf("%d - %s", e.ErrorCode, e.Reason)
}

func (e *InvalidVerificationCodeError) Error() string {
  return fmt.Sprintf("%d - %s", e.ErrorCode, e.Reason)
}

func (e *NoVerificationFoundError) Error() string {
  return fmt.Sprintf("%d - %s", e.ErrorCode, e.Reason)
}

func (e *UnknownError) Error() string {
  return fmt.Sprintf("%d - %s", e.ErrorCode, e.Reason)
}