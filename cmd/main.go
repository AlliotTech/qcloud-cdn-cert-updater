package main

import (
	"flag"
	"log"

	"qcloud-cdn-cert-updater/internal/cdn"
	"qcloud-cdn-cert-updater/internal/config"
	"qcloud-cdn-cert-updater/internal/ssl"
)

func main() {
	configFile := flag.String("f", "config.yaml", "Path to the configuration file")
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	sslClient := ssl.NewClient(cfg.SecretId, cfg.SecretKey)
	cdnClient := cdn.NewClient(cfg.SecretId, cfg.SecretKey)

	certId, err := sslClient.UploadCertificate(cfg.CertPath, cfg.KeyPath, "iots_vip_")

	for _, domain := range cfg.Domains {

		if err != nil {
			log.Fatalf("Failed to upload SSL certificate for domain %s: %v", domain, err)
		}

		log.Printf("New cert ID for domain %s: %s\n", domain, certId)

		err = cdnClient.UpdateDomainConfig(domain, certId)
		if err != nil {
			log.Fatalf("Failed to update CDN domain config for domain %s: %v", domain, err)
		}

		log.Printf("CDN certificate updated successfully for domain %s\n", domain)
	}

}
