package wrap_cloudinary

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/eNViDAT0001/Backend/external/wrap_viper"
)

var (
	mediaServer *cloudinary.Cloudinary
	vp          = wrap_viper.GetViper()
)

func GetMediaServer() *cloudinary.Cloudinary {
	if mediaServer != nil {
		return mediaServer
	}
	mediaServer = newCloudinaryServer()
	return mediaServer
}
func newCloudinaryServer() *cloudinary.Cloudinary {
	cloud := vp.GetString("CLOUDINARY.NAME")
	key := vp.GetString("CLOUDINARY.API_KEY")
	secret := vp.GetString("CLOUDINARY.API_SECRET")
	cld, err := cloudinary.NewFromParams(cloud, key, secret)
	if err != nil {
		panic(err)
	}

	cld.Config.URL.Secure = vp.GetBool("CLOUDINARY.SECURE")
	return cld
}
