---
---

{{ range $project := .Projects }}
    {{ $project.Author }}/{{ $project.Name }}
{{ end }}
