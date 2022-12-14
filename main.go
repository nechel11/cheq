package main

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"log"
)

const (
	Test = "%s test"
)

func main() {
	_, repostStatus, _ := getRepoInfo(".")
	if !repostStatus.IsClean() {
		log.Fatal("directiory not commited")
	}
	fmt.Println("repostStatus is", repostStatus)
	print("ZAZAZA")
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
