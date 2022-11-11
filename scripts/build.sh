#!/bin/bash -e
GITHUB_TAG="${GITHUB_TAG:-local}"
SHA512_CMD="${SHA512_CMD:-sha512sum}"
export BCBSN_DELIVERABLE="${BCBSN_DELIVERABLE:-bcbsn}"

echo "GITHUB_TAG: '$GITHUB_TAG' BCBSN_DELIVERABLE: '$BCBSN_DELIVERABLE'"
cd cmd/bcbsn
go build -buildvcs=false -ldflags "-X main.Version=${GITHUB_TAG}" -o "${BCBSN_DELIVERABLE}"
$SHA512_CMD "${BCBSN_DELIVERABLE}" >"${BCBSN_DELIVERABLE}.sha512.txt"
chmod +x "${BCBSN_DELIVERABLE}"
cd ../..
