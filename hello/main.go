package kubeless

import (
	"fmt"
	"os"

	"github.com/kubeless/kubeless/pkg/functions"
	"github.com/sunfmin/fanout"
)

func Handler(event functions.Event, context functions.Context) (string, error) {
	fmt.Printf("%+v\n", event)
	fmt.Printf("DB_PARAMS: %+v\n", os.Getenv("DB_PARAMS"))

	var inputs []interface{}
	for i := 0; i < 100; i++ {
		inputs = append(inputs, i)
	}
	fanout.ParallelRun(10, func(i interface{}) (r interface{}, err error) {
		fmt.Println(i)
		return
	}, inputs)
	return event.Data, nil
}
