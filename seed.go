package seed

import "fmt"
import "regexp"
import "os"
import "log"
import "bufio"
import "github.com/logicminds/tools/go/vcs"
import "path"

func Sync_Seed_File(dest_path string, seed_file string) (repo_dirs []string) {
    file, err := os.Open(seed_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        repo_path, repo_ref := find_repo(scanner.Text())
        template_path := path.Join(dest_path, repo_path)
        var err error
        repo_dirs = append(repo_dirs, template_path)
        log.Printf("Fetching %s...\n", repo_path)
        vcs_cmd, _, err := vcs.FromDir(template_path, dest_path)
        // repo must not exist yet
        if err != nil {
            repo, _ := vcs.RepoRootForImportPath(repo_path, false)
            repo.Root = template_path
            vcs_cmd = repo.VCS
            if repo_ref != "" {
                err = vcs_cmd.CreateAtRev(repo.Root, repo.Repo, repo_ref)

            } else {
                err = vcs_cmd.Create(repo.Root, repo.Repo)
            }
            if err != nil {
                fmt.Printf("%v", err)
            }
        } else {
            vcs_cmd.DownloadAtRev(template_path, repo_ref)
        }
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return repo_dirs
}
// this is short term solution to having a full blown lexer to parse a seed file
func find_repo(input string) (repo_url string, branch string) {
	re := regexp.MustCompile(`get\s*(.*),\s*ref:\s*(.*)`)
	data := re.FindAllStringSubmatch(input, -1)
    if len(data) < 1 {
        log.Fatal("Error in Seed File")
    }
	if len(data[0]) > 1 {
		repo_path := data[0][1]
		repo_ref := data[0][2]
		return repo_path, repo_ref
	}
	return
}
