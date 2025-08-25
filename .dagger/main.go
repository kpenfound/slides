package main

import (
	"context"
	"dagger/slides/internal/dagger"
)

type Slides struct{}

// Run the slides
func (m *Slides) Slides(
	ctx context.Context,
	// presentation file, e.g. self_healing_ci.md
	presentation string,
	// +defaultPath="/"
	source *dagger.Directory,
) (string, error) {
	return dag.Container().From("cgr.dev/chainguard/wolfi-base:latest").
		With(withSlides).
		With(withMermaidAscii).
		WithWorkdir("/src").
		WithDirectory("/src", source).
		Terminal(
			dagger.ContainerTerminalOpts{
				Cmd: []string{"slides", presentation},
			},
		).
		Stderr(ctx)
}

func withSlides(ctr *dagger.Container) *dagger.Container {
	slides := dag.Git("https://github.com/maaslalani/slides.git").Ref("v0.9.0").Tree()
	build := dag.Container().
		From("golang:latest").
		WithDirectory("/app", slides).
		WithWorkdir("/app").
		WithExec([]string{"go", "build", "-o", "slides"})
	return ctr.WithFile("/bin/slides", build.File("/app/slides"))
}

func withMermaidAscii(ctr *dagger.Container) *dagger.Container {
	slides := dag.Git("https://github.com/AlexanderGrooff/mermaid-ascii.git").Ref("0.7.0").Tree()
	build := dag.Container().
		From("golang:latest").
		WithDirectory("/app", slides).
		WithWorkdir("/app").
		WithExec([]string{"go", "build", "-o", "mermaid-ascii"})
	return ctr.WithFile("/bin/mermaid-ascii", build.File("/app/mermaid-ascii"))
}
