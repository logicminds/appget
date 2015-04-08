package seed

import "testing"

func TestSync_Seed_File(t *testing.T) {
	repos := Sync_Seed_File("/tmp/seeds", "Seedfile.example.txt")
	if  len(repos) != 2 {
		t.Error("Expected 2, got", len(repos))
	}
	if repos[0] != "/tmp/seeds/github.com/logicminds/puppet-retrospec" {
		t.Error("/tmp/seeds/github.com/logicminds/puppet-retrospec got", repos[0])

	}
	if repos[1] != "/tmp/seeds/github.com/logicminds/micro-puppet" {
		t.Error("/tmp/seeds/github.com/logicminds/micro-puppet got", repos[1])

	}
}
