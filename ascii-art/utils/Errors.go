package utils

import (
    "errors"
)

const (
    InputErrorText = "bad input"
    BannerError1Text = "wrong number of banner letters"
    BannerError2Text = "wrong letter weigth"
    BannerError3Text = "wrong letter height"
)
type CustomError struct {
    InputError error
    BannerError1 error
    BannerError2 error
    BannerError3 error
}

var ArtError = CustomError{
    InputError: errors.New(InputErrorText),
    BannerError1: errors.New(BannerError1Text),
    BannerError2: errors.New(BannerError2Text),
    BannerError3: errors.New(BannerError3Text),
}
