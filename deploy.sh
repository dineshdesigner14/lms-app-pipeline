#!/bin/bash
set -euo pipefail
set -x

if [ $# -ne 3 ]; then
   echo "Please pass three args \$1-APP_NAME \$2-APP_PORT and \$3-DEPL_DIR"
   exit
fi

APP_NAME="$1"
APP_PORT="$2"
DEPL_DIR="$3"

BASE_PATH=/home/ec2-user
WORKING_DIR=$BASE_PATH/${DEPL_DIR}
BIN_DIR=$WORKING_DIR/bin
BUILD_BASE_DIR=$BASE_PATH/Build
BACKUP_BASE_DIR=$BASE_PATH/Backup
BIN_FILE=$(ls /tmp/${APP_NAME}_${APP_PORT} 2>/dev/null | head -n 1)

export SEMBASE="$BASE_PATH/${DEPL_DIR}"
PATH=${SEMBASE}/bin:$PATH
export SEMBASE PATH

if [ -z "$BIN_FILE" ]; then
  echo "No lmsapieng binary found in /tmp"
  exit 1
fi

BIN_NAME=$(basename "$BIN_FILE")

DATE=$(date +%Y-%m-%d)
TIME=$(date +%H-%M-%S)

BUILD_DIR=$BUILD_BASE_DIR/$DATE/$TIME/build
BACKUP_DIR=$BACKUP_BASE_DIR/$DATE/$TIME/build

mkdir -p "$BUILD_DIR"
mkdir -p "$BACKUP_DIR"
mkdir -p "$BIN_DIR"
mkdir -p "$BUILD_BASE_DIR"
mkdir -p "$BACKUP_BASE_DIR"
mkdir -p "$BUILD_DIR"
mkdir -p "$BACKUP_DIR"

cp "$BIN_FILE" "$BUILD_DIR/$BIN_NAME"
chmod +x "$BUILD_DIR/$BIN_NAME"

if [ -f "$BIN_DIR/$BIN_NAME" ]; then
  cp "$BIN_DIR/$BIN_NAME" "$BACKUP_DIR/"
fi

script -q -c "stopsem.sh" /dev/null || true
#stopsem.sh
sleep 2

echo "Copying the new BIN file to $BUILD_DIR"
cp "$BUILD_DIR/$BIN_NAME" "$BIN_DIR/$BIN_NAME"
chmod +x "$BIN_DIR/$BIN_NAME"

echo "Running startsem.sh"

startsem.sh
sleep 10

echo "Running semcmd new2 line..."

echo "Running semcmd startup with TTY"
script -q -c "semcmd" /dev/null <<EOF
startup
exit
EOF

ps -eaf | grep -v grep | grep "$BIN_NAME" || {
  echo "LMSAPIENG Service not running"
  # Add rollback strategy here
  #cp "$BACKUP_DIR/$BIN_NAME" "$BIN_DIR/$BIN_NAME"
  exit 1
}

echo "Deployment completed successfully"
