#### Developer miniprogram

https://miniprogram-developer.onrender.com


zip -r upload.zip build

curl -X POST https://api.birdwallet.xyz/v2/miniprogram/17e4f2c119119e18/upload \
 -H "Accept: application/json" \
 -H "X-API-KEY: key_kbuf8TUptAe7nBJTMcfTFU" \
 -F "file=@upload.zip" \
 -F "version=1.0.0"

