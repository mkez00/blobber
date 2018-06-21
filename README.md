Blobber
===================================

Overview
-----------------------------------

Blobber is a CLI program that acts as a wrapper around common cloud storage SDK's which allows users to easily upload/download desired content without having to constantly remember or retrieve access credentials and configuration details to a specific storage container.  

How It Works?
-----------------------------------

A user creates a storage container on the cloud provider of their choice (currently only AWS S3 is supported).  After installing the binary for their respective platform they run `blobber init` which will generate a configuration file that requires the user to update with their credentials and other configuration options.  After they can start using the program to upload/download content to the defined storage container without having to worry about re-entering authentication credentials and other configuration settings.

Why not just use `scp`?  Depending on your network configuration and/or platform this may not be trivial or even possible.

Why not use your platform's native CLI?  AWS CLI (for example) does not allow you to set a default bucket so you must always specify one when accessing any S3 functionality.  Blobber allows you to setup all configuration details prior and never have to worry about them again.

Install and AWS S3 Configuration
-----------------------------------

1) Download `blobber` binary  and copy to `/usr/local/bin/` (<a href="https://s3.amazonaws.com/mk-blobber-storage/blobber-linux.zip" target="_blank">Linux Download</a>, <a href="https://s3.amazonaws.com/mk-blobber-storage/blobber-mac.zip" target="_blank">MacOS Download</a>, <a href="https://s3.amazonaws.com/mk-blobber-storage/blobber-win.zip" target="_blank">Windows Download</a>
)
2) Run `blobber init`
3) Create AWS S3 bucket for use with Blobber
4) Create a new user with programmatic access and give access to S3 bucket (keep Access Key ID and Secret Access Key for next step)
5) Update `.blobber` located at `~/.blobber` updating AWS requiring configuruation options (`AwsAccessKeyId`, `AwsSecretAccessKey`, `AwsS3Bucket` and `AwsRegion`)
6) Start using Blobber

Usage
-----------------------------------

1) List items in bucket: `blobber list`
2) Download item in bucket `blobber get <ITEM_NAME>`
3) Upload item to bucket `blobber put <ITEM_NAME>`
4) Delete item from bucket `blobber delete <ITEM_NAME>`
5) List config options `blobber config`