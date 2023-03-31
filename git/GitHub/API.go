package api

const (
	GitHubAPIName      = "GitHub"
	GitHubAPIProto     = "https"
	GitHubAPIHost      = "api.github.com"
	GitHubAPIReposPath = "repos"
	GitHubAPITagsPath  = "tags"
)

type API struct {
	name      string
	proto     string
	host      string
	reposPath string
	tagsPath  string
}

func GitHubAPI() *API {
	return &API{
		name:      GitHubAPIName,
		proto:     GitHubAPIProto,
		host:      GitHubAPIHost,
		reposPath: GitHubAPIReposPath,
		tagsPath:  GitHubAPITagsPath,
	}
}
