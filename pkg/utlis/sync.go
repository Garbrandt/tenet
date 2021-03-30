package utlis

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Garbrandt/tenet/pkg/config"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"os"
)

func SyncFormLocal(path string, remote string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
	} else if os.IsNotExist(err) {
		return false, err
	} else {
		return false, err
	}

	var (
		src io.ReadCloser
		err error
	)

	src, err = os.Open(path)
	if err != nil {
		return false, err
	}
	defer func() {
		err := src.Close()
		if err != nil {

		}
	}()

	var dest bytes.Buffer
	_, err = io.Copy(&dest, src)
	if err != nil {
		return false, err
	}

	client := setupClient(config.Config.Sync.Token)
	_, err = createUpdateFile(
		client,
		config.Config.Sync.Owner,
		config.Config.Sync.Repo,
		remote,
		dest.Bytes(),
	)
	if err != nil {
		return false, err
	}

	return true, nil
}

func SyncingFormUrl(link string, remote string) (bool, error) {
	resp, err := http.Get(link)
	if err != nil {
		return false, err
	}

	var dest bytes.Buffer
	_, err = io.Copy(&dest, resp.Body)
	if err != nil {
		return false, err
	}

	client := setupClient(config.Config.Sync.Token)
	sha, err := createUpdateFile(
		client,
		config.Config.Sync.Owner,
		config.Config.Sync.Repo,
		remote,
		dest.Bytes(),
	)
	if err == nil {
		fmt.Printf("pushed %s file to %s@%s\n", sha, config.Config.Sync.Repo, sha[:7])
		return false, err
	}

	return true, nil
}

func setupClient(accessToken string) *github.Client {
	token := oauth2.Token{AccessToken: accessToken}
	source := oauth2.StaticTokenSource(&token)
	client := oauth2.NewClient(oauth2.NoContext, source)
	return github.NewClient(client)
}

func createUpdateFile(client *github.Client, owner, name, path string, data []byte) (string, error) {
	sha, err := getSha(client, owner, name, path)
	if err != nil {
		return createFile(client, owner, name, path, data)
	} else {
		return updateFile(client, owner, name, path, sha, data)
	}
}

// helper function to get the MAINTAINER files SHA
func getSha(client *github.Client, owner, name, path string) (string, error) {
	opt := new(github.RepositoryContentGetOptions)
	res, _, _, err := client.Repositories.GetContents(context.Background(), owner, name, path, opt)
	if err != nil {
		return "", err
	}
	return *res.SHA, nil
}

// helper function to update the MAINTAINER file
func updateFile(client *github.Client, owner, name, path, sha string, data []byte) (string, error) {
	res, _, err := client.Repositories.UpdateFile(context.Background(),
		owner,
		name,
		path,
		&github.RepositoryContentFileOptions{
			Content: data,
			Message: github.String("Updated MAINTAINERS file [CI SKIP]"),
			SHA:     github.String(sha),
		},
	)
	if err != nil {
		return "", fmt.Errorf("Error updating MAINTAINERS file at %s. %s", sha, err)
	}
	return *res.SHA, nil
}

// helper function to update the MAINTAINER file
func createFile(client *github.Client, owner, name, path string, data []byte) (string, error) {
	res, _, err := client.Repositories.CreateFile(
		context.Background(),
		owner,
		name,
		path,
		&github.RepositoryContentFileOptions{
			Content: data,
			Message: github.String(fmt.Sprintf("Created %s file", path)),
		},
	)
	if err != nil {
		return "", fmt.Errorf(fmt.Sprintf("Error creating %s file. %s", path, err))
	}
	return *res.SHA, nil
}
