#!/usr/bin/env bash

function build() {
	ROOT=$(dirname "$0")
	NAME="cloud-logger"
	VERSION=$(lookup-version "$ROOT"/../internal/const/const.go)
	DIST=$ROOT/"../dist/${NAME}"

	OS=${1}
	ARCH=${2}

	if [ -z "$OS" ]; then
		echo "usage: build.sh OS ARCH"
		exit
	fi
	if [ -z "$ARCH" ]; then
		echo "usage: build.sh OS ARCH"
		exit
	fi

	echo "checking ..."
	ZIP_PATH=$(which zip)
	if [ -z "$ZIP_PATH" ]; then
		echo "we need 'zip' command to compress files"
		exit
	fi

	echo "building v${VERSION}/${OS}/${ARCH} ..."
	ZIP="${NAME}-${OS}-${ARCH}-v${VERSION}.zip"
	echo "target ${ZIP} ..."

	echo "copying ..."
	if [ ! -d "$DIST" ]; then
		mkdir "$DIST"
		mkdir "$DIST"/bin
		mkdir "$DIST"/configs
		mkdir "$DIST"/logs
	fi

	cp "$ROOT"/configs/config.template.yaml "$DIST"/configs

	echo "building ..."
	env GOOS="${OS}" GOARCH="${ARCH}" go build -trimpath -o "$DIST"/bin/${NAME} -ldflags "-s -w" "$ROOT"/../cmd/cloud-logger/main.go

	if [ ! -f "${DIST}/bin/${NAME}" ]; then
		echo "build failed!"
		exit
	fi

	# delete hidden files
	find "$DIST" -name ".DS_Store" -delete
	find "$DIST" -name ".gitignore" -delete

	echo "zip files"
	cd "${DIST}/../" || exit
	if [ -f "${ZIP}" ]; then
		rm -f "${ZIP}"
	fi
	zip -r -X -q "${ZIP}" ${NAME}/
	rm -rf ${NAME}
	cd - || exit

	echo "OK"
}

function lookup-version() {
	FILE=$1
	VERSION_DATA=$(cat "$FILE")
	re="Version[ ]+=[ ]+\"([0-9.]+)\""
	if [[ $VERSION_DATA =~ $re ]]; then
		VERSION=${BASH_REMATCH[1]}
		echo "$VERSION"
	else
		echo "could not match version"
		exit
	fi
}

build "$1" "$2" "$3"
