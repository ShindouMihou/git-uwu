package utils

import "github.com/go-git/go-git/v5"

func GetWorktree() (*git.Repository, *git.Worktree, error) {
	repository, err := GetRepository()
	if err != nil {
		return nil, nil, err
	}
	worktree, err := repository.Worktree()
	if err != nil {
		return repository, nil, err
	}
	return repository, worktree, nil
}

func GetRepository() (*git.Repository, error) {
	repository, err := git.PlainOpen(GetCurrentDirectory())
	if err != nil {
		return nil, err
	}
	return repository, nil
}
