package db

import (
	"testing"
)

func TestGetDbPoolConn(t *testing.T) {
	tests := []struct {
		name       string
		wantErr    bool
		beforeFunc func()
	}{
		{
			name:    "Success",
			wantErr: false,
		},
		/*{
			name:    "Error",
			wantErr: true,
			beforeFunc: func() {
				os.Setenv("PGUSER", "doesnotexist")
				os.Setenv("PGPASSWORD", "1234567890")
				os.Setenv("PGDATABASE", "dbthatdoesnt exist")
				os.Setenv("PGHOST", "localhost")
			},
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetDBPoolConn()
			if tt.beforeFunc != nil {
				tt.beforeFunc()
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("GetDbPoolConn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
