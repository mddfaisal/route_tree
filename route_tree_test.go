package routetree

import "testing"

func TestAllFineRoutes(t *testing.T) {
	rt := NewRouteTree()
	tests := []struct {
		name     string
		path     string
		resource interface{}
	}{
		{"root", "/", "root_resource"},
		{"user", "/user", "user_resource"},
		{"user_profile", "/user/:id", "user_profile_resource"},
		{"search", "/search", "search_resource"},
		{"search_query", "/search?q=:query", "search_query_resource"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := rt.AddRoute(tt.path, tt.resource)
			if err != nil {
				t.Fatalf("failed to add route: %v", err)
			}

			res, params, err := rt.FindRoute(tt.path)
			if err != nil {
				t.Fatalf("failed to find route: %v", err)
			}

			if res != tt.resource {
				t.Errorf("expected resource %v, got %v", tt.resource, res)
			}

			if len(params) > 0 {
				t.Logf("found params: %v", params)
			}
		})
	}
}

func TestConflictRoutes(t *testing.T) {
	rt := NewRouteTree()
	tests := []struct {
		name     string
		path     string
		resource interface{}
		wantErr  bool
	}{
		{"conflict_user1", "/user/:id", "conflict_user_resource", false},
		{"conflict_user2", "/user/id", "conflict_user_resource", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := rt.AddRoute(tt.path, tt.resource)
			rt.DumpTree()
			if (err != nil) != tt.wantErr {
				t.Errorf("expected conflict error for route %s, but got %v", tt.path, err)
			}
		})
	}
}

func TestFindRoutes(t *testing.T) {
	rt := NewRouteTree()
	tests := []struct {
		name     string
		path     string
		resource interface{}
		wantErr  bool
	}{
		{"find_user", "/user/123", nil, true},
		{"find_search", "/search?q=test", nil, true},
		{"not_found", "/notfound", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, params, err := rt.FindRoute(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("expected error for route %s, but got %v", tt.path, err)
				return
			}

			if res != tt.resource {
				t.Errorf("expected resource %v, got %v", tt.resource, res)
			}

			if len(params) > 0 {
				t.Logf("found params: %v", params)
			}
		})
	}
}
