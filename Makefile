BIN_DIR=bin
BYNARY=detector

BYNARY_OUTPUT= ${BIN_DIR}/${BYNARY}

all:
	go build -o ${BYNARY_OUTPUT}
