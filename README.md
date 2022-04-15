# backend-repo



## Getting started

To make it easy for you to get started with GitLab, here's a list of recommended next steps.

Already a pro? Just edit this README.md and make it your own. Want to make it easy? [Use the template at the bottom](#editing-this-readme)!

## Add your files

- [ ] [Create](https://docs.gitlab.com/ee/user/project/repository/web_editor.html#create-a-file) or [upload](https://docs.gitlab.com/ee/user/project/repository/web_editor.html#upload-a-file) files
- [ ] [Add files using the command line](https://docs.gitlab.com/ee/gitlab-basics/add-file.html#add-a-file-using-the-command-line) or push an existing Git repository with the following command:

```
cd existing_repo
git remote add origin https://gitlab.com/isikaykut/backend-repo.git
git branch -M main
git push -uf origin main
```

## Integrate with your tools

- [ ] [Set up project integrations](https://gitlab.com/isikaykut/backend-repo/-/settings/integrations)

## Collaborate with your team

- [ ] [Invite team members and collaborators](https://docs.gitlab.com/ee/user/project/members/)
- [ ] [Create a new merge request](https://docs.gitlab.com/ee/user/project/merge_requests/creating_merge_requests.html)
- [ ] [Automatically close issues from merge requests](https://docs.gitlab.com/ee/user/project/issues/managing_issues.html#closing-issues-automatically)
- [ ] [Enable merge request approvals](https://docs.gitlab.com/ee/user/project/merge_requests/approvals/)
- [ ] [Automatically merge when pipeline succeeds](https://docs.gitlab.com/ee/user/project/merge_requests/merge_when_pipeline_succeeds.html)

## Test and Deploy

Use the built-in continuous integration in GitLab.

- [ ] [Get started with GitLab CI/CD](https://docs.gitlab.com/ee/ci/quick_start/index.html)
- [ ] [Analyze your code for known vulnerabilities with Static Application Security Testing(SAST)](https://docs.gitlab.com/ee/user/application_security/sast/)
- [ ] [Deploy to Kubernetes, Amazon EC2, or Amazon ECS using Auto Deploy](https://docs.gitlab.com/ee/topics/autodevops/requirements.html)
- [ ] [Use pull-based deployments for improved Kubernetes management](https://docs.gitlab.com/ee/user/clusters/agent/)
- [ ] [Set up protected environments](https://docs.gitlab.com/ee/ci/environments/protected_environments.html)

***

# Editing this README

When you're ready to make this README your own, just edit this file and use the handy template below (or feel free to structure it however you want - this is just a starting point!).  Thank you to [makeareadme.com](https://www.makeareadme.com/) for this template.

## Suggestions for a good README
Every project is different, so consider which of these sections apply to yours. The sections used in the template are suggestions for most open source projects. Also keep in mind that while a README can be too long and detailed, too long is better than too short. If you think your README is too long, consider utilizing another form of documentation rather than cutting out information.

## Name
Choose a self-explaining name for your project.

## Description
Let people know what your project can do specifically. Provide context and add a link to any reference visitors might be unfamiliar with. A list of Features or a Background subsection can also be added here. If there are alternatives to your project, this is a good place to list differentiating factors.

## Badges
On some READMEs, you may see small images that convey metadata, such as whether or not all the tests are passing for the project. You can use Shields to add some to your README. Many services also have instructions for adding a badge.

## Visuals
Depending on what you are making, it can be a good idea to include screenshots or even a video (you'll frequently see GIFs rather than actual videos). Tools like ttygif can help, but check out Asciinema for a more sophisticated method.

## Installation
Within a particular ecosystem, there may be a common way of installing things, such as using Yarn, NuGet, or Homebrew. However, consider the possibility that whoever is reading your README is a novice and would like more guidance. Listing specific steps helps remove ambiguity and gets people to using your project as quickly as possible. If it only runs in a specific context like a particular programming language version or operating system or has dependencies that have to be installed manually, also add a Requirements subsection.

## Usage
Use examples liberally, and show the expected output if you can. It's helpful to have inline the smallest example of usage that you can demonstrate, while providing links to more sophisticated examples if they are too long to reasonably include in the README.

## Support
Tell people where they can go to for help. It can be any combination of an issue tracker, a chat room, an email address, etc.

## Roadmap
If you have ideas for releases in the future, it is a good idea to list them in the README.

## Contributing
State if you are open to contributions and what your requirements are for accepting them.

For people who want to make changes to your project, it's helpful to have some documentation on how to get started. Perhaps there is a script that they should run or some environment variables that they need to set. Make these steps explicit. These instructions could also be useful to your future self.

You can also document commands to lint the code or run tests. These steps help to ensure high code quality and reduce the likelihood that the changes inadvertently break something. Having instructions for running tests is especially helpful if it requires external setup, such as starting a Selenium server for testing in a browser.

## Authors and acknowledgment
Show your appreciation to those who have contributed to the project.

## License
For open source projects, say how it is licensed.

## Project status
If you have run out of energy or time for your project, put a note at the top of the README saying that development has slowed down or stopped completely. Someone may choose to fork your project or volunteer to step in as a maintainer or owner, allowing your project to keep going. You can also make an explicit request for maintainers.

### Provider Test
- After CDC test consumer part is published to packflow, we need to pass to backend side
![Alt text](screenshots/provider.png "Optional Title")
- First of all we need to complete Provider side of the CDC test according to contract which is given from Consumer part.
- We need to arrange this test endpoints and content-type keys according to consumer contract
- Provider test is still failed since there is no server code yet.
- we used the go get gopkg.in/pact-foundation/pact-go.v1 for contract test
### Back-end Unit Test
- Second part of the backend tests is the unit tests, we need to test all business logics in this microservice therefore we write test first.
![Alt text](screenshots/backend.png "Optional Title")
- Tricky point in back-end test we need to use Mock functions of original ones but since there is no function yet we can not used them yet.
### After we finished Back-end
- According to tests, we develop a back-end service and all tests will be fine
![Alt text](screenshots/final.png "Optional Title")


### To set your gopath
- export GOPATH="$HOME/go"
- PATH="\$GOPATH/bin:$PATH"
- https://tutorialedge.net/golang/working-with-environment-variables-in-go/

### We haved used local mongodb docker but after deploy we haved to changed it mongodbatlas
- Previous connection URL: "mongodb://mongoadmin:secret@localhost:27017/?authSource=admin"

### Mock Service 
- For mock service we used mockgen, commands are in Makefile to create mock files.  Another alternative is mockery for go
### go test
- go test -run TestSomethingReallyCool ./folder1/folder2/ -v -count 1

## AMAZON DEPLOY
- We created a docker-compose file to run backend and frontend project in a Amazon EC2 machine
- In pipeline after tests are okay the project is dockerized and push to docker hub.
- Afterwards, in deploy-to-amazon stage maintain a connection to machine in amazon with ssh and docker-compose is pull
and run again with new docker files in docker-hub
- Steps that we follow to accomblish this objective in https://mihajlonesic.gitlab.io/archive/deploy-spring-boot-on-aws-ec2-with-gitlab/
- http://3.85.232.14:8081/ Production


## For update in mongo
	r.TodoElementsCollection.UpdateOne(
		ctx,
		bson.M{"_id": todo.Id},
		bson.D{
			{"$set", bson.D{{"status", todo.Status}}},
		},
	)