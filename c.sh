
SEMSETUP=$SEMBASE/bin
stopsem.sh
go build -tags postgres
cp lmsapieng lmsapieng_4008
cp lmsapieng_4008 $SEMSETUP
