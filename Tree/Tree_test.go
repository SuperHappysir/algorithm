package Tree

import (
	"reflect"
	"testing"
)

func Test_preOrder(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name      string
		args      args
		wantArray []int
	}{
		{
			"二叉树的前序遍历",
			args{
				&Node{
					Value: 1,
					Left: &Node{
						Value: 2,
						Left: &Node{
							Value: 3,
							Left: &Node{
								Value: 4,
							},
							Right: &Node{
								Value: 5,
							},
						},
						Right: &Node{
							Value: 6,
							Left: &Node{
								Value: 7,
								Left: &Node{
									Value: 8,
								},
							},
							Right: &Node{
								Value: 9,
							},
						},
					},
					Right: &Node{
						Value: 10,
					},
				},
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotArray := preOrder(tt.args.node); !reflect.DeepEqual(gotArray, tt.wantArray) {
				t.Errorf("preOrder() = %v, want %v", gotArray, tt.wantArray)
			}
		})
	}
}

func Test_inOrder(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name      string
		args      args
		wantArray []int
	}{
		{
			"二叉树的中遍历",
			args{
				&Node{
					Value: 1,
					Left: &Node{
						Value: 2,
						Left: &Node{
							Value: 3,
							Left: &Node{
								Value: 4,
							},
							Right: &Node{
								Value: 5,
							},
						},
						Right: &Node{
							Value: 6,
							Left: &Node{
								Value: 7,
								Left: &Node{
									Value: 8,
								},
							},
							Right: &Node{
								Value: 9,
							},
						},
					},
					Right: &Node{
						Value: 10,
					},
				},
			},
			[]int{4, 3, 5, 2, 8, 7, 6, 9, 1, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotArray := inOrder(tt.args.node); !reflect.DeepEqual(gotArray, tt.wantArray) {
				t.Errorf("preOrder() = %v, want %v", gotArray, tt.wantArray)
			}
		})
	}
}

func Test_postOrder(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name      string
		args      args
		wantArray []int
	}{
		{
			"二叉树的后遍历",
			args{
				&Node{
					Value: 1,
					Left: &Node{
						Value: 2,
						Left: &Node{
							Value: 3,
							Left: &Node{
								Value: 4,
							},
							Right: &Node{
								Value: 5,
							},
						},
						Right: &Node{
							Value: 6,
							Left: &Node{
								Value: 7,
								Left: &Node{
									Value: 8,
								},
							},
							Right: &Node{
								Value: 9,
							},
						},
					},
					Right: &Node{
						Value: 10,
					},
				},
			},
			[]int{4, 5, 3, 8, 7, 9, 6, 2, 10, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotArray := postOrder(tt.args.node); !reflect.DeepEqual(gotArray, tt.wantArray) {
				t.Errorf("postOrder() = %v, want %v", gotArray, tt.wantArray)
			}
		})
	}
}
