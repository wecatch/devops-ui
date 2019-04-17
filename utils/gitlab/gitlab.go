package gitlab

import (
	"github.com/xanzy/go-gitlab"
	"github.com/wecatch/devops-ui/utils/log"
)

var logger = log.MakeLogger()

func init() {
	logger = log.Logger("utils.gitlab")
}

// GitLabClient for gitlab api client
var GitLabClient *ClientWrapper

// ClientWrapper for gitlab api client
type ClientWrapper struct {
	client *gitlab.Client
}

//NewGitLabClient for git client
func NewGitLabClient(token string, baseURL string) *ClientWrapper {
	client := &ClientWrapper{client: gitlab.NewClient(nil, token)}
	client.SetBaseURL(baseURL)
	return client
}

// SetBaseURL set gitlab api base url
func (gc *ClientWrapper) SetBaseURL(baseURL string) {
	gc.client.SetBaseURL(baseURL)
}

// ListGitlabProjects list all projects
func (gc *ClientWrapper) ListGitlabProjects() []*gitlab.Project {
	opt := &gitlab.ListProjectsOptions{
		Sort:       gitlab.String("asc"),
		Membership: gitlab.Bool(true),
		Simple:     gitlab.Bool(false),
		Search:     gitlab.String("background"),
		Visibility: gitlab.Visibility(gitlab.PrivateVisibility),
	}
	projects, _, err := gc.client.Projects.ListProjects(opt)
	if err != nil {
		logger.Warn(err)
	}
	return projects
}

// ListGitlabProjectTags for project tags
func (gc *ClientWrapper) ListGitlabProjectTags(pid int) []*gitlab.Tag {
	opt := &gitlab.ListTagsOptions{
		OrderBy: gitlab.String("updated"),
		Sort:    gitlab.String("desc"),
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 50,
		},
	}
	tags, _, err := gc.client.Tags.ListTags(pid, opt)
	if err != nil {
		logger.Warn(err)
	}

	return tags
}

// ListGitlabGroups for groups
func (gc *ClientWrapper) ListGitlabGroups() []*gitlab.Group {
	opt := &gitlab.ListGroupsOptions{
		Search: gitlab.String("background"),
	}
	groups, _, err := gc.client.Groups.ListGroups(opt)
	if err != nil {
		logger.Warn(err)
	}

	return groups
}

// ListGitlabGroupProjects  for group projects
func (gc *ClientWrapper) ListGitlabGroupProjects(gid int) []*gitlab.Project {
	opt := &gitlab.ListGroupProjectsOptions{
		Membership: gitlab.Bool(true),
		OrderBy:    gitlab.String("created_at"),
	}
	projects, _, err := gc.client.Groups.ListGroupProjects(gid, opt)
	if err != nil {
		logger.Warn(err)
	}

	return projects
}

// ListGitlabProjectCommits for project commit
func (gc *ClientWrapper) ListGitlabProjectCommits(pid int) []*gitlab.Commit {
	opt := &gitlab.ListCommitsOptions{}
	tags, _, err := gc.client.Commits.ListCommits(pid, opt)
	if err != nil {
		logger.Warn(err)
	}

	return tags
}
