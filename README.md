Blobber
===================================

Overview
-----------------------------------

Blobber is a CLI program that simplifies cloud blob storage access when you just want to securely upload/download files without having to constantly remember or retrieve access credentials and configuration details to a specific storage container.  

How It Works?
-----------------------------------

A user creates a storage container on the cloud provider of their choice.  After installing the binary for their respective platform they run `blobber init` which will generate a configuration file that requires a one time update by the user with their credentials and other configuration options.  After doing so they can start using the program to upload/download content to the defined storage container without having to worry about re-entering authentication credentials and other configuration settings.

Why not just use `scp`?  Depending on your network configuration and/or platform this may not be trivial or even possible.

Why not use your platform's native CLI?  AWS CLI (for example) does not allow you to set a default bucket so you must always specify one when accessing any S3 functionality.  Blobber allows you to setup all configuration details prior and never have to worry about them again.

Install and Initialization
-----------------------------------

1) Download `blobber` binary (<a href="https://s3.amazonaws.com/mk-blobber-storage/blobber-linux.zip" target="_blank">Linux Download</a>, <a href="https://s3.amazonaws.com/mk-blobber-storage/blobber-mac.zip" target="_blank">MacOS Download</a>, <a href="https://s3.amazonaws.com/mk-blobber-storage/blobber-win.zip" target="_blank">Windows Download</a>)
2) If on Linux/MacOS, copy to `/usr/local/bin/`.  If on Windows, update Path environment variable to include directory where executable resides
3) Run `blobber init`

AWS S3 Configuration
---------------------------------

1) Create AWS S3 bucket for use with Blobber
2) Create a new user with programmatic access and give access to S3 bucket (keep Access Key ID and Secret Access Key for next step)
3) Update `.blobber` located at `~/.blobber` setting `StorageService` to `AmazonS3` and  updating AWS required configuruation options (`AwsAccessKeyId`, `AwsSecretAccessKey`, `AwsS3Bucket` and `AwsRegion`)
4) Start using Blobber

Google Cloud Storage Configuration
-----------------------------------

1) Create Google Cloud Storage bucket for use with blobber
2) Create Service Account Key and download JSON credentials file
3) Put credentials file in home directory
4) Update `.blobber` located at `~/.blobber` setting `StorageService` to `GoogleCloudStorage` and updating GCS required configuration options (`GcpCredentialsFile` and `GcpBucketName`)  

Usage
-----------------------------------

1) List items in bucket: `blobber list`
2) Download item in bucket `blobber get <ITEM_NAME>`
3) Upload item to bucket `blobber put <ITEM_NAME>`
4) Delete item from bucket `blobber delete <ITEM_NAME>`
5) List config options `blobber config`