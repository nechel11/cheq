package main

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"log"
)

func main() {
	err := IsCommitted()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("NOT FATAL")
}

func IsCommitted() error {
	_, repostStatus, err := getRepoInfo(".")
	fmt.Println(repostStatus.IsClean(), err)
	if !repostStatus.IsClean() || err != nil {
		return fmt.Errorf("%s", "b;b")
	}
	return nil
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
