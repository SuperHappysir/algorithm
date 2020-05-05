package LimitlessClassification

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_listToTree(t *testing.T) {
	jsonStr := `[
    {
        "channel_id":6719,
        "channel_name":"海际（重庆）",
        "pid":1
    },
    {
        "channel_id":6720,
        "channel_name":"它品商城",
        "pid":6719
    },
    {
        "channel_id":6728,
        "channel_name":"海际（广州）",
        "pid":1
    },
    {
        "channel_id":6733,
        "channel_name":"它品备货库",
        "pid":6719
    },
    {
        "channel_id":6736,
        "channel_name":"芬宠",
        "pid":6719
    },
    {
        "channel_id":6737,
        "channel_name":"积分商城",
        "pid":6719
    },
    {
        "channel_id":6738,
        "channel_name":"E宠专用渠道【已停用】",
        "pid":6719
    },
    {
        "channel_id":6741,
        "channel_name":"彭健代发",
        "pid":6719
    },
    {
        "channel_id":6742,
        "channel_name":"重庆海际1688代发",
        "pid":6719
    },
    {
        "channel_id":6759,
        "channel_name":"降龙库存",
        "pid":6719
    },
    {
        "channel_id":6760,
        "channel_name":"东部大区",
        "pid":6791
    },
    {
        "channel_id":6761,
        "channel_name":"西部大区",
        "pid":6790
    },
    {
        "channel_id":6762,
        "channel_name":"南部大区",
        "pid":6791
    },
    {
        "channel_id":6763,
        "channel_name":"北部大区",
        "pid":6790
    },
    {
        "channel_id":6764,
        "channel_name":"江左盟",
        "pid":6760
    },
    {
        "channel_id":6765,
        "channel_name":"桃花岛",
        "pid":6760
    },
    {
        "channel_id":6766,
        "channel_name":"六扇门",
        "pid":6762
    },
    {
        "channel_id":6767,
        "channel_name":"华山派",
        "pid":6763
    },
    {
        "channel_id":6768,
        "channel_name":"乐天派",
        "pid":6763
    },
    {
        "channel_id":6769,
        "channel_name":"通吃岛",
        "pid":6763
    },
    {
        "channel_id":6770,
        "channel_name":"逍遥派",
        "pid":6763
    },
    {
        "channel_id":6771,
        "channel_name":"独孤派",
        "pid":6760
    },
    {
        "channel_id":6772,
        "channel_name":"少林派",
        "pid":6760
    },
    {
        "channel_id":6773,
        "channel_name":"天地会",
        "pid":6760
    },
    {
        "channel_id":6774,
        "channel_name":"衡山派",
        "pid":6762
    },
    {
        "channel_id":6775,
        "channel_name":"聚贤庄",
        "pid":6762
    },
    {
        "channel_id":6776,
        "channel_name":"情义阁",
        "pid":6762
    },
    {
        "channel_id":6777,
        "channel_name":"铁血盟",
        "pid":6762
    },
    {
        "channel_id":6778,
        "channel_name":"云霄宫",
        "pid":6762
    },
    {
        "channel_id":6779,
        "channel_name":"朝天门",
        "pid":6761
    },
    {
        "channel_id":6780,
        "channel_name":"蜀山派",
        "pid":6761
    },
    {
        "channel_id":6781,
        "channel_name":"武当派",
        "pid":6761
    },
    {
        "channel_id":6782,
        "channel_name":"星宿派",
        "pid":6761
    },
    {
        "channel_id":6783,
        "channel_name":"兄弟会",
        "pid":6761
    },
    {
        "channel_id":6784,
        "channel_name":"线上销售组",
        "pid":6791
    },
    {
        "channel_id":6785,
        "channel_name":"东部大区备货",
        "pid":6760
    },
    {
        "channel_id":6786,
        "channel_name":"西部大区备货",
        "pid":6761
    },
    {
        "channel_id":6787,
        "channel_name":"南部大区备货",
        "pid":6762
    },
    {
        "channel_id":6788,
        "channel_name":"北部大区备货",
        "pid":6763
    },
    {
        "channel_id":6789,
        "channel_name":"降龙备货库",
        "pid":6759
    },
    {
        "channel_id":6790,
        "channel_name":"销售二部",
        "pid":6759
    },
    {
        "channel_id":6791,
        "channel_name":"销售一部",
        "pid":6759
    }
]`
	var list []*Node
	err := json.Unmarshal([]byte(jsonStr), &list)
	if err != nil {
		panic("json 反序列化错误")
	}

	type args struct {
		list []*Node
	}
	tests := []struct {
		name     string
		args     args
		dontWant *TreeNode
	}{
		{
			"列表转树",
			args{
				list,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := listToTree(tt.args.list); got == tt.dontWant {
				t.Errorf("listToTree() = %v, want %v", got, tt.dontWant)
			}
		})
	}
}

