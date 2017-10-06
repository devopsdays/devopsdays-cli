package speaker

const speakerTmpl = `+++
Title = "{{ .Title }}"
type = "speaker"
{{- with .Website }}
website = "{{ . }}"
{{- end }}
{{- with .Twitter }}
twitter = "{{ . }}"
{{- end }}
{{- with .Facebook }}
facebook = "{{ . }}"
{{- end }}
{{- with .Linkedin }}
linkedin = "{{ . }}"
{{- end }}
{{- with .Github }}
github = "{{ . }}"
{{- end }}
{{- with .Gitlab }}
gitlab = "{{ . }}"
{{- end }}
{{- with .ImagePath -}}
image = "{{ . }}"
{{- end }}
+++
{{ with .Bio }}{{.}}{{ end }}
`
