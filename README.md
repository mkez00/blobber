Blobber
===================================

Overview
-----------------------------------

Blobber is a CLI program that acts as a wrapper around common cloud storage SDK's which allows users to easily upload/download desired content without having to constantly remember or retrieve access credentials and configuration details to a specific storage container.  Utilizing storage backends by common providers (AWS S3 supported only for now) a user configures an existing storage endpoint and access credentials once and can then upload/download content to that storage container as desired.

Why not just use scp?  Depending on your network configuration this may not be trivial.

Why not use your platforms native CLI?  AWS CLI (for example) does not have a default bucket so you must always specify one when accessing any S3 functionality.  Blobber allows you to setup all configuration details prior and never have to worry about them again.

Install and AWS S3 Configuration
-----------------------------------

1) Download `blobber` binary  and copy to `/usr/local/bin/` (<a href="https://s3.amazonaws.com/mk-blobber-storage/blobber-linux.zip" target="_blank">Linux Download</a>, <a href="https://s3.amazonaws.com/mk-blobber-storage/blobber-mac.zip" target="_blank">MacOS Download</a>)
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