package main

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
)

const (
	Test = "%s test"
)

func main() {
	repo, repostStatus, err := getRepoInfo(".")
	fmt.Println("repo is", repo)
	fmt.Println("repostStatus is", repostStatus)
	fmt.Println("err is", err)
}

func getRepoInfo(path string) (repo *git.Repository, repoStatus git.Status, err error) {
	repo, err = git.PlainOpen(path)
	gitRepoError := `
Ошибка при получении статуса git репозитория: %w
"scratch update" не предназначен для работы вне git репозиториев.
`
	if err != nil {
		err = fmt.Errorf(gitRepoError, err)
		return
	}
	worktree, err := repo.Worktree()
	if err != nil {
		err = fmt.Errorf(gitRepoError, err)
		return
	}
	repoStatus, err = worktree.Status()
	if err != nil {
		err = fmt.Errorf(gitRepoError, err)
		return
	}
	return
}
