package dbtab_serviceaddrinfo

import (
	"fmt"
	"lmsapieng/include/common/dbdef"
	"lmsapieng/include/dbtabdef"
	"strings"
)

func Load_1_ServiceAddrInfoTable(serviceModule string, ServiceAddrInfoPtr *[]dbtabdef.ServiceAddrInfoTable, dbContext dbdef.DBContextDef, dberr *string, dbrejectreason *string) int {
	if strings.EqualFold(dbContext.DBType, dbdef.DBTypeOracle) {
		if Load_1_ServiceAddrInfoTableOracle(serviceModule, ServiceAddrInfoPtr, dbContext, dberr) < 0 {
			//trace.Lg("Load_1_ServiceAddrInfoTableOracle() failed for serviceModule[%s] with dberr[%s]", serviceModule, *dberr)
			*dbrejectreason = fmt.Sprintf("Load_1_ServiceAddrInfoTableOracle() failed for serviceModule[%s] with dberr[%s]", serviceModule, *dberr)
			return -1
		}
	} else if strings.EqualFold(dbContext.DBType, dbdef.DBTypePostgreSQL) {
		if Load_1_ServiceAddrInfoTablePostgres(serviceModule, ServiceAddrInfoPtr, dbContext, dberr) < 0 {
			//trace.Lg("Load_1_ServiceAddrInfoTablePostgres() failed for serviceModule[%s] with dberr[%s]", serviceModule, *dberr)
			*dbrejectreason = fmt.Sprintf("Load_1_ServiceAddrInfoTablePostgres() failed for serviceModule[%s] with dberr[%s]", serviceModule, *dberr)
			return -1
		}
	} else {
		//trace.Lg("db_type(%s) is Invalid...Load_1_ServiceAddrInfoTable() failed", dbContext.DBType)
		*dbrejectreason = fmt.Sprintf("db_type %s is Invalid...Load_1_ServiceAddrInfoTable() failed", dbContext.DBType)
		return -1
	}
	//trace.Lg("Load_1_ServiceAddrInfoTable() Success for serviceModule[%s]", serviceModule)
	return 1
}
