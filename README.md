# Kyle's Slides

These are the slides for my talks.

Slides in markdown files are run with [github.com/maaslalani/slides](https://github.com/maaslalani/slides)

Mermaid diagrams are rendered with [github.com/AlexanderGrooff/mermaid-ascii](https://github.com/AlexanderGrooff/mermaid-ascii)

Run these slides without any dependencies using [Dagger](https://dagger.io)!

```
dagger call slides --presentation self_healing_ci.md
```

or without cloning the repo:

```
dagger -m github.com/kpenfound/slides call slides --presentation self_healing_ci.md
```

## Presentations

### 12.11.2025 - Capabilities, APIs, and Experiences

- [slides](./kubecon_na_2025.md)

```
dagger -m github.com/kpenfound/slides call slides --presentation kubecon_na_2025.md
```

### 27.08.2025 - Self-Healing CI

- [slides](./self_healing_ci.md)

```
dagger -m github.com/kpenfound/slides call slides --presentation self_healing_ci.md
```

### 03.07.2025 - Getting Started with GitHub Actions and Dagger

- [slides](./gha_intro.md)

```
dagger -m github.com/kpenfound/slides call slides --presentation gha_intro.md
```

- [youtube video](https://www.youtube.com/watch?v=Ahg-3_Ok-rY)
