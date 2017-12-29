# glog
log for zap.Logger wrap

Example
----

import (

    "github.com/sanxia/glog"

)

func main(){
    log := glog.NewLogger("logMsg", glog.DebugLevel)

    log.Info("connect OK")

    log.InfoField("connect OK", "conn", "success")

    log.Infof("ConnectionLost %s\n", "err")


}
