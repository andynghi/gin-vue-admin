package config

type Captcha struct {
	KeyLong            int `mapstructure:"key-long" json:"key-long" yaml:"key-long"`                                     // Verification code length
	ImgWidth           int `mapstructure:"img-width" json:"img-width" yaml:"img-width"`                                  // Verification code width
	ImgHeight          int `mapstructure:"img-height" json:"img-height" yaml:"img-height"`                               // Verification code height
	OpenCaptcha        int `mapstructure:"open-captcha" json:"open-captcha" yaml:"open-captcha"`                         // Explosion-proof verification code turns on this number, 0 means that a verification code is required for each login, other numbers represent incorrect passwords This number, such as 3, means that a verification code will appear after three errors.
	OpenCaptchaTimeOut int `mapstructure:"open-captcha-timeout" json:"open-captcha-timeout" yaml:"open-captcha-timeout"` // Explosion-proof verification code timeout, unit: s (seconds)
}
