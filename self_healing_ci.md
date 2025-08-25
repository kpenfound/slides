---
author: Kyle Penfound, Dagger
date: August 27, 2025
paging: Slide %d / %d
---

# Self Healing CI

Have you ever dreamt of a CI that's always green?

---

## 🤖 The robots have arrived!

Robots are writing code!!

- all the code is production ready      ✅
- human engineers are legacy tech       🧑‍💻
- we can all sit back and drink coffee  ☕

---

### 🔄 The contribution feedback loop

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

- asynchronous between multiple people
- requires context switching

---

### 😅 Pull requests are where dreams die

```
~~~mermaid-ascii
graph TD
Developer1 --> Reviewers
DeveloperN --> Reviewers
~~~
```

---

### 😭 Pull requests are where dreams die

```
~~~mermaid-ascii
graph TD
Developer1 --> Reviewers
DeveloperN --> Reviewers
Agent1 --> Reviewers
AgentN --> Reviewers
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

### ❤️‍🩹 Self healing CI?

A self-healing CI system is one component of "agentic CI"

Why?
- lint errors tell you exactly what the issue is
- some linters can even _just fix it_
- good test failures provide enough detail for an agent to fix

We're adding agents to the development flow, why not also in CI?

---

### 🚧 How do we safely ship it?

A self healing CI system must be
- trustworthy
- reliable
- cost effective

---

### 🤨 Trust

A trustworthy system
- we believe will produce correct output
- cannot ever give us what we've all come to know as "AI slop".

How? The agent in CI needs
- ability to use our actual dev tools for lint, test, build
- context on the task the PR is solving
- context on the code changes in the PR

---

### 🕵️ Reliability

A reliable system is one that we know will consistently fix the problems in CI.

How? Agent evals
- like integration tests for an agentic system
- measure the performance of the agent in a variety of realistic scenarios
- identify where an agent can go wrong
- guide and measure improvements to prompts and tools

---

### 💸 Cost effective

Dont send your entire budget to one of the big AI companies

- Using Anthropic/OpenAI API is usually the most expensive option
- Hosting options in your cloud: AWS Bedrock, Google Vertex, Azure AI Service
  - claude-sonnet-4 on Bedrock is 1:1 pricing with Anthropic's API
  - Amazon Nova Pro costs 10% of claude-sonnet-4
  - gpt-oss-120b costs 5% of claude-sonnet-4
- Run in Kubernetes with GPUs on your nodes
- Local options with Ollama, Docker, llama.cpp, LM Studio

---

## 💻 Demo

       /_\
      /'-'\            |------------*-----------|
    ,'/ ^ ^ \',        |       DEMO TIME!       |
    |  >--@ |          |                        |
    '.\___,/.'         |                        |
                       |________________________|
                      /=========================/
                     /____________===__________/

---

## 🧑‍🎓 Next steps

Check out [dagger.io](https://dagger.io) and join the Discord!

Read my blog post about [the evolution of CI](https://dagger.io/blog/evolution-of-ci)

Find more of my content at [kylepenfound.com](https://kylepenfound.com)

Find these slides at [github.com/kpenfound/slides](https://github.com/kpenfound/slides)