func Test_treeToList(t *testing.T) {
	jsonStr := `[
    {
        "channel_id":6719,
        "channel_name":"海际（重庆）",
        "pid":1
    },
    {
        "channel_id":6720,
        "channel_name":"它品商城",
        "pid":6719
    },
    {
        "channel_id":6728,
        "channel_name":"海际（广州）",
        "pid":1
    },
    {
        "channel_id":6733,
        "channel_name":"它品备货库",
        "pid":6719
    },
    {
        "channel_id":6736,
        "channel_name":"芬宠",
        "pid":6719
    },
    {
        "channel_id":6737,
        "channel_name":"积分商城",
        "pid":6719
    },
    {
        "channel_id":6738,
        "channel_name":"E宠专用渠道【已停用】",
        "pid":6719
    },
    {
        "channel_id":6741,
        "channel_name":"彭健代发",
        "pid":6719
    },
    {
        "channel_id":6742,
        "channel_name":"重庆海际1688代发",
        "pid":6719
    },
    {
        "channel_id":6759,
        "channel_name":"降龙库存",
        "pid":6719
    },
    {
        "channel_id":6760,
        "channel_name":"东部大区",
        "pid":6791
    },
    {
        "channel_id":6761,
        "channel_name":"西部大区",
        "pid":6790
    },
    {
        "channel_id":6762,
        "channel_name":"南部大区",
        "pid":6791
    },
    {
        "channel_id":6763,
        "channel_name":"北部大区",
        "pid":6790
    },
    {
        "channel_id":6764,
        "channel_name":"江左盟",
        "pid":6760
    },
    {
        "channel_id":6765,
        "channel_name":"桃花岛",
        "pid":6760
    },
    {
        "channel_id":6766,
        "channel_name":"六扇门",
        "pid":6762
    },
    {
        "channel_id":6767,
        "channel_name":"华山派",
        "pid":6763
    },
    {
        "channel_id":6768,
        "channel_name":"乐天派",
        "pid":6763
    },
    {
        "channel_id":6769,
        "channel_name":"通吃岛",
        "pid":6763
    },
    {
        "channel_id":6770,
        "channel_name":"逍遥派",
        "pid":6763
    },
    {
        "channel_id":6771,
        "channel_name":"独孤派",
        "pid":6760
    },
    {
        "channel_id":6772,
        "channel_name":"少林派",
        "pid":6760
    },
    {
        "channel_id":6773,
        "channel_name":"天地会",
        "pid":6760
    },
    {
        "channel_id":6774,
        "channel_name":"衡山派",
        "pid":6762
    },
    {
        "channel_id":6775,
        "channel_name":"聚贤庄",
        "pid":6762
    },
    {
        "channel_id":6776,
        "channel_name":"情义阁",
        "pid":6762
    },
    {
        "channel_id":6777,
        "channel_name":"铁血盟",
        "pid":6762
    },
    {
        "channel_id":6778,
        "channel_name":"云霄宫",
        "pid":6762
    },
    {
        "channel_id":6779,
        "channel_name":"朝天门",
        "pid":6761
    },
    {
        "channel_id":6780,
        "channel_name":"蜀山派",
        "pid":6761
    },
    {
        "channel_id":6781,
        "channel_name":"武当派",
        "pid":6761
    },
    {
        "channel_id":6782,
        "channel_name":"星宿派",
        "pid":6761
    },
    {
        "channel_id":6783,
        "channel_name":"兄弟会",
        "pid":6761
    },
    {
        "channel_id":6784,
        "channel_name":"线上销售组",
        "pid":6791
    },
    {
        "channel_id":6785,
        "channel_name":"东部大区备货",
        "pid":6760
    },
    {
        "channel_id":6786,
        "channel_name":"西部大区备货",
        "pid":6761
    },
    {
        "channel_id":6787,
        "channel_name":"南部大区备货",
        "pid":6762
    },
    {
        "channel_id":6788,
        "channel_name":"北部大区备货",
        "pid":6763
    },
    {
        "channel_id":6789,
        "channel_name":"降龙备货库",
        "pid":6759
    },
    {
        "channel_id":6790,
        "channel_name":"销售二部",
        "pid":6759
    },
    {
        "channel_id":6791,
        "channel_name":"销售一部",
        "pid":6759
    }
]`
	var list []*Node
	err := json.Unmarshal([]byte(jsonStr), &list)
	if err != nil {
		panic("json 反序列化错误")
	}
	type args struct {
		nodeList *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []*Node
	}{
		{
			"树转列表",
			args{
				listToTree(list),
			},
			list,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeToList(tt.args.nodeList); reflect.DeepEqual(got, tt.want) {
				t.Errorf("treeToList() = %v, want %v", got, tt.want)
			}
		})
	}
}