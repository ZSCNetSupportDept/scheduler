package signals

import (
	"github.com/golang-module/carbon/v2"
)

type TablePrototype struct {
	LastUpdated carbon.Carbon //值班表最后更新的时间
	NeedUpdate  bool          //值班表是否需要更新
}

var Table TablePrototype
