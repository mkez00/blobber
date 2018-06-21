Blobber
===================================

Overview
-----------------------------------

This program solves the problem of giving a user an easily accessible blob storage when doing day to day activities.  Utilizing storage backends by common providers (AWS S3 supported only for now) a user configures the storage endpoint and access credentials once and can upload/download content to that storage container as desired.

Why not just use scp?  Depending on your network configuration this may not be trivial.
Why not use your platforms native CLI?  AWS CLI (for example) does not have a default bucket so you must always specify one when accessing any S3 functionality.  Blobber allows you to setup all configuration details prior and never have to worry about them again.

Install
-----------------------------------

1) Download `blobber` binary <a href="https://s3.amazonaws.com/mk-blobber-storage/blobber.zip" target="_blank">here</a> and copy to `/usr/local/bin/`
2) Run `blobber init`
3) Modify `.blobber` config which defaults to `~\.blobber`

Usage
-----------------------------------

1) List items in bucket: `blobber list`
2) Download item in bucket `blobber get <ITEM_NAME>`
3) Upload item to bucket `blobber put <ITEM_NAME>`
4) Delete item from bucket `blobber delete <ITEM_NAME>`
5) List config options `blobber config`