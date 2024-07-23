### qcloud-cdn-cert-updater  
一个小项目，用来配合 Certbot 实现自动更新 腾讯云 CDN SSL证书。  



### 配置 renewal-hooks  
vim /etc/letsencrypt/renewal-hooks/deploy/qcloud-cdn-cert-updater.sh
``` shell 
#!/bin/bash
/etc/certbot/cdn/qcloud-cdn-cert-updater -f /etc/certbot/cdn/config.yml  2>&1
```