{{ define "pflag-init" }}
{{ if $.Pflags}}

{{ $CmdName := "Root" }}
{{ if $.CmdName }}
{{ $CmdName = $.CmdName }}
{{ end }}

{{ $Prefix := "Root" }}
{{ if $.Parent.Parent.Parent.Parent }}
  {{ $Prefix = (print $.Parent.Parent.Parent.Parent.Name "__" $.Parent.Parent.Parent.Name "__" $.Parent.Parent.Name "__" $.Parent.Name "__" $.CmdName) }}
{{ else if $.Parent.Parent.Parent }}
  {{ $Prefix = (print $.Parent.Parent.Parent.Name "__" $.Parent.Parent.Name "__" $.Parent.Name "__" $.CmdName) }}
{{ else if $.Parent.Parent }}
  {{ $Prefix = (print $.Parent.Parent.Name "__" $.Parent.Name "__" $.CmdName) }}
{{ else if $.Parent }}
  {{ $Prefix = (print $.Parent.Name "__" $.CmdName) }}
{{ else if $.CmdName }}
  {{ $Prefix = $.CmdName }}
{{ end }}
{{ $Prefix = ( title $Prefix ) }}

func init () {
  {{ range $i, $F := $.Pflags }}
	{{ $CmdName }}Cmd.PersistentFlags().{{- template "cobra-type" $F.Type -}}VarP(&(flags.{{ $Prefix }}Pflags.{{ $F.FlagName }}), "{{ $F.Long }}", "{{ $F.Short }}", {{ if $F.Default}}{{$F.Default}}{{else}}{{template "go-default" $F.Type }}{{end}}, "{{ $F.Help }}")
  {{- end }}
}
{{ end }}
{{ end }}
