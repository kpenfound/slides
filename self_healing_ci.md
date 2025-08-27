---
author: Kyle Penfound, Dagger
date: August 27, 2025
paging: Slide %d / %d
---

# Self Healing CI

Evolving the software factory for the engineers and the robots

👷 🤝 🤖

---

## 🤖 The robots have arrived!

Robots are writing code!!

- all the code is production ready      ✅
- human engineers are legacy tech       🧑‍💻
- we can all sit back and drink coffee  ☕

*Padme face* 👸🤨 right? right??

---

## 💾 The SDLC

1. Planning & Analysis
2. Design
3. Implementation
4. Testing & Integration
5. Deployment
6. Maintenance

---

## 🔄 The contribution feedback loop

```
~~~mermaid-ascii
flowchart LR
Developer -->|writes code| PullRequest
PullRequest --> CI
CI -->|Errors| Developer
CI --> Review
Review -->|Feedback| Developer
Review --> Merged
~~~
```

---

## 😅 Pull requests yesterday

```
~~~mermaid-ascii
graph TD
Developer1 --> Reviewer1
DeveloperN --> ReviewerN
~~~
```

---

## 😭 Pull requests today

```
~~~mermaid-ascii
graph TD
Developer1 --> Reviewer1
DeveloperN --> ReviewerN
Agent1 --> ReviewerN
AgentN --> ReviewerN
~~~
```

---

## 👷 DevOps to the rescue!

We already have the tools, lets put our powers to use

Big organizations have entire teams that solve these kinds of problems
- DevEx teams: empower developers with better tooling
- Platform teams: provide the systems for automation and validation
- RelEng teams: build out advanced CI pipelines

Open source projects address these problems with
- trustworthy test suites
- a high bar for contributions. If its not good, it gets ignored or closed
- give up and abandon the project

---

## 🏭 Self healing systems in industrial factories

A little industrial history lesson

---

## 🏭 Self healing systems in industrial factories

Consider an automated sprinkler system in a factory:

- **Normal operating conditions**: complete lack of fire 😊
- **Incident detection**: trigger mechanism determines if there is a fire 🔥
- **Incident response**: trigger releases high pressure 🧯
- **Incident response assessment**: some professionals make sure we're safe 🧑‍🚒
- **Incident postmortem**: process changes to prevent future fires 👷

All without AI 🤯

---

## ❤️‍🩹 Self healing CI?

A self-healing CI system is one component of a solution

Why?
- lint errors tell you exactly what the issue is
- some linters can even _just fix it_
- good test failures provide enough detail for an agent to fix

We're adding agents to the development flow, why not also in CI?

---

## 🚧 How do we safely ship it?

A self healing CI system must be
- trustworthy
- reliable
- cost effective

---

## 🤨 Trustworthy

A trustworthy system
- we believe will produce correct output
- cannot ever give us what we've all come to know as "AI slop".

How?

The agent in CI needs:
- ability to use our actual dev tools for lint, test, build
- context on the task the PR is solving
- context on the code changes in the PR

Engineers need observability to monitor and debug the behavior of the agent

---

## 🕵️ Reliable

A reliable system is one that we know will consistently fix the problems in CI.

How? Agent evals
- like integration tests for an agentic system
- measure the performance of the agent in a variety of realistic scenarios
- identify where an agent can go wrong
- guide and measure improvements to prompts and tools

---

## 💸 Cost effective

Dont send your entire budget to one of the big AI companies

- Using Anthropic/OpenAI API is usually the most expensive option
- Hosting options in your cloud: AWS Bedrock, Google Vertex AI, Azure AI Service
  - claude-sonnet-4 on Bedrock is 1:1 pricing with Anthropic's API
  - Amazon Nova Pro costs 10% of claude-sonnet-4
  - gpt-oss-120b costs 5% of claude-sonnet-4
- Run in Kubernetes or Nomad with GPUs on your nodes
- Locally with Ollama, Docker, llama.cpp, LM Studio

---

## 💻 Demo

       /_\
      /'-'\            |===========*============|
    ,'/ ^ ^ \',        |       DEMO TIME!       |
    |  >--@ |          |                        |
    '.\___,/.'         |                        |
                       |________________________|
                      /=========================/
                     /____________===__________/

---

## 🧑‍💻 Your CI

This is Dagger:
```typescript
@func()
async test(): Promise<string> {
  return await dag
    .container()
    .from("cypress/included:14.0.3") // containers!
    .withMountedCache("/root/.npm", dag.cacheVolume("npm-cache")) // cache!
    .withServiceBinding("localhost", this.backend) // service dependencies!
    .withServiceBinding("frontend", this.serve())
    .withWorkdir("/app")
    .withDirectory("/app", this.source) // typed artifacts!
    .withExec(["npm", "ci"])
    .withExec(["npm", "run", "test:e2e"]) // run tests!
    .stdout();
}
```

---

## 🤖 Wheres the robot?

In your Dagger code, if the lint fails, call the agent:

```go
// Run the CI Checks for the project
func (g *Greetings) Check( ... ) ( ... ) {
  // Lint
  lintOutput, err := g.Lint(ctx)
  if err != nil {
    // call agent
  }

  // Then Test
  testOutput, err := g.Test(ctx)
  if err != nil {
    // call agent
  }
}
```

---

## 🤖 Whats the robot do?

- it has context on the lint/test failure
- it has the whole repo at the broken commit
- it has the actual containerized lint/test tools to verify solutions

```go
// an instance of a workspace that has my source code and dagger functions
ws := dag.Workspace(
  g.Frontend.Source(),
  g.Frontend.AsWorkspaceCheckable(),
)
// an environment with the workspace as an input requesting the fixed workspace as an output
env := dag.Env().
  WithWorkspaceInput("workspace", ws, "workspace to read, write, and test code").
  WithWorkspaceOutput("fixed", "workspace with fixed tests")
// give an llm the environment and a prompt
return dag.LLM(dagger.LLMOpts{Model: model}).
  WithEnv(env).
  WithPromptFile(prompt)
```

---

## 🗺️ The bigger picture

- address the bottlenecks
- improve DevEx **all** contributors
  - create portable tools
  - maintain developer docs, not rules files
  - make quality easy
- dont lose determinism where it counts
- build workflows, not technologies
- take care of your humans

---

## 🧑‍🎓 Next steps

Check out [dagger.io](https://dagger.io) and join the Discord!

Read my blog post about [the evolution of CI](https://dagger.io/blog/evolution-of-ci)

Find more of my content at [kylepenfound.com](https://kylepenfound.com)

🦋 kylepenfound.com

Find these slides at [github.com/kpenfound/slides](https://github.com/kpenfound/slides)

The demo repo at [github.com/kpenfound/greetings-api](https://github.com/kpenfound/greetings-api)

```bash
qr -t https://github.com/kpenfound/slides/blob/main/self_healing_ci.md -s 10
```
