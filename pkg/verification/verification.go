package verification

import (
    "github.com/styletheory/go.authy/pkg/http"
    "github.com/styletheory/go.authy/pkg/errors"
    "fmt"
    "log"
    "encoding/json"
    "strconv"
    "bytes"
)

type VerificationService struct {
    ApiKey      string
    BaseUrl     string
    HttpClient  http.Http
}

type VerificationCreationRequest struct {
    Via             string  `json:"via"`
    CountryCode     int     `json:"country_code"`
    PhoneNumber     string  `json:"phone_number"`
    CodeLength      int     `json:"code_length"`
    Locale          string  `json:"locale"`
    Message         string  `json:"message"`
}

type VerificationCreationResponse struct {
    Carrier     string  `json:"carrier"`
    IsCellphone bool    `json:"is_cellphone"`
    Message     string  `json:"message"`
    UUID        string  `json:"uuid"`
    IsSuccess   bool    `json:"success"`
}

func NewVerificationService(apiKey string, baseUrl string) *VerificationService {
    return &VerificationService {
        ApiKey:     apiKey,
        BaseUrl:    baseUrl,
        HttpClient: http.NewDefaultHttpClient(),
    }
}

func (s *VerificationService) StartVerification(method string, countryCode int, phoneNumber string, codeLength int, locale string) (*VerificationCreationResponse, error) {
    if method != "sms" && method != "call" {
        return nil, &errors.InvalidVerificationMethodError{1, fmt.Sprintf("Cannot request verification using %s", method)}
    }
    body2 := VerificationCreationRequest {
        Via:            method,
        CountryCode:    countryCode,
        PhoneNumber:    phoneNumber,
        CodeLength:     codeLength,
        Locale:         locale,
        Message:        "asdf",
    }

    _, err := json.Marshal(body2)
    if err != nil {
        log.Println("Failed to marshal json")
        return nil, err
    }

    b := new(bytes.Buffer)
    json.NewEncoder(b).Encode(body2)

    queryParams := make(map[string]string)
    queryParams["api_key"] = s.ApiKey

    res, err := s.HttpClient.Post(s.BaseUrl + "/protected/json/phones/verification/start", "application/json", b, nil, queryParams)
    if err != nil {
        return nil, err
    }

    defer res.Body.Close()

    response := new(VerificationCreationResponse)
    if(res.StatusCode >=200 && res.StatusCode < 300) {
        json.NewDecoder(res.Body).Decode(response)
    }
    return response, err
}

func (s *VerificationService) Verify(countryCode int, phoneNumber string, verificationCode string) (bool, error) {
    request := make(map[string]string)
    request["api_key"] = s.ApiKey
    request["country_code"] = strconv.Itoa(countryCode)
    request["phone_number"] = phoneNumber
    request["verification_code"] = verificationCode

    res, err := s.HttpClient.Get(s.BaseUrl + "/protected/json/phones/verification/check", "application/json", nil, request)
    defer res.Body.Close()
    if err != nil {
        return false, err
    }

    if(res.StatusCode == 200) {
        return true, nil
    } else if (res.StatusCode == 401) {
        return false, &errors.InvalidVerificationCodeError{201, "Invalid verification code used."}
    } else if (res.StatusCode == 404) {
        return false, &errors.NoVerificationFoundError{202, "No verification is pending for this phone number."}
    } else {
        return false, &errors.UnknownError{999, fmt.Sprintf("Unknonw error, verification returns %i.", res.StatusCode)}
    }
}