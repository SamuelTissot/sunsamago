package sunsama

//
// import (
// 	"testing"
// 	"time"
//
// 	"github.com/shurcooL/graphql"
// )
//
// func Test_task_Duration(t1 *testing.T) {
// 	tests := []struct {
// 		name    string
// 		task    task
// 		want    time.Duration
// 		wantErr bool
// 	}{
// 		{
// 			name: "with duration",
// 			task: task{
// 				actualTime: []actualTime{
// 					{
// 						Duration: graphql.Int(10),
// 					},
// 					{
// 						Duration: graphql.Int(30),
// 					},
// 				},
// 			},
// 			want:    time.Minute * 40,
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t1.Run(
// 			tt.name, func(t1 *testing.T) {
// 				t := task{
// 					Text:       tt.fields.Text,
// 					Completed:  tt.fields.Completed,
// 					Subtasks:   tt.fields.Subtasks,
// 					actualTime: tt.fields.actualTime,
// 				}
// 				got, err := t.Duration()
// 				if (err != nil) != tt.wantErr {
// 					t1.Errorf("Duration() error = %v, wantErr %v", err, tt.wantErr)
// 					return
// 				}
// 				if got != tt.want {
// 					t1.Errorf("Duration() got = %v, want %v", got, tt.want)
// 				}
// 			},
// 		)
// 	}
// }
//
// func Test_task_Title(t1 *testing.T) {
// 	type fields struct {
// 		Text          graphql.String
// 		Completed     graphql.Boolean
// 		scheduledTime scheduledTime
// 		Subtasks      []struct {
// 			Title      graphql.String
// 			actualTime []actualTime
// 		}
// 		actualTime []actualTime
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t1.Run(
// 			tt.name, func(t1 *testing.T) {
// 				t := task{
// 					Text:          tt.fields.Text,
// 					Completed:     tt.fields.Completed,
// 					scheduledTime: tt.fields.scheduledTime,
// 					Subtasks:      tt.fields.Subtasks,
// 					actualTime:    tt.fields.actualTime,
// 				}
// 				if got := t.Title(); got != tt.want {
// 					t1.Errorf("Title() = %v, want %v", got, tt.want)
// 				}
// 			},
// 		)
// 	}
// }
