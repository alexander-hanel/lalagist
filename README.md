# lalagist 

`lalagist` is a toy Go project to interact with [Ollama](https://ollama.com/) using Github's Gist.

## Setup 

1. Clone repo 
2. Populate environmental variables within `.env`
```
GITHUB_TOKEN=github_pat_ABCDEFGHIJKLMNOPQRSTUVWXYZ
GITHUB_GIST_ID=75eb740348be801ce17a50eb9559a386
LLM_NAME=llama
MODEL=llama3
AVATAR=https://avatars.githubusercontent.com/u/151674099?s=48&v=3
```
- `GITHUB_TOKEN`:  Create GitHub [token](https://github.com/settings/personal-access-tokens/new), select `
Account permissions`, Navigate to Gists Access, select `Read and Write`
- `GITHUB_GIST_ID`: Create new gist and get the gist id. The gist id can be extracted from the URL. For example, in the URL
`https://gist.github.com/alexander-hanel/75eb740348be801ce17a50eb9559a386` the gist id is `75eb740348be801ce17a50eb9559a386`
- `LLM_NAME` is the name that lalagist should respond to 
- `MODEL` that ollama should use

3. Change directory, `cd lalagist`
4. Build, `go run main.go`
5. Let run 
6. Ask question as a comment, the question must start with the `LLM_NAME` plus `,`. For example,
> llama, can you tell me circumference of the moon?

Wait a minute or so, then lalagist will respond. 

## Example
![](/img/example.png)

## Notes
If browsing on your phone, click on the time stamp of the comment and then press refresh.
