package settings

type QiniuConf struct {
	Zone          string `mapstructure:"zone"`
	Bucket        string `mapstructure:"bucket"`
	ImgPath       string `mapstructure:"img-path"`
	UseHttps      bool   `mapstructure:"use-https"`
	Ak            string `mapstructure:"access-key"`
	Sk            string `mapstructure:"secret-key"`
	UseCdnDomains bool   `mapstructure:"use-cdn-domains"`
}
